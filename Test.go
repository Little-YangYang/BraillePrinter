package main

import (
	"fmt"
)

func Test() {
	//	//str := dict.Sentence().None()
	//	log.Println("str:",testStringArray)
	//	log.Println("part:",resArray)
	testString :=
		`
<p>春晓</p>
<p>孟浩然</p>
<p>春眠不觉晓，</p>
<p>处处闻啼鸟，</p>
<p>夜来风雨声，</p>
<p>花落知多少，</p>
`

	res := Translation(testString)
	var item PrintItem

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
}
