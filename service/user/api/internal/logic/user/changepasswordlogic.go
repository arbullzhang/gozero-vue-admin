package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/request"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangePasswordLogic {
	return ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (l *ChangePasswordLogic) ChangePassword(req request.ChangePasswordStruct) error {
	if err := utils.Verify(req, verify.ChangePasswordVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}

	user := &model.SysUser{Username: req.Username, Password: req.Password}
	if err, _ := model.ChangePassword(user, req.NewPassword); err != nil {
		global.ZapLog.Error("修改失败!", zap.Any("err", err))
		utils.FailWithMessage("修改失败，原密码与当前账户不符", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("修改成功", l.svcCtx.ResponseWriter)
	}

	return nil
}


