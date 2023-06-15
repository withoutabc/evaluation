package logic

import (
	"context"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinJobLogic {
	return &JoinJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinJobLogic) JoinJob(in *pb.JoinJobReq) (*pb.JoinJobResp, error) {
	// todo: add your logic here and delete this line

	return &pb.JoinJobResp{}, nil
}
