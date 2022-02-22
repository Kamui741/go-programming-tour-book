/*
 * @Author: ChZheng
 * @Date: 2022-02-21 22:47:13
 * @LastEditTime: 2022-02-22 23:20:26
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /chatroom/server/websocket.go
 */
package server

import (
	"log"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func WebSocketHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, nil)
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}
	nickname := req.FormValue("nickname")
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal: ", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("非法昵称，昵称长度：4-20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal")
		return
	}
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已经存在：", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("该昵称已经已存在！"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists")
		return
	}
	user := logic.NewUser(conn, nickname, req.RemoteAddr)
	go user.SendMessage(req.Context())
	user.MessageChannel <- logic.NewWelcomeMessage(nickname)

	msg := logic.NewNoticeMessage(nickname + "加入聊天室")
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", nickname, "joins chat")
	err = user.ReceiveMessage(req.Context())
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewNoticeMessage(user.nickname + " 离开了聊天室")
	logic.Broadcaster.Broadcast(msg)
	log.Println("user:", nickname, "leaves chat")
	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")

	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
