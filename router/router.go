package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	."github.com/melonws/goweb/apis"
	."github.com/melonws/goweb/middleware"
)

/**
 * 这个函数在main()中调用,相当于new了一个router对象,
 * 并且构造了一些属性,比如import了中间件,和apis
 * 把整体路由使用了AuthMiddleWare中间件,然后把 POST /person 指向了 AddPersonAPI这个handle函数
 * 可以在apis的person包里找到
 */
func InitRouter() *gin.Engine {

	router := gin.Default()

	//全局中间件
	router.Use(AuthMiddleWare())
	{
		//router.GET("/",IndexApi)

		router.POST("/person",AddPersonApi)
	}

	//群组
	authorized := router.Group("/",AuthMiddleWare())
	authorized.GET("/",IndexApi)


	return router
}

