package main

import (
	"fmt"
	"sync" //there are also sync/atomic
	"time"
)

//mutex = mutual exclusive lock , or using syncronization primitive using sync/atomic
//thread and lock is hard to find and fix the error and sometimes can't be reproduce
func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
