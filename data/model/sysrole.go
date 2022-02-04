package model

import "gorm.io/gorm"

/**
 * @Time   : 2022/1/25 00:30
 * @Author : WindsLeaver
 * @File   : sysrole
 * @Version: 1.0.0
 * @Description:
 **/

//SysRole 角色表
type SysRole struct {
	ID        string         `json:"id" gorm:"column:id;primaryKey;type:varchar(32)"`          //ID
	RoleName  string         `json:"role_name" gorm:"column:role_name;type:varchar(32)"`       //角色名称
	Remark    string         `json:"remark" gorm:"size:512"`                                   //备注
	CreatedAt int64          `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"` //创建时间
	UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
