package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup //usually pointer

func main()  {
	websitelist := []string{
		"https://google.com",
		"https://facebook.com",
		"https://afafa.in",
	}
	for _, v := range websitelist {
		go getStatuscode(v)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatuscode(endpoint string)  {
    defer wg.Done()
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops endpoint ")
	}else{
	fmt.Printf("%d status code for %s", resp.StatusCode, endpoint)
}
}