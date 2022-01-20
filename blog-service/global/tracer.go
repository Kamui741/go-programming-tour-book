/*
 * @Author: ChZheng
 * @Date: 2022-01-20 21:47:07
 * @LastEditTime: 2022-01-20 21:50:00
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/global/tracer.go
 */
package global

import "github.com/opentracing/opentracing-go"

var (
	Tracer opentracing.Tracer
)
