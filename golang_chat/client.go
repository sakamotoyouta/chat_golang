package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// clientはチャットを行なっている1人のユーザーを表します。
type client struct {
	// socket はこのクラアントのためのwebSocketです。
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです。
	send chan *message
	// room はこのクライアントが参加しているチャットルームです。
	room *room
	// userDataはユーザーに関する情報を保持します
	userData map[string]interface{}
}

/* クライアントがWebSocket から ReadMessageを使用してデータを読み込むために使用する */
func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

/* 継続的に sendチャネルからメッセージを受け取り、WebSocket のWriteMessageメソッドを使って書き出す */
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
