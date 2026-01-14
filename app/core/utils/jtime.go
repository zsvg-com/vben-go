package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
	"vben/app/admin/model/constants"
)

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	dataStr := string(data)
	if dataStr == "" {
		return
	}
	now, err := time.ParseInLocation(`"`+constants.TimeFormat+`"`, dataStr, time.Local)
	*t = JsonTime{now}
	return
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(constants.TimeFormat))
	// 如果时间值是空或者0值 返回为null 如果写空字符串会报错
	if &t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(formatted), nil

}

func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return "", nil
	}
	return t.Time, nil
}

func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
