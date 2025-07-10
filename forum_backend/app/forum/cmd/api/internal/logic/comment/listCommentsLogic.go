package comment

import (
	"context"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsLogic {
	return &ListCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentsLogic) ListComments(req *types.ListCommentReq) (resp *types.ListCommentResp, err error) {
	if err = l.checkListCommentsReq(req); err != nil {
		errstr := fmt.Sprintf("check list comments req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 0, 400, errstr), err
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.PostId)
	if err != nil {
		errstr := fmt.Sprintf("find post by id %d failed: %v", req.PostId, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 0, 400, errstr), err
	}
	if post.Status != "published" {
		errstr := fmt.Sprintf("post %d is not published, can't list comments", req.PostId)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 0, 400, errstr), err
	}
	commentList, err := l.svcCtx.CommentModel.FindCommentListByPostId(l.ctx, req.PostId, req.LastIndex, req.PageSize)
	if err != nil {
		errstr := fmt.Sprintf("find comments by post id %d failed: %v", req.PostId, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 0, 400, errstr), err
	}
	lastIndex := int64(0)
	if len(commentList) == int(req.PageSize) {
		lastIndex = req.LastIndex + int64(len(commentList))
	}
	resp = l.generateResp(commentList, lastIndex, 200, "success")
	return
}

func (l *ListCommentsLogic) checkListCommentsReq(req *types.ListCommentReq) error {
	if req.PostId <= 0 {
		return fmt.Errorf("post_id is required, post_id: %d must be greater than 0", req.PostId)
	}
	return nil
}

func (l *ListCommentsLogic) generateResp(data []*comment.Comment, lastIndex int64, code int64, message string) *types.ListCommentResp {
	comments := make([]types.CommentInfo, 0)
	for _, comment := range data {
		comments = append(comments, types.CommentInfo{
			Id:          comment.Id,
			Content:     comment.Content,
			UserId:      comment.UserId,
			PostId:      comment.PostId,
			ParentId:    comment.ParentId.Int64,
			LikeCount:   comment.LikeCount,
			Status:      comment.Status,
			CreatedTime: comment.CreatedAt.Unix(),
			UpdatedTime: comment.UpdatedAt.Unix(),
		})
	}
	return &types.ListCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		Comments:  comments,
		LastIndex: lastIndex,
	}
}
