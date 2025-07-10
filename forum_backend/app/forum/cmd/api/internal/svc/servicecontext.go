package svc

import (
	"forum_backend/app/forum/cmd/api/internal/config"
	"forum_backend/app/forum/model/admin"
	"forum_backend/app/forum/model/category"
	"forum_backend/app/forum/model/comment"
	"forum_backend/app/forum/model/post"
	"forum_backend/app/forum/model/user"
	"forum_backend/app/forum/model/user_like"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	PostModel     post.PostModel
	CommentModel  comment.CommentModel
	CategoryModel category.CategoryModel
	AdminModel    admin.AdminModel
	UserModel     user.UserModel
	UserLikeModel user_like.UserLikeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:        c,
		PostModel:     post.NewPostModel(conn, c.Cache),
		CommentModel:  comment.NewCommentModel(conn, c.Cache),
		CategoryModel: category.NewCategoryModel(conn, c.Cache),
		AdminModel:    admin.NewAdminModel(conn, c.Cache),
		UserModel:     user.NewUserModel(conn, c.Cache),
		UserLikeModel: user_like.NewUserLikeModel(conn, c.Cache),
	}
}
