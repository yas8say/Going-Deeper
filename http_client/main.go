package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	makeRequest("https://jsonplaceholder.typicode.com/posts/1")
	makeRequest("https://jsonplaceholder.typicode.com/invalid-url")
}

func makeRequest(url string) {
	defer func() {
		//r will be return value of recover
		if r := recover(); r != nil {
			fmt.Println("Recovered from error", r)
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Sprintf("HTTP request failed: %s", err))
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status code: %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("failed to read response body: %s", err))
	}

	fmt.Print("Response:", string(body))
}
