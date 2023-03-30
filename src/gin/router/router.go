package router

import (
	"test/src/gin/controller"
	"test/src/gin/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init(ginServer *gin.Engine) {
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("vaname", validatorHandle)
	}
	ginServer.Use(validatorForm())
	//// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
	ginServer.Use(gin.Recovery())

	// ginServer.Use(favicon.New("/var/www/git/test/src/gin/favicon.ico"))
	ginServer.StaticFile("/favicon.ico", "/var/www/git/test/src/gin/favicon.ico")
	ginServer.LoadHTMLGlob("/var/www/git/test/src/gin/html/*")
	ginServer.Static("/static", "/var/www/git/test/src/gin/static")

	ginServer.GET("/index", controller.Index)
	ginServer.POST("/getjson", controller.Getjson)
	ginServer.POST("/data", validatorHanderData(), controller.Data)
	ginServer.POST("/uploadfile", controller.Uploadfile)
	ginServer.GET("/gorm", controller.Gorm)

	rougroup := ginServer.Group("r1").Use(validatorHanderGroup)
	{
		rougroup.GET("/in/:a/*a2", controller.In)
	}
}



// 拦截器(中间件)
func validatorHanderGroup(ctx *gin.Context) {
	ctx.Next()
}

func validatorForm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("type", model.UserInfo{Name: "姓名", Age: 12})
	}
}

func validatorHanderData() gin.HandlerFunc {
	return func(d *gin.Context) {
		d.Set("session", "man")
		d.Next() //中间件处理完后往下走,也可以使用Abort()终止
	}
}

func validatorHandle(fl validator.FieldLevel) bool {
	var allowName []string = []string{"root", "admin"}

	for _, v := range allowName {
		if v == fl.Field().Interface().(string) {
			return false
		}
	}

	return true
}
