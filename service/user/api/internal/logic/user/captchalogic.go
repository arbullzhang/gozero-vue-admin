package logic

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (l *CaptchaLogic) Captcha() error {
	driver := base64Captcha.NewDriverDigit(l.svcCtx.Config.Captcha.ImgHeight, l.svcCtx.Config.Captcha.ImgWidth, l.svcCtx.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	if id, b64s, err := cp.Generate(); err != nil {
		global.ZapLog.Error("验证码获取失败！", zap.Any("err", err))
		utils.FailWithMessage("验证码获取失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(types.CaptchaResp{
			CaptchaId: id,
			PicPath: b64s,
		}, "验证码获取成功",  l.svcCtx.ResponseWriter)
	}

	return nil
}