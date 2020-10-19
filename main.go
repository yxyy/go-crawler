package main

import (
	"lhc.go.crawler/confs"
	"lhc.go.crawler/libs/mysql"
	"lhc.go.crawler/router"
	"log"
)

func main()  {

	//加载配置
	if err := confs.InitConfs();err!=nil{
		log.Fatal(err)
	}
	//加载mysql连接
	if err := mysql.InitMysql();err != nil {
		log.Fatal(err)
	}
	defer mysql.MysqlConnet.Close()

	//加载路由
	r := router.InitRouter()
	if err := r.Run(":8080");err!=nil{
		log.Fatal(err)
	}
}
