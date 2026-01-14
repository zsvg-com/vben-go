package monitor

import (
	"github.com/jinzhu/copier"
	"time"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

type SysJobParam struct {
	JobId          int
	Concurrent     string
	CronExpression string
	InvokeTarget   string
	JobGroup       string
	JobName        string
	MisfirePolicy  string
	Status         string
}
type SysJob struct {
	JobId          int       `json:"jobId" gorm:"column:job_id;primaryKey"` //表示主键
	JobName        string    `json:"jobName" gorm:"job_name"`
	JobGroup       string    `json:"jobGroup" gorm:"job_group"`
	InvokeTarget   string    `json:"invokeTarget" gorm:"invoke_target"`
	CronExpression string    `json:"cronExpression" gorm:"cron_expression"`
	MisfirePolicy  int       `json:"misfirePolicy" gorm:"misfire_policy"`
	Concurrent     int       `json:"concurrent" gorm:"concurrent"`
	Status         string    `json:"status" gorm:"status"`
	CreateBy       string    `json:"createBy" gorm:"create_by"`
	CreateTime     time.Time `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy       string    `json:"updateBy" gorm:"update_by"`
	UpdateTime     time.Time `json:"updateTime" gorm:"column:update_time;type:datetime"`
	Remark         string    `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (SysJob) TableName() string {
	return "sys_job"
}

type SysJobLog struct {
	JobLogId      int       `json:"jobLogId" gorm:"column:job_log_id;primaryKey"` //表示主键
	JobName       string    `json:"jobName" gorm:"job_name"`
	JobGroup      string    `json:"jobGroup" gorm:"job_group"`
	InvokeTarget  string    `json:"invokeTarget" gorm:"invoke_target"`
	JobMessage    string    `json:"jobMessage" gorm:"job_message"`
	ExceptionInfo string    `json:"exceptionInfo"`
	StartTime     string    `json:"startTime"`
	StopTime      string    `json:"stopTime"`
	Status        string    `json:"status" gorm:"status"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time;type:datetime"`
}

// TableName 指定数据库表名称
func (SysJobLog) TableName() string {
	return "sys_job_log"
}

func SelectJobList(params tools.SearchTableDataParam, isPage bool) ([]SysJob, int64) {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysJob := params.Other.(SysJob)

	var jobName = sysJob.JobName
	var jobGroup = sysJob.JobGroup
	var status = sysJob.Status
	var invokeTarget = sysJob.InvokeTarget

	var total int64
	db := mysql.MysqlDb().Model(SysJob{})

	if jobName != "" {
		db.Where("job_name like ?", "%"+jobName+"%")
	}
	if jobGroup != "" {
		db.Where("job_group = ?", jobGroup)
	}
	if status != "" {
		db.Where("status = ?", status)
	}
	if invokeTarget != "" {
		db.Where("invoke_target like concat('%', ?, '%')", invokeTarget)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0
	}

	var list []SysJob
	if isPage {
		offset := (pageNum - 1) * pageSize
		db.Order("job_id DESC").Offset(offset).Limit(pageSize).Find(&list)
	} else {
		db.Order("job_id DESC").Find(&list)
	}
	return list, total
}

func FindJobById(jobId string) SysJob {
	var job SysJob
	err := mysql.MysqlDb().Where("job_id = ?", jobId).First(&job).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return job
}

func SaveJob(jobParam SysJobParam, userId interface{}) R.Result {
	var job SysJob
	err := copier.Copy(&job, jobParam)
	job.CreateTime = time.Now()
	user := system.FindUserById(userId)
	job.CreateBy = user.UserName
	err = mysql.MysqlDb().Model(&SysJob{}).Create(&job).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess("操作成功")
}

func UploadJob(jobParam SysJob, userId interface{}) R.Result {
	var job SysJob
	err := copier.Copy(&job, jobParam)
	user := system.FindUserById(userId)
	job.UpdateBy = user.UserName
	job.UpdateTime = time.Now()
	err = mysql.MysqlDb().Model(&SysJob{}).Create(&jobParam).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess("操作成功")
}

func SelectJobLogList(params tools.SearchTableDataParam) ([]SysJobLog, int64) {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysJobLog := params.Other.(SysJobLog)
	var jobName = sysJobLog.JobName
	var jobGroup = sysJobLog.JobGroup
	var status = sysJobLog.Status

	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	var total int64
	db := mysql.MysqlDb().Model(SysJobLog{})

	if jobName != "" {
		db.Where("job_name like ?", "%"+jobName+"%")
	}
	if jobGroup != "" {
		db.Where("job_group = ?", jobGroup)
	}

	if status != "" {
		db.Where("status = ?", status)
	}
	if beginTime != "" {
		//Loc, _ := time.LoadLocation("Asia/Shanghai")
		//startTime1, _ := time.ParseInLocation("", beginTime, Loc)
		//endTime = endTime + " 23:59:59"
		//endTime1, _ := time.ParseInLocation(constants.TimeFormat, endTime, Loc)
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		//db.Where("create_time >= ? and create_time <= ?", startTime1, endTime1)
		db.Where("create_time >= ?", startTime1)
		db.Where("create_time <= ?", endTime1)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0
	}
	offset := (pageNum - 1) * pageSize
	var list []SysJobLog
	db.Order("job_log_id DESC").Offset(offset).Limit(pageSize).Find(&list)
	return list, total
}

func FindJobLogById(id string) R.Result {
	var jobLog SysJobLog
	err := mysql.MysqlDb().Where("id = ?", id).First(&jobLog).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess(jobLog)
}

func (param *SysJobLog) JobLogAdd() R.Result {
	param.CreateTime = time.Now()
	if err := mysql.MysqlDb().Model(&SysJobLog{}).Create(&param).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DetectJob(jobIds string) {
	if err := mysql.MysqlDb().Where("job_id in ( ? )", jobIds).Delete(&SysJob{}).Error; err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

func ChangeStatus(jobId string, status string) {
	err := mysql.MysqlDb().Model(&SysJob{}).Where("job_id", jobId).Updates("status=" + status).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

func DetectJobLog(ids string) {
	err := mysql.MysqlDb().Where("id in ( ? )", utils2.Split(ids)).Delete(&SysJobLog{}).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

func ClearJobLog() R.Result {
	if err := mysql.MysqlDb().Exec("truncate table sys_job_log").Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}
