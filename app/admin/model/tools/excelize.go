package tools

import (
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// 导出表格

var (
	defaultSheetName = "Sheet1" //默认Sheet名称
	defaultHeight    = 25.0     //默认行高度
)

type lzExcelExport struct {
	file      *excelize.File
	sheetName string //可定义默认sheet名称
}

func NewMyExcel() *lzExcelExport {
	return &lzExcelExport{file: createFile(), sheetName: defaultSheetName}
}

// 导出基本的表格
func (l *lzExcelExport) ExportToPath(params []map[string]string, data []map[string]interface{}, path string) (string, error) {
	l.export(params, data)
	name := createFileName()
	filePath := path + "/" + name
	err := l.file.SaveAs(filePath)
	return filePath, err
}

// 导出到浏览器。此处使用的gin框架 其他框架可自行修改ctx
func (l *lzExcelExport) ExportToWeb(params []map[string]string, data []map[string]interface{}, context *gin.Context) {
	l.export(params, data)
	buffer, _ := l.file.WriteToBuffer()
	//设置文件类型
	context.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	context.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(createFileName()))
	_, _ = context.Writer.Write(buffer.Bytes())
}

// 设置首行
func (l *lzExcelExport) writeTop(params []map[string]string) {
	topStyle, _ := l.file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	var word = 'A'
	//首行写入
	for _, conf := range params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.file.SetCellValue(l.sheetName, line, title)
		//列宽
		_ = l.file.SetColWidth(l.sheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.file.SetCellStyle(l.sheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *lzExcelExport) writeData(params []map[string]string, data []map[string]interface{}) {
	lineStyle, _ := l.file.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		_ = l.file.SetRowHeight(l.sheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.file.SetCellValue(l.sheetName, line, valNum)
			} else {
				_ = l.file.SetCellValue(l.sheetName, line, val[valKey])
			}

			//设置样式
			_ = l.file.SetCellStyle(l.sheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.file.SetRowHeight(l.sheetName, len(data)+1, defaultHeight)
}

func (l *lzExcelExport) export(params []map[string]string, data []map[string]interface{}) {
	l.writeTop(params)
	l.writeData(params, data)
}

func createFile() *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	sheetName := defaultSheetName
	index := f.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}

// excel导出(数据源为Struct) []interface{}
func (l *lzExcelExport) ExportExcelByStruct(titleList []string, data []interface{}, fileName string, sheetName string, context *gin.Context) error {
	l.file.SetSheetName("Sheet1", sheetName)
	header := make([]string, 0)
	for _, v := range titleList {
		header = append(header, v)
	}
	rowStyleID, _ := l.file.NewStyle(`{"font":{"color":"#666666","size":13,"family":"arial"},"alignment":{"vertical":"center","horizontal":"center"}}`)
	_ = l.file.SetSheetRow(sheetName, "A1", &header)
	_ = l.file.SetRowHeight("Sheet1", 1, 30)
	length := len(titleList)
	headStyle := Letter(length)
	var lastRow string
	var widthRow string
	for k, v := range headStyle {

		if k == length-1 {

			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}
	if err := l.file.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		fmt.Print("错误--", err.Error())
	}
	rowNum := 1
	for _, v := range data {

		t := reflect.TypeOf(v)
		fmt.Print("--ttt--", t.NumField())
		value := reflect.ValueOf(v)
		row := make([]interface {
		}, 0)
		for l := 0; l < t.NumField(); l++ {

			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		err := l.file.SetSheetRow(sheetName, "A"+strconv.Itoa(rowNum), &row)
		_ = l.file.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("%s", lastRow), rowStyleID)
		if err != nil {
			return err
		}
	}
	disposition := fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName))
	context.Writer.Header().Set("Content-Type", "application/octet-stream")
	context.Writer.Header().Set("Content-Disposition", disposition)
	context.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	return l.file.Write(context.Writer)
}

// Letter 遍历a-z
func Letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}
