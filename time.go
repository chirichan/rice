package rice

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func TimeString(tm time.Time) string {
	return tm.Format(time.DateTime)
}

func NowTimeString() string {
	return TimeString(time.Now())
}

func NullTimeString(tm sql.NullTime) string {
	if tm.Valid {
		return TimeString(tm.Time)
	}
	return ""
}

func ZeroTime(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

func TodayZeroTime() time.Time {
	return ZeroTime(time.Now())
}

func TodayZeroTimestamp() int64 {
	return ZeroTime(time.Now()).Unix()
}

func ZeroTimestamp(ts int64) int64 {
	tm := time.Unix(ts, 0)
	return ZeroTime(tm).Unix()
}

// CountWeek 获取当前日期为当月第几周
func CountWeek(TimeFormat string) int {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", TimeFormat, time.Local)
	month := t.Month()
	year := t.Year()
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	week := 1
	for i := 1; i <= days; i++ {
		dayString := strconv.Itoa(i)
		if i < 10 {
			dayString = "0" + dayString
		}
		dateString := strings.Split(TimeFormat, "-")[0] + "-" + strings.Split(TimeFormat, "-")[1] + "-" + dayString + " 18:30:50"
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", dateString, time.Local)
		if t.YearDay() > t1.YearDay() {
			if t1.Weekday().String() == "Sunday" {
				week++
			}
		}

	}

	return week
}
