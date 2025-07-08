package category

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/category"

	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) (resp *types.CreateCategoryResp, err error) {
	if err := l.checkCategoryInfo(req); err != nil {
		l.Logger.Infof("check category info error: %v", err)
		return l.generateResp(0, 400, "check category info error"), err
	}
	category := &category.Category{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		SortOrder:   req.SortOrder,
		IsActive:    req.IsActive,
	}
	sqlResult, err := l.svcCtx.CategoryModel.Insert(l.ctx, category)
	if err != nil {
		l.Logger.Errorf("insert category error: %v", err)
		return l.generateResp(0, 400, "insert category error"), err
	}
	categoryId, err := sqlResult.LastInsertId()
	if err != nil {
		l.Logger.Errorf("get last insert id error: %v", err)
		return l.generateResp(0, 400, "get last insert id error"), err
	}
	l.Logger.Infof("create category success!")
	resp = l.generateResp(categoryId, 200, "success")
	return
}

func (l *CreateCategoryLogic) checkCategoryInfo(req *types.CreateCategoryReq) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (l *CreateCategoryLogic) generateResp(categoryId int64, code int64, message string) *types.CreateCategoryResp {
	return &types.CreateCategoryResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		CategoryId: categoryId,
	}
}
