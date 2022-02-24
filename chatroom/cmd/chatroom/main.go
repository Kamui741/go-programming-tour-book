/*
 * @Author: ChZheng
 * @Date: 2022-02-25 00:26:12
 * @LastEditTime: 2022-02-25 00:28:43
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /chatroom/cmd/chatroom/main.go
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"go-programming-tour-book/chatroom/global"
	"go-programming-tour-book/chatroom/server"
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
