package common

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/twbworld/dating/model/db"
)

type InfoResponse struct {
	InfoPost
	TimeStr []string `json:"ts"`
}

type DatingUser struct {
	Id           uint         `db:"id" json:"id"`
	NickName     string       `db:"nick_name" json:"nick_name"`
	AvatarUrl    string       `db:"avatar_url" json:"avatar_url"`
	UtId         uint         `db:"ut_id" json:"ut_id"`
	Info         string       `db:"info" json:"-"`
	InfoResponse InfoResponse `json:"info"`
}

type DatingUserJoin struct {
	UserId       uint `db:"user_id" json:"user_id"`
	DatingId     uint `db:"dating_id" json:"dating_id"`
	CreateUserId uint `db:"create_user_id" json:"create_user_id"`
}

type DatingList struct {
	Id           uint     `db:"id" json:"id"`
	Status       int8     `db:"status" json:"status"`
	AddTime      int64    `db:"add_time" json:"-"`
	AddTimeStr   string   `json:"add_time"`
	UtId         uint     `db:"ut_id" json:"ut_id"`
	CreateUserId uint     `db:"create_user_id" json:"create_user_id"`
	AvatarUrl    []string `json:"avatar_url"`
}

type DatingUserAvatar struct {
	DatingId  uint   `db:"dating_id" json:"dating_id"`
	AvatarUrl string `db:"avatar_url" json:"avatar_url"`
}

type UserInfoWx struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int8   `json:"gender"`
	// Language  string `json:"language" info:""`
	// City      string `json:"city" info:""`
	// Province  string `json:"province" info:""`
	// Country   string `json:"country" info:""`
}

type JwtInfo struct {
	jwt.RegisteredClaims
	//可自定义数据; 加入IP地址/用户名等非敏感数据
}

// json转结构体
func (d *DatingUser) InfoUnmarshal() *db.UserTimeInfo {
	return (&db.UserTime{
		Info: d.Info,
	}).InfoUnmarshal()
}
