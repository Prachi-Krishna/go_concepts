// When a program runs concurrently, the parts of code which modify shared resources
//should not be accessed by multiple Goroutines at the same time. This section of code
//that modifies shared resources is called critical section.

//A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is
//running the critical section of code at any point in time to prevent race conditions
//from happening.

package main

import (
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

}
