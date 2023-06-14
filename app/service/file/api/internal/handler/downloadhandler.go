package handler

import (
	"log"
	"net/http"
	"net/url"
	"rpc/app/service/file/api/internal/logic"
	"rpc/app/service/file/errs"

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
		log.Println(resp, filename, err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			if resp.Status != errs.No {
				httpx.OkJsonCtx(r.Context(), w, resp)
			} else {
				// 设置响应头
				header := w.Header()
				header.Set("Content-Disposition", "attachment; filename*=UTF-8''"+url.PathEscape(filename))
				header.Set("Content-Type", "application/octet-stream")
				w.Write(body)
			}
		}
	}
}
