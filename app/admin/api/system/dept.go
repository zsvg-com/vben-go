package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/arrays"
	"net/http"
	"strconv"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"
)

func ListDept(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	deptName := context.DefaultQuery("deptName", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysDept{
			DeptName: deptName,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var rows, total = system.GetDeptList(param, false)

	if total > 0 {
		for _, dept := range rows {
			dept.Children = []system.SysDeptResult{}
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": rows,
	})
}

/*排除节点*/
func ExcludeDept(context *gin.Context) {
	deptId := context.Param("deptId")
	var param = tools.SearchTableDataParam{}
	var list, _ = system.GetDeptList(param, false)
	var ExcludeList []system.SysDeptResult
	for i := 0; i < len(list); i++ {
		bean := list[i]
		ancestors := bean.Ancestors
		ancestors1 := utils.SplitStr(ancestors)
		index := arrays.ContainsString(ancestors1, deptId)
		if deptId != strconv.Itoa(bean.DeptId) || index == -1 {
			ExcludeList = append(ExcludeList, bean)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": ExcludeList,
	})
}

func GetDept(context *gin.Context) {
	deptId := context.Param("deptId")
	result := system.GetDeptInfo(deptId)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func SaveDept(context *gin.Context) {
	userId, _ := context.Get("userId")
	var deptParam system.SysDept
	if err := context.ShouldBind(&deptParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	deptParam.CreateBy = user.UserName
	result := system.SaveDept(deptParam)
	context.JSON(http.StatusOK, result)
}

func UpDataDept(context *gin.Context) {
	userId, _ := context.Get("userId")
	var deptParam system.SysDept
	if err := context.ShouldBind(&deptParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	deptParam.UpdateBy = user.UserName
	result := system.UpDataDept(deptParam)
	context.JSON(http.StatusOK, result)
}

func DeleteDept(context *gin.Context) {
	var deptId = context.Param("deptId")
	result := system.DeleteDept(deptId)
	context.JSON(http.StatusOK, result)
}
