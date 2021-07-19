package handler

import (
	"net/http"
	logic "gozero-vue-admin/service/user/api/internal/logic/user"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero-vue-admin/service/user/api/internal/svc"
)

func CaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.ResponseWriter = w
		l := logic.NewCaptchaLogic(r.Context(), ctx)
		err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
