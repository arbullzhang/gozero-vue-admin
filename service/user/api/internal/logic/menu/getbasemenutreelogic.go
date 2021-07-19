package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/response"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type GetBaseMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBaseMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBaseMenuTreeLogic {
	return GetBaseMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body utils.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getBaseMenuTree [post]
func (l *GetBaseMenuTreeLogic) GetBaseMenuTree() error {
	if err, menus := model.GetBaseMenuTree(); err != nil {
		global.ZapLog.Error("获取失败!", zap.Any("err", err))
		utils.FailWithMessage("获取失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.SysBaseMenusResponse{Menus: menus}, "获取成功", l.svcCtx.ResponseWriter)
	}
	return nil
}
