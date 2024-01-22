package main

import (
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	data , err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	pp.Println(string(data))
}