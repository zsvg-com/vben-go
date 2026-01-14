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

func ListConfig(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	configName := context.DefaultQuery("configName", "")
	configKey := context.DefaultQuery("configKey", "")
	configType := context.DefaultQuery("configType", "")

	beginTime := context.DefaultQuery("params[beginTime]", "")
	endTime := context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysConfig{
			ConfigName: configName,
			ConfigKey:  configKey,
			ConfigType: configType,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var result = system.SelectConfigList(param, true)
	context.JSON(http.StatusOK, result)
}

func ExportConfig(context *gin.Context) {

	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var param = tools.SearchTableDataParam{
		PageNum:  0,
		PageSize: 10,
		Other: system.SysConfig{
			ConfigName: configParam.ConfigName,
			ConfigKey:  configParam.ConfigKey,
			ConfigType: configParam.ConfigType,
		},
	}
	tab := system.SelectConfigList(param, false)
	list := tab.Rows.([]system.SysConfig)
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "configId",
		"title":  "参数主键",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configName",
		"title":  "参数名称",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configKey",
		"title":  "参数键名",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configValue",
		"title":  "参数键值",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "configType",
		"title":  "系统内置",
		"width":  "10",
		"is_num": "0",
	})

	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			configType := v.ConfigType
			var configTypeStr = ""
			if "Y" == configType {
				configTypeStr = "是"
			}
			if "N" == configType {
				configTypeStr = "否"
			}
			data = append(data, map[string]interface{}{
				"configId":    v.ConfigId,
				"configName":  v.ConfigName,
				"configKey":   v.ConfigKey,
				"configValue": v.ConfigValue,
				"configType":  configTypeStr,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetConfigInfo(context *gin.Context) {
	configId := context.Param("configId")
	result := system.GetConfigInfo(configId)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": result,
	})
}

func GetConfigKey(context *gin.Context) {
	configKey := context.Param("configKey")
	value := cache.GetSysConfigCacheByKey(configKey)
	if len(value) <= 0 {
		var result = system.SelectConfig(system.SysConfig{ConfigKey: configKey})
		value = result.ConfigValue
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  value,
		"code": http.StatusOK,
	})
}

func SaveConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	configParam.CreateBy = user.UserName
	configParam.CreateTime = time.Now()
	result := system.SaveConfig(configParam)
	cache.RefreshSysConfigCache(configParam.ConfigKey)
	context.JSON(http.StatusOK, result)
}

func UploadConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	var configParam system.SysConfig
	if err := context.ShouldBind(&configParam); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	user := system.FindUserById(userId)
	configParam.UpdateBy = user.UserName
	configParam.UpdateTime = time.Now()
	result := system.EditConfig(configParam)
	cache.RefreshSysConfigCache(configParam.ConfigKey)
	context.JSON(http.StatusOK, result)
}

func DetectConfig(context *gin.Context) {
	userId, _ := context.Get("userId")
	println(userId)
	var configIds = context.Param("configIds")
	var ids = utils.Split(configIds)
	var configKeys []string
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		var config = system.GetConfigInfo(id)
		if "Y" == config.ConfigType {
			context.JSON(http.StatusOK, R.ReturnFailMsg("内置参数"+config.ConfigKey+"不能删除 "))
			return
		}
		configKeys = append(configKeys, config.ConfigKey)
	}
	result := system.DelConfig(configIds)
	for _, key := range configKeys {
		cache.DeleteSysConfigCache(key)
	}
	context.JSON(http.StatusOK, result)
}

func RefreshCacheConfig(context *gin.Context) {
	err := cache.DeleteAllSysConfigCache()
	if err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	cache.InitSysConfigCache()
	context.JSON(http.StatusOK, R.ReturnSuccess("操作成功"))
}

func SelectCaptchaEnabled() bool {
	var configValue = cache.GetSysConfigCacheByKey("sys.account.captchaEnabled")
	if len(configValue) <= 0 {
		configValue = system.SelectConfigByKey("sys.account.captchaEnabled")
	}
	boolValue, err := strconv.ParseBool(configValue)
	if err != nil {
		return true
	}
	return boolValue
}
