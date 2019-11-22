package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
)

// demo示例程序
func demo() {

	// Set Serial Port and Baud

	fmt.Println("输入COM口的编号，必须是大写，例如 COM5")
	str := ""
	_, _ = fmt.Scanf("%s", &str)
	fmt.Println(str)
	c := &serial.Config{Name: str, Baud: 115200}

	// Open Serial Port
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	//str := []byte("printerVersion")

	ping := []byte("printerVersion")
	var prPointZero byte = 0x81
	var prPointOne byte = 0x82
	var prPointTwo byte = 0x83
	var prPointThree byte = 0x84
	var prPointFour byte = 0x85
	var prPointFive byte = 0x86
	var prPointSix byte = 0x87
	var prPointSeven byte = 0x88
	var prSpacePoint byte = 0x89
	var prSpaceWord byte = 0x8A
	var prChangeLine byte = 0x8B
	var prMarginTop byte = 0x91
	var prMarginBottom byte = 0x92
	var prMarginLeft byte = 0x93
	//var prMarginRight byte =0x94
	//var adPoint byte = 0xA1
	//var adWord byte = 0xA2
	//var adLine byte = 0xA3
	//var adMarginTop byte =0xA4
	//var adMarginBottom byte =0xA5
	//var adMarginLeft byte=0xA6
	//var adMarginRight byte = 0xA7
	var prtInitialize byte = 0xB1
	////var prtCuInitialize byte =0xB2
	var prtFeedPaper byte = 0xB3
	//var prtPause byte = 0xB4
	//var prtResume byte =0xB5
	//var prtStop byte = 0xB6
	//initialization := []byte{1, 2, adPoint, 3, 9, adWord, 5, 5, adLine,6,6,adMarginTop,6,6,adMarginBottom,3,3,adMarginLeft,3,3,adMarginRight,prtInitialize,
	//	prtFeedPaper,}
	normalWord := []byte{prtFeedPaper, prMarginTop, prMarginLeft, prPointZero, prSpacePoint, prPointZero, prSpaceWord, prPointZero, prSpacePoint, prPointZero, prSpaceWord,
		prPointSeven, prSpacePoint, prPointThree, prSpaceWord, prPointTwo, prSpacePoint, prPointTwo, prSpaceWord, prPointOne, prSpacePoint, prPointZero, prSpaceWord,
		prPointThree, prSpacePoint, prPointTwo, prSpaceWord, prPointFour, prSpacePoint, prPointThree, prSpaceWord, prPointFive, prSpacePoint, prPointZero, prSpaceWord, prSpaceWord, prSpaceWord, prSpaceWord, prSpaceWord, prChangeLine,

		prPointFour, prSpacePoint, prPointSeven, prSpaceWord, prPointThree, prSpacePoint, prPointTwo, prSpaceWord, prPointZero, prSpacePoint, prPointSix, prSpaceWord, prPointTwo, prSpacePoint, prPointSix, prSpaceWord, prPointTwo, prSpacePoint, prPointThree, prSpaceWord, prPointZero, prSpacePoint, prPointSix, prSpaceWord, prPointSeven, prSpacePoint, prPointFour, prSpaceWord, prPointOne, prSpacePoint, prPointFive, prSpaceWord, prPointZero, prSpacePoint, prPointZero, prSpaceWord, prPointZero, prSpacePoint, prPointZero, prSpaceWord, prPointZero, prSpacePoint, prPointZero, prSpaceWord, prPointZero, prSpacePoint, prPointZero, prSpaceWord, prChangeLine,

		prPointSeven, prSpacePoint, prPointThree, prSpaceWord, prPointTwo, prSpacePoint, prPointTwo, prSpaceWord, prPointOne, prSpacePoint, prPointZero, prSpaceWord, prPointFive, prSpacePoint, prPointOne, prSpaceWord, prPointOne, prSpacePoint, prPointFive, prSpaceWord, prPointThree, prSpacePoint, prPointZero, prSpaceWord, prPointFive, prSpacePoint, prPointFour, prSpaceWord, prPointThree, prSpacePoint, prPointThree, prSpaceWord, prPointSix, prSpacePoint, prPointSeven, prSpaceWord, prPointTwo, prSpacePoint, prPointZero, prSpaceWord, prPointThree, prSpacePoint, prPointTwo, prSpaceWord, prPointFour, prSpacePoint, prPointThree, prSpaceWord, prPointFour, prSpacePoint, prPointTwo, prSpaceWord, prChangeLine,

		prPointSix, prSpacePoint, prPointTwo, prSpaceWord, prPointZero, prSpacePoint, prPointFour, prSpaceWord, prPointThree, prSpacePoint, prPointFour, prSpaceWord, prPointThree, prSpacePoint, prPointFive, prSpaceWord, prPointOne, prSpacePoint, prPointTwo, prSpaceWord, prPointThree, prSpacePoint, prPointSix, prSpaceWord, prPointZero, prSpacePoint, prPointTwo, prSpaceWord, prPointTwo, prSpacePoint, prPointTwo, prSpaceWord, prPointZero, prSpacePoint, prPointSix, prSpaceWord, prPointFour, prSpacePoint, prPointFive, prSpaceWord, prPointThree, prSpacePoint, prPointSeven, prSpaceWord, prPointZero, prSpacePoint, prPointFour, prSpaceWord, prPointFour, prSpacePoint, prPointFive, prSpaceWord, prPointThree, prSpacePoint, prPointSeven, prSpaceWord, prChangeLine,
		prPointOne, prSpacePoint, prPointTwo, prSpaceWord, prPointSix, prSpacePoint, prPointZero, prSpaceWord, prPointSeven, prSpacePoint, prPointZero, prSpaceWord, prPointTwo, prSpacePoint, prPointFive, prSpaceWord, prPointTwo, prSpacePoint, prPointZero, prSpaceWord, prPointThree, prSpacePoint, prPointOne, prSpaceWord, prPointFour, prSpacePoint, prPointSeven, prSpaceWord, prPointFour, prSpacePoint, prPointFive, prSpaceWord, prPointFour, prSpacePoint, prPointZero, prSpaceWord, prPointOne, prSpacePoint, prPointSix, prSpaceWord, prPointFour, prSpacePoint, prPointSeven, prSpaceWord, prPointOne, prSpacePoint, prPointTwo, prSpaceWord, prChangeLine,

		prPointZero, prSpacePoint, prPointFour, prSpaceWord, prPointTwo, prSpacePoint, prPointSix, prSpaceWord, prPointSix, prSpacePoint, prPointOne, prSpaceWord, prPointZero, prSpacePoint, prPointOne, prSpaceWord, prPointTwo, prSpacePoint, prPointFive, prSpaceWord, prPointThree, prSpacePoint, prPointOne, prSpaceWord, prPointZero, prSpacePoint, prPointOne, prSpaceWord, prPointOne, prSpacePoint, prPointFour, prSpaceWord, prPointTwo, prSpacePoint, prPointFive, prSpaceWord, prPointZero, prSpacePoint, prPointSeven, prSpaceWord, prPointZero, prSpacePoint, prPointOne, prSpaceWord, prPointSeven, prSpacePoint, prPointSeven, prSpaceWord, prPointTwo, prSpacePoint, prPointThree, prSpaceWord,
		prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prChangeLine, prMarginBottom}
	// Write a String
	prtInit := []byte{prtInitialize}

	n, err := s.Write(ping)
	n, err = s.Write(ping)
	n, err = s.Write(prtInit)
	n, err = s.Write(ping)
	n, err = s.Write(ping)

	for i := 0; i < len(normalWord); i++ {
		byteBuff := []byte{normalWord[i]}
		n, err = s.Write(byteBuff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		buf := make([]byte, 65535)
		n, err = s.Read(buf)

		if err != nil {
			log.Fatal(err)
		}
		n = 0
	}
	fmt.Printf("Write %d Bytes\r\n", n)
	time.Sleep(1 * time.Second)
	// make and set a buff

	//log.Print("%q", buf[:n])
	//fmt.Printf("Read %d Bytes\r\n", n)
	//for i := 0; i < n; i++ {
	//	//fmt.Printf("buf[%d]=%c ", i, buf[i])
	//	fmt.Printf("%c", buf[i])
	//}
	//n = 0
}
