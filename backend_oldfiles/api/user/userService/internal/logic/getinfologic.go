package logic

import (
	"context"
	"errors"

	"forum_backend/api/user/userService/internal/svc"
	"forum_backend/api/user/userService/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo(req *types.GetInfoReq) (resp *types.GetInfoResp, err error) {
	// validate
	if req.UserId <= 0 {
		errResp := &types.GetInfoResp{
			BaseResp: types.BaseResp{
				Code:    400,
				Message: "please check your user id!",
			},
		}
		return errResp, errors.New("invalid user id")
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userInfo := types.UserInfo{
		Id:             user.Id,
		Email:          user.Email,
		Username:       user.Username,
		Avatar:         user.Avatar,
		Signature:      user.Signature,
		BirthdayStr:    user.Birthday.Time.Format("2006-01-02"),
		Role:           user.Role,
		Status:         user.Status,
		IsDeleted:      user.IsDeleted,
		CreatedTimeStr: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedTimeStr: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	resp = &types.GetInfoResp{
		BaseResp: types.BaseResp{
			Code:    200,
			Message: "success",
		},
		UserInfo: userInfo,
	}
	return
}
