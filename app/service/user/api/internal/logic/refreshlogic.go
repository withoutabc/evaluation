package logic

import (
	"context"
	"rpc/app/service/user/errs"
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
	switch response.StatusCode {
	case errs.InternalServer:
		return &types.RefreshResp{
			Status: errs.InternalServer,
			Msg:    errs.ErrorsMap[errs.InternalServer].Error(),
		}, nil
	case errs.WrongTokenType:
		return &types.RefreshResp{
			Status: errs.WrongTokenType,
			Msg:    errs.ErrorsMap[errs.WrongTokenType].Error(),
		}, nil
	case errs.WrongSignature:
		return &types.RefreshResp{
			Status: errs.WrongSignature,
			Msg:    errs.ErrorsMap[errs.WrongSignature].Error(),
		}, nil
	case errs.No:
		return &types.RefreshResp{
			Status:       errs.No,
			Msg:          errs.ErrorsMap[errs.No].Error(),
			AccessToken:  response.AccessToken,
			RefreshToken: response.RefreshToken,
		}, nil
	}
	return &types.RefreshResp{
		Status: errs.Unknown,
		Msg:    errs.ErrorsMap[errs.Unknown].Error(),
	}, nil
}
