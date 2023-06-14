package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rpc/app/service/user/api/internal/logic"
	"rpc/app/service/user/api/internal/svc"
	"rpc/app/service/user/api/internal/types"
)

func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
