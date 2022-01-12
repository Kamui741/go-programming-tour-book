/*
 * @Author: ChZheng
 * @Date: 2022-01-08 22:50:43
 * @LastEditTime: 2022-01-08 23:01:00
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/service/service.go
 */
package service

import (
	"context"
	"go-programming-tour-book/blog-service/global"
	"go-programming-tour-book/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
