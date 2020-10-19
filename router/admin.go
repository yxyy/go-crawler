package router

import (
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/controller/admin/login"
)

func InitAdminRouter(r *gin.Engine)  {

	r.GET("/admin/login", login.Index)

	r.POST("admin/login",login.Login)
}