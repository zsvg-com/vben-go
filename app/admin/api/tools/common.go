package tools

import (
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"vben/app/core/utils/R"
	"vben/common/config"

	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/arrays"
)

/*下载*/
func GetDownload(context *gin.Context) {
	fileName := context.Param("fileName")
	if fileName == "" {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "下载失败",
			"code": http.StatusInternalServerError,
		})
		return
	}
	deleteStr := context.Param("delete")
	var isDelete = false
	if deleteStr != "" {
		isDeletes, err := strconv.ParseBool(deleteStr)
		isDelete = isDeletes
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "下载失败",
				"code": http.StatusInternalServerError,
			})
			return
		}
	}

	filePath := "./static/images/" + fileName
	//打开文件
	fileTmp, errByOpenFile := os.Open(filePath)
	if isDelete {
		//删除文件
		cuowu := os.Remove(filePath)
		if cuowu != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "下载失败",
				"code": http.StatusInternalServerError,
			})
			return
		}
	}
	defer fileTmp.Close()
	if errByOpenFile != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "下载失败",
			"code": http.StatusInternalServerError,
		})
		return
	}

	context.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	context.Header("Content-Disposition", "inline;filename="+fileName)
	context.Header("Content-Transfer-Encoding", "binary")
	context.Header("Cache-Control", "no-cache")

	context.File(filePath)
	return
}

/*单个上传*/
func UploadCommon(context *gin.Context) {
	file, errLoad := context.FormFile("file")
	if errLoad != nil {
		msg := "获取上传文件错误:" + errLoad.Error()
		context.JSON(http.StatusOK, R.ReturnFailMsg(msg))
		return
	}
	//上传图片
	errFile := context.SaveUploadedFile(file, config.FileProfile+file.Filename)
	if errFile != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("上传图片异常，请联系管理员"))
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":              "操作成功",
		"code":             http.StatusOK,
		"url":              config.ShowFileProfile + file.Filename,
		"fileName":         file.Filename,
		"newFileName":      file.Filename,
		"originalFilename": file.Filename,
	})
}

/*批量上传*/
func UploadCommons(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["files"]

	var urls []string
	var fileNames []string
	var newFileNames []string
	var originalFilenames []string
	for _, file := range files {
		context.SaveUploadedFile(file, config.FileProfile+file.Filename)
		urls = append(urls, config.ShowFileProfile+file.Filename)
		fileNames = append(fileNames, file.Filename)
		newFileNames = append(newFileNames, file.Filename)
		originalFilenames = append(originalFilenames, file.Filename)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":         "操作成功",
		"code":        http.StatusOK,
		"url":         strings.Join(urls, ","),
		"fileName":    strings.Join(fileNames, ","),
		"newFileName": strings.Join(newFileNames, ","),
		"strings":     strings.Join(originalFilenames, ","),
	})
}

/*下载*/
func UploadRsource(context *gin.Context) {
	resource := context.Param("resource")
	if strings.Contains(resource, "..") {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "资源文件(" + resource + ")非法，不允许下载。",
			"code": http.StatusInternalServerError,
		})
		return
	}
	var extension = []string{
		"bmp", "gif", "jpg", "jpeg", "png",
		// word excel powerpoint
		"doc", "docx", "xls", "xlsx", "ppt", "pptx", "html", "htm", "txt",
		// 压缩文件
		"rar", "zip", "gz", "bz2",
		// 视频格式
		"mp4", "avi", "rmvb",
		// pdf
		"pdf",
	}
	fileExt := path.Ext(resource)
	if arrays.ContainsString(extension, fileExt) < 0 {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "资源文件(" + resource + ")非法，不允许下载。",
			"code": http.StatusInternalServerError,
		})
		return
	}

	filePath := "./static/images/" + resource
	fileName := "x" + fileExt
	//打开文件
	fileTmp, errByOpenFile := os.Open(filePath)

	defer fileTmp.Close()
	if errByOpenFile != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  "下载失败",
			"code": http.StatusInternalServerError,
		})
		return
	}

	context.Header("Content-Type", "application/octet-stream")
	//强制浏览器下载
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	//浏览器下载或预览
	context.Header("Content-Disposition", "inline;filename="+fileName)
	context.Header("Content-Transfer-Encoding", "binary")
	context.Header("Cache-Control", "no-cache")

	context.File(filePath)
	return
}
