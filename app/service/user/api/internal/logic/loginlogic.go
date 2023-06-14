package logic

import (
	"context"
	"log"
	"rpc/app/service/user/api/internal/svc"
	"rpc/app/service/user/api/internal/types"
	"rpc/app/service/user/errs"
	"rpc/app/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	response, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	switch response.StatusCode {
	case errs.InternalServer:
		return &types.LoginResp{
			Status: errs.InternalServer,
			Msg:    errs.ErrorsMap[errs.InternalServer].Error(),
		}, nil
	case errs.WrongPassword:
		log.Println("")
		return &types.LoginResp{
			Status: errs.WrongPassword,
			Msg:    errs.ErrorsMap[errs.WrongPassword].Error(),
		}, nil
	case errs.NoRecord:
		return &types.LoginResp{
			Status: errs.NoRecord,
			Msg:    errs.ErrorsMap[errs.NoRecord].Error(),
		}, nil
	case errs.No:
		return &types.LoginResp{
			Status: errs.No,
			Msg:    errs.ErrorsMap[errs.No].Error(),
			Data: types.LoginInfo{
				ID:           response.Data.Id,
				Role:         response.Data.Role,
				LoginTime:    response.Data.LoginTime,
				AccessToken:  response.Data.AccessToken,
				RefreshToken: response.Data.RefreshToken,
			},
		}, nil
	}
	return &types.LoginResp{
		Status: errs.Unknown,
		Msg:    errs.ErrorsMap[errs.Unknown].Error(),
	}, nil
}
