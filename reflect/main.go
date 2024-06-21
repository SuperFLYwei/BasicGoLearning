package main

import (
	"fmt"
	"reflect"
)

// 反射获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是:", t)
	// kind()可以获取具体类型
	k := t.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println("a's type is float64")
	case reflect.String:
		fmt.Println("a's type is string")
	}
}

// 反射获取interface类型信息
func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	// kind()可以获取具体类型
	k := v.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println("a是:", v.Float())
	case reflect.String:
		fmt.Println("a是:", v.String())
	}
}

func main() {
	var y = 3.4 // 默认是float64
	reflect_type(y)
	reflect_value(y)

	var x = "asdd"
	reflect_type(x)
	reflect_value(x)
}
