package model

import "gorm.io/gorm"

/**
 * @Time   : 2022/1/24 23:45
 * @Author : WindsLeaver
 * @File   : sys_tenant
 * @Version: 1.0.0
 * @Description:
 **/

// SysTenant 租户表
type SysTenant struct {
	ID             string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`          //ID
	ParentTenantId string         `json:"parentTenantId" gorm:"column:parentTenantId;size:32"`      //父级租户ID
	TenantName     string         `json:"tenantName" gorm:"size:32;not null"`                       //租户名（用户名）
	Name           string         `json:"name" gorm:"size:32;not null"`                             //租户名称
	Description    string         `json:"description" gorm:"size:512"`                              //描述
	IsActive       bool           `json:"isActive" gorm:"not null"`                                 //是否可用 true:可用 false:不可用
	Remark         string         `json:"remark" gorm:"size:512;"`                                  //备注
	CreatedAt      int64          `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` //创建时间
	UpdatedAt      int64          `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"` //更新时间
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
