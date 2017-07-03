package util

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/CardInfoLink/log"
)

const (
	sysEnv = "TXMONI_ENV"
	pkg    = "github.com/wonsikin/dictionary"
)

// Hostname 主机名
var Hostname string

// Env 系统运行环境，须配置环境变量 QUICKPAY_ENV 为 `develop` or `testing` or `product`
var Env string

// LocalIP 本机 IP
var LocalIP string

// WorkDir 程序启动目录
var WorkDir string

func init() {
	hostname()
	localIP()
	workDir()
}

// localIP 本机 IP
func localIP() {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		LocalIP = "127.0.0.1"
	} else {
		LocalIP = strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
	}
	fmt.Printf("local ip:\t %s\n", LocalIP)
}

// workDir 获取程序启动目录
func workDir() {
	var err error
	WorkDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("can not get work directory: %s\n", err)
		os.Exit(2)
	}
	if pos := strings.Index(WorkDir, pkg); pos >= 0 {
		WorkDir = WorkDir[:(pos + len(pkg))]
	}

	fmt.Printf("work directory:\t %s\n", WorkDir)
}

// hostname 取主机名，如果没取到，返回 `unknown`
func hostname() {
	h, err := os.Hostname()
	if err != nil {
		log.Errorf("get hostname error: %s", h)
		Hostname = "unknown"
	}

	Hostname = strings.Replace(h, ".", "_", -1)

	fmt.Printf("hostname:\t %s\n", Hostname)
}
