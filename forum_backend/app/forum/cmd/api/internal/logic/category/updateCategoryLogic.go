package category

import (
	"context"
	"errors"
	"fmt"
	"time"

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
		errstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errstr)
		return l.generateResp(nil, 400, errstr), errors.New(errstr)
	}
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("get category failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	if category.Status == "inactive" {
		errstr := "category is inactive, can't update"
		l.Logger.Infof(errstr)
		return l.generateResp(nil, 400, errstr), errors.New(errstr)
	}
	err = l.svcCtx.CategoryModel.Update(l.ctx, l.generateCategoryInfo(req))
	if err != nil {
		errstr := fmt.Sprintf("update category by id %d failed: %v", req.Id, err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
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
			Status:      category.Status,
			CreatedTime: category.CreatedTime.Unix(),
			UpdatedTime: category.UpdatedTime.Unix(),
		},
	}
}

func (l *UpdateCategoryLogic) generateCategoryInfo(req *types.UpdateCategoryReq) *category.Category {
	return &category.Category{
		Id:          req.Id,
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Status:      req.Status,
		UpdatedTime: time.Now(),
	}
}
