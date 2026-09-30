package auth

import (
	"net/http"
	"vben/app/core/utils/R"
	lolog "vben/modules/mon/login"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var bo LoginBo
	if err := c.ShouldBind(&bo); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	var vo LoginVo
	vo.AccessToken = "11233"
	vo.ExpireIn = 604800
	lolog.Insert(c, "u1", "登录成功", true)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func Code(c *gin.Context) {
	var vo CaptchaVo
	vo.CaptchaEnabled = false
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": vo,
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": 0,
	})
}
