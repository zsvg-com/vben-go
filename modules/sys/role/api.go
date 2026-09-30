package role

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
	var list []SysRole
	queryDB := mysql.MysqlDb().Model(&SysRole{})
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
	var main SysRole
	mysql.MysqlDb().Preload("Users").First(&main, "id = ?", id)
	if main.Id == 0 {
		panic(R.ReturnFailMsg("未找到记录"))
	}

	//成员、菜单与接口
	var actorSql = "select t.id,t.name from sys_actor t inner join sys_role_actor a on a.aid=t.id where a.rid=?"
	err := mysql.MysqlDb().Raw(actorSql, main.Id).Scan(&main.Actors).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	err = mysql.MysqlDb().Raw("select mid id from sys_role_menu where rid=?", main.Id).Scan(&main.Menus).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	err = mysql.MysqlDb().Raw("select aid id from sys_role_api where rid=?", main.Id).Scan(&main.Apis).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	c.JSON(http.StatusOK, r.Ok(main))
}

func Post(c *gin.Context) {
	var main SysRole
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
	insertMapping(main)
	c.JSON(http.StatusOK, r.Ok(main.Id))
}

func insertMapping(main SysRole) {
	actors := make([]SysRoleActor, len(main.Actors))

	for i, src := range main.Actors {
		actors[i] = SysRoleActor{
			Rid: main.Id,
			Aid: src.Id,
		}
	}
	result := mysql.MysqlDb().CreateInBatches(actors, 10)

	if result.Error != nil {
		panic(R.ReturnFailMsg(result.Error.Error()))
	}

	menus := make([]SysRoleMenu, len(main.Menus))

	for i, src := range main.Menus {
		menus[i] = SysRoleMenu{
			Rid: main.Id,
			Mid: src,
		}
	}
	result = mysql.MysqlDb().CreateInBatches(menus, 10)

	if result.Error != nil {
		panic(R.ReturnFailMsg(result.Error.Error()))
	}

	apis := make([]SysRoleApi, len(main.Apis))

	for i, src := range main.Apis {
		apis[i] = SysRoleApi{
			Rid: main.Id,
			Aid: src,
		}
	}
	result = mysql.MysqlDb().CreateInBatches(apis, 10)

	if result.Error != nil {
		panic(R.ReturnFailMsg(result.Error.Error()))
	}
}

func Put(c *gin.Context) {
	var main SysRole
	if err := c.ShouldBindBodyWith(&main, binding.JSON); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
		return
	}
	main.Uptim = timex.Time(time.Now())
	mysql.MysqlDb().Exec("delete from sys_role_actor where rid = ?", main.Id)
	mysql.MysqlDb().Exec("delete from sys_role_menu where rid = ?", main.Id)
	mysql.MysqlDb().Exec("delete from sys_role_api where rid = ?", main.Id)
	insertMapping(main)
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
		err := mysql.MysqlDb().Where("id = ?", arr[i]).Delete(&SysRole{}).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	c.JSON(http.StatusOK, r.Ok(len(arr)))
}

func GetPerms(c *gin.Context) {
	var menuSql = "select id,name,pid,icon,type from sys_menu where avtag=1 order by ornum"
	var menus []MenuVo
	err := mysql.MysqlDb().Raw(menuSql).Scan(&menus).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	var apiSql = "select id,name,menid from sys_api where avtag=1 order by ornum"
	var apis []ApiVo
	err = mysql.MysqlDb().Raw(apiSql).Scan(&apis).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	for i := range menus {
		for j := range apis {
			if apis[j].Menid == menus[i].Id {
				menus[i].Apis = append(menus[i].Apis, apis[j])
			}
		}
	}
	c.JSON(http.StatusOK, r.Ok(menus))
}
