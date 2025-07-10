package handler

import (
	"net/http"

	"forum_backend/api/user/userService/internal/logic"
	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserServiceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserServiceLogic(r.Context(), svcCtx)
		resp, err := l.UserService(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
