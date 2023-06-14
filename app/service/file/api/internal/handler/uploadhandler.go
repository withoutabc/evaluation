package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rpc/app/service/file/api/internal/logic"
	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
