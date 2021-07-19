package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/errorx"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	verify "gozero-vue-admin/service/user/api/internal/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body utils.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (l *DeleteUserLogic) DeleteUser(req utils.GetById) error {
	if err := utils.Verify(req, verify.IdVerify); err != nil {
		utils.FailWithMessage(err.Error(), l.svcCtx.ResponseWriter)
		return err
	}

	token := l.svcCtx.Request.Header.Get("x-token")
	if token == "" {
		return errorx.NewDefaultError("token为空")
	}
	claims, err := l.svcCtx.JwtObject.ParseToken(token)
	if err != nil {
		utils.FailWithMessage("解析token失败", l.svcCtx.ResponseWriter)
		return err
	}

	jwtId := claims.ID
	if jwtId == uint(req.ID) {
		utils.FailWithMessage("删除失败, 自杀失败", l.svcCtx.ResponseWriter)
		return err
	}
	if err := model.DeleteUser(req.ID); err != nil {
		global.ZapLog.Error("删除失败!", zap.Any("err", err))
		utils.FailWithMessage("删除失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("删除成功", l.svcCtx.ResponseWriter)
	}

	return nil
}
