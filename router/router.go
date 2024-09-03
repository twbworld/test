package router

import (
	"net/http"

	"github.com/twbworld/dating/controller"
	"github.com/twbworld/dating/middleware"
	"github.com/twbworld/dating/model/common"

	"github.com/gin-gonic/gin"
)

func Start(ginServer *gin.Engine) {

	ginServer.Use(middleware.CorsHandle()) //全局中间件

	ginServer.StaticFile("/favicon.ico", "static/favicon.ico")
	ginServer.StaticFile("/robots.txt", "static/robots.txt")
	ginServer.LoadHTMLGlob("static/*.html")

	ginServer.NoRoute(func(ctx *gin.Context) {
		//内部重定向
		ctx.Request.URL.Path = "/404.html"
		ginServer.HandleContext(ctx)
		//http重定向
		// ctx.Redirect(http.StatusMovedPermanently, "/404.html")
	})
	ginServer.GET("404.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{"status": "404"})
	})
	ginServer.GET("40x.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{"status": "40x"})
	})
	ginServer.GET("50x.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{"status": "50x"})
	})
	ginServer.POST("404.html", func(ctx *gin.Context) {
		common.FailNotFound(ctx)
	})
	ginServer.POST("40x.html", func(ctx *gin.Context) {
		common.FailNotFound(ctx)
	})
	ginServer.POST("50x.html", func(ctx *gin.Context) {
		common.FailNotFound(ctx)
	})

	ginServer.POST("login", controller.Api.UserApiGroup.BaseApi.Login)
	ginServer.POST("userAdd", controller.Api.UserApiGroup.BaseApi.UserAdd)

	wh := ginServer.Group("wh")
	{
		wh.POST("/tg/:token", middleware.ValidatorTgToken, controller.Api.UserApiGroup.TgApi.Tg)
	}

	auth := ginServer.Use(middleware.JWTAuth)
	{
		auth.POST("getDatingAmount", controller.Api.UserApiGroup.DatingApi.GetDatingAmount)
		auth.POST("getDating", controller.Api.UserApiGroup.DatingApi.GetDating)
		auth.POST("joinDating", controller.Api.UserApiGroup.DatingApi.JoinDating)
		auth.POST("updateUserTime", controller.Api.UserApiGroup.DatingApi.UpdateUserTime)
		auth.POST("getDatingList", controller.Api.UserApiGroup.DatingApi.GetDatingList)
		auth.POST("quitDating", controller.Api.UserApiGroup.DatingApi.QuitDating)
		auth.POST("feedback", controller.Api.UserApiGroup.BaseApi.Feedback)
	}
}
