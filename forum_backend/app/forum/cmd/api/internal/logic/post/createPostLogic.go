package post

import (
	"context"
	"errors"
	"fmt"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/post"
	"time"

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
		errstr := fmt.Sprintf("check post info failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	// userId := l.ctx.Value("userId").(int64)
	userId := req.UserId
	postInfo := l.generatePostInfo(req, userId)
	sqlResult, err := l.svcCtx.PostModel.Insert(l.ctx, postInfo)
	if err != nil {
		errstr := fmt.Sprintf("insert post failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	postId, err := sqlResult.LastInsertId()
	if err != nil {
		errstr := fmt.Sprintf("get post id failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
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

func (l *CreatePostLogic) generatePostInfo(postInfo *types.CreatePostReq, userId int64) *post.Post {
	return &post.Post{
		Title:        postInfo.Title,
		Content:      postInfo.Content,
		UserId:       userId,
		ViewCount:    0,
		LikeCount:    0,
		CommentCount: 0,
		Status:       "published",
		CreatedTime:  time.Now(),
		UpdatedTime:  time.Now(),
	}
}
