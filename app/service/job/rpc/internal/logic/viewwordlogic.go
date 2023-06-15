package logic

import (
	"context"

	"rpc/app/service/job/rpc/internal/svc"
	"rpc/app/service/job/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewWordLogic {
	return &ViewWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewWordLogic) ViewWord(in *pb.ViewWordReq) (*pb.ViewWordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ViewWordResp{}, nil
}
