package middleware

import (
	"net/http"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/model"
)

type CasbinMiddleware struct {
	jwt *JwtMiddleware
}

func NewCasbinMiddleware(jwtObject *JwtMiddleware) *CasbinMiddleware {
	return &CasbinMiddleware{
		jwt: jwtObject,
	}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-token")
		if token == "" {
			return
		}

		claims, err := m.jwt.ParseToken(token)
		if err != nil {
			utils.FailWithMessage("token解析失败", w)
			return
		}

		// 获取请求的URI
		obj := r.URL.RequestURI()
		// 获取请求方法
		act := r.Method
		// 获取用户的角色
		sub := claims.AuthorityId
		e := model.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			utils.FailWithMessage("权限不足", w)
			return;
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
