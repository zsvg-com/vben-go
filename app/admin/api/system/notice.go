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

func ListNotice(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	noticeTitle := context.DefaultQuery("noticeTitle", "")
	createBy := context.DefaultQuery("createBy", "")
	noticeType := context.DefaultQuery("noticeType", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysNotice{
			NoticeTitle: noticeTitle,
			CreateBy:    createBy,
			NoticeType:  noticeType,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectSysNoticeList(param, true)
	context.JSON(http.StatusOK, result)
}

func GetNotice(context *gin.Context) {
	var noticeId = context.Param("noticeId")
	result := system.FindNoticeInfoById(noticeId)
	context.JSON(http.StatusOK, R.ReturnSuccess(result))
}

func SaveNotice(context *gin.Context) {
	userId, _ := context.Get("userId")
	var noticeParam system.SysNotice
	if err := context.ShouldBind(&noticeParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	noticeParam.CreateBy = user.UserName
	noticeParam.CreateTime = time.Now()
	result := system.SaveNotice(noticeParam)
	context.JSON(http.StatusOK, result)
}

func UploadNotice(context *gin.Context) {
	userId, _ := context.Get("userId")
	var noticeParam system.SysNotice
	if err := context.ShouldBind(&noticeParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	noticeParam.UpdateBy = user.UserName
	noticeParam.UpdateTime = time.Now()
	result := system.UploadNotice(noticeParam)
	context.JSON(http.StatusOK, result)
}

func DeleteNotice(context *gin.Context) {
	var noticeIds = context.Param("noticeIds")
	result := system.DeleteNotice(noticeIds)
	context.JSON(http.StatusOK, result)
}
