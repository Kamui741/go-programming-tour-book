/*
 * @Author: ChZheng
 * @Date: 2021-12-31 14:13:07
 * @LastEditTime: 2021-12-31 14:14:25
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/global/db.go
 */
package global

import (
	"github.com/jinzhu/gorm"
)

var (
	DBEngine *gorm.DB
)
