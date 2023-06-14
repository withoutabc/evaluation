package logic

import (
	"context"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/internal/svc"
	"rpc/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	data, code := l.svcCtx.UserModel.Login(l.ctx, in.Username, in.Password)
	return &pb.LoginResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
		Data:       data,
	}, nil

}
