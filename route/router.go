package route

import (
	"blockchain-newsfeed-server/middleware"
	"blockchain-newsfeed-server/module/core/transport"
	"blockchain-newsfeed-server/token"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	engine *gin.Engine,
	trpt *transport.Transport,
	jwtMaker token.IJWTMaker,
) {
	v1Api := engine.Group("/v1")

	publicApi := v1Api.Group("/public")
	publicApi.POST("/register", trpt.Register)
	publicApi.POST("/login", trpt.Login)
	publicApi.POST("/refresh", trpt.Refresh)

	publicApi.Use(middleware.AuthMiddleware(jwtMaker))
	{
		customerApi := publicApi.Group("/customer")
		{
			customerApi.GET("/profile", trpt.GetCustomerProfile)
		}

		postApi := publicApi.Group("/post")
		{
			postApi.POST("", trpt.CreatePost)
		}
	}
}
