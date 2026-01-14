package tools

import "net/http"

type SearchTableDataParam struct {
	PageNum       int    `json:"pageNum"`
	PageSize      int    `json:"pageSize"`
	Other         any    `json:"other"`
	OrderByColumn string `json:"orderByColumn,omitempty"`
	IsAsc         string `json:"isAsc,omitempty"`
	Params        Params `json:"params"`
}

type Params struct {
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
}

type TableDataInfo struct {
	Total int64  `json:"total,omitempty"`
	Rows  any    `json:"rows,omitempty"`
	Code  int    `json:"code,omitempty"`
	Msg   string `json:"msg,omitempty"`
}

func Success(rows any, total int64) TableDataInfo {
	return TableDataInfo{
		Msg:   "查询成功",
		Code:  http.StatusOK,
		Total: total,
		Rows:  rows,
	}
}
func Fail() TableDataInfo {
	return TableDataInfo{
		Msg:  "查询失败",
		Code: http.StatusInternalServerError,
	}
}
