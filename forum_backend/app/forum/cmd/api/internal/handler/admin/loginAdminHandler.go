package admin

import (
	"net/http"

	"forum_backend/app/forum/cmd/api/internal/logic/admin"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewLoginAdminLogic(r.Context(), svcCtx)
		resp, err := l.LoginAdmin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
