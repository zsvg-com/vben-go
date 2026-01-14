package html

import (
	"time"
	"vben/app/html/web"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

var (
	htmlHandler *web.HtmlHandler
)

func Routers(e *gin.Engine) {
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	handlerFunc := cache.CacheByRequestURI(memoryStore, 2*time.Second)
	e.GET("/", htmlHandler.IndexHandler)
	e.GET("/old", htmlHandler.IndexOldHandler)
	e.GET("/protocol.html", handlerFunc, htmlHandler.ProtocolHandler)
}
