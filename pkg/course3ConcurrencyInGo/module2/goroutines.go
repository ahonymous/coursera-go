package module2

import (
    "fmt"
    "sync"
)

var locker sync.Mutex

func incrementCounter(counter *int, waiter *sync.WaitGroup) {
    for i := 0; i < 1000000; i++ {
        locker.Lock()
        *counter++
        locker.Unlock()
    }
    waiter.Done()
}

func Start() {
    counter := 0
    waiter := sync.WaitGroup{}

    fmt.Println("Start two counter incrementing with goroutines... ", counter)

    for i := 0; i < 2; i++ {
        waiter.Add(1)
        go incrementCounter(&counter, &waiter)
    }
    waiter.Wait()

    fmt.Printf("Final Counter Value: %d\n", counter)
}
