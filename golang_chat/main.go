package main

import (
	"os"
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"myapp/trace"
)

/* 構造体 templateHandlerの定義 */
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

/* HTTPリクエストを処理します */
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("./golang_chat/templates", t.filename)))
	})
	// TODO:戻り値をチェックすべき
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	// ルーティング
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/room", r)

	// チャットルームを開始します
	go r.run()

	// Web サーバを開始します
	log.Println("Webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
