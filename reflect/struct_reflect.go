package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

// 修改结构体
func SetValue(i interface{}) {
	v := reflect.ValueOf(i)
	v = v.Elem()
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("akmf")
	}
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

// 查看匿名字段
func ViewAnonymousFields() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

// 查看类型、字段和方法
// 传入interface{}
func ToString(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf(m.Name + "  ")
		fmt.Println(m.Type)
	}

}

func main() {
	u := User{1, "akmf123", 18}
	SetValue(&u)
	// SetValue(u) 错误的写法
	fmt.Println(u)

	// 查看匿名字段
	ViewAnonymousFields()

	//
	fmt.Println("\n下面输出结构体")
	ToString(u)
}
