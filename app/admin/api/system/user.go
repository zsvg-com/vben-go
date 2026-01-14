package system

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	useragent "github.com/wenlng/go-user-agent"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vben/app/admin/model/cache"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/monitor"
	"vben/app/admin/model/system"
	"vben/app/admin/model/tools"
	utils2 "vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/app/core/utils/jwt"
	"vben/config"
	"vben/pkg/cache/redisCache"
	"vben/pkg/mysql"
)

func LoginHandler(context *gin.Context) {
	var param system.LoginParam
	if err := context.ShouldBind(&param); err != nil {
		monitor.LoginInfoAdd(context, param, "登录失败，参数为空", false)
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	if !loginBeforeCheck(param, context) {
		return
	}
	isExist := system.IsExistUser(param.UserName)
	if !isExist {
		monitor.LoginInfoAdd(context, param, "登录失败，用户不存在", false)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "用户不存在",
			"code": http.StatusInternalServerError,
		})
		return
	}
	findUser(param, context)
}
func loginBeforeCheck(param system.LoginParam, context *gin.Context) bool {
	// 验证码验证
	var captchaEnabled = SelectCaptchaEnabled()
	if !captchaEnabled {
		return false
	}
	isVerify := utils2.VerifyCaptcha(param.Uuid, param.Code)
	monitor.LoginInfoAdd(context, param, "登录失败，验证码错误", false)
	if !isVerify {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "请输入正确的验证码",
			"code": http.StatusInternalServerError,
		})
	}
	return isVerify
}

// 判断用户是否存在 返回bool类型
func findUser(param system.LoginParam, context *gin.Context) {
	var loginName = param.UserName
	var pass = param.Password
	var user = system.FindUserByName(loginName)
	if user.UserId <= 0 {
		monitor.LoginInfoAdd(context, param, "登录失败，用户不存在", false)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "用户不存在",
			"code": http.StatusInternalServerError,
		})
		return
	}
	if user.Status == "1" {
		monitor.LoginInfoAdd(context, param, "登录失败，账号已停用", false)
		context.JSON(http.StatusOK, R.ReturnFailMsg("账号已停用"))
		return
	}
	// 验证密码错误次数
	count := cache.GetPasswordTryCount(param.UserName)
	if count >= config.UserPassword.MaxRetryCount {
		monitor.LoginInfoAdd(context, param, "登录失败，用户密码错误最大次数", false)
		context.JSON(http.StatusOK, R.ReturnFailMsg("登录失败，用户密码错误最大次数，请过稍后再试"))
		return
	}
	// 验证 密码是否正确
	if !utils2.PasswordVerify(pass, user.Password) {
		count++
		cache.SetPasswordTryCount(param.UserName, count)
		monitor.LoginInfoAdd(context, param, "登录失败，密码错误", false)
		context.JSON(http.StatusOK, gin.H{
			"msg":  "登录失败，密码错误",
			"code": http.StatusInternalServerError,
		})
		return
	} else {
		cache.DeletePasswordTryCount(param.UserName)
	}
	newUUID, _ := uuid.NewUUID()
	tokenId := newUUID.String()
	dept := system.GetDeptInfo(strconv.Itoa(user.DeptId))
	userAgent := context.Request.Header.Get("User-Agent")
	os := useragent.GetOsName(userAgent)
	browser := useragent.GetBrowserName(userAgent)
	tokenString, err := jwt.CreateToken(user.UserName, user.UserId, user.DeptId, newUUID.String())
	if err != nil {
		monitor.LoginInfoAdd(context, param, "登录失败，"+err.Error(), false)
		context.JSON(http.StatusOK, R.ReturnFailMsg("登录失败"))
		return
	}
	//写入缓存
	loginUser, _ := json.Marshal(&monitor.LoginUserCache{
		TokenId:       tokenId,
		DeptName:      dept.DeptName,
		Ipaddr:        "" + context.ClientIP(),
		LoginLocation: "" + utils2.GetRealAddressByIP(context.ClientIP()),
		UserId:        user.UserId,
		Browser:       browser,
		Os:            os,
		LoginTime:     time.Now().Unix(),
		UserName:      user.UserName,
		Uuid:          tokenId,
		DeptId:        user.DeptId,
	})
	err = redisCache.NewRedisCache().Put(constants.LoginCacheKey+newUUID.String(), string(loginUser), time.Duration(config.Jwt.JwtTtl)*time.Second)
	if err != nil {
		monitor.LoginInfoAdd(context, param, "登录失败，"+err.Error(), false)
		context.JSON(http.StatusOK, R.ReturnFailMsg("登录失败"))
		return
	}
	// 登录日志
	monitor.LoginInfoAdd(context, param, "登录成功", true)
	context.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"code":  http.StatusOK,
		"token": tokenString,
	})
}

func GetInfoHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	var user = system.FindUserById(userId)
	roles := system.GetRolePermission(user)
	permissions := system.GetMenuPermission(user)
	dept := system.GetDeptInfo(strconv.Itoa(user.DeptId))
	context.JSON(http.StatusOK, gin.H{
		"msg":  "获取成功",
		"code": http.StatusOK,
		"user": gin.H{
			"userName":    user.UserName,
			"nickName":    user.NickName,
			"phonenumber": user.Phonenumber,
			"email":       user.Email,
			"avatar":      user.Avatar,
			"sex":         user.Sex,
			"createTime":  user.CreateTime,
			"dept":        dept,
		},
		"roles":       roles,
		"permissions": permissions,
	})
}
func LogoutHandler(context *gin.Context) {
	// 开始删除缓存
	jwtUuid, err := jwt.GetJwtUuid(context)
	if err != nil || len(jwtUuid) <= 0 {
		context.JSON(http.StatusOK, R.ReturnSuccessMsg("退出成功"))
		return
	}
	redisCache.NewRedisCache().Del(constants.LoginCacheKey + jwtUuid)
	context.JSON(http.StatusOK, R.ReturnSuccessMsg("退出成功"))
}

// CaptchaImageHandler 验证码 输出
func CaptchaImageHandler(context *gin.Context) {
	var captchaEnabled = SelectCaptchaEnabled()
	if captchaEnabled {
		id, b64s, err := utils2.CreateImageCaptcha()
		if err != nil {
			context.JSON(http.StatusOK, R.ReturnFailMsg("创建验证码失败，请联系管理员"))
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg":            "操作成功",
			"img":            strings.ReplaceAll(b64s, "data:image/png;base64,", ""),
			"code":           http.StatusOK,
			"captchaEnabled": captchaEnabled,
			"uuid":           id,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg":            "操作成功",
			"code":           http.StatusOK,
			"captchaEnabled": captchaEnabled,
		})
	}
}

// UpdatePwdHandler 修改密码
func UpdatePwdHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	var Password = system.FindUserPassWordById(userId)
	// 没有这个，下面的为空很奇怪
	context.DefaultPostForm("newPassword", "")
	from := context.Request.Form
	OldPassword := from.Get("oldPassword")
	NewPassword := from.Get("newPassword")

	if OldPassword == "" {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	if NewPassword == "" {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}

	// 验证旧密码
	if utils2.PasswordVerify(OldPassword, Password) {
		// 验证新密码
		if utils2.PasswordVerify(NewPassword, Password) {
			context.JSON(http.StatusOK, R.ReturnFailMsg("新密码不能与旧密码相同"))
			return
		}
		// 加密
		passString, err3 := utils2.PasswordHash(NewPassword)
		if err3 != nil {
			context.JSON(http.StatusOK, R.ReturnFailMsg("加密失败"))
			return
		}
		// 更新 密码
		system.UploadUserPassWord(userId, passString)
		context.JSON(http.StatusOK, R.ReturnSuccess("修改密码成功"))
	} else {
		context.JSON(http.StatusOK, R.ReturnFailMsg("旧密码错误"))
		return
	}
}

// ProfileHandler 查询个人信息
func ProfileHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	var user system.SysUser
	err1 := mysql.MysqlDb().Where("user_id = ?", userId).First(&user)

	if err1.Error != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("未找到用户"))
		return
	}
	user.Password = ""
	context.JSON(http.StatusOK, gin.H{
		"msg":  "获取成功",
		"code": http.StatusOK,
		"data": gin.H{
			"userName":    user.UserName,
			"nickName":    user.NickName,
			"phonenumber": user.Phonenumber,
			"email":       user.Email,
			"sex":         user.Sex,
			"createTime":  user.CreateTime.Format(constants.TimeFormat),
		},
		"roleGroup": system.SelectRolesByUserName(user.UserName), // 目前暂时没有用
		"postGroup": system.SelectUserPostGroup(user.UserName),   // 目前暂时没有用
	})
}

// PostProfileHandler 修改个人信息
func PostProfileHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	param := system.Userparam{}
	if err := context.ShouldBindJSON(&param); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var result = system.EditProfileUserInfo(utils2.GetInterfaceToInt(userId), param)
	context.JSON(http.StatusOK, result)
}

// AvatarHandler 上传头像 并更新
func AvatarHandler(context *gin.Context) {
	userId, _ := context.Get("userId")
	var user system.SysUser
	err1 := mysql.MysqlDb().Where("user_id = ?", userId).First(&user)

	if err1.Error != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("未找到用户"))
		return
	}

	file, errLoad := context.FormFile("avatarfile")
	if errLoad != nil {
		msg := "获取上传文件错误:" + errLoad.Error()
		context.JSON(http.StatusOK, R.ReturnFailMsg(msg))
		return
	}

	//fileExt := strings.ToLower(path.Ext(file.Filename))
	//if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
	//	})
	//	return
	//}

	//上传图片
	errFile := context.SaveUploadedFile(file, "./static/images/"+file.Filename)
	if errFile != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("上传图片异常，请联系管理员"))
		return
	}

	// 更新 头像
	err2 := mysql.MysqlDb().Model(&user).Update("avatar", "/profile/"+file.Filename)
	if err2.Error != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("上传图片异常，请联系管理员"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"msg":    "上传头像成功!",
		"imgUrl": "/profile/" + file.Filename,
	})
}

/*-----------用户管理----------------------------*/

func ListUser(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	deptId, _ := strconv.Atoi(context.DefaultQuery("deptId", "0"))
	var userName = context.DefaultQuery("userName", "")
	var status = context.DefaultQuery("status", "")

	var phonenumber = context.DefaultQuery("phonenumber", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysUser{
			DeptId:      deptId,
			UserName:    userName,
			Phonenumber: phonenumber,
			Status:      status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	context.JSON(http.StatusOK, system.SelectUserList(param, true))
}

func ExportExport(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	deptId, _ := strconv.Atoi(context.DefaultQuery("deptId", "0"))
	var userName = context.DefaultQuery("userName", "")
	var status = context.DefaultQuery("status", "")

	var phonenumber = context.DefaultQuery("phonenumber", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: system.SysUser{
			DeptId:      deptId,
			UserName:    userName,
			Phonenumber: phonenumber,
			Status:      status,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	var data1 = system.SelectUserParmList(param, true)
	var list = data1.Rows.([]system.SysUserExcel)

	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "userId",
		"title":  "用户序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "deptId",
		"title":  "部门编号",
		"width":  "15",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "userName",
		"title":  "登录名称",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "nickName",
		"title":  "用户名称",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "email",
		"title":  "用户邮箱",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "phonenumber",
		"title":  "手机号码",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "sex",
		"title":  "用户性别",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "帐号状态",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "loginIp",
		"title":  "最后登录IP",
		"width":  "30",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "loginDate",
		"title":  "最后登录时间",
		"width":  "60",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "deptName",
		"title":  "部门名称",
		"width":  "50",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "leader",
		"title":  "部门负责人",
		"width":  "30",
		"is_num": "0",
	})

	//填充数据
	data := make([]map[string]interface{}, 0)
	if len(list) > 0 {
		for _, v := range list {
			var sexStatus = v.Sex
			var sex = ""
			if sexStatus == "0" {
				sex = "男"
			} else if sexStatus == "1" {
				sex = "女"
			} else {
				sex = "未知"
			}
			var status = v.Status
			var statusStr = ""
			if status == "0" {
				statusStr = "正常"
			} else {
				statusStr = "停用"
			}
			//timeObj, _ := time.Parse(time.RFC3339, v.LoginDate)
			var loginData = v.LoginDate.Format(constants.TimeFormat)
			data = append(data, map[string]interface{}{
				"userId":      v.UserId,
				"deptId":      v.DeptId,
				"userName":    v.UserName,
				"nickName":    v.NickName,
				"email":       v.Email,
				"phonenumber": v.Phonenumber,
				"sex":         sex,
				"status":      statusStr,
				"loginIp":     v.LoginIp,
				"loginDate":   loginData,
				"deptName":    v.DeptName,
				"leader":      v.Leader,
			})
		}
	}
	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func ImportUserData(context *gin.Context) {
	file, _, errLoad := context.Request.FormFile("file")
	if errLoad != nil {
		msg := "获取上传文件错误:" + errLoad.Error()
		context.JSON(http.StatusOK, R.ReturnFailMsg(msg))
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		context.JSON(500, "请选择文件")
		return
	}
	var users []system.SysUserParm

	var updateSupport = context.DefaultQuery("updateSupport", "")

	rows, _ := xlsx.GetRows("Sheet1")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			users = append(users, system.SysUserParm{
				UserName:    data[2],
				NickName:    data[3],
				Phonenumber: data[5],
				Sex:         data[6],
				Email:       data[4],
				Status:      data[7],
				CreateTime:  time.Now(),
				DeptId:      utils2.GetInterfaceToInt(data[1]),
				PostIds:     utils2.Split(data[8]),
				RoleIds:     utils2.Split(data[9]),
			})
		}
	}
	if len(users) == 0 {
		context.JSON(http.StatusOK, R.ReturnFailMsg("请在表格中添加数据"))
		return
	}

	var error, message = system.ImportUserData(users, updateSupport)

	if error == "" {
		context.JSON(http.StatusOK, R.ReturnSuccess(message))
	} else {
		context.JSON(http.StatusOK, R.ReturnFailMsg(message))
	}

}

// 下载模版
func ImportTemplate(context *gin.Context) {
	//定义首行标题
	dataKey := make([]map[string]string, 0)
	dataKey = append(dataKey, map[string]string{
		"key":    "userId",
		"title":  "用户序号",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "deptId",
		"title":  "部门编号",
		"width":  "15",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "userName",
		"title":  "登录名称",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "nickName",
		"title":  "用户名称",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "email",
		"title":  "用户邮箱",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "phonenumber",
		"title":  "手机号码",
		"width":  "20",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "sex",
		"title":  "用户性别",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "帐号状态",
		"width":  "10",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "岗位",
		"width":  "11",
		"is_num": "0",
	})
	dataKey = append(dataKey, map[string]string{
		"key":    "status",
		"title":  "权限",
		"width":  "12",
		"is_num": "0",
	})
	//填充数据
	data := make([]map[string]interface{}, 0)

	ex := tools.NewMyExcel()
	ex.ExportToWeb(dataKey, data, context)
}

func GetUserInfo(context *gin.Context) {
	/*参数用户*/
	userIdStr := context.Param("userId")
	//登录用户
	loginUserId, _ := context.Get("userId")
	userId := utils2.GetInterfaceToInt(loginUserId)

	if userIdStr == "" {
		userIdStr = strconv.Itoa(userId)
	}
	useridP, _ := strconv.Atoi(userIdStr)

	system.CheckUserDataScope(userId, useridP)

	user := system.FindUserById(useridP)

	var roles []system.SysRoles
	var roleIds []int
	// 登录者的权限
	roles = system.SelectRolePermissionByUserId(userId)

	posts := system.SelectSysPostList(tools.SearchTableDataParam{
		Other: system.SysPost{},
	}, false).Rows

	var result = gin.H{
		"msg":   "操作成功",
		"code":  http.StatusOK,
		"data":  user,
		"roles": roles,
		"posts": posts,
	}
	// 判断当期是否为管理员
	if useridP != 0 {
		postIds := system.SelectPostListByUserId(useridP)
		roles2 := system.SelectRolePermissionByUserId(useridP)
		for _, sysRoles := range roles2 {
			roleIds = append(roleIds, sysRoles.RoleId)
		}
		result["postIds"] = postIds
		result["roleIds"] = roleIds
	}

	context.JSON(http.StatusOK, result)
}

func SaveUser(context *gin.Context) {
	userId, _ := context.Get("userId")
	var user system.SysUserParm
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var username = user.UserName
	if system.IsExistUser(username) {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "用户名已存在",
			"code": http.StatusInternalServerError,
		})
		return
	}
	var phonenumber = user.Phonenumber
	if system.IsExistUserByPhoneNumber(phonenumber) {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "手机号已存在",
			"code": http.StatusInternalServerError,
		})
		return
	}
	var email = user.Email
	if system.IsExistUserByEmail(email) {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "邮箱已存在",
			"code": http.StatusInternalServerError,
		})
		return
	}
	var password = user.Password
	var pwd, _ = utils2.PasswordHash(password)
	var user1 = system.FindUserById(userId)
	user.CreateBy = user1.UserName
	user.CreateTime = time.Now()
	user.Password = pwd

	/*用户名、手机号、邮箱不能重复验证*/
	var message = system.SaveUser(user)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": message,
	})
}

func UploadUser(context *gin.Context) {
	userId, _ := context.Get("userId")
	var userParm system.SysUserParm
	if err := context.ShouldBind(&userParm); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var uId = userParm.UserId
	if uId == 0 {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	var user = system.FindUserById(userId)
	var userSql = system.FindUserById(uId)

	var username = userParm.UserName
	if !system.IsExistUser(username) {
		if userSql.UserName != username {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "用户名:" + username + "已存在",
				"code": http.StatusInternalServerError,
			})
			return
		}
	}
	var phonenumber = userParm.Phonenumber
	if !system.IsExistUserByPhoneNumber(phonenumber) {
		if userSql.Phonenumber != phonenumber {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "手机号:" + phonenumber + "已存在",
				"code": http.StatusInternalServerError,
			})
			return
		}
	}
	var email = userParm.Email
	if !system.IsExistUserByEmail(email) {
		if userSql.Email != email {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "邮箱:" + email + "已存在",
				"code": http.StatusInternalServerError,
			})
			return
		}
	}
	userParm.UpdateBy = user.UserName
	userParm.UpdateTime = time.Now()
	system.UploadUser(userParm)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func DeleteUserById(context *gin.Context) {
	var userIds = context.Param("userIds")
	context.JSON(http.StatusOK, system.DeleteUser(utils2.Split(userIds)))
}

// 重设密码
func ResetPwd(context *gin.Context) {
	/*获取为空有可能内部参数错误*/
	var user system.SysUserParm
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	result := system.ResetPwd(user)
	context.JSON(http.StatusOK, result)
}

func ChangeUserStatus(context *gin.Context) {
	var user system.SysUserParm
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
		return
	}
	if user.UserId == constants.AdminId {
		context.JSON(http.StatusOK, R.ReturnFailMsg("管理员数据不开操作"))
		return
	}
	result := system.ChangeUserStatus(user)
	context.JSON(http.StatusOK, result)
}

func GetAuthUserRole(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	uIdStr, _ := context.Get("userId")
	uId := utils2.GetInterfaceToInt(uIdStr)
	// 登录者的权限
	var roles []system.SysRoles
	if system.IsAdminById(uId) {
		roles = system.SelectRolePermissionByUserId(uId)
	} else {
		roles = system.SelectRolePermissionByUserId(userId)
	}

	user := system.FindUserById(userId)

	context.JSON(http.StatusOK, gin.H{
		"msg":   "操作成功",
		"code":  http.StatusOK,
		"user":  user,
		"roles": roles,
	})
}

func PutAuthUser(context *gin.Context) {
	uIdStr, _ := context.GetQuery("userId")
	uId, _ := strconv.Atoi(uIdStr)
	roleIds, _ := context.GetQuery("roleIds")
	uIds := []int{uId}
	system.DeleteRolesByUser(uIds)
	result := system.InsertRolesByUser(uId, utils2.Split(roleIds))
	context.JSON(http.StatusOK, result)
}

// 登录获取菜单
func GetUserDeptTree(context *gin.Context) {
	var list = system.GetUserDeptTree()
	context.JSON(http.StatusOK, R.ReturnSuccess(list))
}
