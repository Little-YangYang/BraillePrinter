package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印机唯一锁，防止多个进程对同一台打印机进行重复操作
var printLock sync.Mutex

// 任务项目
type mission struct {
	progress int
	item     *PrintItem
}

// 待实现的打印任务管理器的大饼
type missionManager struct {
	missionManagerList []*mission
}

// 大饼里面的添加大饼操作
func (m *missionManager) Add(mission *mission) {
	m.missionManagerList = append(m.missionManagerList, mission)
}

// 打印任务创建
func (m *mission) Create(p *PrintItem) {
	// 将任务引用指向已有对象
	m.item = p
}

// 启动打印任务
func (m *mission) On() {
	// 阻塞式锁定打印机
	printLock.Lock()
	defer printLock.Unlock()
	if conn == false {
		fmt.Println("系统无法找到打印机！请确认打印机连接状态后重新启动客户端！")
		return
	}
	// 进纸，调整边距
	m.start()
	// todo 警告：该方法暂未实现第二页至多页的分页打印效果！
	// item.Line为每一[行]的数组
	for _, v := range m.item.Line {
		err := m.send([]byte{lineReset})
		if err != nil {
			fmt.Println("行重置出错！", err)
		}
		// item.Line.LineComm 为每一行的[每列]指令数组，类型均为byte
		for _, val := range v.LineComm {
			err := m.send([]byte{val})
			if err != nil {
				m.COMError()
			}
		}
		// 任务进度+1行
		m.progress += 1
	}
	// lineMax为最大行数,用于检测是否不满一页，若不满则使用回车填满
	for i := m.item.LineNum; i <= lineMax; i++ {
		err := m.send([]byte{prChangeLine})
		if err != nil {
			m.COMError()
		}
	}
	m.End()
}

// 检查任务进度
func (m *mission) CheckProgress() float32 {
	progress := float32(m.progress) / float32(m.item.LineNum+1)
	return progress
}

// 启动打印任务准备工作(进纸,调整页边距)
func (m *mission) start() {
	// 进纸prtFeedPaper, prMarginTop, prMarginLeft
	ping := []byte("printerVersion")
	err := m.send([]byte{prtInitialize})
	err = m.send(ping)
	err = m.send(ping)
	err = m.send([]byte{prtFeedPaper})
	if err != nil {
	}
	// 调整上边距
	err = m.send([]byte{prMarginTop})
	if err != nil {
		m.COMError()
	}
	// 调整左边距
	err = m.send([]byte{prMarginLeft})
	if err != nil {
		m.COMError()
	}
}

// 打印任务收尾工作，目前貌似只有一个调整下边距
func (m *mission) End() {
	// 调整下边距
	err := m.send([]byte{prMarginBottom})
	if err != nil {
		m.COMError()
	}
}

// 向打印机发送数据
func (m *mission) send(data []byte) (err error) {
	n, err := s.Write(data)
	if err != nil {
		m.COMError()
	}
	buf := make([]byte, 65535)
	for true {
		n, err = s.Read(buf)

		if err != nil {
			m.COMError()
		}
		if n == 0 {

			time.Sleep(time.Second * 1)
		} else {
			s.Flush()
			break
		}
	}
	return
}

// 串口错误处理
func (m *mission) COMError() {
	fmt.Println("串口通信错误！")
}
