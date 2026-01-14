package r

import (
	"errors"

	"gorm.io/gorm"
)

// 请求参数（通常由前端传入）
type PaginationQuery struct {
	Page     int `form:"page,default=1" binding:"min=1"`        // 页码，默认第1页
	PageSize int `form:"pageSize,default=10" binding:"max=100"` // 每页数量，限制最大值
}

// 分页响应
type PageData[T any] struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	TotalPages  int   `json:"totalPages"`
	Total       int64 `json:"total"`
	Rows        []T   `json:"rows"`
}

// Paginate 通用分页函数
// db: GORM查询句柄，需已设置好模型（Model）和所有查询条件（Where）
// result: 用于存放查询结果的切片指针，如 &[]User{}
// query: 包含 page 和 pageSize 的结构体
func Paginate[T any](db *gorm.DB, result *[]T, query *PaginationQuery) (*PageData[T], error) {
	if db == nil {
		return nil, errors.New("db connection is nil")
	}

	var total int64
	var totalPages int

	// 1. 克隆一个db来执行COUNT查询，避免后续Limit和Offset影响[citation:4]
	countDB := db.Session(&gorm.Session{})
	if err := countDB.Count(&total).Error; err != nil {
		return nil, err
	}

	// 2. 计算总页数
	if total > 0 {
		totalPages = int((total + int64(query.PageSize) - 1) / int64(query.PageSize))
	}

	// 3. 计算Offset并执行分页查询[citation:1]
	offset := (query.Page - 1) * query.PageSize
	if err := db.Limit(query.PageSize).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	// 4. 组装并返回分页响应
	return &PageData[T]{
		CurrentPage: query.Page,
		PageSize:    query.PageSize,
		Total:       total,
		TotalPages:  totalPages,
		Rows:        *result,
	}, nil
}
