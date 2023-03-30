package model

type UserInfo struct {
	Name string
	Age  int
}

type Info struct {
	Age   int    `json:"jsonage" form:"formage" binding:"required" msg:"年龄必填"`
	Name  string `json:"jsonname" form:"formname" binding:"min=1,max=4,vaname" msg:"姓名最长四位"`
	Other map[string]string
}

type SysUser struct {
	// gorm.Model `json:"-"`
	Nick_name string  `json:"名称"`
	Phone     *string `json:"电话" gorm:"default:110"`
	Enable    int     `json:"状态"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
