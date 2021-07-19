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

type UpdateBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateBaseMenuLogic {
	return UpdateBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /menu/updateBaseMenu [post]
func (l *UpdateBaseMenuLogic) UpdateBaseMenu(req model.SysBaseMenu) error {
	if err := utils.Verify(req, verify.MenuVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := utils.Verify(req.Meta, verify.MenuMetaVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.UpdateBaseMenu(req); err != nil {
		global.ZapLog.Error("更新失败!", zap.Any("err", err))
		utils.FailWithMessage("更新失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("更新成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
