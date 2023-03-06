package main
import "fmt"

//Slices are a dynamic version of arrays

func main()  {
	//A slice does not store any data, it just describes a section of an underlying array
	names := []string{"RAM", "SAM", "JOHN", "DAN"}
	fmt.Println(names)   //[RAM SAM JOHN DAN]
	a := names[:2]
    b := names[1:3]
	c := names[1:]
    fmt.Println(a, b,c)    //[RAM SAM] [SAM JOHN] [SAM JOHN DAN]

	//Changing the elements of a slice modifies the corresponding elements of its underlying array
	b[0] = "XXX"
	fmt.Println(a, b)    //[RAM XXX] [XXX JOHN] 
	fmt.Println(names)   //[RAM XXX JOHN DAN]

	// create a new slice with make
    s := make([]int, 3, 5) // make function takes three arguments: type of slice, length and capacity
	// The capacity is the maximum number of elements that the slice can hold
	fmt.Println(s, len(s), cap(s)) //[0 0 0] 3 5

	iteration() //Iterating over a Slice
	sliceAppend()    //Appending to Slice
	
}

func iteration()  {
	
	mySlice := []string{"apple", "banana", "cherry"}

	for index, value := range mySlice { //use "_" in place of the index or value variable if you don't need it. 
		fmt.Println(index, value) //0 apple
								  //1 banana
								  //2 cherry
	}

}

func sliceAppend()  {
	// Append function use to add elements to a slice
	// Append function takes two or more arguments: 
		//1) slice to which you want to add elements
		//2) elements you want to add

	mySlice := []string{"apple", "banana", "cherry"}
	mySlice = append(mySlice, "date", "elderberry")	
	fmt.Println(mySlice) //[apple banana cherry date elderberry]

	//use the ... syntax to append a slice of elements to another slice.
	MySlice := []int{1,2,3,4}
	MySlice2 := append(MySlice,1000)
	MySlice3 := append(MySlice,MySlice2...)
	fmt.Println(MySlice3) //[1 2 3 4 1 2 3 4 1000]

	// When adding elements to a slice using the append function in Go, a new underlying array may 
	// need to be allocated if there is not enough capacity in the existing slice. If a new array is
	// allocated, modifying elements in the new slice returned by append will not modify the original
	// slice. To ensure that the original slice is modified when appending elements, always use the
	// same variable on the left-hand side of the assignment.

	  //Allocate new capacity  
	  s := make([]int, 5, 5)
	  x := append(s, 1, 2 ,3)
	  x[0] = 1337
	  s[0] = 6800
	  fmt.Println(s,x) //[6800 0 0 0 0] [1337 0 0 0 0 1 2 3]
	
	  //Doesn't allocat new capacity and return reference
	  f := make([]int, 5, 150)
	  z := append(f, 1, 2 ,3)
	  z[0] = 1337
	  f[0] = 6800
	  fmt.Println(f,z) //[6800 0 0 0 0] [6800 0 0 0 0 1 2 3]
					   //notice that 1337 is overwritten

	cut := append(z[:2], z[5:]...) //Cutting elements from a slice
	fmt.Println(cut) // [6800 0 1 2 3]

	copy := append([]int{}, z[2:4]...) //Copy element from slice
	fmt.Println(copy,z) //[1 2] [6800 0 1 2 3 1 2 3]

	deleteslice := append(z[:2], z[2+1:]...) //Deleting elements from a slice
	fmt.Println(deleteslice) //[6800 0 2 3 1 2 3]	

}