package main
import "fmt"

func main()  {
	typicalfunc() //typical function

	a, b := multipleReturn("hello", "world") //Multiple Returns
	fmt.Println(a, b) //prints "world hello"

	x,y := nameReturn(9) //Name Return
	fmt.Println(x,y) //4 5

	avg := average(10,20,30,40,50) // Variadic Functions / Variable arguments list
    println("Average:", avg)

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


	func multipleReturn(x, y string) (string, string) {
		return y, x
	  }
	
	
	 func nameReturn(sum int) (x, y int) {
		//You can declare return variables and name them at the beginning, they are returned in the end.
		//*you can override the returns and return whatever you want at the return statement.
		//Returns x,y at the end.
		x = sum * 4 / 9
		y = sum - x
		return
		//return a,b  <- u can override default return of x,y.
	  }

	  func average(x int, values ...int) float64{
		// x value here has no use for average. just illustrates the idea of having arguments
 		// then a variable number of arguments
		//print values
		fmt.Println("Single argument value: ", x)
		fmt.Println("Variable argument values: ", values)
	
		//calculate average
		total := 0
		for _, value := range values {
		  total += value
		}
	
		return float64(total) / float64(len(values))
	  }
	
