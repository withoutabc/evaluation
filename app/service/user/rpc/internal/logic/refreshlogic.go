package logic

import (
	"context"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/internal/svc"
	"rpc/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *pb.RefreshReq) (*pb.RefreshResp, error) {
	newAToken, newRToken, code := l.svcCtx.UserModel.Refresh(l.ctx, in.RefreshToken)
	switch code {
	case errs.InternalServer:
		return &pb.RefreshResp{
			StatusCode: errs.InternalServer,
		}, nil
	case errs.WrongSignature:
		return &pb.RefreshResp{
			StatusCode: errs.WrongSignature,
		}, nil
	case errs.WrongTokenType:
		return &pb.RefreshResp{
			StatusCode: errs.WrongTokenType,
		}, nil
	case errs.No:
		return &pb.RefreshResp{
			StatusCode:   errs.No,
			AccessToken:  newAToken,
			RefreshToken: newRToken,
		}, nil
	}
	return &pb.RefreshResp{
		StatusCode: errs.Unknown,
		StatusMsg:  errs.ErrorsMap[errs.Unknown].Error()}, nil
}
