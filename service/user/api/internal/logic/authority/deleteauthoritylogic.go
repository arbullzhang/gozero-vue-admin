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

type DeleteAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteAuthorityLogic {
	return DeleteAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (l *DeleteAuthorityLogic) DeleteAuthority(req model.SysAuthority) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.DeleteAuthority(&req); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.ZapLog.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败"+err.Error(), l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("删除成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
