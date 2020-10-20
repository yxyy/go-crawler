package router

import (
	"github.com/gin-gonic/gin"
	"lhc.go.crawler/controller/admin"
	"lhc.go.crawler/controller/admin/login"
	"lhc.go.crawler/controller/admin/user"
	"lhc.go.crawler/controller/admin/user_group"
	"lhc.go.crawler/controller/admin/menu"
)

func InitAdminRouter(r *gin.Engine)  {

	r.GET("/admin/login", login.Index)

	r.POST("admin/login",login.Login)

	AdminGroup := r.Group("/admin")
	{
		AdminGroup.GET("index",admin.Index)


		AdminGroup.GET("menu/index",menu.Index)
		AdminGroup.GET("menu/add",menu.Add)
		AdminGroup.POST("menu/UpdataMenu",menu.UpdateData)

		AdminGroup.GET("user/index",user.Index)
		AdminGroup.GET("user/add",user.Add)
		AdminGroup.POST("user/UpdataMenu",user.Add)

		AdminGroup.GET("user_group/index",user_group.Index)
		AdminGroup.POST("user_group/index",user_group.GetUserGourpList)
		AdminGroup.GET("user_group/add",user_group.Add)
		AdminGroup.POST("user_group/add",user_group.UpdateData)
	}


}