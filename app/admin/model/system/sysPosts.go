package system

import (
	"time"
	"vben/app/admin/model/tools"
	"vben/app/core/utils"
	"vben/app/core/utils/R"
	"vben/pkg/mysql"
)

// SysPost model：数据库字段
type SysPost struct {
	PostId     int       `json:"postId" gorm:"column:post_id;primaryKey"` //表示主键
	PostCode   string    `json:"postCode" gorm:"post_code"`
	PostName   string    `json:"postName" gorm:"post_name"`
	PostSort   int       `json:"postSort" gorm:"post_sort"`
	Status     string    `json:"status" gorm:"status"`
	CreateBy   string    `json:"createBy" gorm:"create_by"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;type:datetime;autoCreateTime"`
	UpdateBy   string    `json:"updateBy" gorm:"update_by"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:update_time;type:datetime;autoCreateTime"`
	Remark     string    `json:"remark" gorm:"remark"`
}

// TableName 指定数据库表名称
func (SysPost) TableName() string {
	return "sys_post"
}

func SelectUserPostGroup(userName string) string {
	var posts []SysPost
	var result = ""
	mysql.MysqlDb().Raw("select p.post_id, p.post_name, p.post_code "+
		"from sys_post p "+
		"left join sys_user_post up on up.post_id = p.post_id "+
		"left join sys_user u on u.user_id = up.user_id "+
		"where u.user_name = ?", userName).Scan(&posts)
	if posts != nil {
		for i := range posts {
			sysPost := posts[i]
			if i == 0 {
				result = sysPost.PostName
			} else {
				result += "," + sysPost.PostName
			}
		}
	}
	return result
}

// 分页查询
func SelectSysPostList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {

	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysPost := params.Other.(SysPost)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysPost

	var db = mysql.MysqlDb().Model(&rows)

	var postCode = sysPost.PostCode
	if postCode != "" {
		db.Where("post_code like ?", "%"+postCode+"%")
	}
	var status = sysPost.Status
	if status != "" {
		db.Where("status = ?", status)
	}
	var postName = sysPost.PostName
	if postName != "" {
		db.Where("post_name like ?", "%"+postName+"%")
	}

	db.Order("post_sort")

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

func SelectPostListByUserId(userId int) []int {
	var result []int
	mysql.MysqlDb().Raw("select p.post_id "+
		"from sys_post p "+
		"left join sys_user_post up on up.post_id = p.post_id "+
		"left join sys_user u on u.user_id = up.user_id "+
		"where u.user_id = ? ", userId).Find(&result)
	return result
}

/*
*添加 岗位和用户关联
 */
func AddPostByUser(user SysUserParm) {
	postIds := user.PostIds
	userId := user.UserId
	AddPostListByUser(postIds, userId)
}

func AddPostListByUser(postIds []int, userId int) {
	var posts []SysUserPost
	for i := 0; i < len(postIds); i++ {
		var post = SysUserPost{
			PostId: postIds[i],
			UserId: userId,
		}
		posts = append(posts, post)
	}
	err := mysql.MysqlDb().CreateInBatches(posts, len(posts)).Error
	if err != nil {
		panic(R.ReturnFailMsg("添加 部门管理用户失败"))
	}
}

// 删除关联用户
func DeletePostByUser(userIds []int) {
	err := mysql.MysqlDb().Exec("delete from sys_user_post where user_id in (?) ", userIds).Error
	if err != nil {
		panic(R.ReturnFailMsg("删除部门关联用户失败"))
	}
}

func FindPostInfoById(postId string) SysPost {
	var post SysPost
	err := mysql.MysqlDb().Where("post_id = ?", postId).First(&post).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}
	return post
}

func SavePost(post SysPost) R.Result {
	err := mysql.MysqlDb().Model(&SysPost{}).Create(&post).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func EditPost(sysPost SysPost) R.Result {
	err := mysql.MysqlDb().Updates(&sysPost).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}

func DeletePost(postIds string) R.Result {
	err := mysql.MysqlDb().Where("post_id in (?)", utils.Split(postIds)).Delete(&SysPost{}).Error
	if err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}
