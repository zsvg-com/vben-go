package groupc

type SysGroupCate struct {
	Id    string `json:"id" gorm:"type:varchar(36);primaryKey"` //主键ID
	Name  string `json:"name" gorm:"type:varchar(255)"`         //名称
	Pid   string `json:"pid" gorm:"type:varchar(36)"`           //分类ID
	Ornum int    `json:"ornum" gorm:"type:int"`                 //排序号
}

func (SysGroupCate) TableName() string {
	return "sys_group_cate"
}
