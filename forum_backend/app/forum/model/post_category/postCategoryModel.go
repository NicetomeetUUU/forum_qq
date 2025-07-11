package post_category

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostCategoryModel = (*customPostCategoryModel)(nil)

type (
	// PostCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostCategoryModel.
	PostCategoryModel interface {
		postCategoryModel
	}

	customPostCategoryModel struct {
		*defaultPostCategoryModel
	}
)

// NewPostCategoryModel returns a model for the database table.
func NewPostCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostCategoryModel {
	return &customPostCategoryModel{
		defaultPostCategoryModel: newPostCategoryModel(conn, c, opts...),
	}
}
