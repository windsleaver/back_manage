package data

import (
	"back_manage/data/model"
	"back_manage/utils/random"
	"back_manage/utils/security"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
 * @Time   : 2022/3/29 00:29
 * @Author : WindsLeaver
 * @File   : init_user_data
 * @Version: 1.0.0
 * @Description:
 **/

// InitUserData 初始化数据库用户信息
func InitUserData() {
	addUserConfig()
}

func addUserConfig() error {
	dsn := "host=localhost user=postgres password=password dbname=back_manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	salt, _ := random.GenerateUuid()
	password, _ := security.GenderPassword(model.DefaultPassword, salt)
	user := model.SysUser{
		ID:             "1",
		UserName:       "WindsLeaver",
		Password:       password,
		Salt:           salt,
		Name:           "WindsLeaver",
		HeadImg:        "",
		Phone:          "",
		Email:          "",
		IsBindWechat:   false,
		UnionId:        "",
		OpenId:         "",
		IsActive:       true,
		IsAdmin:        true,
		LastLoginTime:  0,
		LoginNumberErr: 0,
		IsLock:         false,
	}
	tenant := model.SysTenant{
		ID:             "1",
		ParentTenantId: "",
		TenantName:     "超级管理员",
		Name:           "超级管理员",
		Description:    "超级管理员",
		IsActive:       true,
		Remark:         "",
	}
	tenantAndUser := model.SysTenantAndUser{
		ID:       "1",
		TenantId: "1",
		UserId:   "1",
		Remark:   "超级管理员",
	}
	db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		tx.Create(&tenant)
		tx.Create(&tenantAndUser)
		return nil
	})
	return nil
}
