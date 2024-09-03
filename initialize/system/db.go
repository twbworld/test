package system

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twbworld/dating/dao"
	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/model/db"
	"github.com/twbworld/dating/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type mysql struct{}
type sqlite struct{}
type class interface {
	connect()
	createTable()
	insertData(string, *sqlx.Tx)
	version() string
}

func DbStart() {
	var dbRes class

	switch global.Config.Database.Type {
	case "mysql":
		dbRes = &mysql{}
	case "sqlite":
		dbRes = &sqlite{}
	default:
		dbRes = &sqlite{}
	}

	dbRes.connect()
	dbRes.createTable()
}

func DbClose() {
	dao.DB.Close()
}

func (s *sqlite) connect() {
	var err error

	if dao.DB, err = sqlx.Open("sqlite3", global.Config.Database.SqlitePath); err != nil {
		panic("数据库连接失败[jgadsfgas]: " + err.Error())
	}
	//没有数据库会创建
	if err = dao.DB.Ping(); err != nil {
		panic(fmt.Sprintf("数据库连接失败[khdsfgs]: %s\n%s", global.Config.Database.SqlitePath, err.Error()))
	}

	dao.DB.SetMaxIdleConns(10)
	dao.DB.SetMaxOpenConns(20)

	dao.CanLock = false

	global.Log.Infof("%s版本: %s; 地址: %s", global.Config.Database.Type, s.version(), global.Config.Database.SqlitePath)
}

func (m *mysql) connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.Config.Database.MysqlUsername, global.Config.Database.MysqlPassword, global.Config.Database.MysqlHost, global.Config.Database.MysqlPort, global.Config.Database.MysqlDbname)

	//也可以使用MustConnect连接不成功就panic
	if dao.DB, err = sqlx.Connect("mysql", dsn); err != nil {
		panic(fmt.Sprintf("数据库连接失败[ujefaf]: %s\n%s", dsn, err.Error()))
	}

	dao.DB.SetMaxOpenConns(16)
	dao.DB.SetMaxIdleConns(8)

	if err = dao.DB.Ping(); err != nil {
		panic(fmt.Sprintf("数据库连接失败[ujefddsaf]: %s\n%s", dsn, err.Error()))
	}

	dao.CanLock = true
	global.Log.Infof("%s版本: %s; 地址: @tcp(%s:%s)/%s", global.Config.Database.Type, m.version(), global.Config.Database.MysqlHost, global.Config.Database.MysqlPort, global.Config.Database.MysqlDbname)
}

func (s *sqlite) createTable() {
	var u []string
	err := dao.DB.Select(&u, "SELECT name _id FROM sqlite_master WHERE type ='table'")
	if err != nil {
		panic(err)
	}

	sqls := map[string][]string{
		db.Dating{}.TableName(): {
			`CREATE TABLE "%s" ("id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "create_user_id" INTEGER NOT NULL DEFAULT '', "result" TEXT NOT NULL DEFAULT '', "status" INTEGER(1) NOT NULL DEFAULT 0, "add_time" TEXT(10) NOT NULL DEFAULT '', "update_time" TEXT(10) NOT NULL DEFAULT '');`,
			`CREATE INDEX "idx_create_user_id" ON "%s" ("create_user_id" ASC);`,
		},
		db.User{}.TableName(): {
			`CREATE TABLE "%s" ("id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "nick_name" TEXT(64) NOT NULL DEFAULT '', "avatar_url" TEXT NOT NULL DEFAULT '', "gender" INTEGER(1) NOT NULL DEFAULT 0, "phone_number" TEXT(20) NOT NULL DEFAULT '', "openid" TEXT NOT NULL DEFAULT '', "unionid" TEXT NOT NULL DEFAULT '', "session_key" TEXT NOT NULL DEFAULT '', "add_time" TEXT(10) NOT NULL DEFAULT '', "update_time" TEXT(10) NOT NULL DEFAULT '');`,
			`CREATE INDEX "idx_openid" ON "%s" ("openid");`,
			`CREATE INDEX "idx_unionid" ON "%s" ("unionid");`,
		},
		db.UserTime{}.TableName(): {
			`CREATE TABLE "%s" ("id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "dating_id" INTEGER NOT NULL , "user_id" INTEGER NOT NULL, "info" TEXT NOT NULL DEFAULT '', "status" INTEGER(1) NOT NULL DEFAULT 0, "add_time" TEXT(10) NOT NULL DEFAULT '', "update_time" TEXT(10) NOT NULL DEFAULT '');`,
			`CREATE INDEX "idx_dating_id" ON "%s" ("dating_id" ASC);`,
			`CREATE INDEX "idx_user_id" ON "%s" ("user_id" ASC);`,
		},
		db.Feedback{}.TableName(): {
			`CREATE TABLE "%s" ("id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "user_id" INTEGER NOT NULL, "desc" TEXT(255) NOT NULL DEFAULT '', "media_id" TEXT(255) NOT NULL DEFAULT '', "add_time" TEXT(10) NOT NULL DEFAULT '', "update_time" TEXT(10) NOT NULL DEFAULT '');`,
			`CREATE INDEX "idx_feedback_user_id" ON "%s" ("user_id" ASC);`,
		},
		db.Media{}.TableName(): {
			`CREATE TABLE "%s" ("id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "path" TEXT NOT NULL DEFAULT '', "type" TEXT(10) NOT NULL DEFAULT '', "add_time" TEXT(10) NOT NULL DEFAULT '');`,
		},
	}

	err = dao.Tx(func(tx *sqlx.Tx) (e error) {
		for k, v := range sqls {
			if utils.InSlice(u, k) < 0 {
				for _, val := range v {
					if _, e := tx.Exec(fmt.Sprintf(val, k)); e != nil {
						panic(fmt.Sprintf("错误[ghjbcvgs]:  %s\n%s", val, e.Error()))
					}
				}
				s.insertData(k, tx)
				global.Log.Infof("创建%s表[dkyjh]", k)
			}
		}
		return
	})
	if err != nil {
		panic(err)
	}

}

func (m *mysql) createTable() {
	var u []string
	err := dao.DB.Select(&u, "SHOW TABLES")
	if err != nil {
		panic("错误[hsfds]: " + err.Error())
	}

	sqls := map[string]string{
		db.Dating{}.TableName():   "CREATE TABLE `%s` (`id` int unsigned NOT NULL AUTO_INCREMENT, `create_user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '会面创建者', `result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '推荐结果', `status` tinyint NOT NULL DEFAULT '0' COMMENT '会面状态; 0:结束;1:进行中;', `add_time` int unsigned NOT NULL DEFAULT '0', `update_time` int unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`), KEY `idx_create_user_id` (`create_user_id`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='会面表';",
		db.User{}.TableName():     "CREATE TABLE `%s` (`id` int unsigned NOT NULL AUTO_INCREMENT, `nick_name` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称', `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像', `gender` tinyint NOT NULL DEFAULT '0' COMMENT '性别;0:未知;1:男;2:女', `phone_number` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号', `openid` varchar(255) NOT NULL DEFAULT '' COMMENT '小程序平台的用户识别码', `unionid` varchar(255) NOT NULL DEFAULT '' COMMENT '微信用户识别码', `session_key` varchar(255) NOT NULL DEFAULT '' COMMENT '微信Session_Key', `add_time` int unsigned NOT NULL DEFAULT '0', `update_time` int unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`), KEY `idx_openid` (`openid`), KEY `idx_unionid` (`unionid`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';",
		db.UserTime{}.TableName(): "CREATE TABLE `%s` (`id` int unsigned NOT NULL AUTO_INCREMENT, `dating_id` int unsigned NOT NULL DEFAULT '0' COMMENT 'dating表id', `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT 'user表id', `info` text NOT NULL COMMENT '空闲时间信息;{''t'': [[1706978785,1706978785]]}', `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态; 0:已退出;1:加入', `add_time` int unsigned NOT NULL DEFAULT '0', `update_time` int unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`), KEY `idx_dating_id` (`dating_id`), KEY `idx_user_id` (`user_id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户时间表';",
		db.Feedback{}.TableName(): "CREATE TABLE `%s` (`id` int unsigned NOT NULL AUTO_INCREMENT, `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '反馈人', `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '反馈详情', `media_id` varchar(255) NOT NULL DEFAULT '' COMMENT 'media表id, 逗号相隔', `add_time` int unsigned NOT NULL DEFAULT '0', `update_time` int unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`), KEY `idx_user_id` (`user_id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='反馈信息表';",
		db.Media{}.TableName():    "CREATE TABLE `%s` (`id` int unsigned NOT NULL AUTO_INCREMENT, `path` varchar(255) NOT NULL DEFAULT '' COMMENT '文件路径', `type` varchar(10) NOT NULL DEFAULT '' COMMENT '文件类型,如jpg/mp4等', `add_time` int unsigned NOT NULL DEFAULT '0', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='媒体文件表';",
	}

	err = dao.Tx(func(tx *sqlx.Tx) (e error) {
		for k, v := range sqls {
			if utils.InSlice(u, k) < 0 {
				if _, e := tx.Exec(fmt.Sprintf(v, k)); e != nil {
					panic(fmt.Sprintf("错误[ghjdggs]:  %s\n%s", k, e.Error()))
				}
				global.Log.Infof("创建%s表[dfsjh]", k)
				m.insertData(k, tx)
			}
		}
		return
	})
	if err != nil {
		panic(err)
	}

}

func (m *mysql) insertData(t string, tx *sqlx.Tx) {
	insert(t, tx)
}

func (m *sqlite) insertData(t string, tx *sqlx.Tx) {
	insert(t, tx)
}

func insert(t string, tx *sqlx.Tx) {

	now := time.Now()
	ti := now.Unix()

	var sqls []string

	switch t {
	case db.Dating{}.TableName():
		if gin.Mode() == gin.TestMode {
			ti2 := now.AddDate(0, 0, -1).Unix()
			sqls = []string{
				fmt.Sprintf("INSERT INTO `%s`(`create_user_id`, `status`, `result`, `add_time`, `update_time`) VALUES(2, 1, '', %d, %d)", db.Dating{}.TableName(), ti2, ti2),
			}
		} else if gin.Mode() == gin.DebugMode {
			ti2 := now.AddDate(0, 0, -1).Unix()
			sqls = []string{
				fmt.Sprintf("INSERT INTO `%s`(`create_user_id`, `status`, `result`, `add_time`, `update_time`) VALUES(2, 1, '', %d, %d)", db.Dating{}.TableName(), ti2, ti2),
				fmt.Sprintf("INSERT INTO `%s`(`create_user_id`, `status`, `result`, `add_time`, `update_time`) VALUES(3, 1, '', %d, %d)", db.Dating{}.TableName(), ti2, ti2),
			}
		}
	case db.User{}.TableName():
		n := db.User{}.TableName()
		sqls = []string{
			fmt.Sprintf("INSERT INTO `%s`(`id`, `nick_name`, `avatar_url`, `add_time`, `update_time`) VALUES(%d, '手动', '/static/logo.png', %d, %d)", n, dao.BaseUserId, ti, ti),
		}
		if gin.Mode() == gin.TestMode {
		} else if gin.Mode() == gin.DebugMode {
			sqls = append(sqls,
				fmt.Sprintf("INSERT INTO `%s`(`nick_name`, `avatar_url`, `add_time`, `update_time`) VALUES('test', '/static/logo.png', %d, %d)", n, ti, ti),
				fmt.Sprintf("INSERT INTO `%s`(`nick_name`, `avatar_url`, `add_time`, `update_time`) VALUES('test2', '/static/logo.png', %d, %d)", n, ti, ti),
			)
		}
	case db.UserTime{}.TableName():
		if gin.Mode() == gin.TestMode {
			n := db.UserTime{}.TableName()
			t1, _ := time.ParseInLocation("2006-01-02 15", "2024-02-15 09", global.Tz)
			t2, _ := time.ParseInLocation("2006-01-02 15", "2024-02-15 22", global.Tz)
			t3, _ := time.ParseInLocation("2006-01-02 15", "2024-02-18 08", global.Tz)
			t4, _ := time.ParseInLocation("2006-01-02 15", "2024-02-18 23", global.Tz)
			a, _ := json.Marshal(&[...][2]int64{{t1.Unix(), t2.Unix()}, {t3.Unix(), t4.Unix()}})

			t5, _ := time.ParseInLocation("2006-01-02 15", "2024-02-15 08", global.Tz)
			t6, _ := time.ParseInLocation("2006-01-02 15", "2024-02-15 21", global.Tz)
			b, _ := json.Marshal(&[...][2]int64{{t5.Unix(), t6.Unix()}})

			t7, _ := time.ParseInLocation("2006-01-02 15", "2024-02-18 08", global.Tz)
			t8, _ := time.ParseInLocation("2006-01-02 15", "2024-02-18 23", global.Tz)
			t9, _ := time.ParseInLocation("2006-01-02 15", "2024-02-19 10", global.Tz)
			t10, _ := time.ParseInLocation("2006-01-02 15", "2024-02-19 22", global.Tz)
			c, _ := json.Marshal(&[...][2]int64{{t7.Unix(), t8.Unix()}, {t9.Unix(), t10.Unix()}})

			t11, _ := time.ParseInLocation("2006-01-02 15", "2024-02-19 22", global.Tz)
			t12, _ := time.ParseInLocation("2006-01-02 15", "2024-02-19 23", global.Tz)
			d, _ := json.Marshal(&[...][2]int64{{t11.Unix(), t12.Unix()}})

			sqls = []string{
				fmt.Sprintf("INSERT INTO `%s`(`dating_id`, `user_id`, `info`, `status`, `add_time`, `update_time`) VALUES(1, 2, '{\"t\": %s}', 1, %d, %d)", n, string(a), ti, ti),
				fmt.Sprintf("INSERT INTO `%s`(`dating_id`, `user_id`, `info`, `status`, `add_time`, `update_time`) VALUES(1, 1, '{\"t\": %s}', 1, %d, %d)", n, string(b), ti, ti),
				fmt.Sprintf("INSERT INTO `%s`(`dating_id`, `user_id`, `info`, `status`, `add_time`, `update_time`) VALUES(1, 3, '{\"t\": %s}', 1, %d, %d)", n, string(d), ti, ti),
				fmt.Sprintf("INSERT INTO `%s`(`dating_id`, `user_id`, `info`, `status`, `add_time`, `update_time`) VALUES(1, 1, '{\"t\": %s}', 1, %d, %d)", n, string(c), ti, ti),
			}
		}
	}

	for _, v := range sqls {
		global.Log.Infof("创建数据[dfskkjh]%s", v)
		if _, e := tx.Exec(v); e != nil {
			panic(fmt.Sprintf("错误[gh90iggs]:  %s\n%s", v, e.Error()))
		}
	}

}

func (*sqlite) version() (t string) {
	dao.DB.Get(&t, `SELECT sqlite_version()`)
	return
}

func (*mysql) version() (t string) {
	dao.DB.Get(&t, `SELECT version()`)
	return
}
