package lolog

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type MonLoginLog struct {
	Id int64 `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	//Tenid    string     `json:"tenid" gorm:"type:varchar(16)"`    //租户编号
	Username string     `json:"username" gorm:"type:varchar(36)"` //用户账号
	Clkey    string     `json:"clkey" gorm:"type:varchar(32)"`    //客户端key
	Detyp    string     `json:"detyp" gorm:"type:varchar(8)"`     //设备类型
	Sutag    boolx.Bool `json:"sutag"`                            //登录状态 1成功 0失败
	Loip     string     `json:"loip" gorm:"type:varchar(16)"`     //登录IP地址
	Loloc    string     `json:"loloc" gorm:"type:varchar(64)"`    //登录地点
	Lotim    timex.Time `json:"lotim"  gorm:"type:datetime"`      //登录时间
	Browser  string     `json:"browser" gorm:"type:varchar(64)"`  //浏览器
	Os       string     `json:"os" gorm:"type:varchar(64)"`       //操作系统
	Himsg    string     `json:"himsg" gorm:"type:varchar(255)"`   //提示消息
}

func (MonLoginLog) TableName() string {
	return "mon_login_log"
}
