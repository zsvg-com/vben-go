package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"vben/app/core/utils"
)

type Server struct {
	Cpu     Cpu       `json:"cpu"`
	Jvm     Jvm       `json:"jvm"`
	Mem     Mem       `json:"mem"`
	Sys     Sys       `json:"sys"`
	SysFile []SysFile `json:"sysFile"`
}

type Cpu struct {
	CpuNum int     `json:"cpuNum"`
	Total  int     `json:"total"`
	Sys    string  `json:"sys"`
	Used   float64 `json:"used"`
	Wait   string  `json:"wait"`
	Free   string  `json:"free"`
}

type Jvm struct {
	Total   uint64 `json:"total"`
	Max     uint64 `json:"max"`
	Free    uint64 `json:"free"`
	Version string `json:"version"`
	Home    string `json:"home"`
}

type Mem struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

type Sys struct {
	ComputerName string `json:"computerName"`
	ComputerIp   string `json:"computerIp"`
	UserDir      string `json:"userDir"`
	OsName       string `json:"osName"`
	OsArch       string `json:"osArch"`
}

type SysFile struct {
	DirName     string  `json:"dirName"`
	SysTypeName string  `json:"sysTypeName"`
	TypeName    string  `json:"typeName"`
	Total       string  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	Usage       float64 `json:"usage"`
}

func GetServerInfo(context *gin.Context) Server {
	ip := utils.GetRemoteClientIp(context.Request)

	var server Server
	server.Cpu = getCpu()
	server.Jvm = getJvm()
	server.Mem = getMem()
	server.Sys = getSys(ip)
	server.SysFile = getSysFile()
	return server
}

func getCpu() Cpu {
	cpuCount, _ := cpu.Counts(true) //cpu逻辑数量
	return Cpu{
		CpuNum: runtime.NumCPU(),
		Total:  cpuCount,
		Sys:    "",
		Used:   getCpuPercent(),
		Wait:   "",
		Free:   "",
	}
}
func getCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func getJvm() Jvm {
	info, _ := mem.VirtualMemory()
	return Jvm{
		Total:   info.Total,
		Max:     info.HighTotal,
		Free:    info.Free,
		Version: "",
		Home:    "",
	}
}

func getMem() Mem {
	info, _ := mem.VirtualMemory()
	return Mem{
		Total: info.Total,
		Used:  info.Used,
		Free:  info.Total - info.Used,
	}
}

func getSys(ip string) Sys {
	info, _ := host.Info()
	return Sys{
		ComputerName: info.Hostname,
		ComputerIp:   ip,
		UserDir:      GetAppPath(),
		OsName:       info.OS,
		OsArch:       runtime.GOARCH,
	}
}

// GetAppPath 获取运行时的相对目录
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func getSysFile() []SysFile {
	var sysFiles []SysFile
	infos, _ := disk.Partitions(true) //所有分区
	for i := 0; i < len(infos); i++ {
		var info = infos[i]
		info2, _ := disk.Usage(info.Mountpoint) //指定某路径的硬盘使用情况
		var sysFile = SysFile{
			DirName:     info2.Path,
			SysTypeName: info.Fstype,
			TypeName:    "",
			Total:       utils.FormatFileSize(info2.Total),
			Free:        info2.Free,
			Used:        info2.Used,
			Usage:       info2.UsedPercent,
		}
		sysFiles = append(sysFiles, sysFile)
	}
	return sysFiles
}
