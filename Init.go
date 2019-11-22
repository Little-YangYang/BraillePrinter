package main

import (
	"fmt"
	"github.com/tarm/serial"
	"strconv"
	"time"
)

var s *serial.Port
var conn = false
var COM string
var version string

func Init() {
	findCOM()
}

/**
 * 搜索COM串口
 * Warning：该段代码可能存在隐藏BUG
 *         在双线连接时无法搜索COM3和COM5的区别
 *         ping及返回值疑似无效
 */
func findCOM() {
	fmt.Println("正在查找打印机，开始扫描系统COM口。")
	for i := 1; i <= 18; i++ {
		COMStr := "COM" + strconv.Itoa(i)
		fmt.Println("正在扫描", COMStr)
		c := &serial.Config{Name: COMStr, Baud: 115200, ReadTimeout: time.Second * 5}
		// 开启串口端口
		var err error
		s, err = serial.OpenPort(c)
		// 检查串口是否能被打开
		if err != nil {
			fmt.Println(COMStr, "扫描失败！该COM口无设备！", err)
			continue
		}

		// 该段代码使用ping的方式尝试连接串口设备，返回值无法正常收取，疑似存在BUG
		fmt.Println("在", COMStr, "上发现了设备，正在尝试连接。")
		ping := []byte("printerVersion")
		_, err = s.Write(ping)
		_, err = s.Write(ping)

		// 收取返回值
		time.Sleep(2 * time.Second)
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err == nil {
			// 无错误：连接成功
			fmt.Println("打印机连接成功！COM口为：", COMStr)
			version = string(buf[:n])
			conn = true
			COM = COMStr
			_ = s.Flush()
			break
		} else {
			// 出现错误，该端口无法使用
			fmt.Println("连接失败！该设备非本软件相关设备！", err)
		}
	}
}
