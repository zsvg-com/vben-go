package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"
)

func ListRole(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	roleName := context.DefaultQuery("roleName", "")
	roleKey := context.DefaultQuery("roleKey", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysRoles{
			RoleName: roleName,
			RoleKey:  roleKey,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectRoleList(param, true)
	context.JSON(http.StatusOK, result)
}

func ExportRole(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	roleName := context.DefaultQuery("roleName", "")
	roleKey := context.DefaultQuery("roleKey", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysRoles{
			RoleName: roleName,
			RoleKey:  roleKey,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectRoleList(param, false)
	var list = result.Rows.([]system.SysRoles)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "roleId",
		"title":  "角色序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "roleName",
		"title":  "角色名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "roleKey",
		"title":  "角色权限",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "roleSort",
		"title":  "角色排序",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dataScope",
		"title":  "数据范围",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "角色状态",
		"width":  "10",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			var dataScope = v.DataScope
			dataScopeStr := ""
			if "1" == dataScope {
				dataScopeStr = "所有数据权限"
			}
			if "2" == dataScope {
				dataScopeStr = "自定义数据权限"
			}
			if "3" == dataScope {
				dataScopeStr = "本部门数据权限"
			}
			if "4" == dataScope {
				dataScopeStr = "本部门及以下数据权限"
			}
			if "5" == dataScope {
				dataScopeStr = "仅本人数据权限"
			}
			status := v.Status
			statusStr := ""
			if status == "1" {
				statusStr = "正常"
			}
			if status == "0" {
				statusStr = "停用"
			}
			data = append(data, map[string]interface{}{
				"roleId":    v.RoleId,
				"roleName":  v.RoleName,
				"roleKey":   v.RoleKey,
				"roleSort":  v.RoleSort,
				"dataScope": dataScopeStr,
				"status":    statusStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetRoleInfo(context *gin.Context) {
	roleId := context.Param("roleId")
	userId, _ := context.Get("userId")
	if !system.IsAdminById(utils.GetInterfaceToInt(userId)) {
		var isCheck = system.CheckRoleDataScope(roleId)
		if isCheck {
			context.JSON(http.StatusOK, R.ReturnFailMsg("没有权限访问角色数据！"))
			return
		}
	}
	result := system.FindRoleInfoById(roleId)
	context.JSON(http.StatusOK, R.ReturnSuccess(result))
}

func SaveRole(context *gin.Context) {
	userId, _ := context.Get("userId")
	var rolesParam system.SysRolesParam
	if err := context.ShouldBind(&rolesParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	rolesParam.CreateBy = user.UserName
	result := system.SaveRole(rolesParam)
	context.JSON(http.StatusOK, result)
}

/*修改权限*/
func UploadRole(context *gin.Context) {
	userId, _ := context.Get("userId")
	var rolesParam system.SysRolesParam
	if err := context.ShouldBind(&rolesParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	rolesParam.UpdateBy = user.UserName
	result := system.UploadRole(rolesParam, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, result)
}

func PutDataScope(context *gin.Context) {
	userId, _ := context.Get("userId")
	var rolesParam system.SysRolesParam
	if err := context.ShouldBind(&rolesParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	result := system.PutDataScope(rolesParam, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, result)
}

func ChangeRoleStatus(context *gin.Context) {
	userId, _ := context.Get("userId")
	var rolesParam system.SysRoles
	if err := context.ShouldBind(&rolesParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	rolesParam.UpdateBy = user.UserName
	rolesParam.UpdateTime = time.Now()
	result := system.ChangeRoleStatus(rolesParam, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, result)
}

func DeleteRole(context *gin.Context) {
	userId, _ := context.Get("userId")
	var roleIds = context.Param("roleIds")
	system.DeleteRolesById(roleIds, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}

/*不需要分页*/
func GetRoleOptionSelect(context *gin.Context) {
	result := system.GetRoleOptionSelect()
	context.JSON(http.StatusOK, result)
}

/*分页*/
func GetAllocatedList(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	roleId, _ := strconv.Atoi(context.DefaultQuery("roleId", "0"))
	var userName = context.DefaultQuery("userName", "")
	var phonenumber = context.DefaultQuery("phonenumber", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysUserParm{
			RoleId:      roleId,
			UserName:    userName,
			Phonenumber: phonenumber,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	result := system.GetAllocatedList(param)
	context.JSON(http.StatusOK, result)
}

/*分页*/
func GetUnAllocatedList(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	roleId, _ := strconv.Atoi(context.DefaultQuery("roleId", "0"))
	var userName = context.DefaultQuery("userName", "")
	var phonenumber = context.DefaultQuery("phonenumber", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysUserParm{
			RoleId:      roleId,
			UserName:    userName,
			Phonenumber: phonenumber,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}

	result := system.GetUnAllocatedList(param)
	context.JSON(http.StatusOK, result)
}

func CancelRole(context *gin.Context) {
	var rolesParam system.SysUserRolesParam
	if err := context.ShouldBind(&rolesParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	result := system.CancelRole(rolesParam.UserId, rolesParam.RoleId)
	context.JSON(http.StatusOK, result)
}
func CancelAllRole(context *gin.Context) {
	roleId, _ := context.GetQuery("roleId")
	userIds, _ := context.GetQuery("userIds")

	result := system.CancelAllRole(roleId, userIds)
	context.JSON(http.StatusOK, result)
}

func SelectRoleAll(context *gin.Context) {
	roleId, _ := context.GetQuery("roleId")
	userIds, _ := context.GetQuery("userIds")
	userId, _ := context.Get("userId")
	result := system.SelectRoleAll(roleId, userIds, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, result)
}

func GetDeptTreeRole(context *gin.Context) {
	roleId := context.Param("roleId")
	checkedKeys := system.GetDeptTreeRole(roleId)
	depts := system.SelectDeptTreeList()
	context.JSON(http.StatusOK, gin.H{
		"msg":         "操作成功",
		"code":        http.StatusOK,
		"checkedKeys": checkedKeys,
		"depts":       depts,
	})
}
