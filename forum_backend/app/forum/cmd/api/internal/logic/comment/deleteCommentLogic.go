package comment

import (
	"context"
	"fmt"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	if err = l.checkDeleteCommentReq(req); err != nil {
		errstr := fmt.Sprintf("check delete comment req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), err
	}
	err = l.svcCtx.CommentModel.DeleteCommentByParentId(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("delete children under comment %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), err
	}
	err = l.svcCtx.CommentModel.Delete(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("delete comment failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), err
	}
	l.Logger.Infof("delete comment success")
	resp = l.generateResp(0, "success")
	return
}

func (l *DeleteCommentLogic) checkDeleteCommentReq(req *types.DeleteCommentReq) error {
	if req.Id <= 0 {
		return fmt.Errorf("id is required, id: %d must be greater than 0", req.Id)
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
		return fmt.Errorf("post is not published, can't delete comment")
	}
	return nil
}

func (l *DeleteCommentLogic) generateResp(code int64, message string) *types.DeleteCommentResp {
	return &types.DeleteCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
	}
}
