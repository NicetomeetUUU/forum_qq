package comment

import (
	"context"
	"errors"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// dev ops mysql redis, dockerfile docker-compose, k8s concepts
func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	if req.Id <= 0 {
		resp = l.generateResp(400, "id is required, id must be greater than 0")
		l.Logger.Infof("id is required, id must be greater than 0")
		err = errors.New("id is required, id must be greater than 0")
		return resp, err
	}
	return
}

func (l *DeleteCommentLogic) generateResp(code int64, message string) *types.DeleteCommentResp {
	return &types.DeleteCommentResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
	}
}
