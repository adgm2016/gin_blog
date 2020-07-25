package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routev1 := r.Group("api/v1")
	{
		//用户模块的路由接口
		routev1.POST("user/add", v1.AddUser)
		routev1.GET("users", v1.GetUsers)
		routev1.PUT("user/:id", v1.EditUser)
		routev1.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口

		//文章模块的路由接口
	}
	r.Run(utils.HttpPort)
}
