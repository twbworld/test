package user

import (
	"github.com/gin-gonic/gin"

	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/model/common"
	"github.com/twbworld/dating/model/db"
	"github.com/twbworld/dating/service"
)

type BaseApi struct{}

// 登录注册
func (b *BaseApi) Login(ctx *gin.Context) {
	var (
		data common.LoginPost
		u    db.User
		err  error
	)

	defer func() {
		if p := recover(); p != nil {
			global.Log.Errorln(p)
			common.Fail(ctx, `系统错误[onsds]`)
		}
	}()

	if ctx.ShouldBindJSON(&data) != nil {
		common.Fail(ctx, `参数错误[ddssj]`)
		return
	}
	if err = service.Service.UserServiceGroup.Validator.ValidatorLoginPost(&data); err != nil {
		common.Fail(ctx, err.Error())
		return
	}

	if err = service.Service.UserServiceGroup.DatingService.GetUserByCode(data.Code, &u); err != nil {
		common.Fail(ctx, err.Error())
		return
	}

	token := ""
	if u.Id > 0 {
		token, err = service.Service.UserServiceGroup.BaseService.LoginToken(&u)
		if err != nil {
			global.Log.Error(err)
			common.Fail(ctx, `系统错误[ojhtgnds]`)
			return
		}
	}
	common.SuccessAuth(ctx, token, gin.H{
		`user`: u, //如果这没有user_id, 前端下一步将会请求UserAdd()进行注册
	})
}
