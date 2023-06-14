package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"rpc/app/common/middleware"
	"rpc/app/service/job/api/internal/config"
	"rpc/app/service/job/rpc/job"
)

type ServiceContext struct {
	Config config.Config
	JobRpc job.Job

	CORSMIDDLEWARE rest.Middleware
	JWTMIDDLEWARE  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		JobRpc:         job.NewJob(zrpc.MustNewClient(c.JobRpc)),
		CORSMIDDLEWARE: middleware.NewCORSMIDDLEWAREMiddleware().Handle,
		JWTMIDDLEWARE:  middleware.NewJWTMIDDLEWAREMiddleware().Handle,
	}
}
