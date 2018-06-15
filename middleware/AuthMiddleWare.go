package middleware

import (
	"gopkg.in/gin-gonic/gin.v1"
	//"net/http"
)

/**
 * 中间件函数
 * 通过router里的配置,在这个函数里提前捕捉了 http请求
 * 优先处理一些对应函数前要做的事，比如说身份验证
 */
func AuthMiddleWare() gin.HandlerFunc {



	return func(context *gin.Context) {
		//context.String(http.StatusOK,"No!")
		//context.Abort()
		//中断返回 作为验证等情况
		//return
		//context.Set("request","client_request")
		context.Next()
	}
}