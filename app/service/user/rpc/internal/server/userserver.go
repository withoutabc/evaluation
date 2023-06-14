// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"rpc/app/service/user/rpc/internal/logic"
	"rpc/app/service/user/rpc/internal/svc"
	"rpc/app/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Refresh(ctx context.Context, in *pb.RefreshReq) (*pb.RefreshResp, error) {
	l := logic.NewRefreshLogic(ctx, s.svcCtx)
	return l.Refresh(in)
}