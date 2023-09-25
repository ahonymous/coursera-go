package module4

import (
    "fmt"
    "sync"
    "time"
)

const numPhilosophers = 5
const numMeals = 3
const maxConcurrentEaters = 2

type Chopstick struct {
    sync.Mutex
}

type Philosopher struct {
    id             int
    leftChopstick  *Chopstick
    rightChopstick *Chopstick
    host           chan struct{}
}

func (p Philosopher) eat() {
    for i := 0; i < numMeals; i++ {
        p.host <- struct{}{}

        p.leftChopstick.Lock()
        p.rightChopstick.Lock()

        fmt.Printf("Philosopher %d starting to eat\n", p.id)
        time.Sleep(time.Millisecond * 100) // Simulate eating
        fmt.Printf("Philosopher %d finishing eating\n", p.id)

        p.leftChopstick.Unlock()
        p.rightChopstick.Unlock()

        <-p.host
    }
}

func Start() {
    startTime := time.Now()
    chopsticks := make([]*Chopstick, numPhilosophers)
    philosophers := make([]Philosopher, numPhilosophers)

    for i := 0; i < numPhilosophers; i++ {
        chopsticks[i] = new(Chopstick)
    }

    host := make(chan struct{}, maxConcurrentEaters)

    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = Philosopher{
            id:             i + 1,
            leftChopstick:  chopsticks[i],
            rightChopstick: chopsticks[(i+1)%numPhilosophers],
            host:           host,
        }
    }

    var wg sync.WaitGroup
    for _, philosopher := range philosophers {
        wg.Add(1)
        go func(p Philosopher) {
            p.eat()
            wg.Done()
        }(philosopher)
    }

    wg.Wait()
    close(host)
    fmt.Printf("The whole meal lasted sinse %f seconds.\n", time.Since(startTime).Seconds())
}
