package admin

import (
	"context"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginAdminLogic {
	return &LoginAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginAdminLogic) LoginAdmin(req *types.LoginAdminReq) (resp *types.LoginAdminResp, err error) {
	// todo: add your logic here and delete this line

	return
}
