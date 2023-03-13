package main
import "fmt"
// Receiver is a parameter of a method that associates the method with a specific type.
// Receiver is declared before the method name, and it can be either a "value receiver" or a "pointer receiver".
// Value receiver: Receiver with a non-pointer type. It receives a copy of the struct value and cannot modify the original struct value.
// pointer receiver: Receiver with a pointer type. It receives a pointer to the struct value and can modify the original struct value.

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 { //Value receiver
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 { //Pointer receiver
	return 2 * (r.width + r.height)
}
func main() {
r := Rectangle{width: 10, height: 5}
area := r.Area()
fmt.Println(area)

R := &Rectangle{width: 10, height: 5}
perimeter := R.Perimeter()
fmt.Println(perimeter)

}