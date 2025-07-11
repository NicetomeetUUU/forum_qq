package category

import (
	"context"
	"errors"
	"fmt"
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
		errorstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errorstr)
		return l.generateResp("none", 400, errorstr), errors.New(errorstr)
	}
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("find category failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp("none", 400, errstr), err
	}
	if category.Status == "active" {
		l.Logger.Infof("category is active, the soft delete will be executed.")
		err = l.svcCtx.CategoryModel.SoftDelete(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("soft delete category failed: %v", err)
			l.Logger.Errorf(errstr)
			return l.generateResp("soft", 400, errstr), err
		}
		l.Logger.Infof("soft delete category success, now the category is inactive!")
		resp = l.generateResp("soft", 200, "success")
		return
	} else {
		l.Logger.Infof("category is inactive, the hard delete will be executed.")
		var postCount int64
		postCount, err = l.svcCtx.PostModel.CountPostsByCategoryId(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("count posts by category id failed: %v", err)
			l.Logger.Errorf(errstr)
			return l.generateResp("hard", 400, errstr), err
		}
		if postCount > 0 {
			errstr := fmt.Sprintf("category: %d has %d posts, can't delete, please delete posts first",
				req.Id, postCount)
			l.Logger.Errorf(errstr)
			return l.generateResp("hard", 400, errstr), err
		}
		err = l.svcCtx.CategoryModel.HardDelete(l.ctx, req.Id)
		if err != nil {
			errstr := fmt.Sprintf("hard delete category failed: %v", err)
			l.Logger.Errorf(errstr)
			return l.generateResp("hard", 400, errstr), err
		}
		l.Logger.Infof("hard delete category success, now the category is deleted!")
		resp = l.generateResp("hard", 200, "success")
		return
	}
}

func (l *DeleteCategoryLogic) generateResp(deleteType string, code int64, message string) *types.DeleteCategoryResp {
	return &types.DeleteCategoryResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		DeleteType: deleteType,
	}
}
