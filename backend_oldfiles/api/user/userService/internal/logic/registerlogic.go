package logic

import (
	"context"
	"errors"
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

func (l *RegisterLogic) Validate(req *types.RegisterReq) (err error) {
	if req.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if req.Username == "" {
		return errors.New("用户名不能为空")
	}
	if req.Password == "" {
		return errors.New("密码不能为空")
	}
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		l.Logger.Errorf("查询目标邮箱失败: %v", err)
		return errors.New("查询目标邮箱失败")
	}
	if user != nil {
		l.Logger.Infof("注册失败，用户邮箱已存在：%v", req.Email)
		return errors.New("用户邮箱已存在")
	}
	user, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Errorf("查询目标用户名失败: %v", err)
		return errors.New("查询目标用户名失败")
	}
	if user != nil {
		l.Logger.Infof("注册失败，用户名已存在：%v", req.Username)
		return errors.New("用户名已存在")
	}
	return nil
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	//validate
	l.Logger.Infof("开始验证请求内容: %v", req)
	err = l.Validate(req)
	if err != nil {
		return l.generateResp(400, "请求内容有误!", 0), err
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
		l.Logger.Errorf("创建用户失败: %v", err)
		return l.generateResp(400, "创建用户失败!", 0), err
	}
	l.Logger.Infof("创建用户成功，用户ID: %v", userId)
	return l.generateResp(0, "创建用户成功!", userId), nil
}

func (l *RegisterLogic) generateResp(code int, message string, userId int64) *types.RegisterResp {
	return &types.RegisterResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		UserId: userId,
	}
}
