package logic

import (
	"context"
	"rpc/app/service/user/rpc/user"

	"rpc/app/service/user/api/internal/svc"
	"rpc/app/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp *types.RefreshResp, err error) {
	response, err := l.svcCtx.UserRpc.Refresh(l.ctx, &user.RefreshReq{
		RefreshToken: req.RefreshToken,
	})

	return &types.RefreshResp{
		Status:       response.StatusCode,
		Msg:          response.StatusMsg,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}, nil
}
