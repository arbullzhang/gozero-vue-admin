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

type CreateAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateAuthorityLogic {
	return CreateAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (l *CreateAuthorityLogic) CreateAuthority(req model.SysAuthority) error {
	if err := utils.Verify(req, verify.AuthorityVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, authBack := model.CreateAuthority(req); err != nil {
		global.ZapLog.Error("创建失败!", zap.Any("err", err))
		utils.FailWithMessage("创建失败"+err.Error(), l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
