package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行なっている1人のユーザーを表します。
type client struct {
	// socket はこのクラアントのためのwebSocketです。
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです。
	send chan []byte
	// room はこのクライアントが参加しているチャットルームです。
	room *room
}

/* クライアントがWebSocket から ReadMessageを使用してデータを読み込むために使用する */
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
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
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
