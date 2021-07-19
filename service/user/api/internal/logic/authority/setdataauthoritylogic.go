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

type SetDataAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetDataAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetDataAuthorityLogic {
	return SetDataAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (l *SetDataAuthorityLogic) SetDataAuthority(req model.SysAuthority) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.SetDataAuthority(req); err != nil {
		global.ZapLog.Error("设置失败!", zap.Any("err", err))
		utils.FailWithMessage("设置失败"+err.Error(), l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("设置成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
