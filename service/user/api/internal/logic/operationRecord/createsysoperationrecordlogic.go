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

type CreateSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateSysOperationRecordLogic {
	return CreateSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "创建SysOperationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (l *CreateSysOperationRecordLogic) CreateSysOperationRecord(req model.SysOperationRecord) error {
	if err := model.CreateSysOperationRecord(req); err != nil {
		global.ZapLog.Error("创建失败!", zap.Any("err", err))
		utils.FailWithMessage("创建失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("创建成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
