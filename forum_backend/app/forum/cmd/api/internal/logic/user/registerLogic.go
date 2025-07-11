package user

import (
	"context"
	"errors"
	"fmt"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/user"
	"time"

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
	if err = l.checkRegisterReq(req); err != nil {
		errstr := fmt.Sprintf("check register req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	userInfo := l.generateUserInfo(req)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, userInfo)
	if err != nil {
		errstr := fmt.Sprintf("insert user failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	return
}

func (l *RegisterLogic) checkRegisterReq(req *types.RegisterReq) error {
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	isExist, err := l.svcCtx.UserModel.IsUserExist(l.ctx, req.Username, req.Email, req.Phone)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("user already exists")
	}
	return nil
}

func (l *RegisterLogic) generateResp(userId int64, code int64, message string) *types.RegisterResp {
	return &types.RegisterResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		UserId: userId,
	}
}

func (l *RegisterLogic) generateUserInfo(req *types.RegisterReq) *user.User {
	return &user.User{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		Status:      "active",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
}
