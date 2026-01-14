package role

type ApiVo struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Menid int64  `json:"menid"`
}

type MenuVo struct {
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	Icon string  `json:"icon"`
	Type string  `json:"type"`
	Pid  string  `json:"pid"`
	Apis []ApiVo `json:"apis"  gorm:"-"`
}
