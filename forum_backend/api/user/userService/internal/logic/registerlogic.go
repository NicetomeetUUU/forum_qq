package logic

import (
	"context"

	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"
	"forum_backend/model/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	err = l.svcCtx.UserModel.Validate(l.ctx, user.UserInfo{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	})
	if err != nil {
		resp = &types.RegisterResp{
			BaseResp: types.BaseResp{
				Code:    400,
				Message: "参数校验失败",
			},
			UserId: 0,
		}
		return resp, err
	}
	// 创建用户
	userId, err := l.svcCtx.UserModel.CreateUser(
		l.ctx, user.UserInfo{
			Email:    req.Email,
			Password: req.Password,
			Username: req.Username,
		},
	)
	if err != nil {
		resp = &types.RegisterResp{
			BaseResp: types.BaseResp{
				Code:    401,
				Message: "创建用户失败",
			},
			UserId: 0,
		}
		return resp, err
	}
	resp = &types.RegisterResp{
		BaseResp: types.BaseResp{
			Code:    0,
			Message: "success",
		},
		UserId: userId,
	}
	return
}
