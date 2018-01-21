package main

import (
	"log"
	"crypto/tls"
	"time"
	"io"
	"fmt"
	"net/http"
)

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
	// このエンドポイントでは変更以外は受け付けない
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "MyProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to MyProtocol")

	// 低層のソケットを取得
	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()

	// プロトコルが変わるというレスポンスを送信
	response := http.Response{
		StatusCode: 101,
		Header: make(http.Header),
	}
	response.Header.Set("Upgrade", "MyProtocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	// オリジナルの通信の開始
	for i := 1; i < 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush() // Trigger "chunked" encoding and send a chunk...
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}
}

func main(){
	server := &http.Server{
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			MinVersion: tls.VersionTLS12,
		},
		Addr: ":18443",
	}
	http.HandleFunc("/", handlerUpgrade)
	log.Println("start http listening :18443")
	err := server.ListenAndServeTLS("../tls/server.crt", "../tls/server.key")
	log.Println(err)
}

