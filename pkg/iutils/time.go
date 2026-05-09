package iutils

import (
	"time"
)

// Sec 秒时间戳
func Sec() int64 {
	return time.Now().Unix()
}

// MilliSec 毫秒时间戳
func MilliSec() int64 {
	return time.Now().UnixNano() / 1e6
}

// TimeZero 获取当天0点时间戳
func TimeZero() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t1.Unix()
}

// 当天最后一刻时间戳
func TimeEnd() int64 {
	return TimeZero() + 86399
}

// Week 本周第一天
func WeekZero() int64 {
	week := time.Now().Weekday().String()
	var num int64 = 0
	switch week {
	case "Monday":
		num = 0
	case "Tusday":
		num = 1
	case "Wensday":
		num = 2
	case "Thursday":
		num = 3
	case "Friday":
		num = 4
	case "Sateday":
		num = 5
	case "Sunday":
		num = 6
	}

	fstTime := TimeZero() - num*86400

	return fstTime
}

// 本周最后一天最后一刻时间戳
func WeekEnd() int64 {
	now := time.Now()

	// 今天是周几
	weekday := now.Weekday()

	// 计算距离周日还有几天
	// Go 中周日是 0，周一 1 ... 周六 6
	var daysToSunday int
	if weekday == time.Sunday {
		daysToSunday = 0 // 今天就是周日
	} else {
		daysToSunday = 7 - int(weekday)
	}

	// 本周日的日期
	sunday := now.AddDate(0, 0, daysToSunday)

	// 构造本周日 23:59:59
	weekEndTime := time.Date(
		sunday.Year(),
		sunday.Month(),
		sunday.Day(),
		23, 59, 59, 0, // 时分秒毫秒
		now.Location(), // 使用当前系统时区
	)

	// 返回时间戳（秒）
	return weekEndTime.Unix()
}

// MonthZero 获取本月第一天时间戳
func MonthZero() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	return t1.Unix()
}

// 本月最后一天最后一刻时间戳
func MonthEnd() int64 {
	now := time.Now()
	// 获取当前月份的 1 号
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	// 下个月 1 号减去 1 纳秒 = 本月最后一天 23:59:59
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0).Add(-1 * time.Nanosecond)

	return lastOfMonth.Unix()
}

// SameDay 是否同一天
func SameDay(stamp int64) bool {
	return time.Unix(stamp, 0).Format("2006-01-02") == time.Now().Format("2006-01-02")
}

// 格式日期转unix时间戳
func TimeStrToUnix(timeStr string, loc *time.Location) (int64, error) {
	// 定义时间字符串的格式，必须和输入字符串完全匹配
	layout := "2006-01-02 15:04:05"

	// 解析时间字符串为 time.Time 类型
	t, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return -1, err
	}

	// 转换为 Unix 时间戳（秒级）
	return t.Unix(), nil
}
