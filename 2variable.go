package main
import "fmt"



func main(){
	// unused variables always give an error
    //Uninitialized variables are given its zero value (e.g int = 0, string = "", bool = false)


	//1st method 
    var a int //declaration
	a = 55    //initialization

    //2nd method
	var b int = 5

    //3rd method
	var c = 100

	//4th method
	d := 500 



	fmt.Println(a,b,c,d)
}