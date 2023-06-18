package logic

import (
	"context"
	"rpc/app/common/consts/errs"
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

	return &pb.RefreshResp{
		StatusCode:   code,
		StatusMsg:    errs.ErrorsMap[code].Error(),
		AccessToken:  newAToken,
		RefreshToken: newRToken,
	}, nil

}
