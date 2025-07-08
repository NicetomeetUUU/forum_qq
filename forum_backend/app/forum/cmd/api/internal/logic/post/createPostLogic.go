package post

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/post"

	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	if err := l.checkPostInfo(req); err != nil {
		l.Logger.Infof("checkPostInfo error: %v", err)
		return l.generateResp(0, 400, "checkPostInfo error"), err
	}
	userId := l.ctx.Value("userId").(int64)
	sqlResult, err := l.svcCtx.PostModel.Insert(l.ctx, &post.Post{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId,
		CategoryId: sql.NullInt64{
			Int64: req.CategoryId,
			Valid: req.CategoryId != 0,
		},
	})
	if err != nil {
		l.Logger.Errorf("insert post error: %v", err)
		return l.generateResp(0, 400, "insert post error"), err
	}
	postId, err := sqlResult.LastInsertId()
	if err != nil {
		l.Logger.Errorf("get post id error: %v", err)
		return l.generateResp(0, 400, "get post id error"), err
	}
	l.Logger.Infof("create post success! post id: %d", postId)
	resp = l.generateResp(postId, 200, "create post success!")
	return
}

func (l *CreatePostLogic) checkPostInfo(postInfo *types.CreatePostReq) error {
	if postInfo.Title == "" {
		return errors.New("title is required")
	}
	if postInfo.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func (l *CreatePostLogic) generateResp(postId int64, code int64, message string) *types.CreatePostResp {
	return &types.CreatePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		PostId: postId,
	}
}
