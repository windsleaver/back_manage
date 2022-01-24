package model

import "gorm.io/gorm"

/**
 * @Time   : 2022/1/25 00:22
 * @Author : WindsLeaver
 * @File   : systenantanduser
 * @Version: 1.0.0
 * @Description:
 **/

// SysTenantAndUser 租户和用户关系表
type SysTenantAndUser struct {
	ID        string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`             //ID
	TenantId  string         `json:"tenant_id" gorm:"column:tenant_id;type:varchar(32);not null"` //租户ID
	UserId    string         `json:"user_id" gorm:"column:user_id;type:varchar(32);not null"`     //用户ID
	Remark    string         `json:"remark" gorm:"size:512;"`                                     //备注
	CreatedAt int64          `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`    //创建时间
	UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`    //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
