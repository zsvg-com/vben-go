package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}

// GetRemoteClientIp 获取远程客户端IP
func GetRemoteClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	//本地ip
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	return remoteIp
}

type Tunit struct {
	Pro  string `json:"pro"`
	City string `json:"city"`
}

func GetRealAddressByIP(ip string) string {
	url := "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"
	resp, err := http.Get(url)
	var result = "内网ip"
	if err != nil {
		result = "内网ip"
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			result = "内网ip"
		} else {
			dws := new(Tunit)
			json.Unmarshal(body, &dws)
			result = dws.Pro + " " + dws.City
		}
	}
	return result
}

// 数组转 string,
func Join(sep []int, sp string) string {
	sarr := make([]string, len(sep))
	for i, v := range sep {
		sarr[i] = fmt.Sprint(v)
	}
	return strings.Join(sarr, fmt.Sprint(sp))
}

// string转数组
/* 1,1,2,3,4,5*/
func Split(data string) []int {
	var sa = strings.Split(data, ",")
	var sarr []int
	for i := 0; i < len(sa); i++ {
		var v = sa[i]
		var v1, _ = strconv.Atoi(v)
		sarr = append(sarr, v1)
	}
	return sarr
}

// string转数组
/* 1,1,2,3,4,5*/
func SplitStr(data string) []string {
	var sa = strings.Split(data, ",")
	var sarr []string
	for i := 0; i < len(sa); i++ {
		var v = sa[i]
		sarr = append(sarr, v)
	}
	return sarr
}

// 字节的单位转换 保留两位小数
func FormatFileSize(fileSize uint64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fPB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func GetDbType(dt string) string {
	switch dt {
	case "varchar":
		return "string"
	case "datetime":
		return "time.Time"
	case "timestamp":
		return "time.Time"
	case "double":
		return "float64"
	case "float":
		return "float64"
	case "decimal":
		return "float64"
	case "int":
		return "int64"
	case "tinyint":
		return "int64"
	case "smallint":
		return "int64"
	case "mediumint":
		return "int64"
	case "integer":
		return "int64"
	case "bigint":
		return "int64"
	default:
		return "string"
	}
}

func RemoveEmpty(arr []string) (result []string) {
	for _, value := range arr {
		if len(value) > 0 { // 判断字符串长度大于零则添加到结果数组中
			result = append(result, value)
		}
	}
	return
}
