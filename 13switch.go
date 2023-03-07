package main
import "fmt"

func main()  {
	//Switch statements in GO doesn't require break
	//fallthrough keyword used to go to NEXT statement even if condition doesn't match
	//fallthrough is like a break so no code can be after it.

	var i interface{} = 2
	fmt.Println("Switch for i = ", i, " goes to: ")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
			i = 4
			fallthrough //goes to NEXT case even if doesn't match.
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5,6: //mutiple case
			fmt.Println("five or six")
		default:
			fmt.Println("default")
			//Switch for i =  2  goes to: 
			//two
			//three
}
	switch i.(type) {
	case int:
		fmt.Println("INT")
	case float32:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
			//INT
	}
} 	