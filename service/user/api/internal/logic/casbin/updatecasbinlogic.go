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

type UpdateCasbinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCasbinLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateCasbinLogic {
	return UpdateCasbinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]
func (l *UpdateCasbinLogic) UpdateCasbin(req request.CasbinInReceive) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.UpdateCasbin(req.AuthorityId, req.CasbinInfos); err != nil {
		global.ZapLog.Error("更新失败!", zap.Any("err", err))
		utils.FailWithMessage("更新失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("更新成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
