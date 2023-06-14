package hdfsinit

import (
	"github.com/colinmarc/hdfs/v2"
)

func InitHdfs(addr, user string) *hdfs.Client {
	// 创建一个HDFS客户端配置
	options := hdfs.ClientOptions{
		Addresses: []string{addr},
		User:      user,
	}
	client, err := hdfs.NewClient(options)
	if err != nil {
		panic(err)
	}
	return client
}
