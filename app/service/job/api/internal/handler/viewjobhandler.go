package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rpc/app/service/job/api/internal/logic"
	"rpc/app/service/job/api/internal/svc"
	"rpc/app/service/job/api/internal/types"
)

func ViewJobHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewJobsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewViewJobLogic(r.Context(), svcCtx)
		resp, err := l.ViewJob(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
