package org

type SysOrg struct {
	Id   string `json:"id" gorm:"primaryKey"` //主键ID
	Name string `json:"name"`                 //名称
	Type int    `json:"type"`                 //组织元素类型
}

func (SysOrg) TableName() string {
	return "sys_org"
}
