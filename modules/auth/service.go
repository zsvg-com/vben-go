package auth

import (
	"strings"
	"vben/modules/sys/menu"
)

const (
	TopParentID int64 = 0
	INNER_LINK        = "InnerLink"
	LAYOUT            = "Layout"
	PARENT_VIEW       = "ParentView"
	NO_REDIRECT       = "noRedirect"
)

// 修改函数返回类型为 []SysMenu
func GetChildPerms(list []menu.SysMenu, parentID int64) []menu.SysMenu {
	var result []menu.SysMenu

	// 找到所有直接子节点
	for _, menu := range list {
		if menu.Pid == parentID {
			result = append(result, menu)
		}
	}

	// 为每个直接子节点递归查找其子节点
	for i := range result {
		// 注意：这里要使用索引修改，不能使用 for _, menu := range result
		result[i].Children = getChildrenRecursive(list, result[i].Id)
	}

	return result
}

// 递归获取所有子节点（返回 []SysMenu）
func getChildrenRecursive(list []menu.SysMenu, parentID int64) []menu.SysMenu {
	var children []menu.SysMenu

	// 查找直接子节点
	for _, menu := range list {
		if menu.Pid == parentID {
			// 递归获取当前节点的子节点
			childMenu := menu
			childMenu.Children = getChildrenRecursive(list, menu.Id)
			children = append(children, childMenu)
		}
	}

	return children
}

// 主要转换函数
func buildMenus(menus []menu.SysMenu) []RouterVo {
	routers := make([]RouterVo, 0, len(menus))

	for _, menu := range menus {
		// 构建路由名称
		name := menu.Name + string(menu.Id)

		// 创建路由对象
		router := RouterVo{
			Hidden:    !menu.Shtag,
			Name:      name,
			Path:      buildPath(menu),
			Component: menu.Comp,
			Query:     menu.Param,
			Meta: MetaVo{
				Title:   menu.Name,
				Icon:    menu.Icon,
				NoCache: !menu.Catag,
				Link:    menu.Path,
			},
		}

		// 处理有子菜单的情况
		if isNotEmpty(menu.Children) && menu.Type == "1" {
			router.AlwaysShow = true
			router.Redirect = NO_REDIRECT
			router.Children = buildMenus(menu.Children)

			if menu.Pid == TopParentID {
				router.Component = LAYOUT
			} else {
				router.Component = PARENT_VIEW
			}
		} else if isMenuFrame(menu) {
			// 处理菜单框架情况
			frameName := upperFirst(menu.Path) + string(menu.Id)
			//router.Meta = nil

			children := RouterVo{
				Path:      menu.Path,
				Component: menu.Comp,
				Name:      frameName,
				Meta: MetaVo{
					Title:   menu.Name,
					Icon:    menu.Icon,
					NoCache: !menu.Catag,
					Link:    menu.Path,
				},
				Query: menu.Param,
			}

			router.Children = []RouterVo{children}
			//} else if menu.Pid == TopParentID && menu.InnerLink {
		} else if menu.Pid == TopParentID {
			// 处理内部链接
			router.Meta = MetaVo{
				Title: menu.Name,
				Icon:  menu.Icon,
			}
			router.Path = "/"

			routerPath := innerLinkReplaceEach(menu.Path)
			innerLinkName := upperFirst(routerPath) + string(menu.Id)

			children := RouterVo{
				Path:      routerPath,
				Component: INNER_LINK,
				Name:      innerLinkName,
				Meta: MetaVo{
					Title: menu.Name,
					Icon:  menu.Icon,
					Link:  menu.Path,
				},
			}

			router.Children = []RouterVo{children}
		}

		routers = append(routers, router)
	}

	return routers
}

// 辅助函数：首字母大写
func upperFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 辅助函数：检查切片是否为空
func isNotEmpty[T any](slice []T) bool {
	return slice != nil && len(slice) > 0
}

// 辅助函数：处理内部链接路径（假设有对应实现）
func innerLinkReplaceEach(path string) string {
	// 这里根据实际业务逻辑实现
	// 假设只是简单处理
	if strings.Contains(path, "http") {
		// 移除协议头等处理
		return strings.Replace(path, "http://", "", 1)
	}
	return path
}

// 构建路径的辅助函数
func buildPath(menu menu.SysMenu) string {
	if menu.Pid == TopParentID {
		return "/" + menu.Path
	}
	return menu.Path
}

func isMenuFrame(m menu.SysMenu) bool {
	return m.Pid == 0 && m.Type == "2" && !bool(m.Outag)
}
