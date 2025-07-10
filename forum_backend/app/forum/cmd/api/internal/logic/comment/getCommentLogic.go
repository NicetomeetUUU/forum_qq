package comment

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentReq) (resp *types.GetCommentResp, err error) {
	if req.Id <= 0 {
		errstr := fmt.Sprintf("id is required, id: %d must be greater than 0", req.Id)
		l.Logger.Errorf(errstr)
		return l.generateResp(types.CommentInfo{}, 400, errstr), errors.New(errstr)
	}
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find comment failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(types.CommentInfo{}, 400, errstr), err
	}
	resp = l.generateResp(l.generateCommentInfo(comment), 0, "success")
	return
}

func (l *GetCommentLogic) generateResp(data types.CommentInfo, code int64, message string) *types.GetCommentResp {
	return &types.GetCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		Comment: data,
	}
}

func (l *GetCommentLogic) generateCommentInfo(comment *comment.Comment) types.CommentInfo {
	return types.CommentInfo{
		Id:          comment.Id,
		Content:     comment.Content,
		UserId:      comment.UserId,
		PostId:      comment.PostId,
		ParentId:    comment.ParentId.Int64,
		LikeCount:   comment.LikeCount,
		Status:      comment.Status,
		CreatedTime: comment.CreatedAt.Unix(),
		UpdatedTime: comment.UpdatedAt.Unix(),
	}
}
