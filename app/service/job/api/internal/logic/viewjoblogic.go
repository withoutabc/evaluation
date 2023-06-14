package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
