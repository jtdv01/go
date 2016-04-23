package main

/*
	Counter variable is now locked
	So it should always be correct this time
*/

import(
	"fmt"
	"time"
	"sync"
)

// shared variable now contains a lock
var (
	counter=0
	lock sync.Mutex
)

func inc(){
	// Lock then unlock once the function is finished
	lock.Lock()
	defer lock.Unlock()	
	
	counter++
	fmt.Println(counter)
}

func main(){
	for i:=0; i<100; i++{
		go inc()
	}	
	
	time.Sleep(time.Millisecond * 10)	
}
