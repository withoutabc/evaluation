package handler

import (
	"net/http"
	"path/filepath"
	"rpc/app/service/file/api/internal/logic"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rpc/app/service/file/api/internal/svc"
	"rpc/app/service/file/api/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), svcCtx)
		resp, body, filename, err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			if body == nil {
				httpx.OkJsonCtx(r.Context(), w, resp)
			} else {
				// 设置响应头
				header := w.Header()
				header.Set("Content-Disposition", "attachment; filename="+filepath.Base(filename))
				header.Set("Content-Type", http.DetectContentType(body))
			}
		}
	}
}
