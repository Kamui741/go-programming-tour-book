/*
 * @Author: ChZheng
 * @Date: 2021-12-30 15:28:18
 * @LastEditTime: 2022-01-02 01:40:37
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/model/tag.go
 */
package model

import "go-programming-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
type TagSwagger struct {
	List  []Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
