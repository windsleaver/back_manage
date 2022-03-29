package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Postgres struct {
		Dsn string //链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}
}
