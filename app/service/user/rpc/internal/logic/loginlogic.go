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
	switch code {
	case errs.InternalServer:
		return &pb.LoginResp{
			StatusCode: errs.InternalServer,
		}, nil
	case errs.WrongPassword:
		return &pb.LoginResp{
			StatusCode: errs.WrongPassword,
		}, nil
	case errs.NoRecord:
		return &pb.LoginResp{
			StatusCode: errs.NoRecord,
		}, nil
	case errs.No:
		return &pb.LoginResp{
			StatusCode: errs.No,
			StatusMsg:  errs.ErrorsMap[errs.No].Error(),
			Data:       data,
		}, nil
	}
	return &pb.LoginResp{
		StatusCode: errs.Unknown,
		StatusMsg:  errs.ErrorsMap[errs.Unknown].Error()}, nil
}
