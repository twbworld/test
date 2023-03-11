package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
    // REMOTE_ADDR
    // X-FORWARDED_FOR
    // X-REAL_IP

    ginServer := gin.Default()


    ginServer.GET("/hi", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{"Remote_addr": ctx.Request.Header.Get("Remote_addr"), "X-Appengine-Remote-Addr": ctx.Request.Header.Get("X-Appengine-Remote-Addr"), "X-Forwarded-For": ctx.Request.Header.Get("X-Forwarded-For"), "X-Real-Ip": ctx.Request.Header.Get("X-Real-Ip")})
    })

    ginServer.Run(":8080")

}
