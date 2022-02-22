/*
 * @Author: ChZheng
 * @Date: 2022-02-21 22:46:50
 * @LastEditTime: 2022-02-21 23:03:01
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /chatroom/server/handle.go
 */
package server

import (
	"net/http"
)

func RegisterHandle() {
	inferRootDir()
	go logic.Broadcast.Start()
	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/ws", webSocketHandleFunc)
}
