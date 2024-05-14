package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db        string
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbName    string
	UseUuid   string
	DbCharset string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadData(file *ini.File) {

	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3308")
	DbUser = file.Section("database").Key("Dbuser").MustString("root")
	DbPass = file.Section("database").Key("DbPassWord").MustString("Foxconn88")
	DbName = file.Section("database").Key("DbName").MustString("revenue_report")
	UseUuid = file.Section("database").Key("UseUUID").MustString("true")
	DbCharset = file.Section("database").Key("DbCharset").MustString("utf8mb4")

}
