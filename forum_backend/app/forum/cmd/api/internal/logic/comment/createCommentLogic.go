package comment

import (
	"context"
	"errors"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/comment"

	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	if err = l.checkCreateCommentReq(req); err != nil {
		l.Logger.Infof("check create comment req error: %v", err)
		return l.generateResp(-1, 400, "check create comment req error"), err
	}
	comment := &comment.Comment{
		Content: req.Content,
		PostId:  req.PostId,
		UserId:  req.UserId,
		ParentId: sql.NullInt64{
			Int64: req.ParentId,
			Valid: req.ParentId != 0,
		},
	}
	sqlResult, err := l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Logger.Errorf("insert comment error: %v", err)
		return l.generateResp(-1, 400, "insert comment error"), err
	}
	commentId, err := sqlResult.LastInsertId()
	if err != nil {
		l.Logger.Errorf("get last insert id error: %v", err)
		return l.generateResp(-1, 400, "get last insert id error"), err
	}
	l.Logger.Infof("create comment success, commentId: %d", commentId)
	resp = l.generateResp(commentId, 200, "create comment success")
	return
}

func (l *CreateCommentLogic) checkCreateCommentReq(req *types.CreateCommentReq) error {
	if req.Content == "" {
		return errors.New("content is required")
	}
	currentUserId := l.ctx.Value("currentUserId").(int64)
	if req.UserId != currentUserId {
		return errors.New("user id is not match")
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.PostId)
	if err != nil {
		return errors.New("post not found")
	}
	if post.Status != 1 {
		return errors.New("post is not published")
	}
	if req.ParentId != 0 {
		parentComment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.ParentId)
		if err != nil {
			return errors.New("parent comment not found")
		}
		if parentComment.Status != "published" {
			return errors.New("parent comment is not published")
		}
		if parentComment.PostId != req.PostId {
			return errors.New("parent comment is not in the same post")
		}
	}
	return nil
}

func (l *CreateCommentLogic) generateResp(commentId int64, code int64, message string) *types.CreateCommentResp {
	return &types.CreateCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		CommentId: commentId,
	}
}
