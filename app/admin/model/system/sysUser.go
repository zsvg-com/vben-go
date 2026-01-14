package system

import (
	"strconv"
	"time"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysUser model：数据库字段
type SysUser struct {
	UserId      int       `json:"userId" gorm:"column:user_id;primaryKey"` //表示主键
	DeptId      int       `json:"deptId" gorm:"dept_id"`
	UserName    string    `json:"userName" gorm:"user_name"`
	NickName    string    `json:"nickName" gorm:"nick_name"`
	UserType    string    `json:"userType" gorm:"user_type"`
	Email       string    `json:"email" gorm:"email"`
	Phonenumber string    `json:"phonenumber" gorm:"phonenumber"`
	Sex         string    `json:"sex" gorm:"sex"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Password    string    `json:"password" gorm:"password"`
	Status      string    `json:"status" gorm:"status"`
	DelFlag     string    `json:"delFlag" gorm:"del_flag"`
	LoginIp     string    `json:"loginIp" gorm:"login_ip"`
	LoginDate   time.Time `json:"loginDate" gorm:"column:login_date;type:datetime"`
	CreateBy    string    `json:"createBy" gorm:"create_by"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy    string    `json:"updateBy" gorm:"update_by"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark      string    `json:"remark" gorm:"remark"`
}

type SysUserParm struct {
	UserId      int       `json:"userId" gorm:"column:user_id;primaryKey"` //表示主键
	DeptId      int       `json:"deptId" gorm:"dept_id"`
	UserName    string    `json:"userName" gorm:"user_name"`
	NickName    string    `json:"nickName" gorm:"nick_name"`
	UserType    string    `json:"userType" gorm:"user_type"`
	Email       string    `json:"email" gorm:"email"`
	Phonenumber string    `json:"phonenumber" gorm:"phonenumber"`
	Sex         string    `json:"sex" gorm:"sex"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Password    string    `json:"password" gorm:"password"`
	Status      string    `json:"status" gorm:"status"`
	DelFlag     string    `json:"delFlag" gorm:"del_flag"`
	LoginIp     string    `json:"loginIp" gorm:"login_ip"`
	LoginDate   string    `json:"loginDate" gorm:"column:login_date;type:datetime"`
	Remark      string    `json:"remark" gorm:"remark"`
	CreateBy    string    `json:"createBy" gorm:"create_by"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy    string    `json:"updateBy" gorm:"update_by"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time;type:datetime"`
	DeptName    string    `json:"deptName,omitempty"`
	Leader      string    `json:"leader,omitempty"`
	PostIds     []int     `json:"postIds,omitempty"`
	RoleIds     []int     `json:"roleIds,omitempty"`
	RoleId      int       `json:"roleId,omitempty"`
}

type SysUserExcel struct {
	UserId      int       `json:"userId" gorm:"column:user_id;primaryKey"` //表示主键
	DeptId      int       `json:"deptId" gorm:"dept_id"`
	UserName    string    `json:"userName" gorm:"user_name"`
	NickName    string    `json:"nickName" gorm:"nick_name"`
	UserType    string    `json:"userType" gorm:"user_type"`
	Email       string    `json:"email" gorm:"email"`
	Phonenumber string    `json:"phonenumber" gorm:"phonenumber"`
	Sex         string    `json:"sex" gorm:"sex"`
	Avatar      string    `json:"avatar" gorm:"avatar"`
	Password    string    `json:"password" gorm:"password"`
	Status      string    `json:"status" gorm:"status"`
	DelFlag     string    `json:"delFlag" gorm:"del_flag"`
	LoginIp     string    `json:"loginIp" gorm:"login_ip"`
	LoginDate   time.Time `json:"loginDate" gorm:"column:login_date;type:datetime"`
	CreateBy    string    `json:"createBy" gorm:"create_by"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time;type:datetime"`
	UpdateBy    string    `json:"updateBy" gorm:"update_by"`
	UpdateTime  time.Time `json:"updateTime" gorm:"column:update_time;type:datetime"`
	Remark      string    `json:"remark" gorm:"remark"`
	DeptName    string    `json:"deptName,omitempty"`
	Leader      string    `json:"leader,omitempty"`
}

// LoginParam 登录参数
type LoginParam struct {
	Code     string `json:"code"`
	Password string `json:"password"`
	UserName string `json:"username"`
	Uuid     string `json:"uuid"`
}

// Userparam 请求参数
type Userparam struct {
	NickName    string `json:"nickName"`
	UserName    string `json:"userName"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
	Sex         string `json:"sex"`
}

// TableName 指定数据库表名称
func (SysUser) TableName() string {
	return "sys_user"
}

// IsAdmin 指定数据库表名称
func IsAdmin(user *SysUser) bool {
	return user.UserId == 1
}

// IsAdmin 指定数据库表名称
func IsAdminById(userId int) bool {
	return userId == 1
}

// IsExistUser 判断用户是否存在 返回bool类型
func IsExistUser(loginName string) bool {
	var user SysUser
	err := mysql.MysqlDb().Where("del_flag = 0 and user_name = ?", loginName).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func IsExistUserByPhoneNumber(phonenumber string) bool {
	var user SysUser
	err := mysql.MysqlDb().Where("del_flag = 0 and phonenumber = ?", phonenumber).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func IsExistUserByEmail(email string) bool {
	var user SysUser
	err := mysql.MysqlDb().Where("del_flag = 0 and email = ?", email).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

// 登录用到
func FindUserByName(loginName string) SysUser {
	var user SysUser
	err := mysql.MysqlDb().Where("del_flag != 2 and user_name = ?", loginName).Find(&user).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return user
}

func FindUserById(id any) SysUser {
	userId := utils2.GetInterfaceToInt(id)
	if userId < 1 {
		panic(R.ReturnFailMsg("获取用户信息失败"))
	}
	var user SysUser
	var db = mysql.MysqlDb().Model(&user).
		Select("`sys_user`.*, d.dept_name, d.leader").
		Joins("left join sys_dept d on d.dept_id = `sys_user`.dept_id")
	db.Where("`sys_user`.del_flag = 0")
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	user.Password = ""
	return user
}

// 查找用户密码
func FindUserPassWordById(id any) string {
	userId := utils2.GetInterfaceToInt(id)
	if userId < 1 {
		panic(R.ReturnFailMsg("获取用户信息失败"))
	}
	var user SysUser
	var db = mysql.MysqlDb().Model(&user).
		Select("`sys_user`.*")
	db.Where("`sys_user`.del_flag = 0")
	err := db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return user.Password
}

// 更新用户密码
func UploadUserPassWord(id any, password string) {
	userId := utils2.GetInterfaceToInt(id)
	err := mysql.MysqlDb().Exec("update sys_user set password = ? where user_id = ?", password, userId).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
}

/*校验用户是否有数据权限*/
func CheckUserDataScope(userId int, useridP int) {
	if !IsAdminById(userId) {
		var param = tools.SearchTableDataParam{
			Other: SysUser{
				UserId: useridP,
			},
		}
		var row = SelectUserList(param, false)
		if row.Rows == nil {
			panic(R.ReturnFailMsg("没有权限访问用户数据！"))
		}
	}
}

func SelectUserList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysUser := params.Other.(SysUser)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysUser

	var db = mysql.MysqlDb().Model(&rows).
		Select("`sys_user`.*, d.dept_name, d.leader").
		Joins("left join sys_dept d on d.dept_id = `sys_user`.dept_id")

	db.Where("`sys_user`.del_flag = ?", "0")

	var userId = sysUser.UserId
	if userId != 0 {
		db.Where("`sys_user`.user_id = ?", userId)
	}
	var deptId = sysUser.DeptId
	if deptId != 0 {
		db.Where("`sys_user`.dept_id = ? OR `sys_user`.dept_id IN (SELECT t.dept_id FROM sys_dept t WHERE find_in_set(?, ancestors))", deptId, deptId)
	}

	var userName = sysUser.UserName
	if userName != "" {
		db.Where("`sys_user`.user_name like concat('%', ?, '%')", userName)
	}

	var status = sysUser.Status
	if status != "" {
		db.Where("`sys_user`.status = ? ", status)
	}

	var phonenumber = sysUser.Phonenumber
	if phonenumber != "" {
		db.Where("`sys_user`.phonenumber = ? ", phonenumber)
	}
	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		db.Where("`sys_user`.create_time >= ?", startTime1)
		db.Where("`sys_user`.create_time <= ?", endTime1)
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

func SelectUserParmList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysUser := params.Other.(SysUser)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysUserExcel

	var db = mysql.MysqlDb().Model(&rows).
		Select("`sys_user`.*, d.dept_name, d.leader").
		Joins("left join sys_dept d on d.dept_id = `sys_user`.dept_id")

	db.Where("`sys_user`.del_flag = ?", "0")

	var userId = sysUser.UserId
	if userId != 0 {
		db.Where("`sys_user`.user_id = ?", userId)
	}
	var deptId = sysUser.DeptId
	if deptId != 0 {
		db.Where("`sys_user`.dept_id = ? OR `sys_user`.dept_id IN (SELECT t.dept_id FROM sys_dept t WHERE find_in_set(?, ancestors))", deptId, deptId)
	}

	var userName = sysUser.UserName
	if userName != "" {
		db.Where("`sys_user`.user_name like concat('%', ?, '%')", userName)
	}

	var status = sysUser.Status
	if status != "" {
		db.Where("`sys_user`.status = ? ", status)
	}

	var phonenumber = sysUser.Phonenumber
	if phonenumber != "" {
		db.Where("`sys_user`.phonenumber = ? ", phonenumber)
	}
	var beginTime = params.Params.BeginTime
	var endTime = params.Params.EndTime

	if beginTime != "" {
		startTime1, endTime1 := utils2.GetBeginAndEndTime(beginTime, endTime)
		db.Where("`sys_user`.create_time >= ?", startTime1)
		db.Where("`sys_user`.create_time <= ?", endTime1)
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

func GetUserDeptTree() []SysDeptDto {

	var param = tools.SearchTableDataParam{
		Other: SysDept{},
	}

	var result, total = GetDeptList(param, false)
	var data []SysDeptDto
	for i := 0; i < int(total); i++ {
		var bean = result[i]
		if bean.ParentId == 0 {
			dept := SysDeptDto{
				Id:    bean.DeptId,
				Label: bean.DeptName,
			}
			data = append(data, getDeptChildren(result, dept))
		}
	}

	return data
}

func getDeptChildren(list []SysDeptResult, dept SysDeptDto) SysDeptDto {
	var data []SysDeptDto
	for i := 0; i < len(list); i++ {
		var bean = list[i]
		if bean.ParentId == dept.Id {
			dept := SysDeptDto{
				Id:    bean.DeptId,
				Label: bean.DeptName,
			}
			data = append(data, getDeptChildren(list, dept))
		}
	}
	dept.Children = data
	return dept
}

func ImportUserData(users []SysUserParm, updateSupport string) (string, string) {
	var errList []string
	var error string
	for i := 0; i < len(users); i++ {
		var u1 = users[i]
		u1.CreateTime = time.Now()
		var userName = u1.UserName
		var u2 = FindUserByName(userName)

		password := SelectConfigByKey("sys.user.initPassword")
		var pwd = ""
		if password != "" {
			pwd, _ = utils2.PasswordHash(password)
		} else {
			pwd, _ = utils2.PasswordHash("123456")
		}

		u1.Password = pwd
		var d = strconv.Itoa(i + 1)
		if u2.UserId == 0 {
			err := SaveUser(u1)
			if err != "" {
				error = "error"
				errList = append(errList, d+"、用户名："+userName+"，添加失败<br/>")
			} else {
				errList = append(errList, d+"、用户名："+userName+"，添加成功<br/>")
			}
		} else {
			if updateSupport == "true" {
				userIds := make([]int, 0)
				userIds = append(userIds, u2.UserId)
				DeleteUser(userIds)
				err := SaveUser(u1)
				if err != "" {
					error = "error"
					errList = append(errList, d+"、用户名："+userName+"，添加失败<br/>")
				} else {
					errList = append(errList, d+"、用户名："+userName+"，添加成功<br/>")
				}
			} else {
				error = "error"
				errList = append(errList, d+"、用户名："+userName+"，已存在<br/>")
			}
		}
	}
	var result string
	for _, i := range errList {
		result += i
	}
	return error, result
}

func SaveUser(user SysUserParm) string {
	var result = ""
	var u = SysUser{
		NickName:    user.NickName,
		DeptId:      user.DeptId,
		UserName:    user.UserName,
		UserType:    "00",
		Email:       user.Email,
		Phonenumber: user.Phonenumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Password:    user.Password,
		LoginDate:   user.CreateTime,
		Status:      "0",
		DelFlag:     "0",
		CreateBy:    user.CreateBy,
		CreateTime:  user.CreateTime,
		UpdateTime:  user.CreateTime,
		Remark:      user.Remark,
	}
	err1 := mysql.MysqlDb().Model(&SysUser{}).Create(&u).Error
	if err1 != nil {
		result = "" + err1.Error()
	}
	user.UserId = u.UserId
	AddPostByUser(user)
	AddRolesByUser(user)
	return result
}

func UploadUser(userParm SysUserParm) {
	var userIds = []int{userParm.UserId}
	DeleteRolesByUser(userIds)
	AddRolesByUser(userParm)
	DeletePostByUser(userIds)
	AddPostByUser(userParm)
	var user = SysUser{
		UserId:      userParm.UserId,
		DeptId:      userParm.DeptId,
		NickName:    userParm.NickName,
		Phonenumber: userParm.Phonenumber,
		Email:       userParm.Email,
		Sex:         userParm.Sex,
		Status:      userParm.Status,
		Remark:      userParm.Remark,
	}
	err := mysql.MysqlDb().Updates(&user).Error
	if err != nil {
		panic(R.ReturnFailMsg("更新信息失败"))
	}
}

func EditProfileUserInfo(userId int, param Userparam) R.Result {
	var user SysUser
	err1 := mysql.MysqlDb().Where("user_id = ?", userId).First(&user)

	if err1.Error != nil {
		return R.ReturnFailMsg("未找到用户")
	}

	data := make(map[string]interface{})

	if param.NickName != "" {
		data["nick_name"] = param.NickName
	}

	if param.Email != "" {
		data["email"] = param.Email
	}

	if param.Phonenumber != "" {
		data["phonenumber"] = param.Phonenumber
	}

	if param.Sex != "" {
		data["sex"] = param.Sex
	}
	//等价于: UPDATE `foods` SET `price` = '35', `stock` = '0'  WHERE (user_id = '2')
	err2 := mysql.MysqlDb().Model(&SysUser{}).Where("user_id = ?", userId).Updates(data)

	if err2.Error != nil {
		return R.ReturnFailMsg("修改失败")
	}

	return R.ReturnSuccess("修改成功")
}

func DeleteUser(userId []int) R.Result {
	for _, v := range userId {
		if v == constants.AdminId {
			return R.ReturnFailMsg("管理员数据不开操作")
		}
	}
	// 删除部门管理
	DeletePostByUser(userId)
	// 删除角色管理
	DeleteRolesByUser(userId)
	// 删除用户
	DeleteUserById(userId)
	return R.ReturnSuccess("操作成功")
}

func DeleteUserById(userId []int) R.Result {
	err := mysql.MysqlDb().Exec("update sys_user set del_flag = '2' where user_id in (?) ", userId).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return R.ReturnSuccess("操作成功")
}

func ChangeUserStatus(param SysUserParm) R.Result {
	data := make(map[string]interface{})

	if param.Status != "" {
		data["status"] = param.Status
	}

	//等价于: UPDATE `foods` SET `price` = '35', `stock` = '0'  WHERE (user_id = '2')
	err2 := mysql.MysqlDb().Model(&SysUser{}).Where("user_id = ?", param.UserId).Updates(data)
	if err2.Error != nil {
		return R.ReturnFailMsg("修改失败")
	}

	return R.ReturnSuccess("修改成功")
}

func ResetPwd(param SysUserParm) R.Result {
	data := make(map[string]interface{})

	if param.Password != "" {
		var pwd, _ = utils2.PasswordHash(param.Password)
		data["password"] = pwd
	}

	//等价于: UPDATE `foods` SET `price` = '35', `stock` = '0'  WHERE (user_id = '2')
	err2 := mysql.MysqlDb().Model(&SysUser{}).Where("user_id = ?", param.UserId).Updates(data)
	if err2.Error != nil {
		return R.ReturnFailMsg("修改失败")
	}

	return R.ReturnSuccess("修改成功")
}
