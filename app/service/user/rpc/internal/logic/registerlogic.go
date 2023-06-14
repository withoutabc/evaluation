package logic

import (
	"context"
	"log"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/internal/svc"
	"rpc/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	code := l.svcCtx.UserModel.Register(l.ctx, in.Username, in.Password)
	log.Printf("code:%d", code)

	switch code {
	case errs.InternalServer:
		return &pb.RegisterResp{
			StatusCode: errs.InternalServer,
		}, nil
	case errs.RepeatedUsername:
		return &pb.RegisterResp{
			StatusCode: errs.RepeatedUsername,
		}, nil
	case errs.No:
		return &pb.RegisterResp{
			StatusCode: errs.No,
			StatusMsg:  errs.ErrorsMap[errs.No].Error(),
		}, nil
	}
	return &pb.RegisterResp{
		StatusCode: errs.Unknown,
		StatusMsg:  errs.ErrorsMap[errs.Unknown].Error()}, nil
}
