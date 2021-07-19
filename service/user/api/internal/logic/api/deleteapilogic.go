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

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteApiLogic {
	return DeleteApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /api/deleteApi [post]
func (l *DeleteApiLogic) DeleteApi(req model.SysApi) error {
	if err := utils.Verify(req.BaseModel, verify.IdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.DeleteApi(req); err != nil {
		global.ZapLog.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("删除成功", l.svcCtx.ResponseWriter)
	}
	return nil
}
