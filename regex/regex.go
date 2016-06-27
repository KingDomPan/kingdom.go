package regex

import (
	"fmt"
	"os"
	"regexp"
)

// 三个匹配正则的函数
// regexp.Match(pattern string, b []byte) bool
// regexp.MatchString(pattern string, s string) bool
// regexp.MatchReader(pattern string, r io.RuneReader) bool
func IsIp(ip string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: Regexp [string]")
		os.Exit(0)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not A Number")
	}
	fmt.Println(IsIp("127.0.0.1"))
	fmt.Println(IsIp("1127.0.0.1"))
}
