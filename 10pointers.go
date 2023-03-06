package main
import "fmt"

func main()  {
	//pointer is a variable that stores the memory address of another variable.
	i := 42
 	p := &i 
	//OR
	//pointer variable by prefixing the variable name with an asterisk (*)
	var pointer *int = &*p
	//the & operator is used to get the address of the i variable, and assign it to the pointer variable p.
	fmt.Println(i,p,*p,pointer,*pointer) //42 0xc000016078 42 0xc000016078 42
	(*pointer)++		  			      //43
	*pointer = *pointer + 10	    //53
	println(*pointer)			        //53
	println(*pointer + *pointer)  //53 + 53 = 106
}