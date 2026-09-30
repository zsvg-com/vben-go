package pub

import (
	"time"
	"vben/common/json/boolx"
)

type UserInfoVo struct {
	User  User     `json:"user"`  //用户基本信息
	Perms []string `json:"perms"` //菜单权限
	Roles []string `json:"roles"` //角色权限
}

type User struct {
	Useid  string    `json:"useid"`  //用户ID
	Tenid  string    `json:"tenid"`  //租户ID
	Orgid  string    `json:"orgid"`  //部门ID
	Usena  string    `json:"usena"`  //用户账号
	Nicna  string    `json:"nicna"`  //用户昵称
	Usety  string    `json:"usety"`  //用户类型（sys_user系统用户）
	Email  string    `json:"email"`  //用户邮箱
	Monum  string    `json:"monum"`  //手机号码
	Gender string    `json:"gender"` //用户性别（0男 1女 2未知）
	Avatar string    `json:"avatar"` //头像地址
	Pwd    string    `json:"pwd"`    //密码
	Avtag  string    `json:"avtag"`  //帐号状态（0正常 1停用）
	Loip   string    `json:"loip"`   //最后登录IP
	Lotim  time.Time `json:"lotim"`  //最后登录时间
	Notes  string    `json:"notes"`  //备注
	Crtim  time.Time `json:"ctim"`   //创建时间
	Orgna  string    `json:"orgna"`  //部门名称
	Roids  []string  `json:"roids"`  //角色组
	Poids  []string  `json:"poids"`  //岗位组
	Rolid  int64     `json:"rolid"`  //数据权限 当前角色ID
}

type RouterVo struct {
	Name       string     `json:"name"`       //路由名字
	Path       string     `json:"path"`       //路由地址
	Shtag      boolx.Bool `json:"shtag"`      //是否隐藏路由，当设置 true 的时候该路由不会再侧边栏出现
	Redirect   string     `json:"redirect"`   //重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
	Comp       string     `json:"comp"`       //组件地址
	Param      string     `json:"param"`      //路由参数：如 {"id": 1, "name": "ry"}
	AlwaysShow boolx.Bool `json:"alwaysShow"` //当你一个路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式--如组件页面
	Meta       MetaVo     `json:"meta"`       //其他元素
	Children   []RouterVo `json:"children"`   //子路由
}

type MetaVo struct {
	Title string     `json:"title"` //设置该路由在侧边栏和面包屑中展示的名字
	Icon  string     `json:"icon"`  //设置该路由的图标，对应路径src/assets/icons/svg
	Catag boolx.Bool `json:"catag"` //设置为true，则不会被 <keep-alive>缓存
	Link  string     `json:"link"`  //内链地址（http(s)://开头）
}
