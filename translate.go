package main

import (
	"fmt"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"log"
	"strings"
)

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

//行复位
var lineReset byte = 0xB2

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

// original为HTML字符串，段落以<p></p>划分
func Translation(original string) []string {

	dict := pinyin.NewDict()
	var resArray []string
	InitDictionary()
	//testString := `<p> 我爱北京天AsZz安门，天安门上有1个太阳升</p><p>我还超级爱毛主席呢</p>`

	original = strings.ReplaceAll(original, "<p>", "")
	StringArray := strings.Split(original, "</p>")
	for _, val := range StringArray {
		part := ""
		// 开始解析字符
		for _, font := range val {
			println("font(RUNE编号):", font, "，String：", string(font))
			// 数字解析
			if font >= '0' && font <= '9' {
				part += Dictionary[string(font)]
			} else if font > 64 && font < 123 {
				// 英文字符解析
				str := strings.ToLower(string(font))
				log.Println("发现英文字符", string(font), ":", ENDictionary[str])
				part += ENDictionary[str]
			} else if font == 32 {
				// 空格解析
				log.Println("发现空格:", "00")
				part += "00"
			} else {
				// 其他文本解析
				str := dict.Sentence(string(font)).None()
				res := Translate(str)
				part += res
			}
		}
		// 添加解析结果
		resArray = append(resArray, part)
	}
	return resArray
}

// 翻译文本
func Translate(original string) string {
	log.Println("original:", original)
	if val, found := Dictionary[original]; found {
		log.Println("发现词典内容", original, ":", Dictionary[original])
		return val
	}
	// 若为双字母构成的拼音，则直接用本方法
	if len(original) <= 2 {
		res := ""
		for _, val := range original {
			res += Dictionary[string(val)]
		}
		log.Println("双字符：", original, "：", res)
		return res
	}
	// 直接翻译
	if _, found := Dictionary[original[:2]]; found {
		log.Println("发现词典内容", original[:2], ":", Dictionary[original[:2]], original[2:], ":", Dictionary[original[2:]])
		return Dictionary[original[:2]] + Dictionary[original[2:]]
	}
	log.Println("发现词典内容", original[:1], ":", Dictionary[original[:1]], original[1:], ":", Dictionary[original[1:]])
	return Dictionary[original[:1]] + Dictionary[original[1:]]
}

// 该函数完成从数字到指令的转换
func ToPrintData(original string) []PrintLineItem {
	var res []PrintLineItem
	for len(original) > colMax {
		fmt.Println("单段落长度超出一行，进行折行处理。当前指令文本长度：", len(original))
		res = append(res, Transform(original[:colMax]))
		original = original[colMax:]
	}
	if len(original) > 0 {
		res = append(res, Transform(original))
	}
	return res
}

var left = true

// 到打印任务的转换
func Transform(original string) (res PrintLineItem) {
	length := len(original)
	if length < colMax {
		for len(original) < colMax {
			original += "0"
		}
	}
	fmt.Println("转换为指令码，原文本：", original)
	if left == true {
		fmt.Println("打印头在左侧，开始正向打印")
		for i, val := range original {
			if i%2 == 1 {
				res.LineComm = append(res.LineComm, prSpaceWord)
			} else if i == 0 {
			} else {
				res.LineComm = append(res.LineComm, prSpacePoint)
			}
			switch string(val) {
			case "0":
				res.LineComm = append(res.LineComm, prPointZero)
			case "1":
				res.LineComm = append(res.LineComm, prPointOne)
			case "2":
				res.LineComm = append(res.LineComm, prPointTwo)
			case "3":
				res.LineComm = append(res.LineComm, prPointThree)
			case "4":
				res.LineComm = append(res.LineComm, prPointFour)
			case "5":
				res.LineComm = append(res.LineComm, prPointFive)
			case "6":
				res.LineComm = append(res.LineComm, prPointSix)
			case "7":
				res.LineComm = append(res.LineComm, prPointSeven)
			}
		}
	} else {
		fmt.Println("打印头在右侧，开始反向打印")
		for i := len(original) - 1; i >= 0; i-- {
			if i == len(original)-1 {
			} else if i%2 == 0 {
				res.LineComm = append(res.LineComm, prSpaceWord)
			} else {
				res.LineComm = append(res.LineComm, prSpacePoint)
			}
			switch string(original[i]) {
			case "0":
				res.LineComm = append(res.LineComm, prPointZero)
			case "1":
				res.LineComm = append(res.LineComm, prPointOne)
			case "2":
				res.LineComm = append(res.LineComm, prPointTwo)
			case "3":
				res.LineComm = append(res.LineComm, prPointThree)
			case "4":
				res.LineComm = append(res.LineComm, prPointFour)
			case "5":
				res.LineComm = append(res.LineComm, prPointFive)
			case "6":
				res.LineComm = append(res.LineComm, prPointSix)
			case "7":
				res.LineComm = append(res.LineComm, prPointSeven)
			}
		}
	}

	res.LineComm = append(res.LineComm, prChangeLine)
	if left == true {
		left = true
	} else {
		left = true
	}
	return
}
