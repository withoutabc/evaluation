package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/app/service/job/rpc/pb"

	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectJobLogic {
	return &CollectJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectJobLogic) CollectJob(req *types.CollectJobReq) (resp *types.CollectJobResp, err error) {
	//检查种类是否存在
	if _, ok := maps.WordlistMap[req.Except]; !ok {
		return &types.CollectJobResp{
			Status:  errs.UnknownExcept,
			Msg:     errs.ErrorsMap[errs.UnknownExcept].Error(),
			JobName: "",
		}, nil
	}
	response, err := l.svcCtx.JobRpc.CollectJob(l.ctx, &pb.CollectJobReq{
		Level:  req.Level,
		Begin:  req.Begin,
		End:    req.End,
		Except: req.Except,
	})
	return &types.CollectJobResp{
		Status:  response.StatusCode,
		Msg:     response.StatusMsg,
		JobName: response.JobName,
	}, nil
}
