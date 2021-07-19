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

type DeleteBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteBaseMenuLogic {
	return DeleteBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body utils.GetById true "菜单id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /menu/deleteBaseMenu [post]
func (l *DeleteBaseMenuLogic) DeleteBaseMenu(req utils.GetById) error {
	if err := utils.Verify(req, verify.IdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.DeleteBaseMenu(req.ID); err != nil {
		global.ZapLog.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("删除成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
