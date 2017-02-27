package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/deckarep/golang-set"
	"github.com/shopspring/decimal"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	//var p interface{}
	//data := `[{"server": "panqd"}]`
	//json.Unmarshal([]byte(data), &p)
	//fmt.Println(p)

	set := mapset.NewSet()
	set.Add("Welding")
	set.Add("Music")

	set2 := mapset.NewSet()
	set2.Add("Welding")
	set2.Add("Musi1")
	set2.Add("Automotive")

	fmt.Println(set2.IsSuperset(set))

	fmt.Println(decimal.NewFromFloat(0.083 + 0.062 + 0.013 + 0.007 + 0.030).Div(decimal.NewFromFloat(float64(5))).StringFixed(2))

	tpl_merger_structdata()
	fmt.Println(SliceHost(""))
	fmt.Println(SliceHost("a"))
	fmt.Println(SliceHost("aa"))
	fmt.Println(SliceHost("aaa"))
	fmt.Println(SliceHost("aaaa"))
	fmt.Println(SliceHost("aaaahosta"))
	fmt.Println(SliceHost("aaaahost"))

	fmt.Println(strings.Replace("/home/www/${app}/${app}.log", "${app}", "qqq", -1))
}

func tpl_merger_structdata() {
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	data := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(data, map[string]string{
		"Count":    "17",
		"Material": "Panqd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data.String()))
}

func SliceHost(groupName string) string {
	if groupName == "" {
		return ""
	}
	if len(groupName) <= 4 {
		return groupName
	}
	if groupName[len(groupName)-4:] == "host" {
		return groupName[:len(groupName)-4]
	}
	return groupName
}
