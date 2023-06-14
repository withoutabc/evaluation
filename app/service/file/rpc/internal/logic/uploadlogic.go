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
	switch code {
	case errs.No:
		return &pb.UploadResp{
			StatusCode: errs.No,
			StatusMsg:  errs.ErrorsMap[errs.No].Error()}, nil
	case errs.InternalServer:
		return &pb.UploadResp{
			StatusCode: errs.InternalServer,
			StatusMsg:  errs.ErrorsMap[errs.InternalServer].Error()}, nil
	case errs.FileNotExist:
		return &pb.UploadResp{
			StatusCode: errs.FileNotExist,
			StatusMsg:  errs.ErrorsMap[errs.FileNotExist].Error()}, nil
	}
	return &pb.UploadResp{
		StatusCode: errs.Unknown,
		StatusMsg:  errs.ErrorsMap[errs.Unknown].Error()}, nil
}
