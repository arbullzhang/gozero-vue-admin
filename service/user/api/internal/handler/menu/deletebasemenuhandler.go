package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/logic/menu"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func DeleteBaseMenuHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req utils.GetById
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteBaseMenuLogic(r.Context(), ctx)
		err := l.DeleteBaseMenu(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
