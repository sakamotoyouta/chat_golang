package main

import (
	"log"
	"net/http"
	"myapp/trace"

	"github.com/gorilla/websocket"
)

type room struct {
	// 受け取ったメッセージを全てのクライアントに転送する時に使用
	forward chan []byte
	// チャットルームに参加しようとしているクライアントのためのチャネル
	join chan *client
	// チャットルームから体質しようとしているクライアントのためのチャネル
	leave chan *client
	// 在室している全てのクライアントが保持される
	clients map[*client]bool
	// tracerはチャットルーム上で行われた操作のログを受け取ります
	tracer trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		// FIXME:それぞれの処理を関数に切り出したい
		case client := <-r.join:
			// 参加
			r.clients[client] = true
			r.tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("クライアントが退室しました")
		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました：", string(msg))
			// 全てのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <-msg:
				// TODO:メッセージを送信
				r.tracer.Trace(" -- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)
// HTTP接続をアップグレード
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}
// roomにServerHTTPメソッドを作成することでHTTPハンドラとして使用できるようにする
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}
	r.join <-client
	defer func() { r.leave <-client }()
	go client.write()
	client.read()
}
