package post

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

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
		l.Logger.Infof("id is invalid")
		return l.generateResp(req.Id, 400, "id is invalid"), errors.New("id is invalid")
	}
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("find post error: %v", err)
		return l.generateResp(req.Id, 400, "find post error"), err
	}
	currentUserId := l.ctx.Value("userId").(int64)
	if postInfo.UserId != currentUserId {
		l.Logger.Errorf("user id not match")
		return l.generateResp(req.Id, 400, "user id not match"), errors.New("user id not match")
	}
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
