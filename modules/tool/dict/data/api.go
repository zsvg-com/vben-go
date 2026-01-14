package dictd

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
	//dicid, _ := strconv.ParseInt(c.DefaultQuery("dicid", 0))

	dicidStr := c.DefaultQuery("dicid", "0") // 默认值为 "0"
	dicid, err := strconv.ParseInt(dicidStr, 10, 64)
	if err != nil {
		// 转换错误处理
		c.JSON(400, gin.H{"error": "参数id格式错误"})
		return
	}

	//dicidStr := c.Query("dicid")
	//if dicidStr == "" {
	//	// 处理空值
	//	c.JSON(400, gin.H{"error": "参数id不能为空"})
	//	return
	//}
	//dicid, err := strconv.ParseInt(dicidStr, 10, 64)
	//if err != nil {
	//	// 转换错误处理
	//	c.JSON(400, gin.H{"error": "参数dicid格式错误"})
	//	return
	//}

	var list []ToolDictData
	queryDB := mysql.MysqlDb().Model(&ToolDictData{}).Where("dicid = ?", dicid)
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
	var main ToolDictData
	mysql.MysqlDb().First(&main, "id = ?", id)
	if main.Id == 0 {
		panic(R.ReturnFailMsg("未找到记录"))
	}
	c.JSON(http.StatusOK, r.Ok(main))
}

func Post(c *gin.Context) {
	var main ToolDictData
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
	var main ToolDictData
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
		err := mysql.MysqlDb().Where("id = ?", arr[i]).Delete(&ToolDictData{}).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	c.JSON(http.StatusOK, r.Ok(len(arr)))
}
