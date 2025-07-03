package user

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	UserModel interface {
		userModel
		// batch query
		UsersQuery
		// batch count
		UsersCount
		// create
		UserCreate
		// // single update
		// UserUpdate
		// // batch update
		// UserBatchUpdate
		// // soft delete
		// UserSoftDelete
		// PrintPrivateInfo()
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
