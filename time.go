package rice

import "time"

// TimeFormatDate time to date
func TimeFormatDate(tm time.Time) string { return tm.Format("2006-01-02") }

// TimeParseDate string date to time date
func TimeParseDate(date string) (time.Time, error) { return time.Parse("2006-01-02", date) }

func TimeParseDatetime(datetime string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", datetime)
}

// TimeNowFormat time.Time to "2006-01-02 15:04:05"
func TimeNowFormat() string { return time.Now().Format("2006-01-02 15:04:05") }

// TimeFormat time.Time to "2006-01-02 15:04:05"
func TimeFormat(tm time.Time) string { return tm.Format("2006-01-02 15:04:05") }

// TimeUnixFormatDate timestamp to date
func TimeUnixFormatDate(timestamp int64) string { return time.Unix(timestamp, 0).Format("2006-01-02") }

// TimeUnixFormatDatetime time to datetime
func TimeUnixFormatDatetime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// HMCompare 如果 1 > 2 return true
// 小时和分钟 比较大小
func HMCompare(h1, m1, h2, m2 int) bool {

	t := time.Date(0, 0, 0, h1, m1, 0, 0, time.Local)
	t2 := time.Date(0, 0, 0, h2, m2, 0, 0, time.Local)

	if t.Unix() > t2.Unix() {
		return true
	} else {
		return false
	}
}

// BetweenDays 两个时间之间隔了多少天 startTime >= endTime
func BetweenDays(startTime, endTime time.Time) int64 {

	var days int64

	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.Local)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.Local)

	for i := startTime; i.Before(endTime); i = i.AddDate(0, 0, 1) {
		days += 1
	}
	return days
}
