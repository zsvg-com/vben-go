package menu

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysMenu struct {
	Id       int64      `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	Name     string     `json:"name" gorm:"type:varchar(255)"`    //名称
	Pid      int64      `json:"pid" gorm:"type:bigint"`           //分类ID
	Type     string     `json:"type"  gorm:"type:varchar(32)"`    //类型
	Icon     string     `json:"icon"  gorm:"type:varchar(64)"`    //图标
	Path     string     `json:"path"  gorm:"type:varchar(64)"`    //路由路径
	Param    string     `json:"param"  gorm:"type:varchar(64)"`   //路由参数
	Comp     string     `json:"comp"  gorm:"type:varchar(64)"`    //组件路径
	Notes    string     `json:"notes" gorm:"type:varchar(128)"`   //备注
	Ornum    int        `json:"ornum" gorm:"type:int"`            //排序号
	Crtim    timex.Time `json:"crtim"  gorm:"type:datetime"`      //创建时间
	Uptim    timex.Time `json:"uptim" gorm:"type:datetime"`       //更新时间
	Cruid    string     `json:"cruid" gorm:"type:varchar(36)"`    //创建人ID
	Upuid    string     `json:"upuid" gorm:"type:varchar(36)"`    //更新人ID
	Avtag    boolx.Bool `json:"avtag"`                            //可用标记
	Shtag    boolx.Bool `json:"shtag"`                            //显示标记
	Catag    boolx.Bool `json:"catag"`                            //缓存标记
	Outag    boolx.Bool `json:"outag"`                            //外链标记
	Children []SysMenu  `json:"children" gorm:"-"`                //子菜单
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
