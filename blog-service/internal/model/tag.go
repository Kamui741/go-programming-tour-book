/*
 * @Author: ChZheng
 * @Date: 2021-12-30 15:28:18
 * @LastEditTime: 2021-12-30 20:16:22
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/model/tag.go
 */
package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
