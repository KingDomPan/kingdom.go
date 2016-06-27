package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "I m learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// 查找第一个符合要求的关键字
	match := re.Find([]byte(str))
	fmt.Println(string(match))

	// 查找符合正则的所有slice, 小于0表示返回全部符合的字符串, 否则返回指定的长度
	all := re.FindAll([]byte(str), -1)
	for _, v := range all {
		fmt.Println(string(v))
	}
}
