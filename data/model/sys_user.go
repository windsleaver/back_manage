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
	ID             string         `gorm:"column:id;primaryKey;type:varchar(32)"`      //ID
	TenantId       string         `gorm:"column:tenant_id;type:varchar(32);not null"` //租户ID
	UserName       string         `gorm:"column:user_name;type:varchar(32);not null"` //用户名
	Password       string         `gorm:"column:password;type:varchar(256);not null"` //密码
	Salt           string         `gorm:"column:salt;type:varchar(256);not null"`     //盐值
	Name           string         `gorm:"column:name;type:varchar(32);not null"`      //姓名
	HeadImg        string         `gorm:"column:head_img;type:varchar(512);not null"` //头像
	Phone          string         `gorm:"column:phone;type:varchar(32);not null"`     //联系方式
	Email          string         `gorm:"column:email;type:varchar(32);not null"`     //邮箱
	IsBindWechat   bool           `gorm:"column:is_bind_wechat;not null"`             //是否绑定微信[true:是;false:否]
	UnionId        string         `gorm:"column:union_id;type:varchar(32)"`           //微信UnionId
	OpenId         string         `gorm:"column:open_id;type:varchar(32)"`            //微信OpenId
	IsActive       bool           `gorm:"column:is_active;not null"`                  //是否启用[true:是;false:否]
	IsAdmin        bool           `gorm:"column:is_admin;not null"`                   //是否是管理员[true:是;false:否]
	LastLoginTime  int64          `gorm:"column:last_login_time"`                     //最后一次登录时间
	LoginNumberErr int            `gorm:"column:login_number_err;not null"`           //登录失败次数
	IsLock         bool           `gorm:"column:is_lock;not null"`                    //是否锁定[true:是;false:否]
	CreatedAt      int64          `gorm:"column:created_at;autoCreateTime:milli"`     //创建时间
	UpdatedAt      int64          `gorm:"column:updated_at;autoUpdateTime:milli"`     //更新时间
	CreatedOn      int64          `json:"created_on"`                                 //创建时间
	ModifiedOn     int64          `json:"modified_on"`                                //更新时间
	DeletedOn      gorm.DeletedAt `json:"deleted_on"`                                 //创建时间
}
