package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type SetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetUserInfoLogic {
	return SetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /user/setUserInfo [put]
func (l *SetUserInfoLogic) SetUserInfo(req model.SysUser) error {
	if err := utils.Verify(req, verify.IdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, ReqUser := model.SetUserInfo(req); err != nil {
		global.ZapLog.Error("设置失败!", zap.Any("err", err))
		utils.FailWithMessage("设置失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(utils.H{"userInfo": ReqUser}, "设置成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
