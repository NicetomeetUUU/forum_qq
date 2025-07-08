package category

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoryModel = (*customCategoryModel)(nil)

var (
	cacheQqForumCategoryListPrefix = "cache:qqForum:category:list:"
)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
		FindCategoryList(ctx context.Context, page int64, pageSize int64) ([]*Category, int64, error)
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn, c, opts...),
	}
}

func (m *defaultCategoryModel) FindCategoryList(ctx context.Context, page int64, pageSize int64) ([]*Category, int64, error) {
	cacheKey := fmt.Sprintf("%s%d_%d", cacheQqForumCategoryListPrefix, page, pageSize)
	var resp []*Category
	err := m.CachedConn.QueryRowCtx(ctx, &resp, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, "SELECT * FROM category WHERE is_active = 1 ORDER BY id ASC LIMIT ?,?", (page-1)*pageSize, pageSize)
	})
	if err != nil {
		return nil, 0, err
	}
	return resp, int64(len(resp)), nil
}
