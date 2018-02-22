package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var image []byte

// 画像ファイルを準備
func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// HTMLをブラウザに送信
// 画像をプッシュする
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	// Pusherにキャスト可能であれば（HTTP/2で接続していたら）プッシュする
	pusher, ok := w.(http.Pusher)
	if ok {
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

// 画像をファイルブラウザに送信
func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/png")
	w.Write(image)
}

func main() {
	http.HandleFunc("/", handlerHTML)
	http.HandleFunc("/image", handlerImage)
	fmt.Println("start http listening :18443")
	err := http.ListenAndServeTLS(":18443", "../tls/server.crt", "../tls/server.key", nil)
	fmt.Println(err)
}
