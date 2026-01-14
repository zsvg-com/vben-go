package tools

import (
	"vben/app/core/utils"
)

type GenTable struct {
	TableId        int64          `json:"tableId" gorm:"primaryKey;column:table_id"`
	Table_Name     string         `json:"tableName,omitempty" gorm:"table_name,omitempty"`
	TableComment   string         `json:"tableComment" gorm:"table_comment"`
	SubTableName   string         `json:"subTableName" gorm:"sub_table_name"`
	SubTableFkName string         `json:"subTableFkName" gorm:"sub_table_fk_name"`
	ClassName      string         `json:"className" gorm:"class_name"`
	TplCategory    string         `json:"tplCategory" gorm:"tpl_category"`
	PackageName    string         `json:"packageName" gorm:"package_name"`
	ModuleName     string         `json:"moduleName" gorm:"module_name"`
	BusinessName   string         `json:"businessName" gorm:"business_name"`
	FunctionName   string         `json:"functionName" gorm:"function_name"`
	FunctionAuthor string         `json:"functionAuthor" gorm:"function_author"`
	GenType        string         `json:"genType" gorm:"gen_type"`
	GenPath        string         `json:"genPath" gorm:"gen_path"`
	Options        string         `json:"options" gorm:"options"`
	CreateBy       string         `json:"-" gorm:"create_by"`
	CreateTime     utils.JsonTime `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy       string         `json:"-" gorm:"update_by"`
	UpdateTime     utils.JsonTime `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark         string         `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (GenTable) TableName() string {
	return "gen_table"
}

type GenTableColumn struct {
	ColumnId      int64          `json:"columnId" gorm:"primaryKey;column:column_id"`
	TableId       int64          `json:"tableId" gorm:"column:table_id"`
	ColumnName    string         `json:"columnName" gorm:"column_name"`
	ColumnComment string         `json:"columnComment" gorm:"column_comment"`
	ColumnType    string         `json:"columnType" gorm:"column_type"`
	JavaType      string         `json:"javaType" gorm:"java_type"`
	JavaField     string         `json:"javaField" gorm:"java_field"`
	IsPk          string         `json:"isPk" gorm:"is_pk"`
	IsIncrement   string         `json:"isIncrement" gorm:"is_increment"`
	IsRequired    string         `json:"isRequired" gorm:"is_required"`
	IsInsert      string         `json:"isInsert" gorm:"is_insert"`
	IsEdit        string         `json:"isEdit" gorm:"is_edit"`
	IsList        string         `json:"isList" gorm:"is_list"`
	IsQuery       string         `json:"isQuery" gorm:"is_query"`
	QueryType     string         `json:"queryType" gorm:"query_type"`
	HtmlType      string         `json:"htmlType" gorm:"html_type"`
	DictType      string         `json:"dictType" gorm:"dict_type"`
	Sort          string         `json:"sort" gorm:"sort"`
	CreateBy      string         `json:"createBy" gorm:"create_by"`
	CreateTime    utils.JsonTime `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy      string         `json:"updateBy" gorm:"update_by"`
	UpdateTime    utils.JsonTime `json:"updateTime" gorm:"column:update_time;type:datetime"`
}

// TableName 指定数据库表名称
func (GenTableColumn) TableName() string {
	return "gen_table_column"
}

type GenTableVO struct {
	TableId        string             `json:"tableId" gorm:"table_id"`
	Table_Name     string             `json:"tableName,omitempty" gorm:"table_name,omitempty"`
	TableComment   string             `json:"tableComment" gorm:"table_comment"`
	SubTableName   string             `json:"subTableName" gorm:"sub_table_name"`
	SubTableFkName string             `json:"subTableFkName" gorm:"sub_table_fk_name"`
	ClassName      string             `json:"className" gorm:"class_name"`
	TplCategory    string             `json:"tplCategory" gorm:"tpl_category"`
	PackageName    string             `json:"packageName" gorm:"package_name"`
	ModuleName     string             `json:"moduleName" gorm:"module_name"`
	BusinessName   string             `json:"businessName" gorm:"business_name"`
	FunctionName   string             `json:"functionName" gorm:"function_name"`
	FunctionAuthor string             `json:"functionAuthor" gorm:"function_author"`
	GenType        string             `json:"genType" gorm:"gen_type"`
	GenPath        string             `json:"genPath" gorm:"gen_path"`
	Options        string             `json:"options" gorm:"options"`
	Fields         []GenTableColumnVO `json:"fields"`
}

type GenTableColumnVO struct {
	TableId        string `json:"tableId" gorm:"table_id"`
	Table_Name     string `json:"tableName,omitempty" gorm:"table_name,omitempty"`
	TableComment   string `json:"tableComment" gorm:"table_comment"`
	SubTableName   string `json:"subTableName" gorm:"sub_table_name"`
	SubTableFkName string `json:"subTableFkName" gorm:"sub_table_fk_name"`
	ClassName      string `json:"className" gorm:"class_name"`
	TplCategory    string `json:"tplCategory" gorm:"tpl_category"`
	PackageName    string `json:"packageName" gorm:"package_name"`
	ModuleName     string `json:"moduleName" gorm:"module_name"`
	BusinessName   string `json:"businessName" gorm:"business_name"`
	FunctionName   string `json:"functionName" gorm:"function_name"`
	FunctionAuthor string `json:"functionAuthor" gorm:"function_author"`
	GenType        string `json:"genType" gorm:"gen_type"`
	GenPath        string `json:"genPath" gorm:"gen_path"`
	Options        string `json:"options" gorm:"options"`
	ColumnId       string `json:"columnId" gorm:"column_id"`
	ColumnName     string `json:"columnName" gorm:"column_name"`
	ColumnComment  string `json:"columnComment" gorm:"column_comment"`
	ColumnType     string `json:"columnType" gorm:"column_type"`
	JavaType       string `json:"javaType" gorm:"java_type"`
	FieldName      string `json:"fieldName"`
	JavaField      string `json:"javaField" gorm:"java_field"`
	IsPk           string `json:"isPk" gorm:"is_pk"`
	IsIncrement    string `json:"isIncrement" gorm:"is_increment"`
	IsRequired     string `json:"isRequired" gorm:"is_required"`
	IsInsert       string `json:"isInsert" gorm:"is_insert"`
	IsEdit         string `json:"isEdit" gorm:"is_edit"`
	IsList         string `json:"isList" gorm:"is_list"`
	IsQuery        string `json:"isQuery" gorm:"is_query"`
	QueryType      string `json:"queryType" gorm:"query_type"`
	HtmlType       string `json:"htmlType" gorm:"html_type"`
	DictType       string `json:"dictType" gorm:"dict_type"`
	Sort           string `json:"sort" gorm:"sort"`
}

type EditGenTableVO struct {
	GenTable
	Tree    bool             `json:"tree"`
	Crud    bool             `json:"crud"`
	Sub     bool             `json:"sub"`
	Columns []GenTableColumn `json:"columns"`
}
