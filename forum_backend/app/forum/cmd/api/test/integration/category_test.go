package integration

import (
	"context"
	"forum_backend/app/forum/cmd/api/internal/config"
	"forum_backend/app/forum/cmd/api/internal/logic/category"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"forum_backend/app/forum/cmd/api/internal/types"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	testCategory := &types.CreateCategoryReq{
		Name:        "test",
		Description: "test",
		SortOrder:   1,
		IsActive:    1,
	}
	logic := category.NewCreateCategoryLogic(context.Background(), svc.NewServiceContext(config.Config{}))
	resp, err := logic.CreateCategory(testCategory)
	if err != nil {
		t.Errorf("create category error: %v", err)
	}
	t.Logf("create category success: %v", resp)
}

func TestUpdateCategory(t *testing.T) {
	testCategory := &types.UpdateCategoryReq{
		Id:          1,
		Name:        "test",
		Description: "test",
		SortOrder:   1,
		IsActive:    1,
	}
	logic := category.NewUpdateCategoryLogic(context.Background(), svc.NewServiceContext(config.Config{}))
	resp, err := logic.UpdateCategory(testCategory)
	if err != nil {
		t.Errorf("update category error: %v", err)
	}
	t.Logf("update category success: %v", resp)
}

func TestDeleteCategory(t *testing.T) {

}

func TestGetCategory(t *testing.T) {

}

func TestListCategories(t *testing.T) {

}
