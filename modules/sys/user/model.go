package user

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type SysUser struct {
	Id       string     `json:"id" gorm:"primaryKey"` //主键ID
	Name     string     `json:"name"`                 //名称
	Username string     `json:"username"`             //用户名（账号）
	Password string     `json:"password"`             //用户名（账号）
	Email    string     `json:"email"`                //邮箱
	Monum    string     `json:"monum"`                //手机号
	Gender   string     `json:"gender"`               //性别
	Depid    string     `json:"depid"`                //部门ID
	Depna    string     `json:"depna" gorm:"-"`       //部门名称
	Tier     string     `json:"tier"`                 //层级,以“_”隔开
	Job      string     `json:"job"`                  //职务
	Avatar   string     `json:"avatar"`               //头像URL
	Loip     string     `json:"loip"`                 //登录IP
	Lotim    timex.Time `json:"lotim"`                //登录时间
	Catag    boolx.Bool `json:"catag"`                //缓存标记
	Type     int64      `json:"type"`                 //用户类型
	Label    string     `json:"label"`                //标签
	Notes    string     `json:"notes"`                //备注
	Ornum    int        `json:"ornum"`                //排序号
	Avtag    boolx.Bool `json:"avtag"`                //可用标记
	Crtim    timex.Time `json:"crtim"`                //创建时间
	Uptim    timex.Time `json:"uptim"`                //更新时间
	Cruid    string     `json:"cruid"`                //创建人ID
	Upuid    string     `json:"upuid"`                //更新人ID
}

func (SysUser) TableName() string {
	return "sys_user"
}
