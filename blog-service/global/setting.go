/*
 * @Author: ChZheng
 * @Date: 2021-12-31 01:32:29
 * @LastEditTime: 2021-12-31 16:20:45
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/global/setting.go
 */
package global

import (
	"go-programming-tour-book/blog-service/pkg/logger"
	"go-programming-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
