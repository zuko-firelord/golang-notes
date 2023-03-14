package main

import (
	"fmt"
	"time"
)

// Goroutines are a feature of Go that allows us to execute functions concurrently in a lightweight manner.
func main()  {
	go routine("hello")
	routine("world")
//world hello hello world hello world world hello hello world
//"hello" is not printed because it is executed as a goroutine using the "go" keyword. When a function is called using "go" keyword, it is executed concurrently with the main program.
}

func routine(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond) //pause the execution of the main goroutine for 3 milliseconds
		fmt.Println(s)
	}
}
// To ensure that the "hello" string is printed before the "world" string, synchronization mechanisms such as channels and wait groups can be used.