package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/twbworld/dating/global"
)

// 跨域
func CorsHandle() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = global.Config.Cors
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authorization"}
	return cors.New(config)
}
