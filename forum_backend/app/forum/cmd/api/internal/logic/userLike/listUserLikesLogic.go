package userLike

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/user_like"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLikesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUserLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLikesLogic {
	return &ListUserLikesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUserLikesLogic) ListUserLikes(req *types.ListUserLikesReq) (resp *types.ListUserLikesResp, err error) {
	if err := l.checkListUserLikesreq(req); err != nil {
		errstr := fmt.Sprintf("checkListUserLikesreq error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	userlikes, err := l.svcCtx.UserLikeModel.FindLikeListByUserId(l.ctx, req.UserId)
	if err != nil {
		errstr := fmt.Sprintf("FindLikeListByUserId error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	if len(userlikes) == 0 {
		l.Logger.Infof("the user has no likes yet")
		resp = l.generateResp(userlikes, 0, "no data")
	} else {
		l.Logger.Infof("FindLikeListByUserId success")
		resp = l.generateResp(userlikes, 0, "success")
	}
	return
}

func (l *ListUserLikesLogic) checkListUserLikesreq(req *types.ListUserLikesReq) (err error) {
	if req.UserId <= 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (l *ListUserLikesLogic) generateResp(data []*user_like.UserLike, code int64, msg string) (resp *types.ListUserLikesResp) {
	userLikes := make([]types.UserLikeInfo, 0)
	for _, item := range data {
		userLikes = append(userLikes, types.UserLikeInfo{
			Id:         item.Id,
			UserId:     item.UserId,
			TargetType: item.TargetType,
			TargetId:   item.TargetId,
			CreatedAt:  item.CreatedAt.Unix(),
		})
	}
	resp = &types.ListUserLikesResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: msg,
		},
		UserLikes: userLikes,
	}
	return
}
