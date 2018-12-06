FROM golang:1.8-alpine

#gitをインストール
RUN apk add --no-cache git

ENV GOPATH=$HOME/go

ENV srcDir="/usr/src"
RUN mkdir ${srcDir}
ENV appDir="$srcDir/myapp"
RUN mkdir ${appDir}
RUN chmod 744 ${appDir}
WORKDIR ${appDir}

#パッケージ関連
#websocket
RUN go get github.com/gorilla/websocket
#oauth2（TODO: 作成者曰く、Goth packageがオススメらしい）
RUN go get github.com/stretchr/gomniauth/oauth2
#traceパッケージの準備
ADD "./golang_trace" ${GOPATH}/src/myapp/trace

ENV port=8080
CMD [ "sh", "-c", "$appDir/build_run.sh $port" ]
