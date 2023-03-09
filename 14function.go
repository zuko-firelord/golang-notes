package main
import "fmt"

func main()  {
	typicalfunc() //typical function
	multiplefunc() //Multiple Returns
}

func typicalfunc()  {
	  // return void
	  func add(x int, y int)  {
		fmt.Println("Hello, World!")
	  }
	
	  //-------arguments------return------
	  func add(x int, y int)  int {
		return x + y
	  }
	
	  //-----same type arguments-----------
	  func add(x, y int)  int {
		return x + y
	  }
}

func multiplefunc()  {
	func swap(x, y string) (string, string) {
		return y, x
	  }
	
	  //in main
	  a, b := swap("hello", "world")
	  fmt.Println(a, b) //prints "world hello"
}