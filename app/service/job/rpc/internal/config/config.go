package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Ssh struct {
		Addr     string
		User     string
		Password string
	}
	Hdfs struct {
		Addr string
		User string
	}
	RedisConf struct {
		Addr     string
		Password string
		DB       int
	}
}
