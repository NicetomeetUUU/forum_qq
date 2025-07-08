package post

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

var (
	cacheQqForumPostListPrefix = "cache:qqForum:post:list:"
)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		FindPostList(ctx context.Context, limit int64, lastId int64, orderBy string, orderType string) ([]*Post, error)
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

func (m *customPostModel) FindPostList(ctx context.Context, limit int64, lastId int64, orderBy string, orderType string) ([]*Post, error) {
	cacheKey := fmt.Sprintf("%s%d_%d_%s_%s", cacheQqForumPostListPrefix, limit, lastId, orderBy, orderType)
	var resp []*Post
	err := m.CachedConn.QueryRowCtx(ctx, &resp, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, "SELECT * FROM post WHERE id > ? ORDER BY ? ? LIMIT ?", lastId, orderBy, orderType, limit)
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
