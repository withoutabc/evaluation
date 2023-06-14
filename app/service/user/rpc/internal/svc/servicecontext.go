package svc

import (
	"rpc/app/common/config/database/mysql"
	"rpc/app/service/user/rpc/internal/config"
	"rpc/app/service/user/rpc/internal/model/user"
)

type ServiceContext struct {
	Config config.Config

	UserModel user.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := mysql.InitDB(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: user.NewModel(
			db,
			nil,
		),
	}
}
