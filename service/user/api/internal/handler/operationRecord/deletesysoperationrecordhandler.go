package handler

import (
	"net/http"
	"gozero-vue-admin/service/user/model"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/logic/operationRecord"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func DeleteSysOperationRecordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		var req model.SysOperationRecord
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteSysOperationRecordLogic(r.Context(), ctx)
		err := l.DeleteSysOperationRecord(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
