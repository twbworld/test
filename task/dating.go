package task

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/twbworld/dating/dao"
)

const cleanDayAgo = 7 //清除?天前的数据

func Clean() (err error) {
	agoTime, ids := time.Now().AddDate(0, 0, -cleanDayAgo).Unix(), []uint{}
	if err = dao.App.DatingDb.GetCleanDating(&ids, agoTime); err != nil || len(ids) < 1 {
		return
	}

	return dao.Tx(func(tx *sqlx.Tx) (e error) {
		return dao.App.DatingDb.CloseDating(ids, tx)
	})

}
