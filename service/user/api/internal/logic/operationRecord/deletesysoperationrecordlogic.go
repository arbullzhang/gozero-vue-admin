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

type DeleteSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteSysOperationRecordLogic {
	return DeleteSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (l *DeleteSysOperationRecordLogic) DeleteSysOperationRecord(req model.SysOperationRecord) error {
	if err := model.DeleteSysOperationRecord(req); err != nil {
		global.ZapLog.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("删除成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
