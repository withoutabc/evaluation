package mapreduce

import (
	"golang.org/x/crypto/ssh"
	"rpc/app/service/job/rpc/pb"
)

type (
	Model interface {
		WordCount() int32
		JoinData() int32
		JobList() ([]*pb.Jobs, int32)
	}
	DefaultModel struct {
		Ssh *ssh.Client
	}
)

func (d *DefaultModel) WordCount() int32 {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultModel) JoinData() int32 {
	//TODO implement me
	panic("implement me")
}

func (d *DefaultModel) JobList() ([]*pb.Jobs, int32) {
	//TODO implement me
	panic("implement me")
}

func NewModel(ssh *ssh.Client) *DefaultModel {
	return &DefaultModel{
		Ssh: ssh,
	}
}
