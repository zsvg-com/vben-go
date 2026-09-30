package pub

type IdNameAvatarVo struct {
	Id     string `json:"id"`     //ID
	Name   string `json:"name"`   //名称
	Avatar string `json:"avatar"` //头像
}

type IdNameOrgVo struct {
	Id   string `json:"id"`   //ID
	Name string `json:"name"` //名称
	Org  string `json:"org"`  //部门
}
