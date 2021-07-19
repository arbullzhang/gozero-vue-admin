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

type SetUserAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetUserAuthorityLogic {
	return SetUserAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (l *SetUserAuthorityLogic) SetUserAuthority(req request.SetUserAuth) error {
	if err := utils.Verify(req, verify.SetUserAuthorityVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.SetUserAuthority(req.UUID, req.AuthorityId); err != nil {
		global.ZapLog.Error("修改失败!", zap.Any("err", err))
		utils.FailWithMessage("修改失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("修改成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
