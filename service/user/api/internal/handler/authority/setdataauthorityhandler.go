package handler

import (
	"net/http"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/authority"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func SetDataAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req model.SysAuthority
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSetDataAuthorityLogic(r.Context(), ctx)
		err := l.SetDataAuthority(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
