package post

import (
	"context"
	"errors"

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
		l.Logger.Infof("id is invalid")
		return l.generateResp(nil, 400, "id is invalid"), errors.New("id is invalid")
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("find post error: %v", err)
		return l.generateResp(nil, 400, "find post error"), err
	}
	userId := l.ctx.Value("userId").(int64)
	if post.UserId != userId {
		l.Logger.Errorf("user id not match")
		return l.generateResp(nil, 400, "user id not match"), errors.New("user id not match")
	}
	l.Logger.Infof("get post success! post id: %d", post.Id)
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
