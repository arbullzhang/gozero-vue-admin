package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/errorx"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/response"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMenuLogic {
	return GetMenuLogic{
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
// @Router /menu/getMenu [post]
func (l *GetMenuLogic) GetMenu() error {
	token := l.svcCtx.Request.Header.Get("x-token")
	if token == "" {
		return errorx.NewDefaultError("token为空")
	}
	claims, err := l.svcCtx.JwtObject.ParseToken(token)
	var retErr error
	if err != nil {
		global.ZapLog.Error("解析token出错", zap.Any("err", err))
		retErr = err
	}

	if err, menus := model.GetMenuTree(claims.AuthorityId); err != nil {
		global.ZapLog.Error("获取失败!", zap.Any("err", err))
		utils.FailWithMessage("获取失败", l.svcCtx.ResponseWriter)
		retErr = err
	} else {
		utils.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", l.svcCtx.ResponseWriter)
		retErr = nil
	}

	return retErr
}
