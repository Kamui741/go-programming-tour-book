/*
 * @Author: ChZheng
 * @Date: 2021-12-23 22:14:35
 * @LastEditTime: 2022-01-01 16:46:30
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/cmd/timer.go
 */
package cmd

import (
	"go-programming-tour-book/tour/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := time.LoadLocation("Asia/Shanghai")
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				log.Println(err.Error())
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculatedTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculatedTime err: %v", err)
		}
		log.Printf("输出结果: %s,%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
}
