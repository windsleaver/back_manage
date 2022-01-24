package model

/**
 * @Time   : 2022/1/24 23:45
 * @Author : WindsLeaver
 * @File   : sys_tenant
 * @Version: 1.0.0
 * @Description:
 **/

// SysTenant 租户表
type SysTenant struct {
	ID          uint64 `json:"id" gorm:"auto_increment;primary_key"`
	TenancyName string `json:"tenancyName" gorm:"size:32;not null"` //租户名（用户名）
	Name        string `json:"name" gorm:"size:32;not null"`        //租户名称
	Description string `json:"description" gorm:"size:512"`         //描述
	IsActive    bool   `json:"isActive" gorm:"not null"`            //是否可用 true:可用 false:不可用
	Remark      string `json:"remark" gorm:"size:512;"`             //备注
	CreatedOn   int64  `json:"created_on"`                          //创建时间
	ModifiedOn  int64  `json:"modified_on"`                         //更新时间
	DeletedOn   int64  `json:"deleted_on"`                          //创建时间
}
