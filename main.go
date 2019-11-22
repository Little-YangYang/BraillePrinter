package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os/exec"
	"time"
)

var (
	colMax  = 20
	lineMax = 20
	Version = "Version. Release 1.0"
	signal  = make(chan int)
	status  = 0
)

func main() {
	// 翻译测试用函数
	// Test()

	// 初始化系统
	Init()

	// 创建ECHO服务器
	e := echo.New()
	e.HideBanner = true
	// CORS跨域处理
	e.Use(middleware.CORS())
	fmt.Println("请打开系统浏览器，访问\"http://127.0.0.1:4396\"进入控制页面！")

	// 路由组
	e.Static("/", "www/")
	e.POST("/createMission", CreateMission)
	e.POST("/ping", PingController)
	e.POST("/exit", ExitSystemController)
	// 自动开启浏览器
	go func() {
		time.Sleep(time.Second * 2)

		err := exec.Command(`cmd`, `/c`, `start`, `http://127.0.0.1:4396`).Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	e.Logger.Fatal(e.Start(":4396"))
}
