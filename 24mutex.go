package main
import (
	"fmt"
	"net/http"
	"sync"
)
/*Mutex stands for mutual exclusion, which means that only one goroutine can hold the lock at a time, 
and all other goroutines that attempt to acquire the lock will be blocked until the lock is released by the 
goroutine that holds it.
*/

//EXAMPLE::1
var signal = []string{"test"}
var wg sync.WaitGroup //usually pointer
var mut sync.Mutex //usually pointer

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
	mut.Lock()
		signal = append(signal, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s", resp.StatusCode, endpoint)
}
}

//EXAMPLE::2
var x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()   
}
func main() {  
    var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, &m)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}


