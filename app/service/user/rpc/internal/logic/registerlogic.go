package logic

import (
	"context"
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

	return &pb.RegisterResp{
		StatusCode: code,
		StatusMsg:  errs.ErrorsMap[code].Error(),
	}, nil

}
