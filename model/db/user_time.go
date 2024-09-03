package db

import "encoding/json"

type UserTime struct {
	BaseField
	DatingId uint   `db:"dating_id" json:"dating_id" info:"dating表id"`
	UserId   uint   `db:"user_id" json:"user_id" info:"user表id"`
	Info     string `db:"info" json:"info" info:"空闲时间信息;{'t': [[1706978785,1706978785]]}"`
	Status   int8   `db:"status" json:"status" info:"状态; 0:已退出;1:加入"`
}

func (UserTime) TableName() string {
	return `user_time`
}

// json转结构体
func (d *UserTime) InfoUnmarshal() *UserTimeInfo {
	result := &UserTimeInfo{
		[][]int64{},
	}
	if d.Info == "" {
		return result
	}
	var res UserTimeInfo
	if json.Unmarshal([]byte(d.Info), &res) != nil {
		return result
	}

	if len(res.Time) < 1 {
		//避免转json后,属性为"null"
		res.Time = make([][]int64, 0)
	}
	return &res
}

type UserTimeInfo struct {
	Time [][]int64 `json:"t"`
}

// 结构体转json
func (u *UserTimeInfo) Marshal() string {
	if u.Time == nil {
		//避免转json后,属性为"null"
		u.Time = make([][]int64, 0)
	}

	result, err := json.Marshal(u)
	if err != nil {
		return ``
	}
	return string(result)
}
