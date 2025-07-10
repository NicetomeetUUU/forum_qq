package task

import (
	"context"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataDeleteTask struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataDeleteTask(svcCtx *svc.ServiceContext) *DataDeleteTask {
	return &DataDeleteTask{
		Logger: logx.WithContext(svcCtx.Config.Context),
		svcCtx: svcCtx,
	}
}

func (t *DataDeleteTask) Run() error {
	postIdList, err := t.svcCtx.PostModel.DeletePostByStatusAndTime(t.ctx, "hidden", time.Now().AddDate(0, 0, -7))
	if err != nil {
		t.Logger.Errorf("delete post by status and time failed: %v", err)
		return err
	}
	for _, postId := range postIdList {
		err = t.svcCtx.CommentModel.DeleteCommentByPostId(t.ctx, postId)
		if err != nil {
			t.Logger.Errorf("delete comment by post id %d failed: %v", postId, err)
		}
	}
	return nil
}
