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

type FindSysOperationRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSysOperationRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindSysOperationRecordLogic {
	return FindSysOperationRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysOperationRecord true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (l *FindSysOperationRecordLogic) FindSysOperationRecord(req model.SysOperationRecord) error {
	if err := utils.Verify(req, verify.IdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, resysOperationRecord := model.GetSysOperationRecord(req.ID); err != nil {
		global.ZapLog.Error("查询失败!", zap.Any("err", err))
		utils.FailWithMessage("查询失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(utils.H{"resysOperationRecord": resysOperationRecord}, "查询成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
