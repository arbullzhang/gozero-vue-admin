package handler

import (
	"net/http"
	"gozero-vue-admin/service/user/api/internal/types/request"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/menu"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func AddMenuAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req request.AddMenuAuthorityInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddMenuAuthorityLogic(r.Context(), ctx)
		err := l.AddMenuAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
