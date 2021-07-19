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

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMenuAuthorityLogic {
	return GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body utils.GetAuthorityId true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/GetMenuAuthority [post]
func (l *GetMenuAuthorityLogic) GetMenuAuthority(req utils.GetAuthorityId) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, menus := model.GetMenuAuthority(&req); err != nil {
		global.ZapLog.Error("获取失败!", zap.Any("err", err))
		utils.FailWithDetailed(response.SysMenusResponse{Menus: menus}, "获取失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(utils.H{"menus": menus}, "获取成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
