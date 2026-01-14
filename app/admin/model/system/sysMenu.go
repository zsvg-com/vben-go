package system

import (
	"strconv"
	"strings"
	"time"
	"vben/app/admin/model/tools"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysMenu model：数据库字段
type SysMenu struct {
	MenuId     int       `json:"menuId" gorm:"column:menu_id;primaryKey"` //表示主键
	MenuName   string    `json:"menuName" gorm:"menu_name"`
	ParentId   int       `json:"parentId" gorm:"parent_id"`
	OrderNum   int       `json:"orderNum" gorm:"order_num"`
	MenuType   string    `json:"menuType" gorm:"menu_type"`
	Visible    string    `json:"visible" gorm:"visible"`
	Perms      string    `json:"perms" gorm:"perms"`
	Query      string    `json:"query" gorm:"query"`
	IsFrame    string    `json:"isFrame" gorm:"is_frame"`
	Icon       string    `json:"icon" gorm:"icon"`
	Path       string    `json:"path" gorm:"path"`
	Status     string    `json:"status" gorm:"status"`
	IsCache    string    `json:"isCache" gorm:"is_cache"`
	Component  string    `json:"component" gorm:"component"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark     string    `json:"remark" gorm:"remark"`
}

type MenuVo struct {
	Name       string   `json:"name"`
	Path       string   `json:"path,omitempty"`
	Hidden     bool     `json:"hidden" `
	Redirect   string   `json:"redirect,omitempty"`
	Component  string   `json:"component,omitempty" `
	Query      string   `json:"query,omitempty"`
	AlwaysShow bool     `json:"alwaysShow,omitempty" `
	MetaVo     MetaVo   `json:"meta" `
	Children   []MenuVo `json:"children,omitempty"`
}

type MetaVo struct {
	Title   string `json:"title"`
	Icon    string `json:"icon" `
	NoCache bool   `json:"noCache" `
	Link    string `json:"link,omitempty" `
}

type MenuTreeSelect struct {
	Id       int              `json:"id"`
	Label    string           `json:"label"`
	Children []MenuTreeSelect `json:"children,omitempty"`
}

// TableName 指定数据库表名称
func (SysMenu) TableName() string {
	return "sys_menu"
}

func SelectMenuTreeByUserId(userId int) []SysMenu {
	var menu []SysMenu
	if IsAdminById(userId) {
		err := mysql.MysqlDb().Where("menu_type in (?) ", []string{"M", "C"}).Where("status", "0").Order("parent_id").Order("order_num").Find(&menu).Error
		if err != nil {
			panic(R.ReturnFailMsg("查询错误"))
		}
	} else {
		err := mysql.MysqlDb().Raw("select distinct m.* from sys_menu m left join sys_role_menu rm on m.menu_id = rm.menu_id left join sys_user_role ur on rm.role_id = ur.role_id left join sys_role ro on ur.role_id = ro.role_id left join sys_user u on ur.user_id = u.user_id where u.user_id = ? and m.menu_type in ('M', 'C') and m.status = 0  AND ro.status = 0 order by m.parent_id, m.order_num", userId).Scan(&menu).Error
		if err != nil {
			panic(R.ReturnFailMsg("查询错误"))
		}
	}
	return menu
}

func SelectMenuTree(userId int, menu SysMenu) []SysMenu {
	var menus []SysMenu
	sql := "select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time "
	sql += "from sys_menu m "
	if IsAdminById(userId) {
		var menuName = menu.MenuName
		if menuName != "" {
			sql += "AND m.menu_name like concat('%', " + menuName + ", '%')"
		}
		var visible = menu.Visible
		if visible != "" {
			sql += "AND m.visible = " + visible
		}
		var status = menu.Status
		if status != "" {
			sql += "AND m.status = " + status
		}
		err := mysql.MysqlDb().Raw(sql).Scan(&menus).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	} else {
		sql += "left join sys_role_menu rm on m.menu_id = rm.menu_id "
		sql += "left join sys_user_role ur on rm.role_id = ur.role_id "
		sql += "left join sys_role ro on ur.role_id = ro.role_id "
		sql += "where ur.user_id = " + strconv.Itoa(userId)
		var menuName = menu.MenuName
		if menuName != "" {
			sql += "AND m.menu_name like concat('%', " + menuName + ", '%')"
		}
		var visible = menu.Visible
		if visible != "" {
			sql += "AND m.visible = " + visible
		}
		var status = menu.Status
		if status != "" {
			sql += "AND m.status = " + status
		}
		err := mysql.MysqlDb().Raw(sql).Scan(&menus).Error
		if err != nil {
			panic(R.ReturnFailMsg(err.Error()))
		}
	}
	return menus
}

func BuildMenuTreeSelect(lists []SysMenu) []MenuTreeSelect {
	var menuTreeSelect []MenuTreeSelect
	for i := 0; i < len(lists); i++ {
		var menu = lists[i]
		menuId := menu.MenuId
		parentId := menu.ParentId
		if 0 == parentId {
			var menuVo = MenuTreeSelect{
				Id:    menuId,
				Label: menu.MenuName,
			}
			menuVo.Children = BuildChildMenusTreeSelect(menuId, lists)
			menuTreeSelect = append(menuTreeSelect, menuVo)
		}
	}
	return menuTreeSelect
}

func BuildChildMenusTreeSelect(ParentId int, lists []SysMenu) []MenuTreeSelect {
	var List []MenuTreeSelect
	for i := 0; i < len(lists); i++ {
		var menu = lists[i]
		var menuId = menu.MenuId
		var pId = menu.ParentId
		if pId == ParentId {
			var menuVo = MenuTreeSelect{
				Id:    menuId,
				Label: menu.MenuName,
			}
			menuVo.Children = BuildChildMenusTreeSelect(menuId, lists)
			List = append(List, menuVo)
		}
	}
	return List
}

func BuildMenus(lists []SysMenu) []MenuVo {
	var menuVos []MenuVo
	for i := 0; i < len(lists); i++ {
		var menu = lists[i]
		MenuId := menu.MenuId
		parentId := menu.ParentId
		if 0 == parentId {
			var path = ""
			if isInnerLink(menu.Path) {
				path = menu.Path
			}
			var menuVo = MenuVo{
				Hidden: "1" == menu.Visible,
				Query:  menu.Query,
				MetaVo: MetaVo{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: "1" == menu.IsCache,
					Link:    path,
				},
				Name:      getRouteName(menu),
				Path:      getRouterPath(menu),
				Component: getComponent(menu),
			}
			if "M" == menu.MenuType {
				if !isInnerLink(menu.Path) {
					menuVo.AlwaysShow = true
					menuVo.Redirect = "noRedirect"
					menuVo.Children = BuildChildMenus(MenuId, lists)
				}
			}
			menuVos = append(menuVos, menuVo)
		}

	}
	return menuVos
}

func BuildChildMenus(ParentId int, lists []SysMenu) []MenuVo {
	var List []MenuVo
	for i := 0; i < len(lists); i++ {
		var menu = lists[i]
		var menuId = menu.MenuId
		var pId = menu.ParentId
		if pId == ParentId {
			var path = ""
			if isInnerLink(menu.Path) {
				path = menu.Path
			}
			var menuVo = MenuVo{
				Hidden: "1" == menu.Visible,
				Query:  menu.Query,
				MetaVo: MetaVo{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: "1" == menu.IsCache,
					Link:    path,
				},
				Name:      getRouteName(menu),
				Path:      getRouterPath(menu),
				Component: getComponent(menu),
			}
			if "M" == menu.MenuType {
				if !isInnerLink(menu.Path) {
					menuVo.AlwaysShow = true
					menuVo.Redirect = "noRedirect"
					menuVo.Children = BuildChildMenus(menuId, lists)
				}
			}
			List = append(List, menuVo)
		}
	}
	return List
}

func getRouteName(menu SysMenu) string {
	var name = FirstUpper(menu.Path)
	if isMenuFrame(menu) {
		return ""
	}
	return name
}

func getRouterPath(menu SysMenu) string {
	var routerPath = menu.Path
	if isInnerLink(routerPath) {
		return routerPath
	}
	// 非外链并且是一级目录（类型为目录）
	if 0 == menu.ParentId && "M" == menu.MenuType && "1" == menu.IsFrame {
		routerPath = "/" + menu.Path
	} else if isMenuFrame(menu) {
		routerPath = "/"
	}
	return routerPath
}

func getComponent(menu SysMenu) string {
	var component = "Layout"
	if "" != menu.Component && !isMenuFrame(menu) {
		component = menu.Component
	} else if "" == menu.Component && isInnerLink(menu.Path) {
		component = "InnerLink"
	} else if "" == menu.Component && isParentView(menu) {
		component = "ParentView"
	}
	return component
}

func isParentView(menu SysMenu) bool {
	return menu.ParentId != 0 && "M" == menu.MenuType
}

// 是否为外链
func isInnerLink(path string) bool {
	return strings.Contains(path, "http://") || strings.Contains(path, "https://")
}

func isMenuFrame(menu SysMenu) bool {
	return menu.ParentId == 0 && "C" == menu.MenuType && menu.IsFrame == "1"
}

/*首字母大写*/
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// SelectSysMenuList /*分页查询*/
func SelectSysMenuList(params tools.SearchTableDataParam) []SysMenu {
	sysMenu := params.Other.(SysMenu)
	var rows []SysMenu
	var sql = "select menu_id, menu_name, parent_id, order_num, path, component, `query`, is_frame, is_cache, menu_type, visible, " +
		"status, ifnull(perms,'') as perms, icon, create_time from sys_menu where 1 = 1 "
	var name = sysMenu.MenuName
	if name != "" {
		sql += " AND menu_name LIKE CONCAT('%', '" + name + "', '%')"
	}
	var visible = sysMenu.Visible
	if visible != "" {
		sql += " AND visible = " + visible
	}
	var status = sysMenu.Status
	if status != "" {
		sql += " AND status = " + status
	}
	err := mysql.MysqlDb().Raw(sql).Find(&rows).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return rows
}

func SelectMenuListByRoleId(roleId string, menuCheckStrictly bool) []int {
	var sql = "select m.menu_id from sys_menu m " +
		"left join sys_role_menu rm on m.menu_id = rm.menu_id " +
		"where rm.role_id = " + roleId + " "
	if menuCheckStrictly {
		sql += "and m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = " + roleId + ")"
	}
	sql += "order by m.parent_id, m.order_num"
	var menuIds []int
	err := mysql.MysqlDb().Raw(sql).Scan(&menuIds).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return menuIds
}

func SelectSysMenuListByUserId(userId int, params tools.SearchTableDataParam) []SysMenu {
	sysMenu := params.Other.(SysMenu)
	var sql = "select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.visible," +
		" m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time " +
		"from sys_menu m left join sys_role_menu rm on m.menu_id = rm.menu_id " +
		"left join sys_user_role ur on rm.role_id = ur.role_id " +
		"left join sys_role ro on ur.role_id = ro.role_id "

	sql += "where ur.user_id = " + strconv.Itoa(userId) + " "
	var menuName = sysMenu.MenuName
	if menuName != "" {
		sql += "AND m.menu_name like concat(%" + menuName + "%) "
	}
	var visible = sysMenu.Visible

	if visible != "" {
		sql += "AND m.visible = " + visible + " "
	}
	var status = sysMenu.Status
	if status != "" {
		sql += "AND m.status = " + status + " "
	}

	sql += "order by m.parent_id, m.order_num"
	var list []SysMenu
	err := mysql.MysqlDb().Raw(sql).Scan(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return list
}

func FindMenuInfoById(menuId string) SysMenu {
	sql := "select menu_id, menu_name, parent_id, order_num, path, " +
		"component, `query`, is_frame, is_cache, menu_type, visible, status, ifnull(perms,'') as perms, icon, create_time " +
		"from sys_menu "
	sql = sql + "where menu_id = " + menuId
	var list SysMenu
	err := mysql.MysqlDb().Raw(sql).First(&list).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return list
}

func hasChildByMenuId(menuId string) SysMenu {
	var menu SysMenu
	mysql.MysqlDb().Where("parent_id = ? ", menuId).First(&menu)
	return menu
}

func hasChildCountByMenuId(menuId string) int {
	var menuCount int
	err := mysql.MysqlDb().Where("parent_id = ? ", menuId).Scan(&menuCount).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return menuCount
}

func checkMenuExistRole(menuId string) int {
	var menuCount int
	err := mysql.MysqlDb().Raw("select count(1) from sys_role_menu where menu_id = " + menuId).Scan(&menuCount).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return menuCount
}

func DeleteMenu(menuIds string) R.Result {
	menu := hasChildByMenuId(menuIds)
	if menu.MenuId != 0 {
		return R.ReturnFailMsg("存在子菜单,不允许删除")
	}

	menuCount := checkMenuExistRole(menuIds)
	if menuCount > 0 {
		return R.ReturnFailMsg("菜单已分配,不允许删除")
	}

	err := mysql.MysqlDb().Exec("delete from sys_menu where menu_id in (?) ", menuIds).Error
	if err == nil {
		R.ReturnFailMsg("删除部门关联用户失败")
	}
	return R.ReturnSuccess("操作成功")
}

func checkMenuNameUnique(parentId int, menuName string) int {
	var menuCount int
	err := mysql.MysqlDb().Raw("select count(1) from sys_menu where menu_name = "+menuName+" and parent_id = "+strconv.Itoa(parentId), &menuCount).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return menuCount
}

func AddMenu(sysMenu SysMenu) R.Result {
	count := checkMenuNameUnique(sysMenu.ParentId, sysMenu.MenuName)
	if count > 0 {
		return R.ReturnFailMsg("新增菜单" + sysMenu.MenuName + "失败，菜单名称已存在")
	}
	menuPath := sysMenu.Path
	isFrame := sysMenu.IsFrame
	isPath := isInnerLink(menuPath)
	if isFrame == "true" && !isPath {
		return R.ReturnFailMsg("新增菜单" + sysMenu.MenuName + "失败，地址必须以http(s)://开头")
	}
	err := mysql.MysqlDb().Create(&sysMenu).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess("操作成功")

}

func UpdateMenu(sysMenu SysMenu) R.Result {
	mId := sysMenu.MenuId
	pId := sysMenu.ParentId
	if pId == mId {
		return R.ReturnFailMsg("修改菜单" + sysMenu.MenuName + "失败，上级菜单不能选择自己")
	}
	count := checkMenuNameUnique(pId, sysMenu.MenuName)
	if count > 0 {
		return R.ReturnFailMsg("修改菜单失败，菜单名称已存在" + sysMenu.MenuName + "失败，菜单名称已存在")
	}
	isFrame := sysMenu.IsFrame
	menuPath := sysMenu.Path
	isPath := isInnerLink(menuPath)
	if isFrame == "true" && !isPath {
		return R.ReturnFailMsg("修改菜单" + sysMenu.MenuName + "失败，地址必须以http(s)://开头")
	}
	err := mysql.MysqlDb().Updates(&sysMenu).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess("操作成功")
}
