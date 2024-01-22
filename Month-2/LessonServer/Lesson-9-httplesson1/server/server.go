package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var val string
	fmt.Scan(&val)
	fmt.Fprint(w, val)
	})
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}