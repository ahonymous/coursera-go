package module2

import (
	"fmt"
	"sync"
)

var locker sync.Mutex
var waiter sync.WaitGroup

func incrementCounter(counter *int) {
	for i := 0; i < 1000000; i++ {
		locker.Lock()
		*counter++
		locker.Unlock()
	}
	waiter.Done()
}

func Start() {
	counter := 0
	waiter.Add(2)

	fmt.Println("Start two counter incrementing with goroutines... ", counter)

	go incrementCounter(&counter)
	go incrementCounter(&counter)

	waiter.Wait()

	fmt.Printf("Final Counter Value: %d\n", counter)
}
