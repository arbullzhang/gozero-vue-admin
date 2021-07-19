package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/types/request"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type AddMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMenuAuthorityLogic {
	return AddMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /menu/addMenuAuthority [post]
func (l *AddMenuAuthorityLogic) AddMenuAuthority(req request.AddMenuAuthorityInfo) error {
	if err := utils.Verify(req, verify.AuthorityIdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.AddMenuAuthority(req.Menus, req.AuthorityId); err != nil {
		global.ZapLog.Error("添加失败!", zap.Any("err", err))
		utils.FailWithMessage("添加失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("添加成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
