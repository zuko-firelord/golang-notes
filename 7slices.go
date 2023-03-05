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
	
}