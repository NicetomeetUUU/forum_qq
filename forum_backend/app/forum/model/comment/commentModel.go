package comment

import (
	"context"
	"database/sql"
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
		FindCommentsByParentId(ctx context.Context, parentId int64) ([]*Comment, error)
		FindCommentListByPostId(ctx context.Context, postId int64, lastIndex int64, pageSize int64) ([]*Comment, error)
		CountCommentsByPostId(ctx context.Context, postId int64) (int64, error)
		DeleteCommentByParentId(ctx context.Context, parentId int64) error
		DeleteCommentByPostId(ctx context.Context, postId int64) error
		IncreaseLikeCount(ctx context.Context, id int64) error
		DecreaseLikeCount(ctx context.Context, id int64) error
		UpdateCommentContent(ctx context.Context, id int64, content string) error
		UpdateCommentStatusByPostId(ctx context.Context, postId int64, status string) error
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

func (m *customCommentModel) FindCommentsByParentId(ctx context.Context, parentId int64) ([]*Comment, error) {
	var comments []*Comment
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &comments, "SELECT * FROM comment WHERE parent_id = ?", parentId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *customCommentModel) FindCommentListByPostId(ctx context.Context, postId int64, lastIndex int64, pageSize int64) ([]*Comment, error) {
	var comments []*Comment
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &comments, "SELECT * FROM comment WHERE post_id = ? ORDER BY id DESC LIMIT ? OFFSET ?", postId, pageSize, lastIndex)
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

func (m *customCommentModel) DeleteCommentByParentId(ctx context.Context, parentId int64) error {
	return m.deleteCommentRecursively(ctx, parentId)
}

func (m *customCommentModel) deleteCommentRecursively(ctx context.Context, parentId int64) error {
	comments, err := m.FindCommentsByParentId(ctx, parentId)
	if err != nil {
		return err
	}
	postIds := make([]int64, 0)
	cacheKeys := make([]string, 0)
	for _, comment := range comments {
		err = m.deleteCommentRecursively(ctx, comment.Id)
		if err != nil {
			return err
		}
		cacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, comment.Id)
		cacheKeys = append(cacheKeys, cacheKey)
		postIds = append(postIds, comment.PostId)
	}
	if len(comments) > 0 {
		query := fmt.Sprintf("DELETE FROM %s WHERE parent_id = ?", m.table)
		_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
			return conn.ExecCtx(ctx, query, parentId)
		})
		if err != nil {
			return err
		}
		for _, cacheKey := range cacheKeys {
			m.DelCacheCtx(ctx, cacheKey)
		}
		for _, postId := range postIds {
			countCacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentCountByPostIdPrefix, postId)
			m.DelCacheCtx(ctx, countCacheKey)
		}
	}
	return nil
}

func (m *customCommentModel) DeleteCommentByPostId(ctx context.Context, postId int64) error {
	cacheKeys := make([]string, 0)
	query := fmt.Sprintf("SELECT id FROM %s WHERE post_id = ?", m.table)
	var commentIdList []int64
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &commentIdList, query, postId)
	if err != nil {
		return err
	}
	for _, commentId := range commentIdList {
		cacheKeys = append(cacheKeys, fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, commentId))
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE id IN (?)", m.table)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, commentIdList)
	}, cacheKeys...)
	return err
}

func (m *customCommentModel) IncreaseLikeCount(ctx context.Context, id int64) error {
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, id)
	query := fmt.Sprintf("UPDATE %s SET like_count = like_count + 1 WHERE id = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, id)
	}, cacheKey)
	return err
}

func (m *customCommentModel) DecreaseLikeCount(ctx context.Context, id int64) error {
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, id)
	query := fmt.Sprintf("UPDATE %s SET like_count = GREATEST(like_count - 1, 0) WHERE id = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, id)
	}, cacheKey)
	return err
}

func (m *customCommentModel) UpdateCommentContent(ctx context.Context, id int64, content string) error {
	cacheKey := fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, id)
	query := fmt.Sprintf("UPDATE %s SET content = ? WHERE id = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, content, id)
	}, cacheKey)
	return err
}

func (m *customCommentModel) UpdateCommentStatusByPostId(ctx context.Context, postId int64, status string) error {
	comments, err := m.FindCommentsByPostId(ctx, postId)
	if err != nil {
		return err
	}
	cacheKeys := make([]string, 0)
	for _, comment := range comments {
		cacheKeys = append(cacheKeys, fmt.Sprintf("%s%d", cacheQqForumCommentIdPrefix, comment.Id))
	}
	query := fmt.Sprintf("UPDATE %s SET status = ? WHERE post_id = ?", m.table)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, query, status, postId)
	}, cacheKeys...)
	if err != nil {
		return err
	}
	for _, cacheKey := range cacheKeys {
		m.DelCacheCtx(ctx, cacheKey)
	}
	return nil
}
