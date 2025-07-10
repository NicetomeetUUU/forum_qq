package userLike

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikesCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLikesCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikesCountLogic {
	return &GetUserLikesCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLikesCountLogic) GetUserLikesCount(req *types.GetUserLikesCountReq) (resp *types.GetUserLikesCountResp, err error) {
	if err := l.checkGetUserLikesCountreq(req); err != nil {
		errstr := fmt.Sprintf("checkGetUserLikesCountreq error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	count, err := l.svcCtx.UserLikeModel.CountUserLikeByUserId(l.ctx, req.UserId)
	if err != nil {
		errstr := fmt.Sprintf("CountUserLikeByUserId error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	resp = l.generateResp(count, 0, "success")
	return
}

func (l *GetUserLikesCountLogic) checkGetUserLikesCountreq(req *types.GetUserLikesCountReq) (err error) {
	if req.UserId <= 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (l *GetUserLikesCountLogic) generateResp(data int64, code int64, msg string) (resp *types.GetUserLikesCountResp) {
	resp = &types.GetUserLikesCountResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: msg,
		},
		LikesCount: data,
	}
	return
}
