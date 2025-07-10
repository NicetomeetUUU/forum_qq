package post

import (
	"context"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestorePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestorePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestorePostLogic {
	return &RestorePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestorePostLogic) RestorePost(req *types.RestorePostReq) (resp *types.RestorePostResp, err error) {
	if err = l.checkRestorePostReq(req); err != nil {
		errstr := fmt.Sprintf("check restore post req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	err = l.svcCtx.PostModel.Restore(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("restore post failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	err = l.svcCtx.CommentModel.UpdateCommentStatusByPostId(l.ctx, req.Id, "published")
	if err != nil {
		errstr := fmt.Sprintf("update comment status by post id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(0, 400, errstr), err
	}
	l.Logger.Infof("restore post success! post id: %d", req.Id)
	resp = l.generateResp(req.Id, 200, "restore post success!")
	return
}

func (l *RestorePostLogic) checkRestorePostReq(req *types.RestorePostReq) error {
	if req.Id <= 0 {
		return fmt.Errorf("id is required, id: %d must be greater than 0", req.Id)
	}
	return nil
}

func (l *RestorePostLogic) generateResp(data int64, code int64, message string) *types.RestorePostResp {
	return &types.RestorePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		PostId: data,
	}
}
