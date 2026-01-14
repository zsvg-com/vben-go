package {{.ModuleName}}

import (
    "strconv"
	"github.com/gin-gonic/gin"
	"vben/app/admin/model/tools"
	"log"
	"net/http"
	"vben/app/pkg/utils/R"
    model "vben/app/admin/model/{{.ModuleName}}"
)

func Create{{.ClassName}}(context *gin.Context) {
	userId, _ := context.Get("userId")
	log.Println(userId)
	var req model.{{.ClassName}}
    if err := context.ShouldBindJSON(&req);err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
	} else {
		res := model.Create{{.ClassName}}Service(req)
		context.JSON(http.StatusOK, res)
	}
}

func Delete{{.ClassName}}(context *gin.Context) {
	var id = context.Param("id")
    result := model.Delete{{.ClassName}}Service(id)
    context.JSON(http.StatusOK, result)
}

func Delete{{.ClassName}}ByIds(context *gin.Context) {
    var ids = context.Param("ids")
    result := model.Delete{{.ClassName}}ByIdsService(ids)
    context.JSON(http.StatusOK, result)
}

func Update{{.ClassName}}(context *gin.Context) {
    var req model.{{.ClassName}}
    if err := context.ShouldBindJSON(&req);err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
	} else {
		res := model.Update{{.ClassName}}Service(req)
		context.JSON(http.StatusOK, res)
	}
}

func Get{{.ClassName}}(context *gin.Context) {
    idstr:=context.Param("id")    //查询路径Path参数
    id, err := strconv.ParseInt(idstr, 10, 64)
    if err!=nil{
        context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
    }else{
        res := model.Get{{.ClassName}}Service(id)
    	context.JSON(http.StatusOK, res)
    }
}

func GetPageLimit{{.ClassName}}(context *gin.Context) {
    pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
    pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

    beginTime := context.DefaultQuery("params[beginTime]", "")
    endTime := context.DefaultQuery("params[endTime]", "")

    var param = tools.SearchTableDataParam{
        PageNum:  pageNum,
        PageSize: pageSize,
        Other: model.{{.ClassName}}{
        },
        Params: tools.Params{
            BeginTime: beginTime,
            EndTime:   endTime,
        },
    }

    result := model.GetPageLimit{{.ClassName}}Service(param)
    context.JSON(http.StatusOK, result)
}
