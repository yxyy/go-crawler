package menu

import (
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/model"
	"net/http"
)

func Index(c *gin.Context)  {
	menuList := model.GetMeunTree(0)
	c.HTML(http.StatusOK,"menu/index.html",gin.H{"menuList":menuList})
}

func Add(c *gin.Context)  {
	//定义ico
	menu:=model.GetParentMenu(0)
	ico := []string{"fa-bar-chart-o", "fa-yen", "fa-gamepad", "fa-user", "fa-laptop", "fa-cogs", "fa-dollar", "fa-signal", "fa-align-left"}
	c.HTML(http.StatusOK,"menu/add.html",gin.H{"ico":ico,"menu":menu})
}

func UpdateData(c *gin.Context)  {
	var params model.Menu
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if params.Name=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"名称不为空"})
		return
	}
	if params.Url=="" {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":"url不能为空"})
		return
	}
	if err:=params.UpdateData();err!=nil{
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}