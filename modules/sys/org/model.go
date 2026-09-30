package org

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysOrg struct {
	Id       string     `json:"id"  gorm:"primaryKey"` //主键ID  omitempty
	Name     string     `json:"name"`                  //名称
	Pid      string     `json:"pid"`                   //父ID
	Type     int64      `json:"type"`                  //组织类别 部门在系统参与者sys_actor表中类别为1，这个字段是进行组织细分的，比如分为企业、机构、部门等
	Tier     string     `json:"tier"`                  //层级,以“_”隔开
	Label    string     `json:"label"`                 //标签
	Notes    string     `json:"notes"`                 //备注
	Ornum    int        `json:"ornum"`                 //排序号
	Avtag    boolx.Bool `json:"avtag"`                 //可用标记
	Crtim    timex.Time `json:"crtim"`                 //创建时间
	Uptim    timex.Time `json:"uptim"`                 //更新时间
	Cruid    string     `json:"cruid"`                 //创建人ID
	Upuid    string     `json:"upuid"`                 //更新人ID
	Children []SysOrg   `json:"children" gorm:"-"`     //子部门
}

func (SysOrg) TableName() string {
	return "sys_org"
}
