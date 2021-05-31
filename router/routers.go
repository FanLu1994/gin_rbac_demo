package router

import (
	"gin_rbac/controller"
	_ "gin_rbac/controller"
	"gin_rbac/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine{
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode("debug")


	authApi:=router.Group("auth")
	{
		authApi.POST("/login",controller.GetAuth)
	}


	apis := router.Group("api")
	apis.Use(middleware.JWT())
	{
		apis.GET("/getBook",controller.GetBook)
		apis.GET("/ping",controller.Ping)
	}

	return router
}