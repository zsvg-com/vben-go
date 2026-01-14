package auth

import (
	"time"
	"vben/common/json/boolx"
)

type LoginVo struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ExpireIn        int64  `json:"expire_in"`
	RefreshExpireIn int64  `json:"refresh_expire_in"`
	ClientId        string `json:"client_id"`
	Scope           string `json:"scope"`
	Openid          string `json:"openid"`
}

type CaptchaVo struct {
	CaptchaEnabled bool   `json:"captchaEnabled"`
	Uuid           string `json:"uuid"`
	Img            string `json:"img"`
}

type UserInfoVo struct {
	User        UserVo   `json:"user"`        //用户基本信息
	Permissions []string `json:"permissions"` //菜单权限
	Roles       []string `json:"roles"`       //角色权限
}

type UserVo struct {
	UserId      string    `json:"userId"`      //用户ID
	TenantId    string    `json:"tenantId"`    //租户ID
	DeptId      string    `json:"deptId"`      //部门ID
	UserName    string    `json:"userName"`    //用户账号
	NickName    string    `json:"nickName"`    //用户昵称
	UserType    string    `json:"userType"`    //用户类型（sys_user系统用户）
	Email       string    `json:"email"`       //用户邮箱
	Phonenumber string    `json:"phonenumber"` //手机号码
	Gender      string    `json:"gender"`      //用户性别（0男 1女 2未知）
	Avatar      string    `json:"avatar"`      //头像地址
	Password    string    `json:"password"`    //密码
	Status      string    `json:"status"`      //帐号状态（0正常 1停用）
	LoginIp     string    `json:"loginIp"`     //最后登录IP
	LoginDate   time.Time `json:"loginDate"`   //最后登录时间
	Remark      string    `json:"remark"`      //备注
	CreateTime  time.Time `json:"createTime"`  //创建时间
	DeptName    string    `json:"deptName"`    //部门名称
	RoleIds     []string  `json:"roleIds"`     //角色组
	PostIds     []string  `json:"postIds"`     //岗位组
	RoleId      int64     `json:"roleId"`      //数据权限 当前角色ID
}

type RouterVo struct {
	Name       string     `json:"name"`       //路由名字
	Path       string     `json:"path"`       //路由地址
	Hidden     boolx.Bool `json:"hidden"`     //是否隐藏路由，当设置 true 的时候该路由不会再侧边栏出现
	Redirect   string     `json:"redirect"`   //重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
	Component  string     `json:"component"`  //组件地址
	Query      string     `json:"query"`      //路由参数：如 {"id": 1, "name": "ry"}
	AlwaysShow boolx.Bool `json:"alwaysShow"` //当你一个路由下面的 children 声明的路由大于1个时，自动会变成嵌套的模式--如组件页面
	Meta       MetaVo     `json:"meta"`       //其他元素
	Children   []RouterVo `json:"children"`   //子路由
}

type MetaVo struct {
	Title   string     `json:"title"`   //设置该路由在侧边栏和面包屑中展示的名字
	Icon    string     `json:"icon"`    //设置该路由的图标，对应路径src/assets/icons/svg
	NoCache boolx.Bool `json:"noCache"` //设置为true，则不会被 <keep-alive>缓存
	Link    string     `json:"link"`    //内链地址（http(s)://开头）
}
