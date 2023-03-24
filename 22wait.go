package main
//we use a synchronization primitive called a waitgroup. A waitgroup is used to wait for a collection of goroutines to finish executing their tasks before proceeding further.
import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup //usually pointer
//This is useful in situations where we need to collect results from multiple goroutines, 
//or when we need to make sure that a certain set of goroutines have completed their tasks before we proceed with the next set of tasks.
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