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
	var status string
	var deleteType string = "none"
	status, err = l.checkDeletePostReq(req)
	if err != nil {
		errstr := fmt.Sprintf("check delete post req failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(deleteType, 400, errstr), err
	}
	if status != "hidden" {
		deleteType = "soft"
		err = l.svcCtx.PostModel.SoftDelete(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("soft delete post by id %d failed: %v", req.Id, err)
			l.Logger.Errorf(errstr)
			return l.generateResp(deleteType, 400, errstr), err
		}
		err = l.svcCtx.CommentModel.UpdateCommentStatusByPostId(l.ctx, req.Id, "hidden")
		if err != nil {
			errstr := fmt.Sprintf("update comment status by post id %d failed: %v", req.Id, err)
			l.Logger.Errorf(errstr)
			return l.generateResp(deleteType, 400, errstr), err
		}
	} else {
		infostr := fmt.Sprintf("post is already hidden, no need to delete, post id: %d", req.Id)
		l.Logger.Infof(infostr)
		deleteType = "none"
		return l.generateResp(deleteType, 200, infostr), nil
	}
	l.Logger.Infof("delete post success! post id: %d", req.Id)
	resp = l.generateResp(deleteType, 200, "delete post success!")
	return
}

func (l *DeletePostLogic) checkDeletePostReq(req *types.DeletePostReq) (string, error) {
	if req.Id <= 0 {
		return "", errors.New("id is invalid, id must be greater than 0")
	}
	postInfo, err := l.svcCtx.PostModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find post info by id %d failed: %v", req.Id, err)
		return "", errors.New(errstr)
	}
	// userId := l.ctx.Value("userId").(int64)
	userId := req.UserId
	if postInfo.UserId != userId {
		errstr := "user id not match"
		return "", errors.New(errstr)
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
