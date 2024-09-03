package main

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/twbworld/dating/global"
	initGlobal "github.com/twbworld/dating/initialize/global"
	"github.com/twbworld/dating/initialize/system"
	"github.com/twbworld/dating/model/common"
	"github.com/twbworld/dating/model/db"
	"github.com/twbworld/dating/router"
	"github.com/twbworld/dating/service"
)

func TestMain(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initGlobal.Start()
	system.DbStart()
	defer func ()  {
		time.Sleep(time.Second * 1) //给足够时间处理数据
		system.DbClose()
	}()


	ginServer := gin.Default()
	router.Start(ginServer)

	n := time.Now()
	ni := n.In(global.Tz)

	u := db.User{}
	u.Id = 2 //测试用户
	token, err := service.Service.UserServiceGroup.BaseService.LoginToken(&u)
	if err != nil {
		t.Fatal("jwt错误[isodfji]", err)
	}

	//以下是有执行顺序的, 并且库提前有必要数据
	testCases := [...]struct {
		url      string
		status   int
		postData interface{}
		res      common.Response
	}{
		{url: "/login", postData: common.LoginPost{Code: "aaaaaa"}, res: common.Response{Code: 1}},
		{url: "/userAdd", postData: common.UserInfoPost{
			LoginPost:     common.LoginPost{Code: "aaaaaa"},
			EncryptedData: "aaaaaa",
			Iv:            "aaaaaa",
		}},
		{url: "/userAdd", postData: common.UserInfoPost{
			LoginPost:     common.LoginPost{Code: "aaaaaa"},
			EncryptedData: "aaaaaa",
			Iv:            "aaaaaa",
		}},
		//加入会面(库已存在id为1创建者为2的会面)
		{url: "/joinDating", postData: common.DatingPost{
			Id: 1,
			Info: common.InfoPost{Time: [][]string{
				{ni.Format(time.DateOnly) + ` 9:00:00`, ni.Format(time.DateOnly) + ` 23:00:00`},
				{time.Now().In(global.Tz).AddDate(0, 0, 5).Format(time.DateOnly) + ` 11:00:00`, time.Now().In(global.Tz).AddDate(0, 0, 5).Format(time.DateOnly) + ` 14:00:00`},
			}},
		}},
		//创建会面
		{url: "/joinDating", postData: common.DatingPost{
			Info: common.InfoPost{Time: [][]string{
				{ni.Format(time.DateOnly) + ` 10:00:00`, ni.Format(time.DateOnly) + ` 22:00:00`},
				{time.Now().In(global.Tz).AddDate(0, 0, 10).Format(time.DateOnly) + ` 11:00:00`, time.Now().In(global.Tz).AddDate(0, 0, 10).Format(time.DateOnly) + ` 14:00:00`},
			}},
		}},
		//加入会面(手动虚拟)
		{url: "/joinDating", postData: common.DatingPost{
			Id: 2,
			Info: common.InfoPost{Time: [][]string{
				{ni.Format(time.DateOnly) + ` 10:00:00`, ni.Format(time.DateOnly) + ` 20:00:00`},
				{time.Now().In(global.Tz).AddDate(0, 0, 20).Format(time.DateOnly) + ` 13:00:00`, time.Now().In(global.Tz).AddDate(0, 0, 20).Format(time.DateOnly) + ` 14:00:00`},
			}},
		}},
		{url: "/getDatingList", postData: common.GetDatingListPost{Page: 1, LastId: 0}},
		{url: "/getDating", postData: common.GetDatingPost{Id: 1}},
		{url: "/getDatingAmount", postData: ""},
		//退出会面(失败)
		{url: "/quitDating", postData: common.QuitDatingPost{UtId: 1}, res: common.Response{Code: 1}},
		//退出会面
		{url: "/quitDating", postData: common.QuitDatingPost{Id: 1}},
		//退出虚拟用户
		{url: "/quitDating", postData: common.QuitDatingPost{UtId: 7}},
		//关闭会面
		{url: "/quitDating", postData: common.QuitDatingPost{Id: 2}},
		{url: "/feedback", postData: common.FeedbackPost{Desc: ""}, res: common.Response{Code: 1}},
		{url: "/feedback", postData: common.FeedbackPost{Desc: "测试"}},
	}

	t.Run("Match", func(t *testing.T) {
		defer func() {
			if p := recover(); p != nil {
				t.Fatal("[52mkwer]", p)
			}
		}()

		service.Service.UserServiceGroup.DatingService.Match(1)
	})
	t.Run("MatchFatal", func(t *testing.T) {

		defer func() {
			if p := recover(); p == nil || !strings.Contains(p.(string), "hk942") {
				t.Fatal("[er395i2]", p)
			}
		}()

		service.Service.UserServiceGroup.DatingService.Match(2)
	})

	for k, value := range testCases {
		t.Run(value.url+strconv.FormatInt(int64(k), 10), func(t *testing.T) {

			jsonVal, err := json.Marshal(value.postData)
			if err != nil {
				t.Fatal("json出错[godjg]", err)
			}

			//向注册的路有发起请求
			req, err := http.NewRequest(http.MethodPost, value.url, bytes.NewBuffer(jsonVal))
			if err != nil {
				t.Fatal("请求出错[godkojg]", err)
			}
			req.Header.Set("Authorization", token)
			req.Header.Set("content-type", "application/json")

			res := httptest.NewRecorder() // 构造一个记录
			ginServer.ServeHTTP(res, req) //模拟http服务处理请求

			result := res.Result() //response响应

			status := 200
			if value.status != 0 {
				status = value.status
			}
			assert.Equal(t, status, result.StatusCode)

			body, err := io.ReadAll(result.Body)
			if err != nil {
				t.Fatal(err)
			}
			defer result.Body.Close()

			// fmt.Println("request!!!!!!!!!!", string(jsonVal))
			// fmt.Println("response!!!!!!!!!!", string(body))
			var response common.Response
			if err := json.Unmarshal(body, &response); err != nil {
				t.Fatal("返回错误", err, string(body))
			}

			code := int8(0)
			if value.res != (common.Response{}) {
				code = value.res.Code
			}

			assert.Equal(t, code, response.Code)
			time.Sleep(time.Second)

		})

	}

}
