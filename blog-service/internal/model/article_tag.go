/*
 * @Author: ChZheng
 * @Date: 2021-12-30 15:33:48
 * @LastEditTime: 2021-12-30 15:33:49
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/model/article_tag.go
 */
package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
