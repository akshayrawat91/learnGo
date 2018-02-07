//	readers can be as much you want but writes should be synchronized
package main

import (
	"sync"
	"time"
	"fmt"
)

var (
	counter = 0
	lock sync.Mutex		// using mutex for atomic operations
)

func main(){
	for i:=0; i < 20; i++{
		go inc()		// concurrency
	}
	time.Sleep(time.Millisecond * 10)	// shouldn't use it (used just for testing)
}

func inc(){
	lock.Lock()				// lock when write is being performed
	defer lock.Unlock()		// unlock once the function completes
	counter++
	fmt.Println(counter)

}

