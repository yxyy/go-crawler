package router

import "github.com/gin-gonic/gin"

func InitRouter()  {
	r := gin.Default()

	r.Static("/static","/static")

	r.LoadHTMLGlob("views/***/**/*")

	//后台路由
	InitAdminRouter(r)
}