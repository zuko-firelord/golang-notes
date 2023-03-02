package main
import "fmt"

//iota in Go, is a value used within the const block, its value starts at 0 per block, and increment each time it is used again
const(
	 x = iota + 6 //iota == 0
	 _            // use _ to skip
	 y = iota     //iota == 2
	 z            //iota == 3
	 _            //iota == 4
	 c            //iota == 5
)

const(
	a = iota
	b = iota
)

func main(){
	fmt.Println(x,y,z,c) //6 2 3 5
	fmt.Println(a,b)     //0 1
}