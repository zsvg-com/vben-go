package dictd

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type ToolDictData struct {
	Id int64 `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	//Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Dalab string     `json:"dalab" gorm:"type:varchar(32)"`  //数据标签
	Daval string     `json:"daval" gorm:"type:varchar(32)"`  //数据键值
	Dicid int64      `json:"dicid" gorm:"type:bigint"`       //字典ID
	Shsty string     `json:"shsty" gorm:"type:varchar(32)"`  //显示样式
	Notes string     `json:"notes" gorm:"type:varchar(128)"` //备注
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`    //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`     //更新时间
	//Cruid string     `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	//Cruna string     `json:"cruna" gorm:"-"`                   //创建人姓名
	//Upuid string     `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	//Upuna string     `json:"upuna" gorm:"-"`                   //更新人姓名
	Detag boolx.Bool `json:"detag"`                 //默认标记
	Avtag boolx.Bool `json:"avtag"`                 //可用标记
	Ornum int        `json:"ornum" gorm:"type:int"` //排序号
}

func (ToolDictData) TableName() string {
	return "tool_dict_data"
}
