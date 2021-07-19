package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/api"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func GetAllApisHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		l := logic.NewGetAllApisLogic(r.Context(), ctx)
		err := l.GetAllApis()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
