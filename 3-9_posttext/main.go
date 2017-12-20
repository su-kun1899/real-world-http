package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
