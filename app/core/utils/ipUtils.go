package utils

// import (
// 	"bytes"
// 	"encoding/json"
// 	"golang.org/x/text/encoding/simplifiedchinese"
// 	"golang.org/x/text/transform"
// 	"io/ioutil"
// 	"net"
// 	"net/http"
// 	"time"
// )

// func GetRealAddressByIPText(ip string) string {
// 	toByteIp := net.ParseIP(ip)
// 	if isLocalIp(toByteIp) {
// 		return "服务器登录"
// 	}
// 	if isLANIp(toByteIp) {
// 		return "局域网"
// 	}
// 	return getLocation(ip)
// }

// func isLocalIp(IP net.IP) bool {
// 	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
// 		return true
// 	}
// 	return false
// }

// func isLANIp(IP net.IP) bool {
// 	if ip4 := IP.To4(); ip4 != nil {
// 		switch true {
// 		case ip4[0] == 10:
// 			return true
// 		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
// 			return true
// 		case ip4[0] == 192 && ip4[1] == 168:
// 			return true
// 		default:
// 			return false
// 		}
// 	}
// 	return false
// }
// func getLocation(ip string) string {
// 	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
// 	client := &http.Client{Timeout: 5 * time.Second}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return "未知地址"
// 	}
// 	defer resp.Body.Close()
// 	result, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "未知地址"
// 	}
// 	reader := transform.NewReader(bytes.NewReader(result), simplifiedchinese.GBK.NewDecoder())
// 	d, err := ioutil.ReadAll(reader)
// 	if err != nil {
// 		return "未知地址"
// 	}
// 	m := make(map[string]string)
// 	err = json.Unmarshal(d, &m)
// 	addr := m["addr"]
// 	if addr != "" {
// 		return addr
// 	} else {
// 		return "未知地址"
// 	}
// }
// func GetLocalIP() (ip string, err error) {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return
// 	}
// 	for _, addr := range addrs {
// 		ipAddr, ok := addr.(*net.IPNet)
// 		if !ok {
// 			continue
// 		}
// 		if ipAddr.IP.IsLoopback() {
// 			continue
// 		}
// 		if !ipAddr.IP.IsGlobalUnicast() {
// 			continue
// 		}
// 		return ipAddr.IP.String(), nil
// 	}
// 	return
// }
