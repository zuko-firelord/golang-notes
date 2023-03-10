package main
import "fmt"
// It  used to schedule a function call to be executed at the end of the current function, just before it returns.
// The deferred function call is added to a stack, and is executed in last-in-first-out (LIFO) order.
func main()  {
fmt.Println("One")
defer fmt.Println("Four")
defer fmt.Println("Three")
fmt.Println("Two")
//Prints One Two Three Four

defer fmt.Println("DB Connection close")
fmt.Println("connection open")
fmt.Println("DB manipulation")
// connection open DB manipulation DB connection close
}
