package main

import "time"

// messageは１つのメッセージを表します。
type message struct {
	Name      string    // 送信者名
	Message   string    // 本文
	When      time.Time // 送信時刻
	AvatarURL string    // ユーザー画像
}
