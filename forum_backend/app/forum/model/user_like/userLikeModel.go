package user_like

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLikeModel = (*customUserLikeModel)(nil)

var (
	cacheQqForumUserLikeCountPrefix       = "cache:qqForum:userLike:count:"
	cacheQqForumUserLikeCountTargetPrefix = "cache:qqForum:userLike:count:target:"
)

type (
	// UserLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLikeModel.
	UserLikeModel interface {
		userLikeModel
		FindLikeListByUserId(ctx context.Context, userId int64) ([]*UserLike, error)
		CountUserLikeByUserId(ctx context.Context, userId int64) (int64, error)
		CountUserLikeByTargetTypeTargetId(ctx context.Context, targetType string, targetId int64) (int64, error)
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

func (m *customUserLikeModel) FindLikeListByUserId(ctx context.Context, userId int64) (resp []*UserLike, err error) {
	query := fmt.Sprintf("select %s from %s where 'user_id' = ?", userLikeRows, m.table)
	err = m.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return
}

func (m *customUserLikeModel) CountUserLikeByUserId(ctx context.Context, userId int64) (int64, error) {
	cacheKey := fmt.Sprintf("%s%v", cacheQqForumUserLikeCountPrefix, userId)
	var count int64
	query := fmt.Sprintf("select count(*) from %s where 'user_id' = ?", m.table)
	err := m.CachedConn.QueryRowCtx(ctx, &count, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) (e error) {
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customUserLikeModel) CountUserLikeByTargetTypeTargetId(ctx context.Context, targetType string, targetId int64) (int64, error) {
	cacheKey := fmt.Sprintf("%s%s:%v", cacheQqForumUserLikeCountTargetPrefix, targetType, targetId)
	query := fmt.Sprintf("select count(*) from %s where target_type = ? and target_id = ?", m.table)
	var count int64
	err := m.CachedConn.QueryRowCtx(ctx, &count, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) (e error) {
		return conn.QueryRowCtx(ctx, v, query, targetType, targetId)
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
