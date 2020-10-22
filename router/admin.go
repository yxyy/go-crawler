package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lhc.go.crawler/controller/admin"
	"lhc.go.crawler/controller/admin/category"
	"lhc.go.crawler/controller/admin/login"
	"lhc.go.crawler/controller/admin/user"
	"lhc.go.crawler/controller/admin/user_group"
	"lhc.go.crawler/controller/admin/menu"
	"lhc.go.crawler/middleware"
)

func InitAdminRouter(r *gin.Engine)  {

	//设置存储引擎，
	store := cookie.NewStore([]byte(viper.GetString("session_secret")))
	// 路由使用session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("yxyy.session", store))

	r.GET("/admin/login", login.Index)
	r.POST("admin/login",login.Login)

	AdminGroup := r.Group("/admin").Use(middleware.AdminSessionAuth())
	{
		AdminGroup.GET("logout",login.Logout)
		AdminGroup.GET("index",admin.Index)
		AdminGroup.GET("welcome",admin.Welcome)

		AdminGroup.GET("menu/index",menu.Index)
		AdminGroup.GET("menu/add",menu.Add)
		AdminGroup.POST("menu/UpdataMenu",menu.UpdateData)

		AdminGroup.GET("user/index",user.Index)
		AdminGroup.POST("user/index",user.GetList)
		AdminGroup.GET("user/add",user.Add)
		AdminGroup.POST("user/updateData",user.UpdateData)
		AdminGroup.POST("user/password",user.Update)

		AdminGroup.GET("user_group/index",user_group.Index)
		AdminGroup.POST("user_group/index",user_group.GetUserGourpList)
		AdminGroup.GET("user_group/add",user_group.Add)
		AdminGroup.POST("user_group/add",user_group.UpdateData)

		AdminGroup.GET("category/index",category.Index)
	}


}