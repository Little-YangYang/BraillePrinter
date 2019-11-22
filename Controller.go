package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"time"
)

// 创建打印任务的RESTFul控制器
// 方法：POST
// Body：JSON->{HTML：··}
func CreateMission(c echo.Context) (err error) {
	// 将body中的HTML文本解析到对象中
	uRequest := new(CreateItem)
	if err = c.Bind(uRequest); err != nil {
		return
	}
	fmt.Println("origin HTML TEXT：", uRequest.HTML)
	// 翻译HTML文本
	res := Translation(uRequest.HTML)
	// 创建打印元对象
	var item PrintItem

	// 遍历HTML翻译文本并加入到打印任务元中
	for _, val := range res {
		for _, v := range ToPrintData(val) {
			item.LineNum += 1
			item.Line = append(item.Line, v)
		}
	}

	fmt.Println("打印行数:", item.LineNum)
	fmt.Println("打印Item:")
	for _, val := range item.Line {
		fmt.Println(val)
	}
	// 打印对象
	m := mission{
		progress: 0,
		item:     &item,
	}

	// 检查串口连接
	if conn == false {
		fmt.Println("系统无法找到打印机！请确认打印机连接状态后重新启动客户端！")
		return json.NewEncoder(c.Response()).Encode(struct {
			Code            string      `json:"code"`
			ResponseMessage string      `json:"message"`
			ResponsePayload interface{} `json:"payload,omitempty"`
		}{
			Code:            "NotFound",
			ResponseMessage: "创建新打印任务失败！系统找不到打印机设备，请重新确认打印机连接状态后重启客户端！",
		})
	}

	// 异步启动打印任务
	go m.On()

	// 创建标准HTTP Response
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(struct {
		Code            string      `json:"code"`
		ResponseMessage string      `json:"message"`
		ResponsePayload interface{} `json:"payload,omitempty"`
	}{
		Code:            "OK",
		ResponseMessage: "创建新打印任务成功！",
	})

}

func PingController(c echo.Context) (err error) {
	c.Response().WriteHeader(http.StatusOK)
	printConn := "Stop"
	if conn == true {
		printConn = "OK"
	}
	return json.NewEncoder(c.Response()).Encode(struct {
		Code            string `json:"code"`
		ResponseMessage string `json:"message"`
		PrintConnected  string `json:"printConnected"`
		Version         string `json:"version"`
	}{
		Code:            "OK",
		ResponseMessage: "OK",
		PrintConnected:  printConn,
		Version:         Version,
	})
}

func ExitSystemController(c echo.Context) (err error) {
	go func() {
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}()
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(struct {
		Code            string `json:"code"`
		ResponseMessage string `json:"message"`
		PrintConnected  string `json:"printConnected"`
	}{
		Code:            "OK",
		ResponseMessage: "退出成功！",
	})
}
