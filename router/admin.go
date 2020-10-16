package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAdminRouter(r *gin.Engine)  {

	r.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login/login.html",gin.H{})
	})
}