package userLike

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLikeLogic {
	return &DeleteUserLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLikeLogic) DeleteUserLike(req *types.DeleteUserLikeReq) (resp *types.DeleteUserLikeResp, err error) {
	if err := l.checkDeleteUserLikereq(req); err != nil {
		errstr := fmt.Sprintf("checkDeleteUserLikereq failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), nil
	}
	userLike, err := l.svcCtx.UserLikeModel.FindOneByUserIdTargetTypeTargetId(l.ctx, req.UserId, req.TargetType, req.TargetId)
	if err != nil {
		errstr := fmt.Sprintf("find user_like failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(500, errstr), nil
	}
	if userLike == nil {
		errstr := fmt.Sprintf("user %d like %s: %d not found", req.UserId, req.TargetType, req.TargetId)
		l.Logger.Errorf(errstr)
		return l.generateResp(400, errstr), nil
	}
	err = l.svcCtx.UserLikeModel.Delete(l.ctx, userLike.Id)
	if err != nil {
		errstr := fmt.Sprintf("delete user_like failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(500, errstr), nil
	}
	if req.TargetType == "post" {
		err = l.svcCtx.PostModel.DecreaseLikeCount(l.ctx, req.TargetId)
	} else if req.TargetType == "comment" {
		err = l.svcCtx.CommentModel.DecreaseLikeCount(l.ctx, req.TargetId)
	}
	if err != nil {
		errstr := fmt.Sprintf("decrease like count failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(500, errstr), nil
	}
	l.Logger.Infof("delete user_like success, user_like_id: %d", userLike.Id)
	resp = l.generateResp(200, "success")
	return
}

func (l *DeleteUserLikeLogic) checkDeleteUserLikereq(req *types.DeleteUserLikeReq) (err error) {
	if req.UserId <= 0 {
		return errors.New("user_id is required")
	}
	if req.TargetType == "" {
		return errors.New("target_type is required")
	}
	if req.TargetId <= 0 {
		return errors.New("target_id is required")
	}
	if req.TargetType != "post" && req.TargetType != "comment" {
		return errors.New("target_type is invalid")
	}
	if req.TargetType == "post" {
		post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.TargetId)
		if err != nil {
			return errors.New("find post failed")
		}
		if post.Status != "published" {
			return errors.New("post is hidden, can't delete like")
		}
	} else if req.TargetType == "comment" {
		comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.TargetId)
		if err != nil {
			return errors.New("find comment failed")
		}
		if comment.Status != "published" {
			return errors.New("comment is hidden, can't delete like")
		}
	}
	return nil
}

func (l *DeleteUserLikeLogic) generateResp(code int64, msg string) *types.DeleteUserLikeResp {
	return &types.DeleteUserLikeResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: msg,
		},
	}
}
