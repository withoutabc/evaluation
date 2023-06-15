package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DownloadLogic) Download(in *pb.DownloadReq) (*pb.DownloadResp, error) {
	code := l.svcCtx.FileModel.Download(l.ctx, in.Id, in.Type)
	return &pb.DownloadResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
	}, nil
}
