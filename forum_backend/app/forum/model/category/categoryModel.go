package category

import (
	"context"
	"database/sql"
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
		SoftDelete(ctx context.Context, id int64) error
		HardDelete(ctx context.Context, id int64) error
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
	// cacheKey := fmt.Sprintf("%s%d_%d", cacheQqForumCategoryListPrefix, page, pageSize)
	var resp []*Category
	offset := (page - 1) * pageSize
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, "SELECT * FROM category WHERE is_active = 1 ORDER BY id ASC LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	return resp, int64(len(resp)), nil
}

func (m *defaultCategoryModel) SoftDelete(ctx context.Context, id int64) error {
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s SET is_active = 0 WHERE id = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cacheKey)
	if err != nil {
		return err
	}
	m.DelCacheCtx(ctx, cacheKey)
	return nil
}

func (m *defaultCategoryModel) HardDelete(ctx context.Context, id int64) error {
	return m.Delete(ctx, id)
}
