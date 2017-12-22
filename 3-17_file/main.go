package main

import "net/http"
import "net/http/httputil"
import "log"

func main() {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	client := http.Client{
		Transport: transport,
	}
	resp, err := client.Get("file://./main.go")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
