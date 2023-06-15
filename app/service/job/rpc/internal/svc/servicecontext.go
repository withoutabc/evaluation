package svc

import (
	"github.com/colinmarc/hdfs/v2"
	"rpc/app/common/config/hadoop/hdfsinit"
	"rpc/app/common/config/hadoop/ssh"
	"rpc/app/service/job/rpc/internal/config"
	"rpc/app/service/job/rpc/internal/mapreduce/mapreduce"
)

type ServiceContext struct {
	Config config.Config

	HdfsCli  *hdfs.Client
	JobModel mapreduce.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		HdfsCli:  hdfsinit.InitHdfs(c.Hdfs.Addr, c.Hdfs.User),
		JobModel: mapreduce.NewModel(ssh.InitSsh(c.Ssh.User, c.Ssh.Password, c.Ssh.Addr)),
	}
}
