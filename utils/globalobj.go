package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"zinx-copy/ziface"
)

// 存储一切与zinx有关的全局配置

type GlobalObj struct {

	// 当前server
	TcpServer ziface.IServer

	// 当前主机ip
	Host string

	// 当前端口
	TcpPort int

	// 当前服务器名称
	Name string

	// 当前zinx版本
	Version string

	// 当前服务器允许的最大链接数
	MaxConn int

	// 数据包最大值
	MaxPackageSize uint32
}

// Reload 从配置文件中加载用户自定义的配置
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("config/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

// GlobalObject 定义各个全局对象
var GlobalObject *GlobalObj

func init() {
	// 默认配置
	fmt.Println(os.Getwd())
	GlobalObject = &GlobalObj{
		Name:           "ZinxServer",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 应该从配置文件中加载用户配置
	GlobalObject.Reload()

}
