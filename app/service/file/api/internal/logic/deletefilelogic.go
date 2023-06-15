package logic

import (
	"context"
	"rpc/app/common/consts/errs"
	"rpc/app/common/consts/maps"
	"rpc/app/service/file/rpc/pb"
	"rpc/utils/joinstring"

	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.RemoveFileReq) (resp *types.RemoveFileResp, err error) {
	response, err := l.svcCtx.FileRpc.RemoveFile(l.ctx, &pb.RemoveFileReq{
		Id: req.Id,
	})
	if response.StatusCode != errs.No {
		return &types.RemoveFileResp{
			Status: response.StatusCode,
			Msg:    response.StatusMsg,
		}, nil
	}
	//删除文件
	str := joinstring.Join(response.File.Year, response.File.Month, response.File.Set, response.File.Level)
	mkdirPath := prefix + string(response.File.Year) + "/" + maps.LevelMap[response.File.Level] + "/" + str
	err = l.svcCtx.HdfsCli.Remove(mkdirPath)
	if err != nil {
		return &types.RemoveFileResp{
			Status: errs.FileSystemInternal,
			Msg:    errs.ErrorsMap[errs.FileSystemInternal].Error(),
		}, nil
	}
	return &types.RemoveFileResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
	}, nil
}
