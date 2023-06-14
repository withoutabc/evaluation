package logic

import (
	"context"

	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilesLogic {
	return &GetFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFilesLogic) GetFiles(in *pb.GetFileReq) (*pb.GetFileResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFileResp{}, nil
}
