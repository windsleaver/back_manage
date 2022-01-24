package data

import (
	"fmt"
	"go/data/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitData 初始化数据库
func InitData() {
	tableList := []interface{}{model.SysUser{}, model.SysTenant{}}
	dsn := "host=localhost user=postgres password=b8662279 dbname=back_manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	for _, item := range tableList {
		db.AutoMigrate(item)
	}
	fmt.Println(err)
}
