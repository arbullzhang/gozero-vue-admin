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

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserListLogic {
	return GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body utils.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (l *GetUserListLogic) GetUserList(req utils.PageInfo) error {
	if err := utils.Verify(req, verify.PageInfoVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err, list, total := model.GetUserInfoList(req); err != nil {
		global.ZapLog.Error("获取失败!", zap.Any("err", err))
		utils.FailWithMessage("获取失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", l.svcCtx.ResponseWriter)
	}

	return nil
}