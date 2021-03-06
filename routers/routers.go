package routers

import (
	"bubble/controller"
	"bubble/setting"
	"github.com/gin-gonic/gin"

	_ "bubble/docs"  

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

		// 用户信息
		// 添加
		v1Group.POST("/user", controller.CreateUser)
		// 查看所有的待办事项
		v1Group.GET("/user", controller.GetUserList)
		// 修改某一个待办事项
		v1Group.PUT("/user/:id", controller.UpdateAUser)
		// 删除某一个待办事项
		v1Group.DELETE("/user/:id", controller.DeleteAUser)
	}
	return r
}
