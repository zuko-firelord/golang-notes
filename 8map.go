package main
import "fmt"
//Map is a collection of key-value pairs, where each key is unique and associated with a value.

func main()  {
	// To create an empty map, use the builtin `make`:

	//METHOD 1
	m := make(map[string]int) //The map is currently nil, meaning it has not been initialized yet.
	m = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	//METHOD 2: Declaration  with intialization
	M := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 30,
	}
	
	// Get a value for a key with `name[key]`.
	fmt.Println(m["apple"]) //1

	//value reassign
	m["apple"] = 100
	fmt.Println(m["apple"]) //100

	//add new key value
	m["grape"] = 200

	// The builtin `len` returns the number of key/value
	fmt.Println(len(m)) //4

	// The builtin `delete` removes key/value pairs from
	delete(m, "orange")

	//Note that when we add new key, the changes are reflected in the original map as well.
	// This is because the pointer stored in the map points to the same location.
	newM := M
	newM["grape"] = 52
	fmt.Println(newM,M) //[apple:10 banana:20 grape:52 orange:30] map[apple:10 banana:20 grape:52 orange:30]


	fmt.Println(m,M)
	retrievevalue()
}

//retrieve a value from a map using a key
func retrievevalue()  {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Charlie": 35,
	}
	
	//The first return value is the value associated with the key
	//the second return value is a boolean that indicates whether the key was present in the map or not.
	age, ok := ages["David"]
	if ok {
		fmt.Println("David's age is", age)
	} else {
		fmt.Println("David is not in the map")
	}
	
	//In case where you only care about the second return value, you can use the" _ " to ignore the first return value
	_ , check := ages["Sam"]
	fmt.Println("Sam:", check) //false
}