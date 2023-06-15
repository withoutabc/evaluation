package logic

import (
	"context"
	"rpc/app/common/consts/errs"

	"rpc/app/service/file/rpc/internal/svc"
	"rpc/app/service/file/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFileLogic {
	return &RemoveFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFileLogic) RemoveFile(in *pb.RemoveFileReq) (*pb.RemoveFileResp, error) {
	file, code := l.svcCtx.FileModel.DeleteFile(l.ctx, in.Id)
	return &pb.RemoveFileResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
		File:       file,
	}, nil
}
