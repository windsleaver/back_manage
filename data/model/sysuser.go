package model

import "gorm.io/gorm"

/**
 * @Time   : 2022/1/20 00:07
 * @Author : WindsLeaver
 * @File   : sys_user
 * @Version: 1.0.0
 * @Description:
 **/

// SysUser 用户表
type SysUser struct {
	ID             string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             //ID
	UserName       string         `json:"user_name" gorm:"column:user_name;type:varchar(32);not null"` //用户名
	Password       string         `json:"password" gorm:"column:password;type:varchar(256);not null"`  //密码
	Salt           string         `json:"salt" gorm:"column:salt;type:varchar(256);not null"`          //盐值
	Name           string         `json:"name" gorm:"column:name;type:varchar(32);not null"`           //姓名
	HeadImg        string         `json:"head_img" gorm:"column:head_img;type:varchar(512);not null"`  //头像
	Phone          string         `json:"phone" gorm:"column:phone;type:varchar(32);not null"`         //联系方式
	Email          string         `json:"email" gorm:"column:email;type:varchar(32);not null"`         //邮箱
	IsBindWechat   bool           `json:"isBindWechat" gorm:"column:is_bind_wechat;not null"`          //是否绑定微信[true:是;false:否]
	UnionId        string         `json:"unionId" gorm:"column:union_id;type:varchar(32)"`             //微信UnionId
	OpenId         string         `json:"openId" gorm:"column:open_id;type:varchar(32)"`               //微信OpenId
	IsActive       bool           `json:"isActive" gorm:"column:is_active;not null"`                   //是否启用[true:是;false:否]
	IsAdmin        bool           `json:"isAdmin" gorm:"column:is_admin;not null"`                     //是否是管理员[true:是;false:否]
	LastLoginTime  int64          `json:"lastLoginTime" gorm:"column:last_login_time"`                 //最后一次登录时间
	LoginNumberErr int            `json:"loginNumberErr" gorm:"column:login_number_err;not null"`      //登录失败次数
	IsLock         bool           `json:"isLock" gorm:"column:is_lock;not null"`                       //是否锁定[true:是;false:否]
	CreatedAt      int64          `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`    //创建时间
	UpdatedAt      int64          `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`    //更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
