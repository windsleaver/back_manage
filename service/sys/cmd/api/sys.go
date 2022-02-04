package main

import (
	"back_manage/service/sys/cmd/api/internal/config"
	"back_manage/service/sys/cmd/api/internal/handler"
	"back_manage/service/sys/cmd/api/internal/svc"
	"back_manage/utils/middleware"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/sys-api.yaml", "the config file")

func main() {
	//初始化数据库
	//data.InitData()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//全局中间件
	server.Use(middleware.NewWlAuthMiddleware().Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
