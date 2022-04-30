package svc

import (
	"back_manage/service/sys/cmd/api/internal/config"
	"back_manage/service/sys/cmd/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	DbEngin          *gorm.DB
	WlAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启用gorm支持
	db, err := gorm.Open(postgres.Open(c.Postgres.Dsn), &gorm.Config{
		//NamingStrategy: schema.NamingStrategy{}
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		DbEngin:          db,
		WlAuthMiddleware: middleware.NewWlAuthMiddleware().Handle,
	}
}
