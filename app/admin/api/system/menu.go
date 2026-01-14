package system

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// GetRoutersHandler /*
func GetRoutersHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	var user system.SysUser
	err1 := mysql.MysqlDb().Where("user_id = ?", userId).First(&user)

	if err1.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "未找到用户",
			"code": http.StatusInternalServerError,
		})
		return
	}

	menu := system.SelectMenuTreeByUserId(utils.GetInterfaceToInt(userId))
	var data = system.BuildMenus(menu)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": data,
	})
}

func ListMenu(context *gin.Context) {

	menuName := context.DefaultQuery("menuName", "")
	status := context.DefaultQuery("status", "")

	var param = tools.SearchTableDataParam{
		Other: system.SysMenu{
			MenuName: menuName,
			Status:   status,
		},
		Params: tools.Params{},
	}
	userId, _ := context.Get("userId")
	if system.IsAdminById(utils.GetInterfaceToInt(userId)) {
		var result = system.SelectSysMenuList(param)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "操作成功",
			"code": http.StatusOK,
			"data": result,
		})
	} else {
		var result = system.SelectSysMenuListByUserId(int(userId.(float64)), param)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "操作成功",
			"code": http.StatusOK,
			"data": result,
		})
	}
}

func GetMenuInfo(context *gin.Context) {
	userId, _ := context.Get("userId")
	println(userId)
	var menuId = context.Param("menuId")
	var date = system.FindMenuInfoById(menuId)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": date,
	})
}

func GetTreeSelect(context *gin.Context) {
	userId, _ := context.Get("userId")
	println(userId)
	data, _ := ioutil.ReadAll(context.Request.Body)
	// 这点很重要，把字节流重新放回 body 中
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	var menuParm system.SysMenu
	if err := context.ShouldBind(&menuParm); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	menu := system.SelectMenuTree(utils.GetInterfaceToInt(userId), menuParm)
	var result = system.BuildMenuTreeSelect(menu)

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func TreeSelectByRole(context *gin.Context) {
	userId, _ := context.Get("userId")
	var roleId = context.Param("roleId")
	var menuPerms system.SysMenu
	menu := system.SelectMenuTree(utils.GetInterfaceToInt(userId), menuPerms)
	var result = system.BuildMenuTreeSelect(menu)
	roles := system.FindRoleInfoById(roleId)
	var checkedKeys = system.SelectMenuListByRoleId(roleId, roles.MenuCheckStrictly)
	context.JSON(http.StatusOK, gin.H{
		"msg":         "操作成功",
		"code":        http.StatusOK,
		"menus":       result,
		"checkedKeys": checkedKeys,
	})
}

func SaveMenu(context *gin.Context) {
	userId, _ := context.Get("userId")
	var menuParm system.SysMenu
	if err := context.ShouldBind(&menuParm); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	menuParm.CreateBy = user.UserName
	result := system.AddMenu(menuParm)
	context.JSON(http.StatusOK, result)
}

func UploadMenu(context *gin.Context) {
	userId, _ := context.Get("userId")
	var menuParm system.SysMenu
	if err := context.ShouldBind(&menuParm); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	menuParm.UpdateBy = user.UserName
	result := system.UpdateMenu(menuParm)
	context.JSON(http.StatusOK, result)
}

func DeleteMenu(context *gin.Context) {
	var menuId = context.Param("menuId")
	result := system.DeleteMenu(menuId)
	context.JSON(http.StatusOK, result)
}
