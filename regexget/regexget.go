package regexget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	// 将html标签全部转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除Style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style>\\>")
	src = re.ReplaceAllString(src, "")

	// 去除Script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script>\\>")
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的html代码并换成换行符\n
	re, _ = regexp.Compile("\\<[\\S\\s]+?>")
	src = re.ReplaceAllString(src, "\n")

	// 去除连线的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
