package auth

import (
	"net/http"
	"vben/app/core/utils/R"
	"vben/modules/sys/menu"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var bo LoginBo
	if err := c.ShouldBind(&bo); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	var vo LoginVo
	vo.AccessToken = "11233"
	vo.ExpireIn = 604800

	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func Code(c *gin.Context) {
	var vo CaptchaVo
	vo.CaptchaEnabled = false
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": 0,
	})
}

func GetInfo(c *gin.Context) {
	var vo UserInfoVo
	var user UserVo
	user.UserId = "u1"
	user.UserName = "admin"
	user.DeptId = "d1000"
	user.DeptName = "XX科技"
	user.NickName = "管理员"
	user.Phonenumber = "13812345678"
	user.Avatar = "https://unpkg.com/@vbenjs/static-source@0.1.7/source/avatar-v1.webp"
	vo.User = user
	vo.Roles = append(vo.Roles, "superadmin")
	vo.Permissions = append(vo.Permissions, "*:*:*")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func GetRouters(c *gin.Context) {
	//var routerVoList []RouterVo

	var menuList []menu.SysMenu
	var sql = "select m.shtag,m.type,m.name,m.path,m.icon,m.comp,m.ornum,m.id,m.pid,m.outag,m.catag from sys_menu m where m.avtag=1 order by m.pid,m.ornum"
	err := mysql.MysqlDb().Raw(sql).Scan(&menuList).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	menuList = GetChildPerms(menuList, 0)
	menuVoList := buildMenus(menuList)

	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": menuVoList,
	})
}
