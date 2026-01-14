package monitor

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/monitor"
	"vben/app/core/utils/R"
	"vben/pkg/cache/redisCache"
)

func ListOnLine(context *gin.Context) {
	//TODO 搜索条件
	var ipaddr = context.DefaultQuery("ipaddr", "")
	println(ipaddr)
	var userName = context.DefaultQuery("userName", "")
	println(userName)

	key := constants.LoginCacheKey + "*"
	keyList, _, _ := redisCache.NewRedisCache().Scan(0, key, 0)
	rows := []monitor.LoginUserCache{}
	for i := range keyList {
		keyString := keyList[i]
		result, _ := redisCache.NewRedisCache().Get(keyString)
		var loginUser monitor.LoginUserCache
		json.Unmarshal([]byte(result), &loginUser)
		rows = append(rows, loginUser)
	}

	search := []monitor.LoginUserCache{}

	for i := range rows {
		row := &rows[i]

		if userName != "" || row.UserName == userName {
			if strings.Contains(userName, row.UserName) {
				search = append(search, *row)
			}
		}

		if ipaddr != "" || row.Ipaddr == ipaddr {
			if strings.Contains(ipaddr, row.Ipaddr) {
				search = append(search, *row)
			}
		}
	}

	if userName != "" || ipaddr != "" {
		context.JSON(http.StatusOK, gin.H{
			"msg":   "操作成功",
			"code":  http.StatusOK,
			"total": len(search),
			"rows":  search,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg":   "操作成功",
		"code":  http.StatusOK,
		"total": len(rows),
		"rows":  rows,
	})
}

func DetectOnLine(context *gin.Context) {
	var tokenId = context.Param("tokenId")
	var key = constants.LoginCacheKey + tokenId
	_, error := redisCache.NewRedisCache().Del(key)
	if error != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg(error.Error()))
		return
	}
	context.JSON(http.StatusOK, R.ReturnSuccessMsg("操作成功"))
}
