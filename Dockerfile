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

#websocketパッケージの取得
RUN go get github.com/gorilla/websocket
#traceパッケージの準備
ADD "./golang_trace" ${GOPATH}/src/myapp/trace

ENV port=8080
CMD [ "sh", "-c", "$appDir/build_run.sh $port" ]
