package admin

import (
	"context"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminLogic {
	return &DeleteAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAdminLogic) DeleteAdmin(req *types.DeleteAdminReq) (resp *types.DeleteAdminResp, err error) {
	// todo: add your logic here and delete this line

	return
}
