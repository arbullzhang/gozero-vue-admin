package logic

import (
	"context"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/request"
	"gozero-vue-admin/service/user/api/internal/types/response"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type GetPolicyPathByAuthorityIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyPathByAuthorityIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPolicyPathByAuthorityIdLogic {
	return GetPolicyPathByAuthorityIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (l *GetPolicyPathByAuthorityIdLogic) GetPolicyPathByAuthorityId(req request.CasbinInReceive) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	paths := model.GetPolicyPathByAuthorityId(req.AuthorityId)
	utils.OkWithDetailed(response.PolicyPathResponse{Paths: paths}, "获取成功", l.svcCtx.ResponseWriter)

	return nil
}
