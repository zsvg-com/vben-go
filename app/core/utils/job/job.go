package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vben/config"
)

func GetXxlJobCookie() []byte {
	var urlP = config.XxlJob.AdminAddress + "/login"
	data := make(map[string]interface{})
	data["userName"] = config.XxlJob.JobUName
	data["password"] = config.XxlJob.JobUPass
	// 序列化
	bytesData, _ := json.Marshal(data)

	resp, err := http.Post(urlP, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body
}

func AddXxlJob() {

}

func UpdateXxlJob() {

}

func deleteXxlJob() {

}

func startXxlJob() {

}

func stopXxlJob() {

}

func getGroupId() {

}
