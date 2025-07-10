package userLike

import (
	"context"
	"errors"
	"fmt"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/user_like"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLikeLogic {
	return &CreateUserLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLikeLogic) CreateUserLike(req *types.CreateUserLikeReq) (resp *types.CreateUserLikeResp, err error) {
	if err := l.checkUserLikereq(req); err != nil {
		errstr := fmt.Sprintf("checkUserLikereq failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), nil
	}
	sqlResult, err := l.svcCtx.UserLikeModel.Insert(l.ctx, l.generateUserLike(req))
	if err != nil {
		errstr := fmt.Sprintf("insert user_like failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 500, errstr), nil
	}
	userLikeId, err := sqlResult.LastInsertId()
	if err != nil {
		errstr := fmt.Sprintf("get user_like id failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 500, errstr), nil
	}
	if req.TargetType == "post" {
		err = l.svcCtx.PostModel.IncreaseLikeCount(l.ctx, req.TargetId)
	} else if req.TargetType == "comment" {
		err = l.svcCtx.CommentModel.IncreaseLikeCount(l.ctx, req.TargetId)
	}
	if err != nil {
		errstr := fmt.Sprintf("increase like count failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 500, errstr), nil
	}
	l.Logger.Infof("create user_like success, user_like_id: %d", userLikeId)
	resp = l.generateResp(userLikeId, 200, "success")
	return
}

func (l *CreateUserLikeLogic) checkUserLikereq(req *types.CreateUserLikeReq) (err error) {
	if req.UserId <= 0 {
		return errors.New("user_id is required")
	}
	if req.TargetType == "" {
		return errors.New("target_type is required")
	} else if req.TargetType != "post" && req.TargetType != "comment" {
		return errors.New("target_type is invalid")
	}
	if req.TargetId <= 0 {
		return errors.New("target_id is required")
	}
	if req.TargetType == "post" {
		post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.TargetId)
		if err != nil {
			return errors.New("select post failed")
		}
		if post == nil {
			return errors.New("post not found")
		}
		if post.Status != "published" {
			return errors.New("post is hidden, can't like")
		}
	} else if req.TargetType == "comment" {
		comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.TargetId)
		if err != nil {
			return errors.New("select comment failed")
		}
		if comment == nil {
			return errors.New("comment not found")
		}
		if comment.Status != "published" {
			return errors.New("comment is hidden, can't like")
		}
	}
	return nil
}

func (l *CreateUserLikeLogic) generateResp(userLikeId int64, code int64, msg string) *types.CreateUserLikeResp {
	return &types.CreateUserLikeResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: msg,
		},
		UserLikeId: userLikeId,
	}
}

func (l *CreateUserLikeLogic) generateUserLike(req *types.CreateUserLikeReq) (userLike *user_like.UserLike) {
	userLike = &user_like.UserLike{
		UserId:     req.UserId,
		TargetType: req.TargetType,
		TargetId:   req.TargetId,
		CreatedAt:  time.Now(),
	}
	return
}
