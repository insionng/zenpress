package helper

import (
	"fmt"
	"strconv"
	"time"
)

// FMT_TYPE_NOMAL
const (
	DATE_TIME_FMT = "2006-01-02 15:04:05"

	DATE_FMT = "2006-01-02"

	TIME_FMT = "15:04:05"

	DATE_TIME_FMT_CN = "2006年01月02日 15时04分05秒"

	DATE_FMT_CN = "2006年01月02日"

	TIME_FMT_CN = "15时04分05秒"
)

const SecondInNano = 1000 * 1000 * 1000

//return 1441006057 in sec
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//return 1441006057 in sec
func GetTimestampString() string {
	return strconv.FormatInt(GetTimestamp(), 10)
}

// return 1441007112776 in millisecond
func GetTimestampInMilli() int64 {
	return int64(time.Now().UnixNano() / (1000 * 1000)) // ms
}

// return 1441007112776 in millisecond
func GetTimestampInMilliString() string {
	return strconv.FormatInt(GetTimestampInMilli(), 10)
}

//微秒
func GetTimestampInMicro() int64 {
	return int64(time.Now().UnixNano() / 1000) // ms
}

// 微秒
func GetTimestampInMicroString() string {
	return strconv.FormatInt(GetTimestampInMicro(), 10)
}

//format
func GetCurrentTimeFormat(format string) string {
	return GetTimeFormat(GetTimestamp(), format)
}

//
func GetTimeFormat(second int64, format string) string {
	return time.Unix(second, 0).Format(format)
}

// Timing the cost of function call, unix nano was returned
func Elapse(f func()) int64 {
	now := time.Now().UnixNano()
	f()
	return time.Now().UnixNano() - now
}

// Timing the cost of function call, unix nano was returned
func ElapseString(f func()) string {
	return strconv.FormatInt(Elapse(f), 10)
}

// GetMonthDays return days of the month/year
func GetMonthDays(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		panic(fmt.Sprintf("Illegal month:%d", month))
	}
}

// IsLeapYear check whether a year is leay
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}

	return year%4 == 0
}
