package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/menu"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func GetMenuHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		ctx.Request = r
		l := logic.NewGetMenuLogic(r.Context(), ctx)
		err := l.GetMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
