package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"rpc/app/common/middleware"
	"rpc/app/service/user/api/internal/config"
	"rpc/app/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.User

	CORSMIDDLEWARE rest.Middleware
	JWTMIDDLEWARE  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CORSMIDDLEWARE: middleware.NewCORSMIDDLEWAREMiddleware().Handle,
		JWTMIDDLEWARE:  middleware.NewJWTMIDDLEWAREMiddleware().Handle,
	}
}
