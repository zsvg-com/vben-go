package lolog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"vben/app/core/utils/R"
	"vben/common/json/boolx"
	"vben/common/json/timex"
	"vben/common/utils"
	"vben/pkg/mysql"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	useragent "github.com/wenlng/go-user-agent"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Insert(c *gin.Context, username string, message string, loginSucess bool) {
	userAgent := c.Request.Header.Get("User-Agent")
	Os := useragent.GetOsName(userAgent)
	browser := useragent.GetBrowserName(userAgent)
	var log = MonLoginLog{
		Id:       utils.NextId(),
		Username: username,
		//Tenid:    "0000",
		Himsg:   message,
		Loip:    c.ClientIP(),
		Loloc:   GetRealAddressByIP(c.ClientIP()),
		Browser: browser,
		Os:      Os,
		Sutag:   boolx.Bool(loginSucess),
		Lotim:   timex.Time(time.Now()),
	}

	//Clkey    string     `json:"clkey" gorm:"type:varchar(32)"`    //客户端key
	//Detyp    string     `json:"detyp" gorm:"type:varchar(8)"`     //设备类型

	err := mysql.MysqlDb().Create(&log).Error
	if err != nil {
		panic(R.ReturnFailMsg(err.Error()))
	}

}

type Tunit struct {
	IP   string `json:"ip"`
	Pro  string `json:"pro"`
	City string `json:"city"`
}

func GetRealAddressByIP(ip string) string {
	url := "http://whois.pconline.com.cn/ipJson.jsp?ip=" + ip + "&json=true"

	resp, err := http.Get(url)
	if err != nil {
		return "内网IP"
	}
	defer resp.Body.Close()

	// 读取原始数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "内网IP"
	}

	// 将 GBK 转换为 UTF-8
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	utf8Body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "内网IP"
	}

	// 解析 JSON
	var dws Tunit
	if err := json.Unmarshal(utf8Body, &dws); err != nil {
		return "内网IP"
	}

	// 检查是否为空
	if dws.Pro == "" && dws.City == "" {
		return "内网IP"
	}

	return fmt.Sprintf("%s %s", dws.Pro, dws.City)
}
