package post

import (
	"context"
	"errors"
	"fmt"
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
	var status int64
	var deleteType string = "none"
	status, err = l.checkDeletePostReq(req)
	if err != nil {
		errstr := fmt.Sprintf("check delete post req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(deleteType, 400, errstr), err
	}
	if status != 2 {
		deleteType = "soft"
		err = l.svcCtx.PostModel.SoftDelete(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("soft delete post by id %d failed: %v", req.Id, err)
			l.Logger.Errorf(errstr)
			return l.generateResp(deleteType, 400, errstr), err
		}
	} else {
		deleteType = "hard"
		commentCount, err := l.svcCtx.CommentModel.CountCommentsByPostId(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("count comments by post id %d failed: %v", req.Id, err)
			l.Logger.Errorf(errstr)
			return l.generateResp(deleteType, 400, errstr), err
		}
		if commentCount > 0 {
			errstr := "post has comments, can't hard delete, please delete comments first"
			l.Logger.Infof(errstr)
			return l.generateResp(deleteType, 400, errstr), errors.New(errstr)
		}
		err = l.svcCtx.PostModel.HardDelete(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("hard delete post by id %d failed: %v", req.Id, err)
			l.Logger.Errorf(errstr)
			return l.generateResp(deleteType, 400, errstr), err
		}
	}
	l.Logger.Infof("delete post success! post id: %d", req.Id)
	resp = l.generateResp(deleteType, 200, "delete post success!")
	return
}

func (l *DeletePostLogic) checkDeletePostReq(req *types.DeletePostReq) (int64, error) {
	if req.Id <= 0 {
		return 0, errors.New("id is invalid, id must be greater than 0")
	}
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find post info by id %d failed: %v", req.Id, err)
		return 0, errors.New(errstr)
	}
	// userId := l.ctx.Value("userId").(int64)
	userId := req.UserId
	if postInfo.UserId != userId {
		errstr := "user id not match"
		return 0, errors.New(errstr)
	}
	return postInfo.Status, nil
}

func (l *DeletePostLogic) generateResp(deleteType string, code int64, message string) *types.DeletePostResp {
	return &types.DeletePostResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		DeleteType: deleteType,
	}
}
