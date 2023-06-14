package logic

import (
	"context"
	"rpc/app/service/file/errs"

	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *pb.UploadReq) (*pb.UploadResp, error) {
	code := l.svcCtx.FileModel.Upload(l.ctx, in.Level, in.Year, in.Month, in.Set)
	return &pb.UploadResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
	}, nil

}
