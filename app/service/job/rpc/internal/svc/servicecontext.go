package svc

import (
	"rpc/app/common/config/hadoop/ssh"
	"rpc/app/service/job/rpc/internal/config"
	"rpc/app/service/job/rpc/internal/mapreduce"
)

type ServiceContext struct {
	Config config.Config

	JobModel mapreduce.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		JobModel: mapreduce.NewModel(ssh.InitSsh(c.Ssh.User, c.Ssh.Password, c.Ssh.Addr)),
	}
}
