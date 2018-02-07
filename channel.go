// channel is used to send data from one goroutine to another
// when data is more than it can be processed we use either buffer, select or timeout

package main

import (
	"math/rand"
	"time"
	"fmt"
)

//send  CHANNEL <- DATA
//receive VAR:= <-CHANNEL

func main()  {
	c := make(chan int)		// make(chan int, 100) building a buffer queue of size 100
	for i := 0; i < 5; i++ {
		worker := &Worker{id:i}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}

	// OR
	// using "select" when the number of incoming message is larger than the worker can handle
	//for{
	//	select{
	//	case c <- rand.Int():
	//		//optional code
	//	default:
	//		// can be empty
	//		fmt.Println("dropped")
	//	}
	//	time.Sleep(time.Millisecond * 50)
	//}

	//OR
	// using timeout
	//for{
	//	select {
	//	case c <- rand.Int():
	//	case <- time.After(time.Millisecond * 100):		// the value returns channel, which that we don't care
	//		fmt.Println("timeout")					// we just select that channel
	//	}
	//	time.Sleep(time.Millisecond * 50)
	//}

}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int)  {
	for {
		data := <- c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}

	//OR
	//for{
	//	select {
	//	case data := <- c:
	//		fmt.Printf("worker %d got %d\n", w.id, data)
	//	case <- time.After(time.Millisecond * 10):
	//		fmt.Println("Break time")
	//		time.Sleep(time.Second)
	//	}
	//}
}