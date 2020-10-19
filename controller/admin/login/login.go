package login

import (
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/model"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"admin/login/index.html",gin.H{})
}

func Login(c *gin.Context)  {
	var params model.User
	if err := c.ShouldBind(&params);err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"服务器繁忙，请稍候重试！"})
		return
	}
	if params.UserName=="" {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"账号错误"})
		return
	}
	if params.PassWord=="" {
		c.JSON(http.StatusOK,gin.H{"code":500,"msg":"密码错误"})
		return
	}
}