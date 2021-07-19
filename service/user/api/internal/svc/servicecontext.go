package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"net/http"
	"gozero-vue-admin/service/user/api/internal/config"
	"gozero-vue-admin/service/user/api/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config

	JwtObject  	    *middleware.JwtMiddleware
	Jwt  	   	    rest.Middleware

	Casbin    	    rest.Middleware

	OperationRecord rest.Middleware

	Request         *http.Request
	ResponseWriter  http.ResponseWriter
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwtObject := middleware.NewJwtMiddleware(c)
	return &ServiceContext{
		Config: c,
		JwtObject: jwtObject,
		Jwt: jwtObject.Handle,
		Casbin: middleware.NewCasbinMiddleware(jwtObject).Handle,
		OperationRecord: middleware.NewOperationRecordMiddleware(jwtObject).Handle,
	}
}
