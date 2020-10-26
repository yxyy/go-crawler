package main

import (
	"fmt"
	"lhc.go.crawler/confs"
	"lhc.go.crawler/libs/mysql"
	"lhc.go.crawler/libs/es"
	"lhc.go.crawler/router"
	"lhc.go.crawler/logs"
)

func main()  {

	fmt.Println("加载日志组件.....")
	if err:=logs.InitLogs();err!=nil {
		logs.Fatal.Fatal(err)
	}
	logs.Error.Info("log init")

	//加载配置
	fmt.Println("加载配置组件.....")
	if err := confs.InitConfs();err!=nil{
		logs.Fatal.Fatal(err)
	}
	//加载mysql连接
	fmt.Println("加载mysql连接组件.....")
	if err := mysql.InitMysql();err != nil {
		logs.Fatal.Fatal(err)
	}
	defer mysql.MysqlConnet.Close()

	//加载mysql连接
	fmt.Println("加载es连接组件.....")
	if err := es.InitEs();err != nil {
		logs.Fatal.Fatal(err)
	}

	//加载路由
	fmt.Println("加载路由组件.....")
	r := router.InitRouter()
	if err := r.Run(":8080");err!=nil{
		logs.Fatal.Fatal(err)
	}
}
