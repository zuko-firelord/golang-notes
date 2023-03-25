//When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time. 
//This section of code that modifies shared resources is called critical section.

/*EXAMPLE
	let's assume that we have some piece of code that increments a variable x by 1.
		x = x + 1  
 	If this code accessed by a single Goroutine, there shouldn't be any problem. But, this code will fail when there are multiple Goroutines running concurrently

	The code will be executed by the system in the following steps:- 	get the current value of x
																		compute x + 1
																		assign the computed value in step 2 to x

	SCENARIO_1:
	We have assumed the initial value of x to be 0. Goroutine 1 gets the initial value of x, computes x + 1 and 
	before it could assign the computed value to x, the system context switches to Goroutine 2. Now Goroutine 2 
	gets the initial value of x which is still 0, computes x + 1. After this, the system context switches again to Goroutine 1.
	Now Goroutine 1 assigns its computed value 1 to x and hence x becomes 1. Then Goroutine 2 starts execution again and
	then assigns its computed value, which is again 1 to x and hence x is 1 after both Goroutines execute.

	SCENARIO_2:
	Goroutine 1 starts execution and finishes all its three steps and hence the value of x becomes 1. 
	Then Goroutine 2 starts execution. Now the value of x is 1 and when Goroutine 2 finishes execution, the value of x is 2.

	So from the two cases, you can see that the final value of x is 1 or 2 depending on how context switching happens. 
	This type of undesirable situation where the output of the program depends on the sequence of execution of Goroutines is called race condition.

The race condition could have been avoided if only one Goroutine was allowed to access the critical section of the code 
at any point in time. This is made possible by using MUTEX.

	 */



package main  
import (  
    "fmt"
    "sync"
    )
var x  = 0  
func increment(wg *sync.WaitGroup) {  
    x = x + 1
    wg.Done()
}
func main() {  
    var w sync.WaitGroup
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}

/*
Run this code in your local as the playground is deterministic and the race condition will not occur in the playground.
Run this program multiple times in your local machine and you can see that the output will be different for each time 
because of race condition. Some of the outputs which I encountered are final value of x 941, final value of x 928, 
final value of x 922 and so on.
*/
