package {{.ModuleName}}

import (
    "vben/app/pkg/utils"
    "vben/app/pkg/utils/R"
)

// Model {{.Table_Name}}  {{.TableComment}}
type {{.ClassName}} struct {
     {{- range .Fields}}
     {{.FieldName}} {{.JavaType}} `json:"{{.JavaField}}" gorm:"{{.ColumnName}}"`//{{.ColumnComment}}{{ end }}
}


func ({{.ClassName}}) TableName() string {
  return "{{.Table_Name}}"
}

//创建
func Create{{.ClassName}}Service(data {{.ClassName}}) any {
    err := mysql.MysqlDb().Create(&data).Error
	if err!=nil{
        return R.ReturnFailMsg("保存失败")
	}
	return R.ReturnSuccess(nil)
}

//根据ID删除
func Delete{{.ClassName}}Service(id string) any {
    err := mysql.MysqlDb().Delete(&{{.ClassName}}{},"id = ?",id).Error
	if err!=nil{
        return R.ReturnFailMsg("删除失败")
    }

    return R.ReturnSuccess(nil)
}

//根据ID批量删除
func Delete{{.ClassName}}ByIdsService(ids string) any {
	if len(ids) == 0{
         return R.ReturnFailMsg("参数获取失败")
    }

    err := mysql.MysqlDb().Delete(&{{.ClassName}}{},"id in ?",ids).Error
    if err!=nil{
       return R.ReturnFailMsg("批量删除失败")
    }

    return R.ReturnSuccess(nil)
}

//根据id 更新 ，排除零值
func Update{{.ClassName}}Service(data {{.ClassName}}) any {
     err := mysql.MysqlDb().Updates(&data).Error
    if err!=nil{
        return R.ReturnFailMsg("更新失败")
    }
    return R.ReturnSuccess(nil)
}

//根据id获取model
func Get{{.ClassName}}Service(id int64) any {

    if id == 0{
        return R.ReturnFailMsg("参数获取失败")
    }
    var data {{.ClassName}}
    err := mysql.MysqlDb().Where("id = ?", id).First(&data).Error
    if err!=nil{
        return R.ReturnFailMsg("获取数据失败")
    }else{
        return R.ReturnSuccess(data)
    }

}

//获取所有的model
func GetList{{.ClassName}}Service() any {
    var list []{{.ClassName}}
    err := mysql.MysqlDb().Find(&list).Error
	if err!=nil {
        return R.ReturnFailMsg("获取数据失败")
    }else{
        return R.ReturnSuccess(list)
    }
}

//按条件分页查询 limit offset ,参数用指针&, 数据会自动填充到req对象
func GetPageLimit{{.ClassName}}Service(params tools.SearchTableDataParam) tools.TableDataInfo {

    var pageNum = params.PageNum
	var pageSize = params.PageSize

	var total int64
	db := mysql.MysqlDb().Model({{.ClassName}}{})
	// 可以自定义搜索方式

	var rows []{{.ClassName}}

	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}
	offset := (pageNum - 1) * pageSize
	db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&rows)
	if rows == nil {
        return tools.Fail()
    } else {
        return tools.Success(rows, total)
    }
}