#!/bin/ash

#バイナリ作成
go build -o ./chat ./golang_chat

#権限付
chmod 744 ./chat

#ファイル実行
./chat -addr=":${1:-8080}"
