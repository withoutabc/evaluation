package logic

import (
	"context"

	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountJobLogic {
	return &CountJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountJobLogic) CountJob(req *types.CountJobReq) (resp *types.CountJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
