package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
