package userLike

import (
	"net/http"

	"forum_backend/app/forum/cmd/api/internal/logic/userLike"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListUserLikesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUserLikesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userLike.NewListUserLikesLogic(r.Context(), svcCtx)
		resp, err := l.ListUserLikes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
