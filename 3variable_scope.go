package main
import "fmt"

//VARIABLE SCOPE: Global,Local,Package_level

    //global_variable(pascal_case)
    var Global_variable int = 555

    //package_variable(camel_case)
    var package_Variable = 333

    //must use var keyword for variable declaration outside the function
	var c string = "hi"
	// d := "hello"(not work)

//Shadowing
var a int = 100 ; var b int

func main(){

    //local_variable(small_case)
    local_variable := 111

	var a = 101
	b = 202

fmt.Println(local_variable,a,b)
fmt.Printf("%v,%T",c,c)
}