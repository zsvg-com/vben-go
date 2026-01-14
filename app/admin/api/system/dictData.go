package system

import (
	"net/http"
	"strconv"
	"time"
	"vben/app/admin/model/cache"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"

	"github.com/gin-gonic/gin"
)

func ListDict(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	dictType := context.DefaultQuery("dictType", "")
	status := context.DefaultQuery("status", "")
	dictLabel := context.DefaultQuery("dictLabel", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysDictData{
			DictType:  dictType,
			Status:    status,
			DictLabel: dictLabel,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	result := system.SelectDictDataList(param, false)
	context.JSON(http.StatusOK, gin.H{
		"msg":   "操作成功",
		"code":  http.StatusOK,
		"rows":  result.Rows,
		"total": result.Total,
	})
}

func ExportDict(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other:    system.SysDictData{},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	result := system.SelectDictDataList(param, false)
	var list = result.Rows.([]system.SysDictData)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "dictCode",
		"title":  "字典编码",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictSort",
		"title":  "字典排序",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictLabel",
		"title":  "字典标签",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictValue",
		"title":  "字典键值",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictType",
		"title":  "字典类型",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "isDefault",
		"title":  "是否默认",
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
			defaults := v.IsDefault
			var df = ""
			if "Y" == defaults {
				df = "是"
			}
			if "N" == defaults {
				df = "否"
			}
			var status = v.Status
			statusStr := ""
			if status == "0" {
				statusStr = "正常"
			}
			if status == "1" {
				statusStr = "停用"
			}
			data = append(data, map[string]interface{}{
				"dictCode":  v.DictCode,
				"dictSort":  v.DictSort,
				"dictLabel": v.DictLabel,
				"dictValue": v.DictValue,
				"dictType":  v.DictType,
				"isDefault": df,
				"status":    statusStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetDictCode(context *gin.Context) {
	dictCode := context.Param("dictCode")
	result := system.FindDictCodeById(dictCode)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func GetDictDataByDictType(context *gin.Context) {
	dictType := context.Param("dictType")
	result := cache.GetDictDataCacheByType(dictType)
	if result == nil {
		relResult := system.SelectDictDataByType(dictType)
		result = relResult
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func SaveDictData(context *gin.Context) {
	userId, _ := context.Get("userId")
	var dictDataParam system.SysDictData
	if err := context.ShouldBind(&dictDataParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	dictDataParam.CreateBy = user.UserName
	dictDataParam.CreateTime = time.Now()
	dictDataParam.UpdateTime = time.Now()
	result := system.SaveDictData(dictDataParam)
	cache.RefreshDictDataCache(dictDataParam.DictType)
	context.JSON(http.StatusOK, result)
}

func UpdateDictData(context *gin.Context) {
	userId, _ := context.Get("userId")
	var dictDataParam system.SysDictData
	if err := context.ShouldBind(&dictDataParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	dictDataParam.UpdateBy = user.UserName
	dictDataParam.UpdateTime = time.Now()
	result := system.EditDictData(dictDataParam)
	cache.RefreshDictDataCache(dictDataParam.DictType)
	context.JSON(http.StatusOK, result)
}

func DeleteDictData(context *gin.Context) {
	var dictCodes = context.Param("dictCodes")
	var dictTypes []string
	for _, dictCode := range utils.SplitStr(dictCodes) {
		oneRow := system.FindDictCodeById(dictCode)
		dictTypes = append(dictTypes, oneRow.DictType)
	}
	result := system.DeleteDictData(dictCodes)
	//刷新缓存
	for _, dictType := range dictTypes {
		cache.RefreshDictDataCache(dictType)
	}
	context.JSON(http.StatusOK, result)
}

// ListDictType 分页接口
func ListDictType(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	dictName := context.DefaultQuery("dictName", "")
	dictType := context.DefaultQuery("dictType", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysDictType{
			DictName: dictName,
			DictType: dictType,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectSysDictTypeList(param, true)
	context.JSON(http.StatusOK, result)
}

func ExportType(context *gin.Context) {

	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	dictName := context.DefaultQuery("dictName", "")
	dictType := context.DefaultQuery("dictType", "")
	status := context.DefaultQuery("status", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysDictType{
			DictName: dictName,
			DictType: dictType,
			Status:   status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectSysDictTypeList(param, false)

	var list = result.Rows.([]system.SysDictType)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "dictId",
		"title":  "字典主键",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictName",
		"title":  "字典名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "dictType",
		"title":  "字典类型",
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
			status = v.Status
			var statusStr = ""
			if status == "0" {
				statusStr = "正常"
			}
			if status == "1" {
				statusStr = "停用"
			}
			data = append(data, map[string]interface{}{
				"dictId":   v.DictId,
				"dictName": v.DictName,
				"dictType": v.DictType,
				"status":   statusStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetTypeDict(context *gin.Context) {
	dictId := context.Param("dictId")
	result := system.FindTypeDictById(dictId)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func SaveType(context *gin.Context) {
	var dictTypeParam system.SysDictType
	if err := context.ShouldBind(&dictTypeParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	result := system.SaveType(dictTypeParam)
	cache.RefreshDictDataCache(dictTypeParam.DictType)
	context.JSON(http.StatusOK, result)
}

func UpdateType(context *gin.Context) {
	var dictTypeParam system.SysDictType
	if err := context.ShouldBind(&dictTypeParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	result := system.UpdateType(dictTypeParam)
	cache.RefreshDictDataCache(dictTypeParam.DictType)
	context.JSON(http.StatusOK, result)
}

func DeleteDataType(context *gin.Context) {
	dictIds := context.Param("dictIds")
	var dictTypes []string
	for _, dictId := range utils.SplitStr(dictIds) {
		dictById := system.FindTypeDictById(dictId)
		dictTypes = append(dictTypes, dictById.DictType)
	}
	result := system.DeleteDataType(dictIds)
	for _, dictType := range dictTypes {
		cache.DeleteDictDataCache(dictType)
	}
	context.JSON(http.StatusOK, result)
}

func RefreshCache(context *gin.Context) {
	/*删除缓存*/
	err := cache.DeleteAllDictCache()
	if err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	/*重新赋值初始化参数*/
	cache.InitDictCache()
	context.JSON(http.StatusOK, R.ReturnSuccess("操作成功"))
}

func GetOptionSelect(context *gin.Context) {
	result := system.GetOptionSelect()
	context.JSON(http.StatusOK, result)
}
