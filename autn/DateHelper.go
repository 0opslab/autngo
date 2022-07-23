package autn

import (
	"fmt"
	"strings"
	"time"
)

type DateHelper struct {
	Location time.Location
}

//时区
var Location = time.FixedZone("Asia/Shanghai", 8*60*60)

//获取当前时间 年-月-日 时:分:秒
func (gt *DateHelper) Now() string {
	return gt.NowTime().Format(CST_TIME_TT)
}

//获取当前时间戳
func (gt *DateHelper) NowUnix() int64 {
	return gt.NowTime().Unix()
}

//获取当前时间Time
func (gt *DateHelper) NowTime() time.Time {
	return time.Now().In(Location)
}

//获取年月日
func (gt *DateHelper) GetYmd() string {
	return gt.NowTime().Format(CST_TIME_YMD)
}

//获取时分秒
func (gt *DateHelper) GetHms() string {
	return gt.NowTime().Format(CST_TIME_HMS)
}

//获取当天的开始时间, eg: 2018-01-01 00:00:00
func (gt *DateHelper) NowStart() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, Location)
	return tm.Format(CST_TIME_TT)
}

//获取当天的结束时间, eg: 2018-01-01 23:59:59
func (gt *DateHelper) NowEnd() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, Location)
	return tm.Format(CST_TIME_TT)
}

//当前时间 减去 多少秒
func (gt *DateHelper) Before(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()-beforeSecond, 0).Format(CST_TIME_TT)
}

//当前时间 加上 多少秒
func (gt *DateHelper) Next(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()+beforeSecond, 0).Format(CST_TIME_TT)
}

//2006-01-02T15:04:05Z07:00 转 时间戳
func (gt *DateHelper) RfcToUnix(layout string) int64 {
	//转化所需模板
	//使用模板在对应时区转化为time.time类型
	tm, err := time.ParseInLocation(CST_TIME_RFC3339, layout, Location)
	if err != nil {
		return int64(0)
	}
	return tm.Unix()
}

//2006-01-02 15:04:05 转 时间戳
func (gt *DateHelper) ToUnix(layout string) int64 {
	theTime, _ := time.ParseInLocation(CST_TIME_TT, layout, Location)
	return theTime.Unix()
}

//获取RFC3339格式
func (gt *DateHelper) GetRFC3339() string {
	return gt.NowTime().Format(CST_TIME_RFC3339)
}

//转换成RFC3339格式
func (gt *DateHelper) ToRFC3339(layout string) string {
	tm, err := time.ParseInLocation(CST_TIME_TT, layout, Location)
	if err != nil {
		return ""
	}
	return tm.Format(CST_TIME_RFC3339)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func (gt *DateHelper) Format(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}
