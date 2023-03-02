// type conversion is the process of converting a value from one type to another.
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var i float64 = 127.4
	f := int8(i)
	//OR
	var g float32 = float32(f)
	fmt.Println(reflect.TypeOf(f),f)
    fmt.Println(reflect.TypeOf(g),g)
	intconv()
	stringconv()
	floatconv()
}

//int to string
func stringconv(){
	integer := 42
    i := strconv.Itoa(integer)
	fmt.Println(reflect.TypeOf(i),i)
}

// string to int
func intconv(){
	str := "42"
    i, err := strconv.Atoi(str)
	if err != nil {
        panic(err)
    }
	fmt.Println(reflect.TypeOf(i),i)
}

//STRING TO FLOAT
func floatconv(){
	str := "3.14"
	f, err := strconv.ParseFloat(str, 64) //The second argument 64 specifies the bit size of the resulting floating-point value.
	if err != nil {
		// handle error
	}
	fmt.Println(reflect.TypeOf(f),f)
}



// GO PRIMITIVE TYPES
	// The possible values for bool are true and false.
	// uint8 : unsigned 8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed 8-bit integers (-128 to 127)
	// int16 : signed 16-bit integers (-32768 to 32767)
	// int32 : signed 32-bit integers (-2147483648 to 2147483647)
	// int64 : signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	// int is either int64 or int32 depends on the implementation.
	// float32 : set of all IEEE-754 32-bit floating-point numbers
	// float64 : set of all IEEE-754 64-bit floating-point numbers
	// complex64 the set of all complex numbers with float32 real and imaginary parts
	// complex128 the set of all complex numbers with float64 real and imaginary parts
	// byte alias for uint8 rune alias for int32