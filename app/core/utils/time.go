package utils

import (
	"time"
	"vben/app/admin/model/constants"
)

func GetBeginAndEndTime(beginTime string, endTime string) (time.Time, time.Time) {
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	startTime1, _ := time.ParseInLocation(constants.DateFormat, beginTime, Loc)
	endTime = endTime + " 23:59:59"
	endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
	return startTime1, endTime1
}
