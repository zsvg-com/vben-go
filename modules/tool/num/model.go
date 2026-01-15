package num

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type ToolNum struct {
	Id    string     `json:"id" gorm:"type:varchar(32);primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`         //名称
	Label string     `json:"label" gorm:"type:varchar(32)"`         //标签
	Numod string     `json:"numod" gorm:"type:varchar(32)"`         //生成模式
	Nupre string     `json:"nupre" gorm:"type:varchar(32)"`         //编号前缀
	Nunex string     `json:"nunex" gorm:"type:varchar(32)"`         //下一个编号
	Nulen int        `json:"nulen" gorm:"type:int"`                 //编号长度
	Cudat int        `json:"cudat" gorm:"type:varchar(32)"`         //当前日期
	Notes string     `json:"notes" gorm:"type:varchar(128)"`        //备注
	Ornum int        `json:"ornum" gorm:"type:int"`                 //排序号
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`           //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`            //更新时间
	Nflag boolx.Bool `json:"nflag"`                                 //是否被修改过或新添加的
	Avtag boolx.Bool `json:"avtag"`
}

func (ToolNum) TableName() string {
	return "tool_num"
}
