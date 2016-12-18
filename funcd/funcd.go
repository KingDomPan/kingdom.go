package main

import (
	"fmt"
	"reflect"
)

type s struct {
	Name string
}

func say(text string) s {
	fmt.Println(text)
	return s{
		Name: "panqd",
	}
}

func Q(name string) string {
	return "panqd"
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	var funcMap = make(map[string]interface{})
	funcMap["say"] = say
	funcMap["Q"] = Q
	r := Call(funcMap, "say", "hello")
	fmt.Println(r[0].FieldByName("Name").String())

	q := funcMap["Q"]
	if v, ok := q.(func(string) string); ok != false {
		fmt.Println(v("q"))
	}

}
