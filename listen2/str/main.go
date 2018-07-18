package main

import (
	"fmt"
	"strings"

	"github.com/study/listen2/asscess"
)

func main() {
	// 获取字符串长度 len(str)
	var a string = "hello"
	var alen = len(a)
	fmt.Println(alen)

	// 字符串拼接 + fmt.Sprintf
	b := "hello "
	c := "world"
	fmt.Printf("b+c = %s\n", b+c)
	e := fmt.Sprintf("%s%s", b, c)
	fmt.Printf("b+c = %s\n", e)

	// 字符串分割 strings.Split
	f := "10.194.94.234;102.50.23.55;13.12.123.42"
	g := strings.Split(f, ";")
	fmt.Println(g)

	// 包含 strings.Contains
	h := strings.Contains(f, "10.194.94.234")
	fmt.Println(h)

	// 前缀和后缀判断  strings.HasPrefix strings.HasSuffix
	j := "hello world"
	k := strings.HasPrefix(j, "h")
	l := strings.HasSuffix(j, "d")
	fmt.Println(k)
	fmt.Println(l)

	// 子字符串出现的位置 strings.Index strings.LastIndex
	m := strings.Index(j, "e")     //1
	n := strings.LastIndex(j, "l") //9
	fmt.Println(m, n)

	// join
	var arr []string = []string{"104,194,94,249", "127.0.0.1", "22.123.44.231"}
	o := strings.Join(arr, ";")
	fmt.Println(o)

	// 包外部引用必须首字母大写
	fmt.Println(asscess.A)
}
