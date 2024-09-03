package db

type User struct {
	BaseField
	UserInfo
	PhoneNumber string `db:"phone_number" json:"phoneNumber" info:"手机号"`
	OpenId      string `db:"openid" json:"-" info:"小程序平台的用户识别码"`
	UnionId     string `db:"unionid" json:"-" info:"微信用户识别码"`
	SessionKey  string `db:"session_key" json:"-" info:"微信Session_Key"`
}

type UserInfo struct {
	NickName  string `db:"nick_name" json:"nick_name" info:"昵称"`
	AvatarUrl string `db:"avatar_url" json:"avatar_url" info:"头像"`
	Gender    int8   `db:"gender" json:"gender" info:"性别;0:未知;1:男;2:女"`
}

func (User) TableName() string {
	return `user`
}
