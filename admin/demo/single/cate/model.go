package singlec

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type DemoSingleCate struct {
	Id    int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Pid   int64      `json:"pid" gorm:"type:bigint"`           //分类ID
	Ornum int        `json:"ornum" gorm:"type:int"`            //排序号
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	Cruid string     `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	Cruna string     `json:"cruna" gorm:"-"`                   //创建人名称
	Upuid string     `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	Upuna string     `json:"upuna" gorm:"-"`                   //更新人名称
	Avtag boolx.Bool `json:"avtag"`
	Notes string     `json:"notes" gorm:"type:varchar(128)"` //备注
}

func (DemoSingleCate) TableName() string {
	return "demo_single_cate"
}
