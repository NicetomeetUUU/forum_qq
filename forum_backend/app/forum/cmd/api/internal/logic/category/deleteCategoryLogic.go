package category

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryReq) (resp *types.DeleteCategoryResp, err error) {
	if req.Id <= 0 {
		l.Logger.Infof("id is invalid")
		return l.generateResp(400, "id is invalid"), errors.New("id is invalid")
	}
	err = l.svcCtx.CategoryModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("delete category error: %v", err)
		return l.generateResp(400, "delete category error"), err
	}
	l.Logger.Infof("delete category success!")
	resp = l.generateResp(200, "success")
	return
}

func (l *DeleteCategoryLogic) generateResp(code int64, message string) *types.DeleteCategoryResp {
	return &types.DeleteCategoryResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
	}
}
