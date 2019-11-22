package main

type PrintItem struct {
	LineNum int
	Line    []PrintLineItem
}

type PrintLineItem struct {
	LineComm []byte
}

type CreateItem struct {
	HTML string `json:"html"`
}
