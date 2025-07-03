package user

import (
	"context"
	"errors"
	"time"
)

type UserInfo struct {
	Email    string
	Password string
	Username string
}

type UserCreate interface {
	Validate(ctx context.Context, userInfo UserInfo) error
	CreateUser(ctx context.Context, userInfo UserInfo) (int64, error)
}

func (m *customUserModel) Validate(ctx context.Context, userInfo UserInfo) error {
	if userInfo.Email == "" {
		return errors.New("email is required")
	}
	if userInfo.Password == "" {
		return errors.New("password is required")
	}
	if userInfo.Username == "" {
		return errors.New("username is required")
	}
	existingUser, err := m.FindOneByEmail(ctx, userInfo.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}
	existingUser, err = m.FindOneByUsername(ctx, userInfo.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}
	return nil
}

func (m *customUserModel) CreateUser(ctx context.Context, userInfo UserInfo) (int64, error) {
	// 1. 验证输入参数
	if err := m.Validate(ctx, userInfo); err != nil {
		return 0, err
	}
	// 4. 创建用户对象，设置默认值
	now := time.Now()
	user := &User{
		Email:     userInfo.Email,
		Password:  userInfo.Password, // 注意：实际项目中应该加密密码
		Username:  userInfo.Username,
		Avatar:    "",       // 默认空头像
		Signature: "",       // 默认空签名
		Role:      "user",   // 默认用户角色
		Status:    "active", // 默认激活状态
		IsDeleted: 0,        // 默认未删除
		CreatedAt: now,      // 创建时间
		UpdatedAt: now,      // 更新时间
	}

	// 5. 插入数据库
	result, err := m.Insert(ctx, user)
	if err != nil {
		return 0, err
	}
	// 6. 返回用户ID
	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
