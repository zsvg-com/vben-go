package pub

import (
	"net/http"
	"vben/app/core/utils/R"
	"vben/common/r"
	"vben/modules/sys/menu"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func GetActor(c *gin.Context) {
	//typ := c.Param("type")
	//pid := c.Param("pid")
	var sql = "select id,name from sys_actor"
	var list []IdNameAvatarVo
	err := mysql.MysqlDb().Raw(sql).Scan(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(list))
}

func GetRece(c *gin.Context) {
	//var list []IdNameOrgVo
	c.JSON(http.StatusOK, r.Ok([0]IdNameOrgVo{}))
}

func GetGroupTree(c *gin.Context) {
	//var list []IdNameOrgVo
	c.JSON(http.StatusOK, r.Ok([0]r.Stree{}))
}

func PostRece(c *gin.Context) {
	c.JSON(http.StatusOK, r.Ok(0))
}

func GetUser(c *gin.Context) {
	var vo UserInfoVo
	var user User
	user.Useid = "u1"
	user.Usena = "admin"
	user.Orgid = "o1000"
	user.Orgna = "XX科技"
	user.Nicna = "管理员"
	user.Monum = "13812345678"
	user.Avatar = "https://unpkg.com/@vbenjs/static-source@0.1.7/source/avatar-v1.webp"
	vo.User = user
	vo.Roles = append(vo.Roles, "superadmin")
	vo.Perms = append(vo.Perms, "*:*:*")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func GetMenus(c *gin.Context) {
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
