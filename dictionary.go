package main

var Dictionary = make(map[string]string)
var ENDictionary = make(map[string]string)

func InitDictionary() {
	// *------声母表-------*
	Dictionary["b"] = "30"
	Dictionary["p"] = "71"
	Dictionary["m"] = "51"
	Dictionary["f"] = "31"
	Dictionary["d"] = "13"
	Dictionary["t"] = "63"
	Dictionary["n"] = "53"
	Dictionary["l"] = "70"
	Dictionary["g"] = "33"
	Dictionary["j"] = "33"
	Dictionary["k"] = "50"
	Dictionary["q"] = "50"
	Dictionary["h"] = "32"
	Dictionary["x"] = "32"
	Dictionary["r"] = "23"
	Dictionary["z"] = "56"
	Dictionary["c"] = "11"
	Dictionary["s"] = "61"
	Dictionary["zh"] = "41"
	Dictionary["ch"] = "73"
	Dictionary["sh"] = "16"

	// *------韵母表------*
	Dictionary["a"] = "42"
	Dictionary["o"] = "24"
	Dictionary["wo"] = "24"
	Dictionary["e"] = "24"
	Dictionary["i"] = "21"
	Dictionary["u"] = "54"
	Dictionary["v"] = "45"
	Dictionary["er"] = "72"
	Dictionary["ai"] = "25"
	Dictionary["ao"] = "62"
	Dictionary["ei"] = "65"
	Dictionary["ou"] = "76"
	Dictionary["ia"] = "35"
	Dictionary["iao"] = "43"
	Dictionary["ie"] = "12"
	// iou => iu
	Dictionary["iu"] = "36"
	Dictionary["ua"] = "77"
	Dictionary["uai"] = "57"
	// uei => ui
	Dictionary["ui"] = "27"
	Dictionary["uo"] = "52"
	// ve => ue
	Dictionary["ue"] = "67"
	Dictionary["an"] = "74"
	Dictionary["ang"] = "64"
	Dictionary["en"] = "46"
	Dictionary["eng"] = "47"
	Dictionary["ian"] = "15"
	// iang => yang
	Dictionary["yang"] = "55"
	Dictionary["in"] = "34"
	Dictionary["ing"] = "14"
	// uan => wan
	Dictionary["wan"] = "37"
	// uang => wang
	Dictionary["wang"] = "66"
	// un => wen
	Dictionary["un"] = "22"
	Dictionary["wun"] = "22"
	// uong => wong
	Dictionary["ong"] = "26"
	Dictionary["wong"] = "26"
	// van => uan
	Dictionary["uan"] = "75"
	// vun => yun
	Dictionary["yun"] = "07"
	Dictionary["iong"] = "17"

	// *------数字表------*
	Dictionary["1"] = "4710"
	Dictionary["2"] = "4720"
	Dictionary["3"] = "4711"
	Dictionary["4"] = "4713"
	Dictionary["5"] = "4712"
	Dictionary["6"] = "4731"
	Dictionary["7"] = "4733"
	Dictionary["8"] = "4732"
	Dictionary["9"] = "4721"
	Dictionary["0"] = "4723"

	// *------标点符号------*
	Dictionary[","] = "0200"
	Dictionary["，"] = "0200"
	Dictionary["."] = "0260"
	Dictionary["。"] = "0260"
	Dictionary["?"] = "0240"
	Dictionary["!"] = "0620"
	Dictionary[":"] = "0440"
	Dictionary["："] = "0440"
	Dictionary[";"] = "0600"
	Dictionary["；"] = "0600"
	Dictionary["\""] = "0620"
	Dictionary["“"] = "0620"
	Dictionary["”"] = "0460"
	Dictionary["'"] = "0624"
	Dictionary["‘"] = "0624"
	Dictionary["’"] = "4260"
	Dictionary["("] = "0644"
	Dictionary["（"] = "0644"
	Dictionary[")"] = "4460"
	Dictionary["）"] = "4460"
	Dictionary["["] = "0660"
	Dictionary["【"] = "0660"
	Dictionary["]"] = "0660"
	Dictionary["】"] = "0660"
	Dictionary["·"] = "0220"
	Dictionary["…"] = "020202"

	// *------英文字母表------*
	ENDictionary["a"] = "10"
	ENDictionary["b"] = "30"
	ENDictionary["c"] = "11"
	ENDictionary["d"] = "13"
	ENDictionary["e"] = "12"
	ENDictionary["f"] = "31"
	ENDictionary["g"] = "33"
	ENDictionary["h"] = "32"
	ENDictionary["i"] = "21"
	ENDictionary["j"] = "23"
	ENDictionary["k"] = "50"
	ENDictionary["l"] = "70"
	ENDictionary["m"] = "51"
	ENDictionary["n"] = "53"
	ENDictionary["o"] = "52"
	ENDictionary["p"] = "71"
	ENDictionary["q"] = "73"
	ENDictionary["r"] = "72"
	ENDictionary["s"] = "61"
	ENDictionary["t"] = "63"
	ENDictionary["u"] = "54"
	ENDictionary["v"] = "71"
	ENDictionary["w"] = "27"
	ENDictionary["x"] = "55"
	ENDictionary["y"] = "57"
	ENDictionary["z"] = "56"
}
