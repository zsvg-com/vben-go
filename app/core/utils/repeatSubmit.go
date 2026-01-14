package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
	"vben/app/admin/model/constants"
	"vben/app/core/utils/R"
	"vben/app/core/utils/jwt"
	"vben/pkg/cache/redisCache"
)

type RequestInfo struct {
	RepeatParams string `json:"repeatParams"`
	RepeatTime   int64  `json:"repeatTime"`
}

// RepeatSubmitMiddleware 防重复提交组件 interval 单位 毫秒
func RepeatSubmitMiddleware(intervalMs int64) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		uuid, err := jwt.GetJwtUuid(ctx)
		if len(uuid) <= 0 || err != nil {
			return
		}
		bodyBytes, _ := io.ReadAll(ctx.Request.Body)
		param := string(bodyBytes)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		if bodyBytes == nil {
			param = ctx.Request.URL.Query().Encode()
		}
		requestURL := ctx.Request.URL.String()
		hash := sha256.New()
		hash.Write([]byte(param))
		hashString := hex.EncodeToString(hash.Sum(nil))
		nowDataMap := RequestInfo{
			RepeatParams: hashString,
			RepeatTime:   time.Now().UnixNano() / int64(time.Millisecond),
		}

		key := constants.RepeatSubmitCacheKey + strings.ReplaceAll(requestURL[1:], "/", "-") + "-" + uuid
		get, err := redisCache.NewRedisCache().Get(key)
		if len(get) > 0 {
			sessionMap := make(map[string]RequestInfo)
			err := json.Unmarshal([]byte(get), &sessionMap)
			if err == nil {
				if oldDataMap, exists := sessionMap[requestURL]; exists {
					if oldDataMap.RepeatParams == nowDataMap.RepeatParams && (nowDataMap.RepeatTime-oldDataMap.RepeatTime <= intervalMs) {
						ctx.JSON(http.StatusOK, R.ReturnFailMsg("不允许重复提交，请稍候再试"))
						ctx.Abort()
						return
					}
				}
			}
		}

		cacheMap := map[string]RequestInfo{
			requestURL: nowDataMap,
		}
		jsonData, err := json.Marshal(cacheMap)
		redisCache.NewRedisCache().Put(key, string(jsonData), time.Duration(intervalMs)*time.Millisecond)

	}
}
