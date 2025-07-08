package admin

import (
	"context"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAdminLogic {
	return &UpdateAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAdminLogic) UpdateAdmin(req *types.UpdateAdminReq) (resp *types.UpdateAdminResp, err error) {
	// todo: add your logic here and delete this line

	return
}
