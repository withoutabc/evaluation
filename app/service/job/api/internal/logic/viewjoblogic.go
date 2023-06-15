package logic

import (
	"context"
	"rpc/app/service/job/rpc/pb"

	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewJobLogic {
	return &ViewJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewJobLogic) ViewJob(req *types.ViewJobsReq) (resp *types.ViewJobsResp, err error) {
	response, err := l.svcCtx.JobRpc.ViewJobs(l.ctx, &pb.ViewJobsReq{Id: req.Id})
	return &types.ViewJobsResp{
		Status:  response.StatusCode,
		Msg:     response.StatusMsg,
		JobList: response.Jobs,
	}, nil
}
