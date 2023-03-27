package main
/*
If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. 
If this does not happen, then the program will panic at runtime with Deadlock
*/
import (
	"fmt"
	"sync"
)

//Channels are the way thought which Goroutines communicate. Data can be sent from one end and received from the other end using channels.

/*
Channels are unbuffered by default, which means that sending a value to a channel will block the sender until the value is received
by the receiver. Channels can also be buffered, which allows sending multiple values without blocking the sender until the buffer is full.
*/

func main()  {
	ch := make(chan int) // Create an unbuffered channel of type int, which means that it can only hold one value at a time.

	/*fmt.Println(<-ch)  
	When we try to send a value to an unbuffered channel using the <- operator, it will block until the value is received by another goroutine.
	*/
	go func() {
		ch <- 42 // Send/write  a value to the channel
	}()
	result := <- ch // Receive/read the value from the channel
	fmt.Println(result) // Output: 42

}

func main()  {
	mych:= make(chan int,2) //Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer
	wg := &sync.WaitGroup{}
wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup )  {
		fmt.Println(<-ch) //42
		fmt.Println(<-ch) //50
		fmt.Println(<-ch) //100
		wg.Done()
	}(mych,wg)

	go func(ch chan int, wg *sync.WaitGroup )  {
		ch <- 42
		ch <- 50
		ch <- 100
		wg.Done()
	}(mych,wg)

	wg.Wait()
}

	
func main() {  
	// Closing channels and for range loops on channels
		ch := make(chan int) //unbuffered channel of type int 
		go func (chnl chan int) {  //Goroutine writes 0 to 9 to the chnl channel 
			for i := 0; i < 10; i++ {
				chnl <- i
			}
			close(chnl) // then closes the channel
		}(ch)
		for {
			v, ok := <-ch
			if ok == false { // the status of the channel (whether it's still open or closed)
				break
			}
			fmt.Println("Received ", v, ok) 
			/*Received  0 true  
			Received  1 true  
			Received  2 true  
			Received  3 true  
			Received  4 true  
			Received  5 true  
			Received  6 true  
			Received  7 true  
			Received  8 true  
			Received  9 true  
			*/
		}
	
	}

func main()  {
	//A uni-directional channel is a channel that is restricted to send or receive operations only.
	 
		sendch := make(chan int)

		go func (sendch chan<- int) {  // chan<- int syntax indicates a channel that can only be used for sending integers
			sendch <- 10
		}(sendch)
		fmt.Println(<-sendch)// <-chan int indicates a channel that can only be used for receiving integers.	
}


	