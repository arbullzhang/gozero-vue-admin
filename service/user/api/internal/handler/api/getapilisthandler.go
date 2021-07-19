package handler

import (
	"net/http"
	"gozero-vue-admin/service/user/api/internal/types/request"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/api"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func GetApiListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req request.SearchApiParams
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetApiListLogic(r.Context(), ctx)
		err := l.GetApiList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
