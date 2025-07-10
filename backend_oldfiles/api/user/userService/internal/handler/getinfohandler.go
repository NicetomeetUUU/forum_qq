package handler

import (
	"net/http"

	"forum_backend/api/user/userService/internal/logic"
	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
