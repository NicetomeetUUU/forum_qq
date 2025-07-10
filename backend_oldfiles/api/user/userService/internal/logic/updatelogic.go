package logic

import (
	"context"
	"errors"

	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateReq) (resp *types.UpdateResp, err error) {
	l.Logger.Infof("开始更新用户信息: %v", req)
	err = l.Validate(req)
	if err != nil {
		l.Logger.Infof("请求内容有误: %v", err)
		return l.generateResp(400, "请求内容有误!", types.UserInfo{}), err
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		l.Logger.Errorf("查询用户信息失败: %v", err)
		return l.generateResp(400, "查询用户信息失败!", types.UserInfo{}), err
	}
	if userInfo == nil {
		l.Logger.Infof("用户不存在，用户ID：%v", req.UserId)
		return l.generateResp(400, "用户不存在!", types.UserInfo{}), errors.New("用户不存在")
	}

	return
}

func (l *UpdateLogic) Validate(req *types.UpdateReq) (err error) {
	if req.UserId <= 0 {
		return errors.New("用户ID不合法")
	}
	if req.UserInfo.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if req.UserInfo.Username == "" {
		return errors.New("用户名不能为空")
	}
	if req.UserInfo.Password == "" {
		return errors.New("密码不能为空")
	}
	return nil
}

func (l *UpdateLogic) generateResp(code int, message string, userInfo types.UserInfo) *types.UpdateResp {
	return &types.UpdateResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		UserInfo: userInfo,
	}
}
