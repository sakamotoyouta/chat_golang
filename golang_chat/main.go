package main

import (
	"flag"
	"log"
	"myapp/trace"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
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
	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("hogefugahogefuga") // 自分で決めた文字列
	gomniauth.WithProviders(
		facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		github.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/github"),
		google.New(GoogleAuth["clientId"], GoogleAuth["secretKey"], "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	// ルーティング
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// チャットルームを開始します
	go r.run()

	// Web サーバを開始します
	log.Println("Webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
