package menu

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"vben/app/core/utils/R"
	"vben/common/json/timex"
	"vben/common/r"
	"vben/common/utils"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Get(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var list []SysMenu
	queryDB := mysql.MysqlDb().Model(&SysMenu{})
	pagination := &r.PaginationQuery{
		Page:     pageNum,
		PageSize: pageSize,
	}
	paginate, err := r.Paginate(queryDB, &list, pagination)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, r.Ok(paginate))
}

func GetList(c *gin.Context) {
	var sql = "select * from sys_menu"
	var list []SysMenu
	err := mysql.MysqlDb().Raw(sql).Scan(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(list))
}

func GetTree(c *gin.Context) {
	var sql = "select id,name,pid from sys_menu"
	var list []r.Stree
	err := mysql.MysqlDb().Raw(sql).Scan(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	tree := r.TreeBuild(list)
	c.JSON(http.StatusOK, r.Ok(tree))
}

func GetInfo(c *gin.Context) {
	id := c.Param("id")
	var main SysMenu
	mysql.MysqlDb().First(&main, "id = ?", id)
	if main.Id == 0 {
		panic(R.ReturnFailMsg("未找到记录"))
	}
	c.JSON(http.StatusOK, r.Ok(main))
}

func Post(c *gin.Context) {
	var main SysMenu
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Id = utils.NextId()
	main.Crtim = timex.Time(time.Now())
	main.Uptim = main.Crtim
	err := mysql.MysqlDb().Create(&main).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(main.Id))

}

func Put(c *gin.Context) {
	var main SysMenu
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Uptim = timex.Time(time.Now())
	err := mysql.MysqlDb().Select("*").Omit("crtim,cruid").Updates(&main).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(main.Id))

}

func Delete(c *gin.Context) {
	var ids = c.Param("ids")
	arr := strings.Split(ids, ",")
	for i := 0; i < len(arr); i++ {
		err := mysql.MysqlDb().Where("id = ?", arr[i]).Delete(&SysMenu{}).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	c.JSON(http.StatusOK, r.Ok(len(arr)))
}
