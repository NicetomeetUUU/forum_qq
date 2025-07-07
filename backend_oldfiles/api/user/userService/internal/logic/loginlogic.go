package logic

import (
	"context"
	"errors"
	"forum_backend/model/user"

	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Validate(req *types.LoginReq) (err error) {
	if req.Email == "" && req.Username == "" {
		return errors.New("邮箱或用户名不能为空")
	}
	if req.Password == "" {
		return errors.New("密码不能为空")
	}
	return nil
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	err = l.Validate(req)
	if err != nil {
		l.Logger.Error("validate failed: %v", err)
		return nil, err
	}
	var user *user.User
	if req.Email != "" {
		user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
		if err != nil {
			l.Logger.Error("find user by email failed: %v", err)
			return nil, err
		}
	} else {
		user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
		if err != nil {
			return nil, err
		}
	}
	if user.Password != req.Password {
		return nil, errors.New("password is incorrect")
	}
	return
}
