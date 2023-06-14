package logic

import (
	"context"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewJobsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewJobsLogic {
	return &ViewJobsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewJobsLogic) ViewJobs(in *pb.ViewJobsReq) (*pb.ViewJobsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ViewJobsResp{}, nil
}
