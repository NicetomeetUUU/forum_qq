package post

import (
	"context"
	"errors"
	"fmt"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostLogic) GetPost(req *types.GetPostReq) (resp *types.GetPostResp, err error) {
	if req.Id <= 0 {
		errstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errstr)
		return l.generateResp(nil, 400, errstr), errors.New(errstr)
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find post by id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	userId := l.ctx.Value("userId").(int64)
	if post.UserId != userId {
		errstr := "user id not match"
		l.Logger.Infof(errstr)
		return l.generateResp(nil, 400, errstr), errors.New(errstr)
	}
	l.Logger.Infof("get post success! post id: %d", post.Id)
	err = l.svcCtx.PostModel.UpdateViewCount(l.ctx, post.Id)
	if err != nil {
		errstr := fmt.Sprintf("update view count by post id %d failed: %v", post.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	resp = l.generateResp(post, 200, "get post success!")
	return
}

func (l *GetPostLogic) generateResp(post *post.Post, code int64, message string) *types.GetPostResp {
	var postInfo types.PostInfo
	if post != nil {
		postInfo = types.PostInfo{
			Id:          post.Id,
			Title:       post.Title,
			Content:     post.Content,
			UserId:      post.UserId,
			CategoryId:  post.CategoryId.Int64,
			Status:      post.Status,
			IsTop:       post.IsTop,
			IsHot:       post.IsHot,
			CreatedTime: post.CreatedTime.Unix(),
			UpdatedTime: post.UpdatedTime.Unix(),
		}
	}
	return &types.GetPostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		Post: postInfo,
	}
}
