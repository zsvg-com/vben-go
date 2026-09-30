package actor

type SysActor struct {
	Id   string `json:"id" gorm:"primaryKey"` //主键ID
	Name string `json:"name"`                 //名称
	Type int    `json:"type"`                 //组织元素类型
}

func (SysActor) TableName() string {
	return "sys_actor"
}
