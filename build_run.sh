#!/bin/ash

#バイナリ削除
if [ -e ./chat ]; then
  rm ./chat
fi

#再度自作モジュールの移動（変更があった場合、リフレッシュさせるため）
#FIXME:更新ファイルがあれば更新するように修正
rm ${GOPATH}/src/myapp/trace/*
\cp ./golang_trace/* ${GOPATH}/src/myapp/trace/

#バイナリ作り
go build -o ./chat ./golang_chat 2> ./error_log.txt
if [ -e ./error_log.txt ]; then
  echo -e "\e[31m$(cat ./error_log.txt)\e[m"
  rm ./error_log.txt
fi

#権限付
chmod 744 ./chat

#ファイル実行
./chat -addr=":${1:-8080}"
