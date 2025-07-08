package category

import (
	"context"
	"errors"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/category"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoriesLogic) ListCategories(req *types.ListCategoryReq) (resp *types.ListCategoryResp, err error) {
	if err := l.checkReq(req); err != nil {
		l.Logger.Infof("check req error: %v", err)
		return l.generateResp(nil, 0, req.Page, req.PageSize, 400, "check req error"), err
	}
	categoryList, total, err := l.svcCtx.CategoryModel.FindCategoryList(l.ctx, req.Page, req.PageSize)
	if err != nil {
		l.Logger.Errorf("list categories error: %v", err)
		return l.generateResp(nil, 0, req.Page, req.PageSize, 400, "list categories error"), err
	}
	l.Logger.Infof("list categories success!")
	resp = l.generateResp(categoryList, total, req.Page, req.PageSize, 200, "success")
	return
}

func (l *ListCategoriesLogic) checkReq(req *types.ListCategoryReq) (err error) {
	if req.Page <= 0 {
		return errors.New("page is invalid")
	}
	if req.PageSize <= 0 {
		return errors.New("page size is invalid")
	}
	return nil
}

func (l *ListCategoriesLogic) generateResp(categoryList []*category.Category, total int64, page int64, pageSize int64, code int64, message string) *types.ListCategoryResp {
	return &types.ListCategoryResp{
		BaseResp: types.BaseResp{
			Code:    code,
			Message: message,
		},
		CategoryList: l.generateCategoryList(categoryList),
		Total:        total,
		TotalPages:   int64(total / pageSize),
		CurrentPage:  page,
		HasNextPage:  total > page*pageSize,
		HasPrevPage:  page > 1 && page*pageSize < total,
	}
}

func (l *ListCategoriesLogic) generateCategoryList(categoryList []*category.Category) []types.CategoryInfo {
	var categoryInfos []types.CategoryInfo
	for _, category := range categoryList {
		categoryInfos = append(categoryInfos, types.CategoryInfo{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description.String,
			SortOrder:   category.SortOrder,
			IsActive:    category.IsActive,
			CreatedTime: category.CreatedTime.Unix(),
			UpdatedTime: category.UpdatedTime.Unix(),
		})
	}
	return categoryInfos
}
