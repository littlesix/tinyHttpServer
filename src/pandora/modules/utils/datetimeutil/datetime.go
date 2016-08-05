package datetimeutil

import (
	"strings"
	"time"
	"fmt"
)

//返回 2015-10-10
func FormatDateNow() string {
	return FormatDateWithSplitNow("-")
}

//设置一个日期分割字符  如 /  返回2015/10/10
func FormatDateWithSplitNow(split string) string {
	t := time.Now()
	return t.Format("2006" + split + "01" + split + "02")
}
func FormatDateTimeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func FormatDateTimeMillsNow() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}
func FormatDateTimeWithSplitNow(dateSplit string, timeSplit string) string {
	t := time.Now()
	return t.Format("2006" + dateSplit + "01" + dateSplit + "02 15" + timeSplit + "04" + timeSplit + "05")
}
func FormatDate(timeStampOrMillis int64) string {
	return FormatDateWithSplit("-", timeStampOrMillis)
}
func FormatDateWithSplit(split string, timeStampOrMillis int64) string {
	if timeStampOrMillis > 100000000000 {
		timeStampOrMillis = timeStampOrMillis / 1000
	}
	t := time.Unix(timeStampOrMillis, 0)
	return t.Format("2006" + split + "01" + split + "02")
}
func FormatDateTime(timeStampOrMillis int64) string {
	return FormatDateTimeWithSplit("-", ":", timeStampOrMillis)
}
func FormatDateTimeWithSplit(dateSplit string, timeSplit string, timeStampOrMillis int64) string {
	if timeStampOrMillis > 100000000000 {
		timeStampOrMillis = timeStampOrMillis / 1000
	}
	t := time.Unix(timeStampOrMillis, 0)
	return t.Format("2006" + dateSplit + "01" + dateSplit + "02 15" + timeSplit + "04" + timeSplit + "05")
}

func ParseDateWithSplit(dateSplit, dateStr string) int64 {
	dateStr = strings.Trim(dateStr, " ")
	t, e := time.Parse(dateStr, "2006"+dateSplit+"01"+dateSplit+"02")
	if e != nil {
		return 0
	} else {
		return t.Unix()
	}
}
func ParseDateTimeWithSplitNow(dateSplit, timeSplit, dateTimeStr string) int64 {
	dateTimeStr = strings.Trim(dateTimeStr, " ")
	t, e := time.Parse(dateTimeStr, "2006"+dateSplit+"01"+dateSplit+"02 15"+timeSplit+"04"+timeSplit+"05")
	if e != nil {
		return 0
	} else {
		return t.Unix()
	}
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
func GetCurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
func GetCurrentTimeMillisStr() string {
	return fmt.Sprint(GetCurrentTimeMillis())
}
