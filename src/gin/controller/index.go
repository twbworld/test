package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"test/src/gin/dao"
	"test/src/gin/model"
	"test/src/gin/service"
)

func Index(ctx *gin.Context) {

	//JWT===========================begin
	type MapClaims struct {
		UserName string `json:"user_name"`
		jwt.StandardClaims
	}

	claims := MapClaims{
		"twb",
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 5,
			Issuer:    "twb",
		},
	}

	mySigningKey := []byte("wositanweibiao")

	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	c, err := myToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	d, err := jwt.ParseWithClaims(c, &MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
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

// 绑定参数
func Getjson(ctx *gin.Context) {
	// /getjson?formname=twb&formage=12&arr[a]=2&arr[b]=5

	zjjData, _ := ctx.Get("type") //接收中间件传输的数据

	_zjjJson, _ := zjjData.(model.UserInfo) //需要断言才能取出里面具体值

	fmt.Println(_zjjJson.Name)

	var form model.Info

	err := ctx.ShouldBind(&form) //获取值
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "msg": service.GetError(err, &form)})
		return
	}
	arr := ctx.QueryMap("arr") //map获取(PostFormMap)
	form.Other = arr
	ctx.JSON(http.StatusOK, form) //{"jsonage": 12, "jsonname": "twb"}

}

func Data(ctx *gin.Context) {
	sdata := ctx.MustGet("session")
	fmt.Println("===========", sdata)

	data, _ := ctx.GetRawData()

	var j map[string]any
	_ = json.Unmarshal(data, &j)

	ctx.JSON(http.StatusOK, j)
}

// 接收文件
func Uploadfile(ctx *gin.Context) {
	formData, _ := ctx.MultipartForm()

	files := formData.File["upload[]"]

	for _, v := range files {
		ctx.SaveUploadedFile(v, "upload/"+v.Filename)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("成功接收%d个文件", len(files))})

}

func In(ctx *gin.Context) {
	str, _ := json.Marshal(gin.H{"X-Forwarded-For": ctx.Request.Header.Get("X-Forwarded-For"), "X-Real-Ip": ctx.Request.Header.Get("X-Real-Ip"), "Remote_addr": ctx.RemoteIP(), "go-ip": ctx.ClientIP()})
	a := ctx.Param("a")
	ctx.HTML(http.StatusOK, "index.html", gin.H{"ip": string(str), "Param": a})
}

// 文档: https://www.cnblogs.com/davis12/p/16365264.html
func Gorm(c *gin.Context) {


	var sysUsers []model.SysUser

	tx := dao.DB.Begin() //开启事务
	defer func() {
		if recover() != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return
	}

	res := dao.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Table("sys_users").Where("id = ?", 1).Find(&sysUsers)
	if res.RowsAffected == 0 || res.Error != nil {
		tx.Rollback()
		fmt.Println("数据出错")
		fmt.Println(res.Error)
		return
	}

	fmt.Println("=========================")
	fmt.Println("成功查询")
	json, _ := json.Marshal(sysUsers[0])
	fmt.Println(string(json))

	/**
	//插入
	u2 := sysUser{Nick_name: "twb", Enable: sysUsers[0].Enable, Phone: new(string)}
	result := dao.DB.Create(&u2)

	if result.Error != nil {
		tx.Rollback()
		fmt.Println("插入出错:")
		fmt.Println(result.Error)
		return
	}

	fmt.Println("成功插入")
	fmt.Println(result.RowsAffected)
	**/

	/**
	//更新
	// dao.DB.Table("sys_users").Where("id = ?", 6).Updates(&sysUser{Nick_name: ""}) //!!!!!!!!这个更新会失败,不能更新为零值字段(0, nil, "", false), 要使用map方式
	upUser := make(map[string]interface{})
	upUser["Nick_name"] = ""
	db.Table("sys_users").Where("id = ?", 6).Updates(&upUser)
	**/

	//删除
	// dao.DB.Where("id = ?", 5).Delete(&sysUser{})

	tx.Commit()

	c.JSON(http.StatusOK, string(json))
	fmt.Println("=========================" + strconv.FormatInt(time.Now().Unix(), 10))

}
