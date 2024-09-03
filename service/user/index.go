package user

import (
	"bytes"
	"context"
	"database/sql"
	"strconv"
	"time"

	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/twbworld/dating/dao"
	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/model/common"
	"github.com/twbworld/dating/model/db"
	"github.com/twbworld/dating/utils"
)

const (
	minTime, maxTime = 8, 23
)

type DatingService struct{}

func (d *DatingService) MatchGoroutine(datingId uint) {
	if datingId < 1 {
		return
	}
	go func() {
		defer func() {
			//避免协程内panic影响到外层主程序
			if p := recover(); p != nil {
				global.Log.Error("[gji075]", p)
			}
		}()
		d.Match(datingId)
	}()
}

// 匹配最合适的时间
func (d *DatingService) Match(datingId uint) {
	var (
		du        []common.DatingUser
		duLen     int
		TimeIds   map[int64][]uint = map[int64][]uint{}
		maxSecond int              = (maxTime - minTime) * 3600
		numTime   map[int][]int    = map[int][]int{}
		res       db.DatingResult
	)
	if datingId < 1 {
		panic("系统错误[odfjgo]")
	}
	if err := dao.App.DatingDb.GetDatingUsers(&du, datingId); err != nil {
		panic(err)
	} else if duLen = len(du); duLen < 1 {
		panic("数据为空[hk942]" + strconv.FormatInt(int64(datingId), 10))
	}

	//汇总每个时间段(小时)下, 有哪些人空闲; 如: {"下午13点" : ["张三"]}
	for _, value := range du {
		if len(value.Info) < 1 {
			continue
		}
		info := *value.InfoUnmarshal()
		if info.Time == nil || len(info.Time) < 1 {
			continue
		}

		for _, val := range info.Time {
			val = val[:2] //取切片内前两个时间戳数据
			if val[0] >= val[1] || int(val[1]-val[0]) > maxSecond {
				continue
			}
			ts := utils.SpreadPeriodToHour(val[0], val[1])
			for _, v := range ts {
				if _, ok := TimeIds[v]; !ok {
					TimeIds[v] = make([]uint, 0, duLen)
				}
				TimeIds[v] = append(TimeIds[v], value.UtId)
			}

		}
	}

	if len(TimeIds) < 1 {
		return
	}

	//汇总空闲时间段相同的用户数下, 有哪些时间段(小时); 如: {"有三人空闲" : ["下午13点"]}
	for key, value := range TimeIds {
		n := len(value)
		if _, ok := numTime[n]; !ok {
			numTime[n] = make([]int, 0)
		}
		numTime[n] = append(numTime[n], int(key))
	}

	if len(numTime) < 1 {
		return
	}

	//从匹配度(空闲次数)最高的开始计算; 如共5个用户参与, 在numTime中从5开始往下找
	for i := duLen; i > 0; i-- {
		if _, ok := numTime[i]; !ok {
			continue
		}

		if i == 1 && duLen != 1 {
			//多个用户也没有共同时间, 则匹配失败
			goto SAVE
		}

		res.Date = d.SimplePeriod(numTime[i])

		if i == duLen {
			//首次遍历并没continue, 证明所有用户都有共同的空闲时间, 匹配成功 !
			res.Res = true
		}

	SAVE:

		err := dao.Tx(func(tx *sqlx.Tx) (e error) {
			return dao.App.DatingDb.DatingUpdate(datingId, &res, tx)
		})
		if err != nil {
			global.Log.WithError(err)
			return
		}
		//已匹配最优结果, 直接退出
		return
	}
}

// 简化时间表达; 如: ["02-15(8-10时|12-13时)"]
func (d *DatingService) SimplePeriod(unixTimes []int) []string {
	date := make([]string, 0)
	if len(unixTimes) < 1 {
		return date
	}

	dateTime := utils.UnixGroup(unixTimes)

	for _, val := range dateTime {

		tim := time.Unix(int64(val[0]), 0).In(global.Tz)

		fo := "01-02"
		if time.Now().In(global.Tz).Year() != tim.Year() {
			fo = time.DateOnly
		}

		var recommendText bytes.Buffer //节省资源的字符串拼接方式
		recommendText.WriteString(tim.Format(fo))

		//如果是具体时间段才空闲(非一整天都空闲), 需要括号内标出具体时间(小时)
		if l := len(val); l < maxTime-minTime {
			recommendText.WriteString("(")
			recommendText.WriteString(strconv.Itoa(tim.Hour()))
			recommendText.WriteString("-")
			for k, v := range val {
				tv := time.Unix(int64(v), 0).In(global.Tz)
				if l == k+1 {
					//如果是最后一个(小时)值, 直接拼接
					recommendText.WriteString(strconv.Itoa(tv.Hour() + 1))
					recommendText.WriteString("时")
				} else if tvlast := time.Unix(int64(val[k+1]), 0).In(global.Tz); tv.Hour()+1 != tvlast.Hour() {
					//判断当前(小时)值和下一个(小时)值是否相同, 避免数据错误
					recommendText.WriteString(strconv.Itoa(tv.Hour() + 1))
					recommendText.WriteString("时")
					recommendText.WriteString("|") //因为不是最后一个(小时)值, 所以还有其他时间段拼接
					recommendText.WriteString(strconv.Itoa(tvlast.Hour()))
					recommendText.WriteString("-")
				}
			}
			recommendText.WriteString(")")
		}

		date = append(date, recommendText.String())

	}

	return date
}

func (d *DatingService) GetUserByCode(code string, u *db.User) (err error) {
	if len(code) < 1 {
		return errors.New("系统错误[iodgj]")
	}

	if len(global.Config.Weixin.XcxAppid) < 1 {
		err = errors.New("没配置小程序Appid[nbvkpl]")
		return
	}

	rs, err := global.MiniProgramApp.Auth.Session(context.Background(), code)
	if err != nil {
		return
	}
	if len(rs.OpenID) < 1 {
		err = errors.New("请刷新[nb09]")
		return
	}

	if err = dao.App.UserDb.GetUserByOpenId(u, rs.OpenID); err != nil {
		if err == sql.ErrNoRows {
			//新用户
			u.OpenId = rs.OpenID
			u.UnionId = rs.UnionID
			u.SessionKey = rs.SessionKey
			return nil
		}
		return
	}
	if u.Id == dao.BaseUserId {
		return errors.New("出现测试账号[rtyoij]")
	}

	if u.SessionKey != rs.SessionKey {
		go func() {
			defer func() {
				//避免协程内panic影响到外层主程序
				if p := recover(); p != nil {
					global.Log.Error(p)
				}
			}()
			err := dao.Tx(func(tx *sqlx.Tx) (e error) {
				return dao.App.UserDb.UpdateSessionKey(u.Id, rs.SessionKey, tx)
			})
			if err != nil {
				panic(err)
			}
		}()
	}
	return
}
