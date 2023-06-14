package logic

import (
	"context"
	"rpc/app/service/user/api/internal/svc"
	"rpc/app/service/user/api/internal/types"
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
	return &types.LoginResp{
		Status: response.StatusCode,
		Msg:    response.StatusMsg,
		Data: types.LoginInfo{
			ID:           response.Data.Id,
			Role:         response.Data.Role,
			LoginTime:    response.Data.LoginTime,
			AccessToken:  response.Data.AccessToken,
			RefreshToken: response.Data.RefreshToken,
		},
	}, nil
}
