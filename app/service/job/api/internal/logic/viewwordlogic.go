package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/service/job/rpc/pb"

	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewWordLogic {
	return &ViewWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewWordLogic) ViewWord(req *types.ViewWordReq) (resp *types.ViewWordResp, err error) {
	//判断是否为46级,1-3套，年份4位数<2024，月份1-12
	if req.Level != 4 && req.Level != 6 {
		return &types.ViewWordResp{
			Status: errs.WrongLevel,
			Msg:    errs.ErrorsMap[errs.WrongLevel].Error(),
		}, nil
	}
	response, err := l.svcCtx.JobRpc.ViewWord(l.ctx, &pb.ViewWordReq{
		Level: req.Level,
		Begin: req.Begin,
		End:   req.End,
		Count: req.Count,
	})
	return &types.ViewWordResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
		Words:  response.Words,
	}, nil

}
