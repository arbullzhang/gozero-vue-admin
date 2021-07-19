package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type DeleteSysOperationRecordByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysOperationRecordByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSysOperationRecordByIdsLogic {
	return DeleteSysOperationRecordByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body utils.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (l *DeleteSysOperationRecordByIdsLogic) DeleteSysOperationRecordByIds(req utils.IdsReq) error {
	if err := model.DeleteSysOperationRecordByIds(req); err != nil {
		global.ZapLog.Error("批量删除失败!", zap.Any("err", err))
		utils.FailWithMessage("批量删除失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("批量删除成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
