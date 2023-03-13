package main
import "fmt"
// A new method with the same name as an existing method but associated with a different type of receiver.
// When this happens, the new method replaces the old one, effectively "overriding" it.
type Person struct {
    First string
    Last string
    Age int
   }

   type Employee struct {
    Person
    ID string
    Salary int
   }

   func (p Person) FullName() string{
    return p.First + " " + p.Last
   }

   //Override
   func (p Employee) FullName() string{
    return p.ID + " " + p.First + " " + p.Last
   }


   func main() {
    x := Employee{
      Person{
        "sach",
        "tedwa",
        22},
      "0ID12000ID",
      99999,
    }

fmt.Println(x)
fmt.Println(x.Person.FullName()) //sachin tedwa		
fmt.Println(x.FullName()) 		   //0ID12000ID sachin tedwa
   }