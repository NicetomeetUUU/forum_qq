package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		IsUserExist(ctx context.Context, username string, email string) (bool, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *customUserModel) IsUserExist(ctx context.Context, username string, email string) (bool, error) {
	userInfo, err := m.FindOneByUsername(ctx, username)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	if userInfo != nil {
		return true, nil // username already exists
	}
	userInfo, err = m.FindOneByEmail(ctx, email)
	if err != nil && err != ErrNotFound {
		return false, err
	}
	if userInfo != nil {
		return true, nil // email already exists
	}
	return false, nil
}
