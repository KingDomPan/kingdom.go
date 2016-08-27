package main

import (
	"fmt"
	"math"
)

// 来源于第三方包, 处理程序不返回错误, 但是会抛出一个异常
// 因此需要程序进行panic到error的转换
func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out od the int 32 range", x))
}

// 需要使用命名的返回参数
func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e) // 命名返回参数作为返回值
		}
	}()
	i = ConvertInt64ToInt(x)
	return i, nil
}

func main() {}
