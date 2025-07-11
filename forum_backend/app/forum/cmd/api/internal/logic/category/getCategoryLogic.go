package category

import (
	"context"
	"errors"
	"fmt"

	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"forum_backend/app/forum/model/category"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.GetCategoryReq) (resp *types.GetCategoryResp, err error) {
	if req.Id <= 0 {
		errorstr := "id is invalid, id must be greater than 0"
		l.Logger.Infof(errorstr)
		return l.generateResp(nil, 400, errorstr), errors.New(errorstr)
	}
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
	if err != nil {
		errstr := fmt.Sprintf("get category failed: %v", err)
		l.Logger.Errorf(errstr)
		return l.generateResp(nil, 400, errstr), err
	}
	l.Logger.Infof("get category success!")
	resp = l.generateResp(category, 200, "success")
	return
}

func (l *GetCategoryLogic) generateResp(category *category.Category, code int64, message string) *types.GetCategoryResp {
	return &types.GetCategoryResp{
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
