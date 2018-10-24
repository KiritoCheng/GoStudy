package main

import "fmt"
import "sync"
import "runtime"

var counter = 0

func Count(lock *sync.Mutex,x,y int){
	lock.Lock()
	counter++
	Add(x,y)
	lock.Unlock()
}

func Add(x,y int){
	z := x + y
	fmt.Println(z)
}

func main(){
	lock := &sync.Mutex{} 

	for i := 0;i<10;i++ {
		go Count(lock,i,i)
	}

	for{
		lock.Lock() 
		c:=counter
		lock.Unlock()

		runtime.Gosched()
		if c>10{
			break
		}
	}
}