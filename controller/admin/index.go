package admin

import (
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/model"
	"net/http"
)

func Index(c *gin.Context)  {

	menuList := model.GetMeunTree(0)
	c.HTML(http.StatusOK,"admin/index.html",gin.H{"menuList":menuList})
}
