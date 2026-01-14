package notice

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysNotice struct {
	Id    int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name  string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Cont  string     `json:"cont" gorm:"type:varchar(128)"`    //公告内容
	Notes string     `json:"notes" gorm:"type:varchar(128)"`   //备注
	Ornum int        `json:"ornum" gorm:"type:int"`            //排序号
	Type  int        `json:"type" gorm:"type:int"`             //公告类型（1通知 2公告）
	Crtim timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	Avtag boolx.Bool `json:"avtag"`
}

func (SysNotice) TableName() string {
	return "sys_notice"
}
