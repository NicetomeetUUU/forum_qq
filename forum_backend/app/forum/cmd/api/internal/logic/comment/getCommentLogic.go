package comment

import (
	"context"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentReq) (resp *types.GetCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
