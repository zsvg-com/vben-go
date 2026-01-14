package dict

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type ToolDictMain struct {
	Id    int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Code  string     `json:"code" gorm:"type:varchar(32)"`     //字典代码
	Catid int64      `json:"catid" gorm:"type:bigint"`         //字典分类ID
	Notes string     `json:"notes" gorm:"type:varchar(128)"`   //备注
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	//Cruid string     `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	//Cruna string     `json:"cruna" gorm:"-"`                   //创建人姓名
	//Upuid string     `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	//Upuna string     `json:"upuna" gorm:"-"`                   //更新人姓名
	Avtag boolx.Bool `json:"avtag"`                 //可用标记
	Ornum int        `json:"ornum" gorm:"type:int"` //排序号
}

func (ToolDictMain) TableName() string {
	return "tool_dict_main"
}
