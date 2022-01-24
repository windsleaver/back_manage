package model

import "gorm.io/gorm"

/**
 * @Time   : 2022/1/25 00:44
 * @Author : WindsLeaver
 * @File   : systenantandrole
 * @Version: 1.0.0
 * @Description:
 **/

// SysTenantAndRole 租户和角色关系表
type SysTenantAndRole struct {
	ID        string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             //ID
	TenantId  string         `json:"tenant_id" gorm:"column:tenant_id;type:varchar(32);not null"` //租户ID
	RoleId    string         `json:"role_id" gorm:"column:role_id;type:varchar(32);not null"`     //角色ID
	Remark    string         `json:"remark" gorm:"size:512;"`                                     //备注
	CreatedAt int64          `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`    //创建时间
	UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`    //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
