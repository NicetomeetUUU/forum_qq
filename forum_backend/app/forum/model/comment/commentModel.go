package comment

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

var (
	cacheQqForumCommentCountByPostIdPrefix = "cache:qqForum:comment:countByPostId:"
)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindCommentsByPostId(ctx context.Context, postId int64) ([]*Comment, error)
		CountCommentsByPostId(ctx context.Context, postId int64) (int64, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c, opts...),
	}
}

func (m *customCommentModel) FindCommentsByPostId(ctx context.Context, postId int64) ([]*Comment, error) {
	var comments []*Comment
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &comments, "SELECT * FROM comment WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *customCommentModel) CountCommentsByPostId(ctx context.Context, postId int64) (int64, error) {
	var count int64
	targetTable := m.table
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentCountByPostIdPrefix, postId)
	err := m.CachedConn.QueryRowCtx(ctx, &count, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, "SELECT COUNT(*) FROM "+targetTable+" WHERE post_id = ?", postId)
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
