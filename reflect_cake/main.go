package main

import (
	"fmt"
	"io/ioutil"
)

// 解析文件
func parseFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var conf Config
	err = UnMarshall(data, &conf)
	if err != nil {
		return
	}
	fmt.Printf("反序列化成功 conf: %#v\n port: %#v\n", conf, conf.ServerConfig.Port)
}

func main() {
	parseFile("config.ini")
}
