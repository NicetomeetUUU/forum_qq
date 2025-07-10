package logic

import (
	"context"

	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserServiceLogic {
	return &UserServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserServiceLogic) UserService(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	return
}
