package handler

import (
	"net/http"
	"rpc/app/service/file/api/internal/logic"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), svcCtx)
		resp, err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
