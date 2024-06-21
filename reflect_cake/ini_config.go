package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// 文件读取数据，反序列化
func UnMarshallFile(filename string, result interface{}) (err error) {
	//  1. 文件读取
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	// 2. 进行反序列化
	return UnMarshall(data, result)
}

// 反序列化
func UnMarshall(data []byte, result interface{}) (err error) {
	// 先判断是否是指针
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		return
	}
	// 判断是否是结构体
	if typeInfo.Elem().Kind() != reflect.Struct {
		return
	}
	// 转类型，按行切割
	lineArr := strings.Split(string(data), "\n")
	// 定义全局标签名  也就是server 和 mysql
	var myFiledName string

	for _, line := range lineArr {
		line = strings.TrimSpace(line)
		// 处理有注释的情况
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		// 按照括号去判断
		if line[0] == '[' {
			// 按照大标签去处理
			myFiledName, err = myLabel(line, typeInfo.Elem())
			if err != nil {
				return
			}
			continue
		}
		// 按照正常数据处理
		err = myField(myFiledName, line, result)
		if err != nil {
			return
		}
	}
	return
}

// 处理大标签
func myLabel(line string, typeInfo reflect.Type) (fieldName string, err error) {
	// 去除字符串的头和尾
	labelName := line[1 : len(line)-1]
	// 循环去结构体找tag, 对应上进行解析
	for i := 0; i < typeInfo.NumField(); i++ {
		filed := typeInfo.Field(i)
		tagValue := filed.Tag.Get("ini")
		// 判断tag
		if labelName == tagValue {
			fieldName = filed.Name
			break
		}
	}
	return
}

// 解析属性
// 参数：大标签名，行数据，对象
func myField(fieldName string, line string, result interface{}) (err error) {
	fmt.Println(line)
	key := strings.TrimSpace(line[0:strings.Index(line, "=")])
	val := strings.TrimSpace(line[strings.Index(line, "=")+1:])
	// 解析到结构体
	//resultType := reflect.TypeOf(result)
	resultValue := reflect.ValueOf(result)
	// 拿到字段值，这里直接设置不知道类型
	labelValue := resultValue.Elem().FieldByName(fieldName)
	// 拿到该字段类型
	fmt.Println(labelValue)
	labelType := labelValue.Type()
	// 第一次进来应该是server
	// 存放取到的字段名
	var keyName string
	// 遍历server结构体的所有字段
	for i := 0; i < labelType.NumField(); i++ {
		// 获取结构体字段
		field := labelType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyName = field.Name
			break
		}
	}

	// 给字段赋值
	// 取字段值
	filedValue := labelValue.FieldByName(keyName)
	// 修改值
	switch filedValue.Type().Kind() {
	case reflect.String:
		filedValue.SetString(val)
	case reflect.Int:
		i, err2 := strconv.ParseInt(val, 10, 64)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		filedValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			fmt.Println(err)
			return err
		}
		filedValue.SetUint(i)
	case reflect.Float32:
		f, _ := strconv.ParseFloat(val, 64)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		filedValue.SetFloat(f)
	}
	return
}
