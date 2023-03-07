package main
import "fmt"

func main()  {
	//There are 3 forms of for loops, also there is no a while loop syntax in GO (instead it is a form of for loops), also there is no do-while at all


	//For loop
	for j := 7; j <= 9; j++ { /*stuff*/ }

	//While like for loop
	i := 1
	for i <= 3 { /*stuff*/ i++ }

	//Infinite Loop : While(true)
	for { /*stuff*/ if (/*stuff*/) break }

	//For-Each loop
	for i, v := range arr {	/*stuff*/ }
	for _, v := range arr {	/*stuff*/ }
	for i, _ := range arr {	/*stuff*/ }

}