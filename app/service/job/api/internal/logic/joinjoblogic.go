package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinJobLogic {
	return &JoinJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinJobLogic) JoinJob(req *types.JoinJobReq) (resp *types.JoinJobResp, err error) {
	//判断是否为46级,1-3套，年份4位数<2024，月份1-12
	if req.Level != 4 && req.Level != 6 {
		return &types.JoinJobResp{
			Status: errs.WrongLevel,
			Msg:    errs.ErrorsMap[errs.WrongLevel].Error(),
		}, nil
	}
	response, err := l.svcCtx.JobRpc.JoinJob(l.ctx, &pb.JoinJobReq{
		Level: req.Level,
		Begin: req.Begin,
		End:   req.End,
	})
	return &types.JoinJobResp{
		Status:  response.StatusCode,
		Msg:     response.StatusMsg,
		JobName: response.JobName,
	}, nil
}
