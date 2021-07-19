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

type CopyAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) CopyAuthorityLogic {
	return CopyAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (l *CopyAuthorityLogic) CopyAuthority(req model.SysAuthorityCopyResponse) error {
	if err := utils.Verify(req, verify.OldAuthorityVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := utils.Verify(req.Authority, verify.AuthorityVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, authBack := model.CopyAuthority(req); err != nil {
		global.ZapLog.Error("拷贝失败!", zap.Any("err", err))
		utils.FailWithMessage("拷贝失败"+err.Error(), l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "拷贝成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
