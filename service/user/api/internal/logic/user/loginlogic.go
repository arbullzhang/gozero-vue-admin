package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
	"time"
	"gozero-vue-admin/common/errorx"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/middleware"
	"gozero-vue-admin/service/user/api/internal/svc"
	"gozero-vue-admin/service/user/api/internal/types/request"
	"gozero-vue-admin/service/user/api/internal/types/response"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (l *LoginLogic) Login(req request.Login) error {
	if err := utils.Verify(req, verify.LoginVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}

	if base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.Captcha, true) {
		user := model.SysUser{Username: req.Username, Password: req.Password}
		req.Password = utils.MD5V([]byte(req.Password))
		err := global.GormDB.Where("username = ? AND password = ?", req.Username, req.Password).Preload("Authority").First(&user).Error
		if err != nil {
			global.ZapLog.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			utils.FailWithMessage("用户名不存在或者密码错误", l.svcCtx.ResponseWriter)
			return err
		} else {
			//now := time.Now().Unix()
			//jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, user.ID, user.UUID, user.NickName, user.Username)
			//if err != nil {
			//	utilResponse.FailWithMessage("获取token错误", l.svcCtx.ResponseWriter)
			//	return err
			//} else {
			//	utilResponse.OkWithDetailed(response.LoginResponse{
			//		User:      user,
			//		Token:     jwtToken,
			//		ExpiresAt: l.svcCtx.Config.Auth.AccessExpire * 1000,
			//	}, "登录成功", l.svcCtx.ResponseWriter)
			//	return nil
			//}
			return l.tokenNext(user)
		}
	} else {
		utils.FailWithMessage("验证码错误", l.svcCtx.ResponseWriter)
		return errorx.NewDefaultError("验证码错误")
	}
}

// 使用go-zero本身自带的jwt中间件
//func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId uint, uuid uuid.UUID, nickName string, username string) (string, error) {
//	claims := make(jwt.MapClaims)
//	claims["exp"] = iat + seconds
//	claims["iat"] = iat
//	claims["user"] = userId
//	claims["uuid"] = uuid
//	claims["nickName"] = nickName
//	claims["username"] = username
//	token := jwt.New(jwt.SigningMethodHS256)
//	token.Claims = claims
//	return token.SignedString([]byte(secretKey))
//}

func (l *LoginLogic) tokenNext(user model.SysUser) error {
	claims := middleware.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		BufferTime:  l.svcCtx.Config.Auth.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + l.svcCtx.Config.Auth.AccessExpire, // 过期时间 7天
			Issuer: "gozero-vue-admins",												  // 签名的发行者
		},
	}

	token, err := l.svcCtx.JwtObject.CreateToken(claims)
	if err != nil {
		global.ZapLog.Error("获取token失败!", zap.Any("err", err))
		utils.FailWithMessage("获取token失败", l.svcCtx.ResponseWriter)
		return err
	}

	if !l.svcCtx.Config.System.UseMultipoint {
		utils.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", l.svcCtx.ResponseWriter)

		return nil
	}

	if err, jwtStr := model.GetRedisJWT(user.Username); err == redis.Nil {
		if err := model.SetRedisJWT(token, user.Username, l.svcCtx.Config.Auth.AccessExpire); err != nil {
			global.ZapLog.Error("设置登录状态失败!", zap.Any("err", err))
			utils.FailWithMessage("设置登录状态失败", l.svcCtx.ResponseWriter)
			return err
		}
		utils.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", l.svcCtx.ResponseWriter)
	} else if err != nil {
		global.ZapLog.Error("设置登录状态失败!", zap.Any("err", err))
		utils.FailWithMessage("设置登录状态失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := model.JsonInBlacklist(blackJWT); err != nil {
			utils.FailWithMessage("jwt作废失败", l.svcCtx.ResponseWriter)
			return err
		}
		if err := model.SetRedisJWT(token, user.Username, l.svcCtx.Config.Auth.AccessExpire); err != nil {
			utils.FailWithMessage("设置登录状态失败", l.svcCtx.ResponseWriter)
			return err
		}

		utils.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", l.svcCtx.ResponseWriter)
	}

	return nil
}