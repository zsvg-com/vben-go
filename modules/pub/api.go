package pub

import (
	"net/http"
	"vben/app/core/utils/R"
	"vben/common/r"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
)

func GetOrg(c *gin.Context) {
	//typ := c.Param("type")
	//pid := c.Param("pid")
	var sql = "select id,name from sys_org"
	var list []IdNameAvatarVo
	err := mysql.MysqlDb().Raw(sql).Scan(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(list))
}

func GetRece(c *gin.Context) {
	//var list []IdNameDeptVo
	c.JSON(http.StatusOK, r.Ok([0]IdNameDeptVo{}))
}

func GetGroupTree(c *gin.Context) {
	//var list []IdNameDeptVo
	c.JSON(http.StatusOK, r.Ok([0]r.Stree{}))
}

func PostRece(c *gin.Context) {
	c.JSON(http.StatusOK, r.Ok(0))
}
