package dept

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
	var list []SysDept
	// 1. 构建基础查询
	queryDB := mysql.MysqlDb().Model(&SysDept{})
	//if name != "" {
	//	queryDB = queryDB.Where("name LIKE ?", "%"+name+"%")
	//}
	// 可以继续链式添加其他条件，如 .Where("status = ?", "active")

	// 2. 封装分页请求参数
	pagination := &r.PaginationQuery{
		Page:     pageNum,
		PageSize: pageSize,
	}

	// 3. 调用通用分页函数
	paginate, err := r.Paginate(queryDB, &list, pagination)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, r.Ok(paginate))
}

func GetTree(c *gin.Context) {
	var sql = "select id,name,pid from sys_dept"
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
	var main SysDept
	mysql.MysqlDb().First(&main, "id = ?", id)
	if main.Id == "" {
		panic(R.ReturnFailMsg("未找到记录"))
	}
	c.JSON(http.StatusOK, r.Ok(main))
}

func Post(c *gin.Context) {
	var main SysDept
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
	sysOrg.Type = 1
	err = mysql.MysqlDb().Create(&sysOrg).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(main.Id))

}

func Put(c *gin.Context) {
	var main SysDept
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Uptim = timex.Time(time.Now())
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
		err := mysql.MysqlDb().Where("id = ?", arr[i]).Delete(&SysDept{}).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	c.JSON(http.StatusOK, r.Ok(len(arr)))
}
