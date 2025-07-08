package user_like

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLikeModel = (*customUserLikeModel)(nil)

type (
	// UserLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLikeModel.
	UserLikeModel interface {
		userLikeModel
	}

	customUserLikeModel struct {
		*defaultUserLikeModel
	}
)

// NewUserLikeModel returns a model for the database table.
func NewUserLikeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserLikeModel {
	return &customUserLikeModel{
		defaultUserLikeModel: newUserLikeModel(conn, c, opts...),
	}
}
