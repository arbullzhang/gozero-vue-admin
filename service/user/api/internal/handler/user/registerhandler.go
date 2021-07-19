package handler

import (
	"net/http"
	"gozero-vue-admin/service/user/api/internal/types/request"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/user"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req request.Register
		if err := httpx.Parse(r, &req); err != nil  {
			httpx.Error(w, err)
			return
		}
		l := logic.NewRegisterLogic(r.Context(), ctx)
		err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
