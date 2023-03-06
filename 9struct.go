package main
import "fmt"
//struct is a composite data type that groups different types together into a single entity.
//Each value in a struct is called a field, and the fields are defined with a name and a type.


func main()  {

	type Person struct {
		Name    string
		Age     int
		Address struct {
			Street  string
			City    string
			State   string
			ZipCode string
		}
	}
	
	//Method 1
	var john Person
	john.Name = "John"
	john.Age = 30
	john.Address.Street = "123 Main St"
	john.Address.City = "Anytown"
	john.Address.State = "CA"
	john.Address.ZipCode = "12345"

	//Method 2
	jane := Person{
		Name: "Jane",
		Age: 25,
		Address: struct { //Address struct is defined as an anonymous struct inside the Person struct
			Street  string
			City    string
			State   string
			ZipCode string
		}{
			Street: "456 Elm St",
			City: "Othertown",
			State: "NY",
			ZipCode: "67890",
		},
	}

	gane := jane

	fmt.Println(john.Name)           //John
	fmt.Println(jane.Address.State)  //NY
	fmt.Println(gane) //this will not use same memory location of jane variable

}