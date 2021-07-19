package logic

import (
	"context"
	"go.uber.org/zap"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

type JsonInBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonInBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) JsonInBlacklistLogic {
	return JsonInBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (l *JsonInBlacklistLogic) JsonInBlacklist() error {
	token := l.svcCtx.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := model.JsonInBlacklist(jwt); err != nil {
		global.ZapLog.Error("jwt作废失败!", zap.Any("err", err))
		utils.FailWithMessage("jwt作废失败", l.svcCtx.ResponseWriter)
		return err
	} else {
		utils.OkWithMessage("jwt作废成功", l.svcCtx.ResponseWriter)
	}
	return nil
}
