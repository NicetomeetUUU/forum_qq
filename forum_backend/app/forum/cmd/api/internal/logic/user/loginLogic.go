package user

import (
	"context"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		errstr := fmt.Sprintf("find user by email %s failed: %v", req.Email, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, "", 0, "", 400, errstr), err
	}
	if userInfo.Password != req.Password {
		errstr := "password is incorrect"
		l.Logger.Errorf(errstr)
		return l.generateResp(0, "", 0, "", 400, errstr), err
	}
	return l.generateResp(userInfo.Id, "", 0, "", 200, "login success"), nil
}

func (l *LoginLogic) generateResp(userId int64, token string, expire int64,
	refreshToken string, code int64, message string) *types.LoginResp {
	return &types.LoginResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		UserId:       userId,
		Token:        token,
		Expire:       expire,
		RefreshToken: refreshToken,
	}
}
