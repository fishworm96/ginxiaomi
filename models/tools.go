package models

import (
	"time"
)

//时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02T15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	} else {
		return t.Unix()
	}
}

	func GetUnix() int64 {
		return time.Now().Unix()
	}

	func getDate() string {
		template := "2006-01-02 15:04:05"
		return time.Now().Format(template)
	}

	func GetDay() string {
		template := "2006-01-02 15:04:05"
		return time.Now().Format(template)
	}