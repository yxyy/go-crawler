package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"user/index.html",gin.H{})
}

func Add(c *gin.Context)  {
	c.HTML(http.StatusOK,"user/add.html",gin.H{})
}