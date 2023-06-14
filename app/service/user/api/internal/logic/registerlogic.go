package logic

import (
	"context"
	"rpc/app/service/user/api/internal/svc"
	"rpc/app/service/user/api/internal/types"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	//发给rpc
	response, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	switch response.StatusCode {
	case errs.RepeatedUsername:
		return &types.RegisterResp{
			Status: errs.RepeatedUsername,
			Msg:    errs.ErrorsMap[errs.RepeatedUsername].Error(),
		}, nil
	case errs.InternalServer:
		return &types.RegisterResp{
			Status: errs.InternalServer,
			Msg:    errs.ErrorsMap[errs.InternalServer].Error(),
		}, nil
	case errs.No:
		return &types.RegisterResp{
			Status: errs.No,
			Msg:    errs.ErrorsMap[errs.No].Error(),
		}, nil
	}
	return &types.RegisterResp{
		Status: errs.Unknown,
		Msg:    errs.ErrorsMap[errs.Unknown].Error(),
	}, nil
}
