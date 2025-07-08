package post

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("find post error: %v", err)
		return l.generateResp(400, "find post error"), err
	}
	if postInfo.UserId != l.ctx.Value("userId").(int64) {
		l.Logger.Errorf("user id not match")
		return l.generateResp(400, "user id not match"), errors.New("user id not match")
	}
	err = l.svcCtx.PostModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("delete post error: %v", err)
		return l.generateResp(400, "delete post error"), err
	}
	l.Logger.Infof("delete post success! post id: %d", req.Id)
	resp = l.generateResp(200, "delete post success!")
	return
}

func (l *DeletePostLogic) generateResp(code int64, message string) *types.DeletePostResp {
	return &types.DeletePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
	}
}
