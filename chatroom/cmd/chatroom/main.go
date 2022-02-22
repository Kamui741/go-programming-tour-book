/*
 * @Author: ChZheng
 * @Date: 2022-02-21 22:43:29
 * @LastEditTime: 2022-02-21 23:02:54
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /chatroom/cmd/chatroom/main.go
 */
package main

import (
	"fmt"
	"go-programming-tour-book/chatrom/server"
	"log"
	"net/http"

	_ "net/http/pprof"
)

var (
	addr   = ":2022"
	banner = `
	 ____              _____
	|    |    |   /\     |
	|    |____|  /  \    |
	|    |    | /----\   |
	|____|    |/      \  |
 Go语言编程之旅 —— 一起用Go做项目：ChatRoom，start on：%s
 `
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf(banner, addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
