package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/app/core/utils/R"
)

func ListPost(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	postName := context.DefaultQuery("postName", "")
	postCode := context.DefaultQuery("postCode", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysPost{
			PostName: postName,
			PostCode: postCode,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectSysPostList(param, true)
	context.JSON(http.StatusOK, result)
}
func ExportPost(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	postName := context.DefaultQuery("postName", "")
	postCode := context.DefaultQuery("postCode", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysPost{
			PostName: postName,
			PostCode: postCode,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectSysPostList(param, false)
	var list = result.Rows.([]system.SysPost)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "postId",
		"title":  "岗位序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "postCode",
		"title":  "岗位编码",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "postName",
		"title":  "岗位名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "postSort",
		"title":  "岗位排序",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "状态",
		"width":  "10",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			var status = v.Status
			var statusStr = ""
			if status == "0" {
				statusStr = "正常"
			} else {
				statusStr = "停用"
			}
			data = append(data, map[string]interface{}{
				"postId":   v.PostId,
				"postCode": v.PostCode,
				"postName": v.PostName,
				"postSort": v.PostSort,
				"status":   statusStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}
func GetPostInfo(context *gin.Context) {
	var postId = context.Param("postId")
	result := system.FindPostInfoById(postId)
	context.JSON(http.StatusOK, R.ReturnSuccess(result))
}

func SavePost(context *gin.Context) {
	userId, _ := context.Get("userId")
	var postParam system.SysPost
	if err := context.ShouldBind(&postParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	postParam.CreateBy = user.UserName
	postParam.CreateTime = time.Now()
	result := system.SavePost(postParam)
	context.JSON(http.StatusOK, result)
}

func UploadPost(context *gin.Context) {
	userId, _ := context.Get("userId")
	var postParam system.SysPost
	if err := context.ShouldBind(&postParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	postParam.UpdateBy = user.UserName
	postParam.UpdateTime = time.Now()
	result := system.EditPost(postParam)
	context.JSON(http.StatusOK, result)
}

func DeletePost(context *gin.Context) {
	var postIds = context.Param("postIds")
	result := system.DeletePost(postIds)
	context.JSON(http.StatusOK, result)
}

func GetPostOptionSelect(context *gin.Context) {
	var param = tools.SearchTableDataParam{}
	var result = system.SelectSysPostList(param, false)
	context.JSON(http.StatusOK, result.Rows)
}
