package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func Patch(link string)  {
	url := link
	payload := map[string]string{"name": "New Name"}

	// encode payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// create new HTTP PATCH request with JSON payload
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	// set content-type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// create HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	pp.Println(req)
}

func Put(link string) {
	url := link
    data := []byte(`{"name": "John", "age": 30}`)

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
    if err != nil {
        // handle error
		return 
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // handle error
		return
    }
    defer resp.Body.Close()

    // do something with the response
	pp.Println(req)
}

func Post(link string)  {
    url := link

    jsonStr := []byte(`{"name":"John","age":30,"city":"New York"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // Handle response
	pp.Println(req)
}

func Delete(link string)  {
	// create a new HTTP client
    client := &http.Client{}

    // create a new DELETE request
    req, err := http.NewRequest("DELETE", link, nil)
    if err != nil {
        panic(err)
    }

    // send the request
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    // print the response body
    fmt.Println(string(body))
}

func Get(link string) {
	// make http get request
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	// print the the request data
	pp.Println(resp)
}

func main() {
	// Get("https://httpbin.org/get")
	// Post("https://httpbin.org/post")
	// Delete("https://httpbin.org/delete")
	// Patch("https://httpbin.org/patch/")
	// Put("https://httpbin.org/put")
}
