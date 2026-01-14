package tools

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
	"vben/pkg/strcase"

	"github.com/jinzhu/copier"
	"github.com/wxnacy/wgo/file"
)

func GenDbList(params tools.SearchTableDataParam) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	genTable := params.Other.(tools.GenTable)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []tools.GenTable

	sqlCount := "select count(*) as count from information_schema.tables where table_schema = (select database()) AND table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%' AND table_name NOT IN (select table_name from gen_table) "
	sql := "select table_name as Table_Name, table_comment as TableComment, create_time as CreateTime, update_time as UpdateTime from information_schema.tables where table_schema = (select database()) AND table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%' AND table_name NOT IN (select table_name from gen_table) "

	var tableName = genTable.Table_Name
	if tableName != "" {
		sql = sql + "AND lower(table_name) like lower(concat('%', " + tableName + ", '%'))"
	}
	var tableComment = genTable.TableComment
	if tableComment != "" {
		sql = sql + "AND lower(table_comment) like lower(concat('%', " + tableComment + ", '%'))"
	}
	sql = sql + " order by create_time desc"

	if err := mysql.MysqlDb().Raw(sqlCount).Scan(&total).Error; err != nil {
		return tools.Fail()
	}

	sql = sql + " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(pageSize) + ""
	if err := mysql.MysqlDb().Raw(sql).Scan(&rows).Error; err != nil {
		return tools.Fail()
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}
}

func SelectGenList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	genTable := params.Other.(tools.GenTable)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []tools.GenTable

	var db = mysql.MysqlDb().Model(&rows)

	var tableName = genTable.Table_Name
	if tableName != "" {
		db.Where("table_name like concat('%', ?, '%')", tableName)
	}
	var tableComment = genTable.TableComment
	if tableComment != "" {
		db.Where("table_comment like concat('%', ?, '%')", tableComment)
	}

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		db.Where("create_time >= ?", startTime1)
		db.Where("create_time <= ?", endTime1)
	}

	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}

	if isPage {
		if err := db.Limit(pageSize).Offset(offset).Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	} else {
		if err := db.Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}
}

// 批量生成代码zip
func GenBatchCode(table []string, fileName string) string {
	// 运行路径
	pwd, _ := os.Getwd()
	dir := pwd + "/view/template"
	dst := pwd + "/view/template/vben-go"

	if err := DirExistAndMake(dst); err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	// 获取所有表的数据
	autoData := make([]tools.GenTableVO, 0, len(table))

	for _, v := range table {
		log.Printf(v)
		genTable := SelectGenTableByName(v)
		/*测试数据*/
		//genTable := tools.GenTable{
		//	Table_Name:   "xxx",
		//	SubTableName: fileName,
		//	ModuleName:   "auth",
		//	BusinessName: "member",
		//	ClassName:    "RyMember",
		//}
		autoData = append(autoData, genTable)
	}

	// 获取所有的模板
	allTempFile := getTemplateList(dir)

	filePath := dir + "/" + fileName

	archive, err := os.Create(filePath)
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	// 测试
	//for _, filePath := range allTempFile {
	//	log.Printf("模板=" + filePath)
	//	f, _ := os.Open(filePath)
	//	file_name := filepath.Base(filePath)
	//
	//	w1, _ := zipWriter.Create(file_name)
	//	if _, err := io.Copy(w1, f); err != nil {
	//		panic(err)
	//	}
	//}

	// 实际数据
	for _, tv := range autoData { //数据列表
		for _, fv := range allTempFile { // 文件列表
			if err := autocodeFile(tv, fv, zipWriter, dst); err != nil {
				panic(R.ReturnFailMsg(err.Error()))
			}
		}
	}

	zipWriter.Close()

	err = DirExistAndDel(dst)
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	return filePath
}

// 获取所有模板
func getTemplateList(dir string) []string {

	files, err := os.ReadDir(dir) // 找出所有模板文件
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	var allTempFile []string
	for _, v := range files {
		if !v.IsDir() {
			if strings.HasSuffix(v.Name(), ".tpl") {
				allTempFile = append(allTempFile, dir+"/"+v.Name())
			}
		}
	}

	return allTempFile
}

// 赋值到模板里面数据
// 模板生成新的文件
// 文件导入到writer里面
func autocodeFile(tv tools.GenTableVO, fv string, zipWriter *zip.Writer, dst string) error {

	files, err := template.ParseFiles(fv)
	if err != nil {
		return err
	}
	index := strings.LastIndex(fv, "/") + 1

	s := fv[index:strings.Index(fv, ".")]

	/*生成代码的目录*/
	f := dst + "/" + s + "/"

	if err := DirExistAndMake(f); err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	file, _ := os.OpenFile(f+tv.BusinessName+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_RDWR, 0644)
	err = files.Execute(file, tv)
	if err != nil {
		return err
	}

	f1, _ := os.Open(f + tv.BusinessName + ".go")
	/*模块+业务名*/
	fileName := tv.ModuleName + "/" + s + "/" + tv.BusinessName + ".go"

	defer file.Close()
	defer f1.Close()

	w1, _ := zipWriter.Create(fileName)
	if _, err := io.Copy(w1, f1); err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	return nil
}

// 文件不存在就创建文件
func DirExistAndMake(autoPath string) error {
	if !utils2.Exists(autoPath) { // 检查 文件夹是否存在
		if err := os.MkdirAll(autoPath, os.ModePerm); err != nil {
			//logger.Log.WithFields(logrus.Fields{"data": err}).Warn("文件夹不存在，创建文件夹出错")
			return err
		}
	}
	return nil
}

// 文件存在就删除文件
func DirExistAndDel(autoPath string) error {
	if utils2.Exists(autoPath) { // 检查文件是否存在
		if file.IsDir(autoPath) { // 检查是否是文件夹
			if err := os.RemoveAll(autoPath); err != nil {
				return err
			}
		} else {
			if err := os.Remove(autoPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func SelectGenTableByName(tableName string) tools.GenTableVO {
	const sql = "SELECT t.table_id as TableId, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name," +
		" t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark, c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, " +
		" c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort " +
		" FROM gen_table t " +
		" LEFT JOIN gen_table_column c ON t.table_id = c.table_id" +
		" where t.table_name = ? order by c.sort"
	var rows []tools.GenTableColumnVO

	err := mysql.MysqlDb().Model(&tools.GenTableColumnVO{}).Raw(sql, tableName).Scan(&rows).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	var tab tools.GenTableVO

	if rows != nil {
		body := rows[0]
		tab.Table_Name = body.Table_Name
		tab.TableComment = body.TableComment
		tab.SubTableName = body.SubTableName
		tab.SubTableFkName = body.SubTableFkName
		tab.ClassName = body.ClassName
		tab.TplCategory = body.TplCategory
		tab.PackageName = body.PackageName
		tab.ModuleName = body.ModuleName
		tab.BusinessName = body.BusinessName
		tab.FunctionName = body.FunctionName
		tab.FunctionAuthor = body.FunctionAuthor
		tab.GenType = body.GenType
		tab.GenPath = body.GenPath
		tab.Options = body.Options
	}

	list := make([]tools.GenTableColumnVO, 0, len(rows))

	for _, row := range rows {
		row.FieldName = utils2.FirstUpper(row.JavaField)
		javaType := row.JavaType
		if javaType == "String" {
			row.JavaType = "string"
		} else if javaType == "Long" {
			row.JavaType = "int64"
		} else if javaType == "Date" {
			row.JavaType = "utils.JsonTime"
		} else if javaType == "int" {
			row.JavaType = "int"
		} else if javaType == "string" {
			row.JavaType = "string"
		} else if javaType == "int64" {
			row.JavaType = "int64"
		} else if javaType == "utils.JsonTime" {
			row.JavaType = "utils.JsonTime"
		} else if javaType == "time.Time" {
			row.JavaType = "time.Time"
		} else if javaType == "float64" {
			row.JavaType = "float64"
		} else {
			row.JavaType = "string"
		}
		list = append(list, row)
	}

	tab.Fields = list
	return tab
}

func SelectDbTableListByNames(tableName []string) []tools.GenTable {
	var tableVo []tools.GenTable

	query := mysql.MysqlDb().
		Table("information_schema.tables").
		Select("table_name as Table_Name, table_comment as TableComment, create_time as CreateTime, update_time").
		Where("table_name NOT LIKE ?", "qrtz_%").
		Where("table_name NOT LIKE ?", "gen_%").
		Where("table_schema = (SELECT DATABASE())").
		Where("table_name IN ?", tableName)

	if err := query.Scan(&tableVo).Error; err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return tableVo
}

func selectDbTableColumnsByName(tableName string) []tools.GenTableColumn {
	var tableColumn []tools.GenTableColumn
	sql := "SELECT column_name as ColumnName, (case when (is_nullable = 'NO' && column_key != 'PRI') then '1' else null end) as is_required," +
		" (case when extra = 'auto_increment' then '1' else '0' end) as is_increment," +
		" (case when column_key = 'PRI' then '1' else '0' end) as is_pk, DATA_TYPE as ColumnType, ordinal_position as sort, data_type," +
		" COLUMN_COMMENT column_comment " +
		" FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = (select database()) order by ordinal_position"
	err := mysql.MysqlDb().Raw(sql, tableName).Scan(&tableColumn).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return tableColumn
}

func ImportGenTable(genTables []tools.GenTable, userId int) {
	var user = system.FindUserById(userId)

	// 开启事务
	tx := mysql.MysqlDb().Begin()

	for _, genTable := range genTables {
		// 先tableName
		genTable.ClassName = strcase.UpperSnakeCase(genTable.Table_Name) // 首字母大写
		genTable.TplCategory = "crud"
		// 写死
		genTable.PackageName = "com.ruoyi.go"
		genTable.ModuleName = "system"
		genTable.BusinessName = "user"
		genTable.FunctionName = "用户"
		genTable.FunctionAuthor = "OptimisticDevelopers"

		genTable.GenType = "0"
		genTable.GenPath = "/"

		genTable.CreateBy = user.UserName
		genTable.UpdateBy = user.UserName
		genTable.CreateTime = utils2.JsonTime{Time: time.Now()}
		genTable.UpdateTime = utils2.JsonTime{Time: time.Now()}
		// 新增
		result := tx.Create(&genTable)
		if result.Error != nil {
			tx.Rollback()
			panic(R.ReturnFailMsg(result.Error.Error()))
		}

		if result.RowsAffected < 1 {
			tx.Rollback()
			panic(R.ReturnFailMsg("添加失败"))
		}
		// 再genTableColumns
		// 获取
		tableName := genTable.Table_Name
		if tableName == "" {
			tx.Rollback()
			panic(R.ReturnFailMsg("tableName为空"))
		}

		tableColumns := selectDbTableColumnsByName(tableName)
		// 新增
		for _, tableColumn := range tableColumns {
			tableColumn.CreateBy = user.UserName
			tableColumn.UpdateBy = user.UserName
			tableColumn.CreateTime = utils2.JsonTime{Time: time.Now()}
			tableColumn.UpdateTime = utils2.JsonTime{Time: time.Now()}

			TableId := genTable.TableId
			if TableId <= 0 {
				tx.Rollback()
				panic(R.ReturnFailMsg("TableId为空"))
			}
			columnName := tableColumn.ColumnName

			tableColumn.TableId = genTable.TableId

			dataType := utils2.GetDbType(tableColumn.ColumnType)
			tableColumn.JavaType = dataType
			tableColumn.JavaField = strcase.SnakeCase(columnName)
			if tableColumn.IsPk == "1" {
				tableColumn.ColumnName = "primaryKey;column:" + columnName
			}
			err := tx.Model(&tools.GenTableColumn{}).Create(&tableColumn).Error
			if err != nil {
				tx.Rollback()
				panic(R.ReturnFailMsg(err.Error()))
			}
		}
	}
	//提交事务
	tx.Commit()
}

func DeleteGenTableByIds(tableIds string) R.Result {
	if err := mysql.MysqlDb().Model(&tools.GenTable{}).Delete("table_id in (?)", tableIds).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DeleteGenTableColumnByIds(tableIds string) R.Result {
	sql := "delete from gen_table_column where table_id in ( ? )"
	if err := mysql.MysqlDb().Exec(sql, tableIds).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func PreviewGenTableCode(tableId string) map[string]string {
	genTableVO := selectGenTableById(tableId)
	println("" + genTableVO.Table_Name)
	pwd, _ := os.Getwd()
	dir := pwd + "/view/template"
	// 获取所有的模板
	allTempFile := getTemplateList(dir)
	m := make(map[string]string)

	for _, fv := range allTempFile {
		index := strings.LastIndex(fv, "/") + 1
		s := fv[index:strings.Index(fv, ".tpl")]
		str := "vm/go/" + s + ".vm"
		m[str] = getAutocodeFileData(genTableVO, fv, dir)
	}

	return m
}

func selectGenTableById(tableId string) tools.GenTableVO {
	sql := "SELECT t.table_id as TableId, t.table_name as Table_Name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, " +
		" t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, " +
		" t.options, t.remark, c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment," +
		" c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort " +
		" FROM gen_table t" +
		" LEFT JOIN gen_table_column c ON t.table_id = c.table_id where t.table_id = ? order by c.sort"

	var rows []tools.GenTableColumnVO

	err := mysql.MysqlDb().Model(&tools.GenTableColumnVO{}).Raw(sql, tableId).Scan(&rows).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	var tab tools.GenTableVO

	if rows != nil {
		body := rows[0]
		tab.Table_Name = body.Table_Name
		tab.TableComment = body.TableComment
		tab.SubTableName = body.SubTableName
		tab.SubTableFkName = body.SubTableFkName
		tab.ClassName = body.ClassName
		tab.TplCategory = body.TplCategory
		tab.PackageName = body.PackageName
		tab.ModuleName = body.ModuleName
		tab.BusinessName = body.BusinessName
		tab.FunctionName = body.FunctionName
		tab.FunctionAuthor = body.FunctionAuthor
		tab.GenType = body.GenType
		tab.GenPath = body.GenPath
		tab.Options = body.Options
	}

	list := make([]tools.GenTableColumnVO, 0, len(rows))

	for _, row := range rows {
		row.FieldName = utils2.FirstUpper(row.JavaField)
		javaType := row.JavaType
		if javaType == "String" {
			row.JavaType = "string"
		} else if javaType == "Long" {
			row.JavaType = "int64"
		} else if javaType == "Date" {
			row.JavaType = "utils.JsonTime"
		} else if javaType == "int" {
			row.JavaType = "int"
		} else if javaType == "string" {
			row.JavaType = "string"
		} else if javaType == "int64" {
			row.JavaType = "int64"
		} else if javaType == "utils.JsonTime" {
			row.JavaType = "utils.JsonTime"
		} else if javaType == "time.Time" {
			row.JavaType = "time.Time"
		} else if javaType == "float64" {
			row.JavaType = "float64"
		} else {
			row.JavaType = "string"
		}
		list = append(list, row)
	}

	tab.Fields = list
	return tab
}

// 赋值到模板里面数据
// 模板生成新的文件
// 文件导入到writer里面
func getAutocodeFileData(tv tools.GenTableVO, fv string, dst string) string {

	files, err := template.ParseFiles(fv)
	if err != nil {
		return err.Error()
	}
	index := strings.LastIndex(fv, "/") + 1

	// 模板名字
	s := fv[index:strings.Index(fv, ".")]

	/*生成代码的目录*/
	f := dst + "/cache/" + s + "/"

	if err := DirExistAndMake(f); err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	file, _ := os.OpenFile(f+tv.BusinessName+".go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	/*初始化*/
	err = files.Execute(file, tv)
	if err != nil {
		return err.Error()
	}

	defer file.Close()
	f1, _ := os.Open(f + tv.BusinessName + ".go")

	var buf bytes.Buffer
	io.Copy(&buf, f1)
	asString := string(buf.Bytes())

	defer f1.Close()

	return asString
}

func GenInfoService(tableId string) map[string]any {
	m := make(map[string]any)
	genTableVO := selectGenTableInfoById(tableId)
	var rows []tools.GenTableColumn
	mysql.MysqlDb().Order("sort").Where(" table_id = ?", tableId).Model(&tools.GenTableColumn{}).Find(&rows)
	m["info"] = genTableVO
	m["rows"] = rows
	m["tables"] = selectGenTableAllInfoById()
	return m
}

func selectGenTableInfoById(tableId string) tools.EditGenTableVO {

	var editGenTableVO tools.EditGenTableVO
	var genTable tools.GenTable

	err := mysql.MysqlDb().Model(&tools.GenTable{}).Where("table_id = ?", tableId).First(&genTable).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	copier.Copy(&editGenTableVO, genTable)

	var columnVO = selectGenTableColumnById(int64(utils2.GetInterfaceToInt(tableId)))
	editGenTableVO.Columns = columnVO
	editGenTableVO.Tree = false
	editGenTableVO.Crud = true
	editGenTableVO.Sub = false

	return editGenTableVO
}

func selectGenTableAllInfoById() []tools.EditGenTableVO {

	var genTable []tools.GenTable

	err := mysql.MysqlDb().Model(&tools.GenTable{}).Find(&genTable).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	list := make([]tools.EditGenTableVO, 0, len(genTable))

	for _, genTable := range genTable {
		var editGenTableVO tools.EditGenTableVO
		copier.Copy(&editGenTableVO, genTable)
		tableId := editGenTableVO.TableId
		columnVO := selectGenTableColumnById(tableId)
		editGenTableVO.Columns = columnVO
		editGenTableVO.Tree = false
		editGenTableVO.Crud = true
		editGenTableVO.Sub = false
		list = append(list, editGenTableVO)
	}

	return list
}

func selectGenTableColumnById(tableId int64) []tools.GenTableColumn {
	sql := "SELECT t.table_id as TableId, t.table_name as Table_Name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, " +
		" t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, " +
		" t.options, t.remark, c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment," +
		" c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort " +
		" FROM gen_table t" +
		" LEFT JOIN gen_table_column c ON t.table_id = c.table_id where t.table_id = ? order by c.sort"

	var rows []tools.GenTableColumn

	err := mysql.MysqlDb().Model(&tools.GenTableColumn{}).Raw(sql, tableId).Scan(&rows).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

	return rows
}

/*有bug*/
func UpdateGenTableService(tableParam tools.EditGenTableVO) R.Result {
	tx := mysql.MysqlDb().Begin()
	genTable := tools.GenTable{}
	err := copier.Copy(&genTable, tableParam)
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}

	genTable.UpdateTime = utils2.JsonTime{Time: time.Now()}
	err = tx.Model(&tools.GenTable{}).Where("table_id = ?", genTable.TableId).Updates(&genTable).Error
	if err != nil {
		tx.Rollback()
		return R.ReturnFailMsg(err.Error())
	}
	genTableColumnVOs := tableParam.Columns
	for _, genTableColumnVO := range genTableColumnVOs {
		genTableColumn := tools.GenTableColumn{}
		err := copier.Copy(&genTableColumn, genTableColumnVO)
		if err != nil {
			tx.Rollback()
			return R.ReturnFailMsg(err.Error())
		}
		genTableColumn.UpdateTime = utils2.JsonTime{Time: time.Now()}
		err = tx.Model(&tools.GenTableColumn{}).Where("column_id = ?", genTableColumnVO.ColumnId).Updates(&genTableColumn).Error
		if err != nil {
			tx.Rollback()
			return R.ReturnFailMsg(err.Error())
		}
	}

	tx.Commit()
	return R.ReturnSuccess("")
}
