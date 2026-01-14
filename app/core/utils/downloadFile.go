package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/file"
	"io"
	"net/http"
	"os"
)

// 下载文件 到本地，返回文件路径
func DownloadFileForUrl(url, path, name string) (string, error) {

	if !Exists(path) {
		//logger.Log.WithFields(logrus.Fields{"data": path}).Warn("文件夹不存在，创建文件夹")
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return "", err
		}
	}
	if Exists(path + name) {
		//logger.Log.WithFields(logrus.Fields{"data": path + name}).Warn("文件已经存在", path+name)
		return "已经存在", nil
	} else {
		//logger.Log.WithFields(logrus.Fields{"data": path}).Warn("文件不存在，下载文件", path+name)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.Create(path + name)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return path + name, nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 这样可行
func DonwloadFile(context *gin.Context, fileName string, filePath string) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("AAccess-Control-Expose-Headers", "Content-Disposition")
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	context.Header("Content-Type", "application/octet-stream")

	context.File(filePath)
}

// 文件不存在就创建文件
func DirExistAndMake(autoPath string) error {
	if !Exists(autoPath) { // 检查 文件夹是否存在
		if err := os.MkdirAll(autoPath, os.ModePerm); err != nil {
			//logger.Log.WithFields(logrus.Fields{"data": err}).Warn("文件夹不存在，创建文件夹出错")
			return err
		}
	}
	return nil
}

// 文件存在就删除文件
func DirExistAndDel(autoPath string) error {
	if Exists(autoPath) { // 检查文件是否存在
		if file.IsDir(autoPath) { // 检查是否是文件夹
			if err := os.RemoveAll(autoPath); err != nil {
				return err
			}
		} else {
			if err := os.Remove(autoPath); err != nil {
				return err
			}
		}
	}
	return nil
}
