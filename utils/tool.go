package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/twbworld/dating/global"
)

type timeNumber interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(strings.TrimSpace(strings.Trim(str, "\n"))))
}
func Base64Decode(str string) string {
	bstr, err := base64.StdEncoding.DecodeString(strings.TrimSpace(strings.Trim(str, "\n")))
	if err != nil {
		return str
	}
	return string(bstr)
}
func Hash(str string) string {
	b := sha256.Sum224([]byte(str))
	return hex.EncodeToString(b[:])
}

func TimeFormat[T timeNumber](t T) string {
	return time.Unix(int64(t), 0).In(global.Tz).Format(time.DateTime)
}

// 四舍五入保留小数位
func NumberFormat[T ~float32 | ~float64](f T, n ...uint) float64 {
	var num uint
	if len(n) == 0 {
		num = 2
	} else {
		num = n[0]
	}
	nu := math.Pow(10, float64(num))
	return math.Floor(float64(f)*nu+0.5) / nu
}

func CreateFile(path string) (err error) {
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		paths, _ := filepath.Split(path)

		if _, err = os.Stat(paths); err != nil {
			if err = os.MkdirAll(paths, os.ModePerm); err != nil {
				return
			}
		}

		fi, e := os.Create(path)
		if e != nil {
			return e
		}
		fi.Close()
	}
	file.Close()

	return
}

// 类似php的array_column($a, null, 'key')
func ListToMap(list interface{}, key string) map[string]interface{} {
	data := make([]interface{}, 0)
	if v := reflect.ValueOf(list); v.Kind() != reflect.Slice {
		data = append(data, list)
	} else {
		for i := range v.Len() {
			data = append(data, v.Index(i).Interface())
		}
	}

	res := make(map[string]interface{}, len(data))
	for _, value := range data {
		res[reflect.ValueOf(value).FieldByName(key).String()] = value
	}

	return res
}

func InSlice(slice []string, value string) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}
	return -1
}

// 时间戳按日期分组; 例: {[1707100000,1707100000], [170720000]}
func UnixGroup(times []int) [][]int {
	unixGroup := [][]int{}
	if len(times) < 1 {
		return unixGroup
	}
	sort.Ints(times)
	dateTime := map[string][]int{}
	for _, val := range times {
		if val < 1 {
			continue
		}
		d := time.Unix(int64(val), 0).In(global.Tz).Format(time.DateOnly)
		if _, ok := dateTime[d]; !ok {
			dateTime[d] = make([]int, 0, len(times))
		}
		dateTime[d] = append(dateTime[d], val)
	}
	if len(dateTime) < 1 {
		return unixGroup
	}
	for _, v := range dateTime {
		unixGroup = append(unixGroup, v)
	}

	return unixGroup
}

// 打散时间段(粒度为1小时) ; 如: "1-4点" 转为 ["1点", "2点", "3点"] 三个时间段
func SpreadPeriodToHour[T timeNumber](start, end T) (res []T) {
	add := T(3600)
	for start < end {
		//这不使用"<=", 不算最后的时间戳,是因为: 往后的一个时间戳值,代表当前时间戳的后一小时, 而不是当前秒
		res = append(res, start)
		start += add
	}
	return
}
