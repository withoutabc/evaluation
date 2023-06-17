package svc

import (
	"github.com/colinmarc/hdfs/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"rpc/app/common/config/hadoop/hdfsinit"
	"rpc/app/common/middleware"
	"rpc/app/service/file/api/internal/config"
	"rpc/app/service/file/rpc/file"
)

type ServiceContext struct {
	Config         config.Config
	FileRpc        file.FileZrpcClient
	HdfsCli        *hdfs.Client
	CORSMIDDLEWARE rest.Middleware
	JWTMIDDLEWARE  rest.Middleware
	AuthMIDDLEWARE rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	c.MaxBytes = 104857600
	return &ServiceContext{
		Config:         c,
		FileRpc:        file.NewFileZrpcClient(zrpc.MustNewClient(c.FileRpc)),
		HdfsCli:        hdfsinit.InitHdfs(c.Hdfs.Addr, c.Hdfs.User),
		CORSMIDDLEWARE: middleware.NewCORSMIDDLEWAREMiddleware().Handle,
		JWTMIDDLEWARE:  middleware.NewJWTMIDDLEWAREMiddleware().Handle,
		AuthMIDDLEWARE: middleware.NewAuthMIDDLEWAREMiddleware().Handle,
	}
}
