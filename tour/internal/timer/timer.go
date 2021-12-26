/*
 * @Author: ChZheng
 * @Date: 2021-12-23 17:28:35
 * @LastEditTime: 2021-12-24 01:43:16
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/internal/timer/timer.go
 */

package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}
func GetCalculatedTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}
