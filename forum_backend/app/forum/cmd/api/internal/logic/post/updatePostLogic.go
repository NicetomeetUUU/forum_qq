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
	if req.Id <= 0 {
		errstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errstr)
		return l.generateResp(req.Id, 400, errstr), errors.New(errstr)
	}
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find post by id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(req.Id, 400, errstr), err
	}
	// currentUserId := l.ctx.Value("userId").(int64)
	currentUserId := req.UserId
	if postInfo.UserId != currentUserId {
		errstr := "user id not match"
		l.Logger.Infof(errstr)
		return l.generateResp(req.Id, 400, errstr), errors.New(errstr)
	}
	if postInfo.Status != 1 {
		errstr := "post is not published, can't update"
		l.Logger.Infof(errstr)
		return l.generateResp(req.Id, 400, errstr), errors.New(errstr)
	}
	postInfo = l.generatePostInfo(req)
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

func (l *UpdatePostLogic) generateResp(postId int64, code int64, message string) *types.UpdatePostResp {
	return &types.UpdatePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		PostId: postId,
	}
}

func (l *UpdatePostLogic) generatePostInfo(req *types.UpdatePostReq) *post.Post {
	return &post.Post{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		UpdatedTime: time.Now(),
	}
}
