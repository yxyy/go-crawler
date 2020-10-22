package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/model"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"category/index.html",gin.H{})
}

func Add(c *gin.Context)  {
	c.HTML(http.StatusOK,"category/add.html",gin.H{"group_list":model.GetUserGourpList()})
}

func GetList(c *gin.Context)  {
	var params model.User
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	data, total, err := params.GetList()
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}

	fmt.Printf("%#v\n",data)

	c.JSON(http.StatusOK,gin.H{"code":200,"data":data,"recordsTotal":total,"recordsFiltered":total})
}

func UpdateData(c *gin.Context){
	var params model.User
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.UpdateData();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}

func Update(c *gin.Context){
	var params model.User
	if err:=c.ShouldBind(&params);err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	if err:=params.Update();err!=nil {
		c.JSON(http.StatusOK,gin.H{"code":400,"msg":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"msg":"操作成功","url":"index"})
}
