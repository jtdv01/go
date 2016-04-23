package main

/*
	Since counter variable is shared and not locked
	this program may not always be correct
*/

import(
	"fmt"
	"time"
)

// shared variable
var counter =0

func inc(){
	counter ++
	fmt.Println(counter)
}

func main(){
	for i:=0; i<100; i++{
		go inc()
	}	
	
	time.Sleep(time.Millisecond * 10)	
}
