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
			Words:  nil,
		}, nil
	}
	//检查种类是否存在
	if _, ok := maps.WordlistMap[req.Except]; !ok {
		return &types.ViewWordResp{
			Status: errs.UnknownExcept,
			Msg:    errs.ErrorsMap[errs.UnknownExcept].Error(),
			Words:  nil,
		}, nil
	}
	response, err := l.svcCtx.JobRpc.ViewWord(l.ctx, &pb.ViewWordReq{
		Level:  req.Level,
		Begin:  req.Begin,
		End:    req.End,
		Count:  req.Count,
		Except: req.Except,
	})
	return &types.ViewWordResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
		Words:  response.Words,
	}, nil

}
