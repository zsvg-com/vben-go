package timex

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 1. 定义自定义时间类型
type Time time.Time

// 2. 实现 MarshalJSON 方法，定义序列化（Go -> JSON）的格式
func (t Time) MarshalJSON() ([]byte, error) {
	// 格式化为 "2006-01-02 15:04:05"
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// 3. 实现 UnmarshalJSON 方法，定义反序列化（JSON -> Go）的格式
func (t *Time) UnmarshalJSON(data []byte) error {
	// 去除JSON字符串两端的引号
	if string(data) == "null" {
		return nil
	}
	str := string(data)
	if len(str) > 2 {
		str = str[1 : len(str)-1]
	}
	// 解析时间字符串
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

// 4. (可选，用于GORM) 实现数据库的 Scan 和 Value 方法
func (t *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*t = Time(v)
	case []byte:
		parsedTime, _ := time.Parse("2006-01-02 15:04:05.000", string(v))
		*t = Time(parsedTime)
	case string:
		parsedTime, _ := time.Parse("2006-01-02 15:04:05.000", v)
		*t = Time(parsedTime)
	case nil:
		*t = Time(time.Time{})
	default:
		return fmt.Errorf("无法扫描类型 %T 到 LocalTime", value)
	}
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t).Format("2006-01-02 15:04:05.000"), nil
}
