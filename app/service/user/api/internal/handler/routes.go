// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"rpc/app/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CORSMIDDLEWARE},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/register",
					Handler: RegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/login",
					Handler: LoginHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CORSMIDDLEWARE, serverCtx.JWTMIDDLEWARE},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/refresh",
					Handler: RefreshHandler(serverCtx),
				},
			}...,
		),
	)
}
