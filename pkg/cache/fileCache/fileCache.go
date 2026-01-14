package fileCache

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type fileCache struct {
}

type structure struct {
	Data   string `json:"data"`
	Expire int    `json:"expire"`
}

func NewFileCache() *fileCache {

	return &fileCache{}
}

func (f fileCache) Put(key string, value string, ttl time.Duration) error {
	//key的md5作为文件名
	file, err := os.OpenFile("storage/"+fmt.Sprintf("%x", md5.Sum([]byte(key))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		return err
	}

	defer file.Close()

	var v structure

	v.Data = value

	//设置过期时间
	if ttl == 0 {

		v.Expire = 0
	} else {

		v.Expire = int(time.Now().Add(ttl).Unix())
	}

	jsonStr, err := json.Marshal(v)

	if err != nil {

		return err
	}

	_, err = file.Write(jsonStr)

	return err

}

func (f fileCache) Get(key string) (string, error) {
	filePath := "storage/" + fmt.Sprintf("%x", md5.Sum([]byte(key)))

	file, err := os.Open(filePath)

	if err != nil {

		return "", err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {

		return "", err
	}

	var v structure

	err = json.Unmarshal(data, &v)

	if err != nil {

		return "", err
	}

	now := time.Now().Unix()

	//判断过期时间
	if v.Expire != 0 && (now > int64(v.Expire)) {
		//删除过期文件
		defer func(f *os.File, ff string) {
			f.Close()
			os.Remove(ff)

		}(file, filePath)

		return "", errors.New("缓存已超时")
	}

	return v.Data, err

}

func (r fileCache) Del(key string) (string, error) {
	filePath := "storage/" + fmt.Sprintf("%x", md5.Sum([]byte(key)))
	os.Remove(filePath)
	return "", nil
}

func (r fileCache) Clear() (string, error) {
	filePath := "storage/"
	os.RemoveAll(filePath)
	return "", nil
}
