package oplog

import (
	"vben/common/json/boolx"
	"vben/common/json/timex"
)

type MonOperLog struct {
	Id int64 `json:"id" gorm:"type:bigint;primaryKey"` //主键ID
	//Tenid    string     `json:"tenid" gorm:"type:varchar(16)"`    //租户编号
	Opmod   string     `json:"clkey" gorm:"type:varchar(32)"`   //操作模块
	Butyp   int        `json:"butyp"`                           //业务类型（0其它 1新增 2修改 3删除）
	Remet   string     `json:"remet" gorm:"type:varchar(64)"`   //请求方法
	Reway   string     `json:"reway" gorm:"type:varchar(8)"`    //请求方式
	Optyp   int        `json:"optyp"`                           //操作类别（0其它 1后台用户 2手机端用户）
	Opuna   string     `json:"opuna" gorm:"type:varchar(16)"`   //操作人员
	Opdna   string     `json:"opdna" gorm:"type:varchar(64)"`   //部门名称
	Sutag   boolx.Bool `json:"sutag"`                           //登录状态 1成功 0失败
	Opip    string     `json:"opip" gorm:"type:varchar(16)"`    //登录IP地址
	Oploc   string     `json:"oploc" gorm:"type:varchar(64)"`   //登录地点
	Reurl   string     `json:"reurl" gorm:"type:varchar(64)"`   //请求url
	Repar   string     `json:"repar" gorm:"type:varchar(2000)"` //请求参数
	Bapar   string     `json:"bapar" gorm:"type:varchar(2000)"` //返回参数
	Optim   timex.Time `json:"optim"  gorm:"type:datetime"`     //登录时间
	Browser string     `json:"browser" gorm:"type:varchar(64)"` //浏览器
	Os      string     `json:"os" gorm:"type:varchar(64)"`      //操作系统
	Himsg   string     `json:"himsg" gorm:"type:varchar(255)"`  //提示消息
	Ermsg   string     `json:"ermsg" gorm:"type:varchar(2000)"` //错误消息
	Cotim   int64      `json:"cotim" gorm:"type:bigint"`        //主键ID
}

func (MonOperLog) TableName() string {
	return "mon_oper_log"
}
