/*
 * @Author: ChZheng
 * @Date: 2021-12-30 15:31:24
 * @LastEditTime: 2022-01-02 01:39:58
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/model/article.go
 */
package model

import "go-programming-tour-book/blog-service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}
type ArticleSwagger struct {
	List  []Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
