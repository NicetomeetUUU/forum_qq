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
		errstr := fmt.Sprintf("check list post req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, false, 0, 400, errstr), err
	}
	postList, err := l.svcCtx.PostModel.FindPostList(l.ctx, req.PageSize, req.LastIndex, req.OrderBy, req.OrderType)
	if err != nil {
		errstr := fmt.Sprintf("find more post failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, false, 0, 400, errstr), err
	}
	hasMore := false
	if len(postList) == int(req.PageSize) {
		hasMore = true
	}
	lastIndex := int64(0)
	if len(postList) > 0 {
		lastIndex = req.LastIndex + int64(len(postList))
	}
	resp = l.generateResp(postList, hasMore, lastIndex, 200, "list posts success!")
	return
}

func (l *ListPostsLogic) checkListPostReq(req *types.ListPostReq) error {
	if req.PageSize <= 0 {
		return errors.New("pageSize is required, pageSize must be greater than 0")
	}
	if req.LastIndex < 0 {
		return errors.New("lastIndex is required, lastIndex must be greater than 0")
	}
	if req.CategoryId <= 0 {
		return errors.New("categoryId is required, categoryId must be greater than 0")
	}
	if req.OrderBy == "" || (req.OrderBy != "created_time" && req.OrderBy != "updated_time" && req.OrderBy != "view_count" && req.OrderBy != "like_count" && req.OrderBy != "comment_count") {
		return errors.New("orderBy is required, orderBy must be 'created_time' or 'updated_time' or 'view_count' or 'like_count' or 'comment_count'")
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
			Id:           post.Id,
			Title:        post.Title,
			Content:      post.Content,
			UserId:       post.UserId,
			ViewCount:    post.ViewCount,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			Status:       post.Status,
			CreatedTime:  post.CreatedTime.Unix(),
			UpdatedTime:  post.UpdatedTime.Unix(),
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
