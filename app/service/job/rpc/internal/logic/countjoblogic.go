package logic

import (
	"context"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountJobLogic {
	return &CountJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountJobLogic) CountJob(in *pb.CountJobReq) (*pb.CountJobResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CountJobResp{}, nil
}
