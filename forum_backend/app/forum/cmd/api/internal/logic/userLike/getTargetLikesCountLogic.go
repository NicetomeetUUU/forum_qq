package userLike

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTargetLikesCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTargetLikesCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTargetLikesCountLogic {
	return &GetTargetLikesCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTargetLikesCountLogic) GetTargetLikesCount(req *types.GetTargetLikesCountReq) (resp *types.GetTargetLikesCountResp, err error) {
	if err := l.checkGetTargetLikesCountreq(req); err != nil {
		errstr := fmt.Sprintf("checkGetTargetLikesCountreq error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	count, err := l.svcCtx.UserLikeModel.CountUserLikeByTargetTypeTargetId(l.ctx, req.TargetType, req.TargetId)
	if err != nil {
		errstr := fmt.Sprintf("CountUserLikeByTargetTypeTargetId error: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	l.Logger.Infof("GetTargetLikesCount success")
	resp = l.generateResp(count, 0, "success")
	return
}

func (l *GetTargetLikesCountLogic) checkGetTargetLikesCountreq(req *types.GetTargetLikesCountReq) (err error) {
	if req.TargetType == "" {
		return errors.New("target_type is required")
	}
	if req.TargetType != "post" && req.TargetType != "comment" {
		return errors.New("target_type is invalid")
	}
	if req.TargetId <= 0 {
		return errors.New("target_id is required")
	}
	return nil
}

func (l *GetTargetLikesCountLogic) generateResp(data int64, code int64, msg string) (resp *types.GetTargetLikesCountResp) {
	resp = &types.GetTargetLikesCountResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: msg,
		},
		LikesCount: data,
	}
	return
}
