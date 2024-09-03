package user

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/twbworld/dating/global"
	"github.com/twbworld/dating/model/common"
)

type Validator struct{}

// 检验LoginPost参数
func (v *Validator) ValidatorLoginPost(data *common.LoginPost) error {
	if len(data.Code) < 1 {
		return errors.New("参数错误[dotsd]")
	}
	return nil
}

// 检验UserAddPost参数
func (v *Validator) ValidatorUserAddPost(data *common.UserInfoPost) error {
	//nick_name可能为空, 不做判断
	if len(data.Code) < 1 || len(data.EncryptedData) < 1 || len(data.Iv) < 1 {
		return errors.New("参数错误[dofsd]")
	}
	return nil
}

// 检验FeedbackPost参数
func (v *Validator) ValidatorFeedbackPost(data *common.FeedbackPost) error {
	//nick_name可能为空, 不做判断
	if len(data.Desc) < 1 || utf8.RuneCountInString(data.Desc) > 100 {
		return errors.New("参数错误[dofsd]")
	}
	return nil
}

// 检验UserAddPost参数
func (v *Validator) ValidatorEncryptedData(data *common.UserInfoWx) error {
	//nick_name可能为空, 不做判断
	if len(data.AvatarUrl) < 1 || data.Gender > 2 {
		return errors.New("参数错误[dosd]")
	}
	return nil
}

// 检验InfoPost参数
func (v *Validator) ValidatorInfo(data *common.InfoPost) error {
	if len(data.Time) < 1 {
		return errors.New("参数错误[doifj]")
	}

	y1, y2, dates := time.Now().AddDate(1, 0, 0), time.Now().AddDate(-1, 0, 0), make(map[string]bool) //允许两年时间跨度
	for _, value := range data.Time {
		te, err := time.ParseInLocation(time.DateTime, value[1], global.Tz)
		if err != nil {
			return errors.New("时间选择错误[oilsng]")
		}
		ts, err := time.ParseInLocation(time.DateTime, value[0], global.Tz)
		if err != nil {
			return errors.New("时间选择错误[oildng]")
		}

		if ts.Format("04:05") != "00:00" || te.Format("04:05") != "00:00" || ts.After(te) || ts.Before(y2) || te.After(y1) {
			return errors.New("时间错误[odfibj]")
		}
		if ts.Hour() == te.Hour() {
			return errors.New("前后时间不能相同[dojinv]")
		}
		if !(ts.Hour() >= minTime && ts.Hour() < maxTime) {
			return fmt.Errorf("开始时间必须%d点 - %d点[oifjd]", minTime, maxTime-1)
		}
		if !(te.Hour() > minTime && te.Hour() <= maxTime) {
			return fmt.Errorf("结束时间必须%d点 - %d点[oifgjka]", minTime+1, maxTime)
		}
		if fts := ts.Format(time.DateOnly); fts != te.Format(time.DateOnly) {
			return errors.New("时间选择错误[odfkng]")
		} else if _, ok := dates[fts]; ok {
			return errors.New("时间选择错误[odfdfskng]")
		} else {
			dates[fts] = true
		}
	}
	return nil
}
