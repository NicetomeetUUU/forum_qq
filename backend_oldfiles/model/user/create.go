package user

import (
	"context"
	"time"
)

type UserInfo struct {
	Email    string
	Password string
	Username string
}

type UserCreate interface {
	CreateUser(ctx context.Context, userInfo UserInfo) (int64, error)
}

func (m *customUserModel) CreateUser(ctx context.Context, userInfo UserInfo) (int64, error) {
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
	//session jwt 认证 鉴权
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
