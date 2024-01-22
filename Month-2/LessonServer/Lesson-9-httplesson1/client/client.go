package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
	}

	defer resp.Body.Close()

	mm, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(mm))
}