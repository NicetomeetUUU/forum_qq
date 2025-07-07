package svc

import (
	"forum_backend/api/user/userService/internal/config"
	"forum_backend/model/user"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
