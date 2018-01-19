package main

import (
	"log"
	"net/http/httputil"
	"net/http"
)

func main(){
	resp, err := http.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
