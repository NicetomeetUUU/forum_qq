package task

import (
	"context"
	"forum_backend/app/forum/cmd/api/internal/config"
	"forum_backend/app/forum/model/comment"
	"forum_backend/app/forum/model/post"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DataDeleteTask struct {
	logx.Logger
	ctx    context.Context
	config config.Config
}

func NewDataDeleteTask(ctx context.Context, config config.Config) *DataDeleteTask {
	return &DataDeleteTask{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		config: config,
	}
}

func (t *DataDeleteTask) Run() error {
	// 创建独立的数据库连接
	conn := sqlx.NewMysql(t.config.DataSource)

	// 创建模型实例
	postModel := post.NewPostModel(conn, t.config.Cache)
	commentModel := comment.NewCommentModel(conn, t.config.Cache)

	postIdList, err := postModel.DeletePostByStatusAndTime(t.ctx, "hidden", time.Now().AddDate(0, 0, -7))
	if err != nil {
		t.Logger.Errorf("delete post by status and time failed: %v", err)
		return err
	}

	for _, postId := range postIdList {
		err = commentModel.DeleteCommentByPostId(t.ctx, postId)
		if err != nil {
			t.Logger.Errorf("delete comment by post id %d failed: %v", postId, err)
		}
	}

	return nil
}
