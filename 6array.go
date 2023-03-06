package main
import "fmt"

func main()  {

	//method 1
	var a [5]int // declare an array of 5 integers
    a = [5]int{1, 2, 3, 4, 5} // initialize the array with values
	
	//method 2
	b := [5]int{6, 7, 3, 4, 9} // declare and initialize an array of integers

	//method 3 (undefine size)
	c := [...]int{1, 2, 3, 4, 5} //the size of the array from the number of elements in the initialization list
   	//OR
	letters := []int{1, 5, 3, 5} //slice, not a array
	c[3] = 30 //reassign value

	newarr := c
	c[2] = 40 //this will not reflect on newarray 


	fmt.Println(a,b)
	fmt.Println(len(a)) //length of array
	fmt.Println(c[3]) //print 3rd index of array "c"
	fmt.Println(c) // 1 2 40 30 5
	fmt.Println(letters) // 1 5 3 5
	fmt.Println(newarr) // 1 2 3 30 5

	matrix()
}

func matrix() 	{
// Multi-dimensional arrays
	matrix := [2][2]int{{1,5},{4,7}}
	fmt.Println(matrix)

}