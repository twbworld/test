package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twbworld/dating/global"
)

// 验证TG的token
func ValidatorTgToken(ctx *gin.Context) {
	token := ctx.Param("token")
	if !global.Config.Debug && global.Config.Telegram.Token == token {
		//业务前执行
		ctx.Next()
		//业务后执行
		return
	}

	ctx.Abort()

	if global.Config.Telegram.Token != token {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		errMsg, _ := json.Marshal(map[string]string{"error": "Lack of token"})
		_, _ = ctx.Writer.Write(errMsg)
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "/404.html")
}
