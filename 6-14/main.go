package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"bufio"
	"time"
	"net"
)

func main() {
	// TCPソケットをオープン
	dialer := &net.Dialer{
		Timeout: 30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// リクエストを作成し、ソケットに直接書き込み
	request, _ := http.NewRequest("GET", "http://localhost:18888/upgrade", nil)
	request.Header.Set("Connection", "Upgrade")
	request.Header.Set("Upgrade", "MyProtocol")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	// ソケットから直接データを読み込んでレスポンス解析
	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)

	// オリジナルの通信を開始
	counter := 10
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Println("<-", string(bytes.TrimSpace(data)))
		fmt.Fprintf(conn, "%d\n", counter)
		fmt.Println("->", counter)
		counter--
	}
}
