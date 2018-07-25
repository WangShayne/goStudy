package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

func example() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成风格
		num: 纯数字 [0-9]
		char:只使用英文字母[a-zA-Z]
		mix:使用数字和字母
		advance:使用数字,字母以及特殊字符
		`)
	flag.Parse()
}

func generatePWD() string {
	example()
	rand.Seed(time.Now().Unix())
	const (
		NUMSTR   = "0123456789"
		STRSTR   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		SPACESTR = "!@#$%^&*"
	)
	var passwrd []byte = make([]byte, length, length)
	var strList string
	switch charset {
	case "num":
		strList = NUMSTR
	case "char":
		strList = STRSTR
	case "mix":
		strList = fmt.Sprintf("%s%s", NUMSTR, STRSTR)
	case "advance":
		strList = fmt.Sprintf("%s%s%s", NUMSTR, STRSTR, SPACESTR)
	}
	for i := 0; i < length; i++ {
		x := rand.Intn(len(strList))
		passwrd[i] = strList[x]
	}
	return string(passwrd)
}

func main() {
	fmt.Println(generatePWD())
}
