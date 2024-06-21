package main

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"Host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

type Config struct {
	ServerConfig `ini:"server"`
	MysqlConfig  `ini:"mysql"`
}
