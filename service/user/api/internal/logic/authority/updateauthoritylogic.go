package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/response"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type UpdateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateAuthorityLogic {
	return UpdateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (l *UpdateAuthorityLogic) UpdateAuthority(req model.SysAuthority) error {
	if err := utils.Verify(req, verify.AuthorityVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, authority := model.UpdateAuthority(req); err != nil {
		global.ZapLog.Error("更新失败!", zap.Any("err", err))
		utils.FailWithMessage("更新失败"+err.Error(), l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.SysAuthorityResponse{Authority: authority}, "更新成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
