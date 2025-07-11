package post

import (
	"context"
	"errors"
	"fmt"
	"time"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostReq) (resp *types.UpdatePostResp, err error) {
	originalPostInfo, err := l.checkUpdatePostReq(req)
	if err != nil {
		errstr := fmt.Sprintf("check update post req failed, err: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(req.Id, 400, errstr), err
	}
	postInfo := l.generatePostInfo(originalPostInfo, req)
	err = l.svcCtx.PostModel.Update(l.ctx, postInfo)
	if err != nil {
		errstr := fmt.Sprintf("update post by id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(req.Id, 400, errstr), err
	}
	l.Logger.Infof("update post success! post id: %d", req.Id)
	resp = l.generateResp(req.Id, 200, "update post success!")
	return
}

func (l *UpdatePostLogic) checkUpdatePostReq(req *types.UpdatePostReq) (*post.Post, error) {
	if req.Id <= 0 {
		errstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errstr)
		return nil, errors.New(errstr)
	}
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find post by id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return nil, err
	}
	if postInfo == nil {
		errstr := fmt.Sprintf("post not found, id: %d", req.Id)
		l.Logger.Infof(errstr)
		return nil, errors.New(errstr)
	}
	if postInfo.Status != "published" {
		errstr := "post is not published, can't update"
		l.Logger.Infof(errstr)
		return nil, errors.New(errstr)
	}
	if postInfo.UserId != req.UserId {
		errstr := "user id not match"
		l.Logger.Infof(errstr)
		return nil, errors.New(errstr)
	}
	return postInfo, nil
}

func (l *UpdatePostLogic) generateResp(postId int64, code int64, message string) *types.UpdatePostResp {
	return &types.UpdatePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		PostId: postId,
	}
}

func (l *UpdatePostLogic) generatePostInfo(originalPostInfo *post.Post, req *types.UpdatePostReq) *post.Post {
	updatedPostInfo := originalPostInfo
	if req.Title != "" {
		updatedPostInfo.Title = req.Title
	}
	if req.Content != "" {
		updatedPostInfo.Content = req.Content
	}
	updatedPostInfo.UpdatedTime = time.Now()
	return updatedPostInfo
}
