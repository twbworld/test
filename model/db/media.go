package db

// (实质上表并没有update_time字段)
type Media struct {
	BaseField
	Path string `db:"path" json:"path" info:"文件路径"`
	Type string `db:"type" json:"type" info:"文件类型,如jpg/mp4等"`
}

func (Media) TableName() string {
	return `media`
}
