package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/request"
	"gozero-vue-admin/service/user/api/internal/types/response"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (l *RegisterLogic) Register(req request.Register) error {
	if err := utils.Verify(req, verify.RegisterVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	user := &model.SysUser{Username: req.Username, NickName: req.NickName, Password: req.Password, HeaderImg: req.HeaderImg, AuthorityId: req.AuthorityId}
	err, userReturn := model.Register(*user)
	if err != nil {
		global.ZapLog.Error("注册失败!", zap.Any("err", err))
		utils.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", l.svcCtx.ResponseWriter)
	}
	return nil
}
