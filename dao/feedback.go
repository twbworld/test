package dao

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/twbworld/dating/model/db"
)

type FeedbackDb struct{}

func (u *FeedbackDb) AddFeedback(f *db.Feedback, tx *sqlx.Tx) (uint, error) {
	if tx == nil {
		return 0, errors.New(`请使用事务[iodhja]`)
	}
	sql, args := utils.getInsertSql(f, map[string]interface{}{
		"user_id":     f.UserId,
		"desc":        f.Desc,
		"add_time":    time.Now().Unix(),
		"update_time": time.Now().Unix(),
	})

	res, err := tx.Exec(sql, args...)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return uint(id), err
}
