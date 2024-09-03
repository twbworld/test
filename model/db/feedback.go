package db

type Feedback struct {
	BaseField
	Desc    string `db:"desc" json:"desc" info:"反馈详情"`
	UserId  uint   `db:"user_id" json:"user_id" info:"反馈人"`
	MediaId string `db:"media_id" json:"media_id" info:"media表id, 逗号相隔"`
}

func (Feedback) TableName() string {
	return `feedback`
}
