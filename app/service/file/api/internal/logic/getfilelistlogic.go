package logic

import (
	"context"
	"rpc/app/service/file/rpc/pb"

	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileListLogic {
	return &GetFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileListLogic) GetFileList(req *types.GetFileListReq) (resp *types.GetFileListResp, err error) {
	response, err := l.svcCtx.FileRpc.GetFiles(l.ctx, &pb.GetFileReq{
		Level: req.Level,
		Month: req.Month,
		Year:  req.Year,
		Set:   req.Set,
	})
	return &types.GetFileListResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
		List:   response.Data,
	}, nil
}
