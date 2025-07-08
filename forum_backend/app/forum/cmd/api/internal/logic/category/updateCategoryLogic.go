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

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) (resp *types.UpdateCategoryResp, err error) {
	if req.Id <= 0 {
		l.Logger.Infof("id is invalid")
		return l.generateResp(nil, 400, "id is invalid"), errors.New("id is invalid")
	}
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("get category error: %v", err)
		return l.generateResp(nil, 400, "get category error"), err
	}
	category.Name = req.Name
	category.Description = sql.NullString{String: req.Description, Valid: req.Description != ""}
	category.SortOrder = req.SortOrder
	category.IsActive = req.IsActive
	err = l.svcCtx.CategoryModel.Update(l.ctx, category)
	if err != nil {
		l.Logger.Errorf("update category error: %v", err)
		return l.generateResp(nil, 400, "update category error"), err
	}
	l.Logger.Infof("update category success!")
	resp = l.generateResp(category, 200, "success")
	return
}

func (l *UpdateCategoryLogic) generateResp(category *category.Category, code int64, message string) *types.UpdateCategoryResp {
	return &types.UpdateCategoryResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		CategoryInfo: types.CategoryInfo{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description.String,
			SortOrder:   category.SortOrder,
			IsActive:    category.IsActive,
			CreatedTime: category.CreatedTime.Unix(),
			UpdatedTime: category.UpdatedTime.Unix(),
		},
	}
}
