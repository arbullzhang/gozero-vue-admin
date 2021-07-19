package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/user"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func JsonInBlacklistHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		ctx.Request = r
		l := logic.NewJsonInBlacklistLogic(r.Context(), ctx)
		err := l.JsonInBlacklist()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
