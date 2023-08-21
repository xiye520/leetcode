package main

import (
	"fmt"
	"time"
)

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "20060102 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "20060102"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func main() {
	dates := GetBetweenDates("19700101", time.Now().Format("20060102"))
	s2 := make([]string, 0, 10)
	for i := range dates {
		if isValid(dates[i]) {
			s2 = append(s2, dates[i])
		}
	}

	fmt.Println(s2)
}
