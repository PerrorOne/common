package utils

import (
	"github.com/jmcvetta/randutil"
	"time"
)

// 获取今天0点的时间
func GetTodayZeroTime(addDay ...int) time.Time {
	timeStr := time.Now().Format("06-01-02")
	local, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("06-01-02", timeStr, local)
	if len(addDay) != 0 {
		return t.AddDate(0, 0, addDay[0])
	}
	return t
}

// 从当前时间减去last_time时间中随机一个时间
func RandTimeByDy(lastTime time.Duration) (time.Time, error) {
	now := time.Now()
	lastDay := now.Add(-lastTime)
	send, err := randutil.IntRange(int(lastDay.Unix()), int(now.Unix()))
	if err != nil {
		return now, err
	}
	return time.Unix(int64(send), 0), nil
}
