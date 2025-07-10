package comment

import (
	"context"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *types.UpdateCommentReq) (resp *types.UpdateCommentResp, err error) {
	if err = l.checkUpdateCommentReq(req); err != nil {
		errstr := fmt.Sprintf("check update comment req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), err
	}
	err = l.svcCtx.CommentModel.UpdateCommentContent(l.ctx, req.Id, req.Content)
	if err != nil {
		errstr := fmt.Sprintf("update comment failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), err
	}
	resp = l.generateResp(0, "success")
	return
}

func (l *UpdateCommentLogic) checkUpdateCommentReq(req *types.UpdateCommentReq) error {
	if req.Id <= 0 {
		return fmt.Errorf("id is required, id: %d must be greater than 0", req.Id)
	}
	if req.Content == "" {
		return fmt.Errorf("content is required, content: %s must be not empty", req.Content)
	}
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return fmt.Errorf("find comment failed: %v", err)
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, comment.PostId)
	if err != nil {
		return fmt.Errorf("find post failed: %v", err)
	}
	if post.Status != "published" {
		return fmt.Errorf("post is not published, can't update comment")
	}
	return nil
}

func (l *UpdateCommentLogic) generateResp(code int64, message string) *types.UpdateCommentResp {
	return &types.UpdateCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
	}
}
