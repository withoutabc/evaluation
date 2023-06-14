package svc

import (
	"rpc/app/common/config/database/mysql"
	"rpc/app/service/file/rpc/internal/config"
	"rpc/app/service/file/rpc/internal/model/file"
)

type ServiceContext struct {
	Config config.Config

	FileModel file.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		FileModel: file.NewModel(
			mysql.InitDB(c.Mysql.DataSource),
			nil,
		),
	}
}
