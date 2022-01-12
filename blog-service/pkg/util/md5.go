/*
 * @Author: ChZheng
 * @Date: 2022-01-09 23:02:24
 * @LastEditTime: 2022-01-09 23:09:02
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/pkg/util/md5.go
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
