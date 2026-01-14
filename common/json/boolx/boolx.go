package boolx

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// BoolFromInt 是一个自定义类型，用于处理 1/0 到 true/false 的转换
type Bool bool

// Scan - 实现 sql.Scanner 接口，从数据库读取时调用
func (b *Bool) Scan(value interface{}) error {
	if value == nil {
		*b = false
		return nil
	}
	// 处理各种可能的数据库返回值
	switch v := value.(type) {
	case int64:
		*b = v > 0
	case []byte:
		// 处理 MySQL 驱动可能返回的字节切片，如 []byte{1}
		*b = len(v) > 0 && v[0] == 1
	default:
		// 如果已经是bool，直接赋值
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			*b = rv.Int() > 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			*b = rv.Uint() > 0
		case reflect.Bool:
			*b = Bool(rv.Bool())
		default:
			return fmt.Errorf("无法将 %T 扫描为 BoolFromInt", value)
		}
	}
	return nil
}

// Value - 实现 driver.Valuer 接口，写入数据库时调用
func (b Bool) Value() (driver.Value, error) {
	if b {
		return int64(1), nil // 存为 1
	}
	return int64(0), nil // 存为 0
}
