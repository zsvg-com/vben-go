package post

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"vben/app/core/utils/R"
	"vben/common/json/timex"
	"vben/common/r"
	"vben/common/utils"
	"vben/modules/sys/org"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Get(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	var list []SysPost
	queryDB := mysql.MysqlDb().Model(&SysPost{})
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

func GetInfo(c *gin.Context) {
	id := c.Param("id")
	var main SysPost
	mysql.MysqlDb().Preload("Users").First(&main, "id = ?", id)
	if main.Id == "" {
		panic(R.ReturnFailMsg("未找到记录"))
	}
	c.JSON(http.StatusOK, r.Ok(main))
}

func Post(c *gin.Context) {
	var main SysPost
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Id = utils.NextIdStr()
	main.Crtim = timex.Time(time.Now())
	main.Uptim = main.Crtim
	err := mysql.MysqlDb().Create(&main).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	var sysOrg org.SysOrg
	sysOrg.Id = main.Id
	sysOrg.Name = main.Name
	sysOrg.Type = 4
	err = mysql.MysqlDb().Create(&sysOrg).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(main.Id))
}

func Put(c *gin.Context) {
	var main SysPost
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Uptim = timex.Time(time.Now())
	mysql.MysqlDb().Exec("delete from sys_post_org where pid = ?", main.Id)
	err := mysql.MysqlDb().Select("*").Omit("crtim,cruid").Updates(&main).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	mysql.MysqlDb().Model(&org.SysOrg{}).Where("id = ?", main.Id).Update("name", main.Name)
	c.JSON(http.StatusOK, r.Ok(main.Id))

}

func Delete(c *gin.Context) {
	var ids = c.Param("ids")
	arr := strings.Split(ids, ",")
	for i := 0; i < len(arr); i++ {
		err := mysql.MysqlDb().Where("id = ?", arr[i]).Delete(&SysPost{}).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	c.JSON(http.StatusOK, r.Ok(len(arr)))
}
