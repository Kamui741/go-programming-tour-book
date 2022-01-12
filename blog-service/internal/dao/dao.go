/*
 * @Author: ChZheng
 * @Date: 2022-01-08 22:30:47
 * @LastEditTime: 2022-01-09 22:39:29
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/dao/dao.go
 */
package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
