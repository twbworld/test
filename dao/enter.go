package dao

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

const (
	BaseUserId uint = 1 //"手动添加会面"用到的的虚拟用户
)

var (
	DB      *sqlx.DB
	CanLock bool     //是否支持锁(FOR UPDATE)
	App     *DbGroup = new(DbGroup)
	utils   *dbUtils = new(dbUtils)
)

type DbGroup struct {
	DatingDb
	UserDb
	FeedbackDb
}

func Tx(fc func(tx *sqlx.Tx) error) (err error) {
	panicked := true

	tx, err := DB.Beginx()
	if err != nil {
		err = errors.New("系统错误[jjhokp9]" + err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if panicked || err != nil {
			tx.Rollback()
		}
	}()

	if err = fc(tx); err == nil {
		panicked = false
		return tx.Commit()
	}

	panicked = false
	return
}
