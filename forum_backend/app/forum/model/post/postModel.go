package post

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

var (
	cacheQqForumPostListPrefix              = "cache:qqForum:post:list:"
	cacheQqForumPostCountByCategoryIdPrefix = "cache:qqForum:post:countByCategoryId:"
)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		FindPostList(ctx context.Context, pageSize int64, lastIndex int64, orderBy string, orderType string) ([]*Post, error)
		CountPostsByCategoryId(ctx context.Context, categoryId int64) (int64, error)
		UpdateViewCount(ctx context.Context, postId int64) error
		SoftDelete(ctx context.Context, id int64) error
		HardDelete(ctx context.Context, id int64) error
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn, c, opts...),
	}
}

func (m *customPostModel) FindPostList(ctx context.Context, pageSize int64, lastIndex int64, orderBy string, orderType string) ([]*Post, error) {
	var resp []*Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE status = 1 ORDER BY %s %s LIMIT ? OFFSET ?", m.table, orderBy, orderType)
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, pageSize, lastIndex)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customPostModel) CountPostsByCategoryId(ctx context.Context, categoryId int64) (int64, error) {
	var count int64
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumPostCountByCategoryIdPrefix, categoryId)
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE category_id = ?", m.table)
	err := m.CachedConn.QueryRowCtx(ctx, &count, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, categoryId)
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customPostModel) UpdateViewCount(ctx context.Context, postId int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s SET view_count = view_count + 1 WHERE id = ?", m.table)
		return conn.ExecCtx(ctx, query, postId)
	}, "postId")
	if err != nil {
		return err
	}
	return nil
}

func (m *customPostModel) SoftDelete(ctx context.Context, id int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("UPDATE %s SET status = 2 WHERE id = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, "id")
	return err
}

func (m *customPostModel) HardDelete(ctx context.Context, id int64) error {
	return m.Delete(ctx, id)
}
