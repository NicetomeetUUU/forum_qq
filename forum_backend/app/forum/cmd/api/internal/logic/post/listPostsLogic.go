package post

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPostsLogic) ListPosts(req *types.ListPostReq) (resp *types.ListPostResp, err error) {
	if err = l.checkListPostReq(req); err != nil {
		l.Logger.Errorf("check list post req error: %v", err)
		return l.generateResp(nil, false, 0, 400, "check list post req error"), err
	}
	return
}

func (l *ListPostsLogic) checkListPostReq(req *types.ListPostReq) error {
	if req.Limit <= 0 {
		return errors.New("limit is required, limit must be greater than 0")
	}
	if req.LastId < 0 {
		return errors.New("lastId is required, lastId must be greater than 0")
	}
	if req.CategoryId <= 0 {
		return errors.New("categoryId is required, categoryId must be greater than 0")
	}
	if req.OrderBy == "" || (req.OrderBy != "created_at" && req.OrderBy != "updated_at" && req.OrderBy != "view_count" && req.OrderBy != "like_count" && req.OrderBy != "comment_count") {
		return errors.New("orderBy is required, orderBy must be 'created_at' or 'updated_at' or 'view_count' or 'like_count' or 'comment_count'")
	}
	if req.OrderType == "" || (req.OrderType != "asc" && req.OrderType != "desc") {
		return errors.New("orderType is required, orderType must be 'asc' or 'desc'")
	}
	return nil
}

func (l *ListPostsLogic) generateResp(postList []*post.Post, hasMore bool, lastId int64, code int64, message string) *types.ListPostResp {
	var postInfoList []types.PostInfo
	for _, post := range postList {
		postInfoList = append(postInfoList, types.PostInfo{
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
		})
	}
	return &types.ListPostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		Posts:   postInfoList,
		HasMore: hasMore,
		LastId:  lastId,
	}
}
