package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


func main() {

    // 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()
    fi, _ := os.Create("gin.log") //记录到文件。
    gin.DefaultWriter = io.MultiWriter(fi)









	ginServer := gin.Default()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("vaname", validatorHandle)
	}
    ginServer.Use(hander())
	//// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
    ginServer.Use(gin.Recovery())

	// ginServer.Use(favicon.New("/var/www/git/test/src/gin/favicon.ico"))
	ginServer.StaticFile("/favicon.ico", "/var/www/git/test/src/gin/favicon.ico")
	ginServer.LoadHTMLGlob("/var/www/git/test/src/gin/html/*")
	ginServer.Static("/static", "/var/www/git/test/src/gin/static")

	ginServer.GET("/index", index)
	ginServer.POST("/getjson", getjson)
    ginServer.POST("/data", handerData(), data)
    ginServer.POST("/uploadfile", uploadfile)
    ginServer.POST("/gorm", Gorm)

	rougroup := ginServer.Group("r1").Use(handerGroup)
	{
	    rougroup.GET("/in/:a/*a2", in)
	}

	// ginServer.Run(":8081")










	//平滑优雅关闭服务======================================begin
	server := http.Server{Addr: ":8081", Handler: ginServer}

	//协程启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen err:%s", err)
		}
	}()


	sb := make(chan os.Signal, 1)
	signal.Notify(sb, syscall.SIGINT, syscall.SIGTERM) //监听关闭(ctrl+C)指令
	<-sb //阻塞等待

	//来到这 证明有关闭指令
	c, f := context.WithTimeout(context.Background(), 5 * time.Second) //如果有连接就超时5s后关闭
	defer f() //关闭当前请求

	//关闭监听端口
	if err := server.Shutdown(c); nil != err {
		log.Fatal("server listen err:%s", err)
	}
	log.Println("server exiting...")
	//平滑优雅关闭服务======================================end

}

type UserInfo struct{
    Name string
    Age int
}

// 拦截器(中间件)
func handerGroup(ctx *gin.Context){
    ctx.Next()
}
func hander() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Set("type", UserInfo{Name: "姓名", Age: 12})
    }
}
func handerData() gin.HandlerFunc {
	return func(d *gin.Context) {
		d.Set("session", "man")
		d.Next() //中间件处理完后往下走,也可以使用Abort()终止
	}
}


func in(ctx *gin.Context) {
	str, _ := json.Marshal(gin.H{"X-Forwarded-For": ctx.Request.Header.Get("X-Forwarded-For"), "X-Real-Ip": ctx.Request.Header.Get("X-Real-Ip"), "Remote_addr": ctx.RemoteIP(), "go-ip": ctx.ClientIP()})
	a := ctx.Param("a")
	ctx.HTML(http.StatusOK, "index.html", gin.H{"ip": string(str), "Param": a})
}


func index(ctx *gin.Context) {





	//JWT===========================begin
	type MapClaims struct{
		UserName string `json:"user_name"`
		jwt.StandardClaims
	}


	claims := MapClaims{
		"twb",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 5,
			Issuer: "twb",
		},
	}

	mySigningKey := []byte("wositanweibiao")

	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	c, err := myToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)


	d, err := jwt.ParseWithClaims(c, &MapClaims{}, func(token *jwt.Token) (interface{}, error){
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(d.Claims.(*MapClaims).UserName)

	}
	//JWT===========================end




	cookie, err := ctx.Cookie("gin_cookie")
	if err != nil || cookie == "" {
		ctx.SetCookie("gin_cookie", c, 3600, "/", "go.cc.cc", false, true)
	}


	page := ctx.DefaultQuery("id", "0") //取uri值(PostForm取post数据)
	aa := map[string]interface{}{
		"c": page,
	}

	// time.Sleep(5 * time.Second)


	ctx.JSONP(http.StatusOK, aa)
	bb := []string{"a", "v"}


	ctx.SecureJSON(http.StatusOK, bb)


}

func data(ctx *gin.Context) {
	sdata := ctx.MustGet("session")
	fmt.Println("===========", sdata)

	data, _ := ctx.GetRawData()

	var j map[string]any
	_ = json.Unmarshal(data, &j)

	ctx.JSON(http.StatusOK, j)
}

type Info struct {
	Age  int    `json:"jsonage" form:"formage" binding:"required" msg:"年龄必填"`
	Name string `json:"jsonname" form:"formname" binding:"min=1,max=4,vaname" msg:"姓名最长四位"`
	Other map[string]string
}

func GetError(err error, form *Info) string {
	//断言.判断是否为err类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		formType := reflect.TypeOf(form)
		for _, v := range errs {
			if StructField, ok := formType.Elem().FieldByName(v.Field()); ok {
				return StructField.Tag.Get("msg")
			}
		}

	}
	return err.Error()
}

//绑定参数
func getjson(ctx *gin.Context) {
	// /getjson?formname=twb&formage=12&arr[a]=2&arr[b]=5

    zjjData, _ := ctx.Get("type") //接收中间件传输的数据

    _zjjJson, _ := zjjData.(UserInfo) //需要断言才能取出里面具体值

    fmt.Println(_zjjJson.Name)

	var form Info

	err := ctx.ShouldBind(&form) //获取值
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": GetError(err, &form)})
		return
	}
	arr := ctx.QueryMap("arr") //map获取(PostFormMap)
	form.Other = arr
	ctx.JSON(http.StatusOK, form) //{"jsonage": 12, "jsonname": "twb"}

}

//接收文件
func uploadfile(ctx *gin.Context) {
    formData, _ := ctx.MultipartForm()

    files := formData.File["upload[]"]

    for _, v := range files{
        ctx.SaveUploadedFile(v, "upload/" + v.Filename)
    }

    ctx.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("成功接收%d个文件", len(files))})

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



func Gorm(ctx *gin.Context) {


}
