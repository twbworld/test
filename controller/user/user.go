package user

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/twbworld/dating/dao"
	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/model/common"
	"github.com/twbworld/dating/model/db"
	"github.com/twbworld/dating/service"
)

// 用户注册
func (b *BaseApi) UserAdd(ctx *gin.Context) {
	var (
		data       common.UserInfoPost
		u          db.User
		userInfoWx common.UserInfoWx
		token      string
	)

	defer func() {
		if p := recover(); p != nil {
			global.Log.Errorln(p)
			common.Fail(ctx, `系统错误[oinds]`)
		}
	}()

	if ctx.ShouldBindJSON(&data) != nil {
		common.Fail(ctx, `参数错误[ddssj]`)
		return
	}
	if err := service.Service.UserServiceGroup.Validator.ValidatorUserAddPost(&data); err != nil {
		common.Fail(ctx, err.Error())
		return
	}

	if gin.Mode() == gin.TestMode {
		userInfoWx.AvatarUrl = "/static/logo.png"
		userInfoWx.Gender = 0
		userInfoWx.NickName = "测试"
		u.OpenId = ""
		u.UnionId = ""
		u.SessionKey = ""
	} else {

		// 下一步是数据解密和用户注册; 所以提前向微信官方获取解密要用到的SessionKey(如原本已存且不过期, 可省略这步直接本地解密)
		if err := service.Service.UserServiceGroup.DatingService.GetUserByCode(data.Code, &u); err != nil {
			panic(err)
		}
		if u.Id > 0 {
			common.Success(ctx, gin.H{
				"user": u,
			})
			return
		}
		if len(u.SessionKey) < 1 {
			panic("错误[sidf]")
		}

		//解密用户信息
		//(用户信息有可能变化, 可加上修改用户信息的逻辑)
		ed, cryptError := global.MiniProgramApp.Encryptor.DecryptData(data.EncryptedData, u.SessionKey, data.Iv)
		if cryptError != nil {
			panic(cryptError)
		}

		if err := json.Unmarshal(ed, &userInfoWx); err != nil {
			panic(err)
		}
	}

	if err := service.Service.UserServiceGroup.Validator.ValidatorEncryptedData(&userInfoWx); err != nil {
		common.Fail(ctx, err.Error())
		return
	}

	u.Id = 0
	u.NickName = userInfoWx.NickName
	u.AvatarUrl = userInfoWx.AvatarUrl
	u.Gender = userInfoWx.Gender

	err := dao.Tx(func(tx *sqlx.Tx) (e error) {
		id, e := dao.App.UserDb.AddUser(&u, tx)
		if e != nil {
			panic("[gdfjinf]" + e.Error())
		}

		if dao.App.UserDb.GetUserById(&u, id, tx) == sql.ErrNoRows {
			panic(`系统错误[o8s6nm]`)
		}

		token, e = service.Service.UserServiceGroup.BaseService.LoginToken(&u)

		return
	})
	if err != nil {
		panic(err)
	}

	common.SuccessAuth(ctx, token, gin.H{
		"user": u,
	})
}

// 用户反馈
func (b *BaseApi) Feedback(ctx *gin.Context) {
	var (
		data common.FeedbackPost
	)

	defer func() {
		if p := recover(); p != nil {
			global.Log.Errorln(p)
			common.Fail(ctx, `系统错误[oin7ds]`)
		}
	}()

	userId := ctx.MustGet(`userId`).(uint)
	if userId < 1 {
		common.Fail(ctx, `系统错误[th9pi]`)
		return
	}

	if userId == dao.BaseUserId {
		common.Fail(ctx, `参数错误[6f7j]`)
		return
	}

	if ctx.ShouldBindJSON(&data) != nil {
		common.Fail(ctx, `参数错误[dds6sj]`)
		return
	}
	if err := service.Service.UserServiceGroup.Validator.ValidatorFeedbackPost(&data); err != nil {
		common.Fail(ctx, err.Error())
		return
	}

	err := dao.Tx(func(tx *sqlx.Tx) (e error) {
		id, e := dao.App.FeedbackDb.AddFeedback(&db.Feedback{
			Desc:   data.Desc,
			UserId: userId,
		}, tx)
		if id < 1 {
			panic(`系统错误[ond0sm]`)
		}
		return
	})
	if err != nil {
		panic(err)
	}

	go service.Service.UserServiceGroup.TgService.TgSend(fmt.Sprintf("用户反馈通知:\n%s", data.Desc))

	common.SuccessOk(ctx, `成功`)
}
