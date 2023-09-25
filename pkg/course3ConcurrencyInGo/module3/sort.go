package module3

import (
    "fmt"
    "math"
    "sort"
    "sync"
)

const n = 4

// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
//
// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

func Start() {
    arr := getMainArray()

    chunks := chunkSlice(arr, int(math.Ceil(float64(len(arr))/float64(n))))
    ch := make(chan any, n)

    wg := &sync.WaitGroup{}

    wg.Add(n)
    for i, chunk := range chunks {
        go sortChunk(chunk, i, ch, wg)
    }
    wg.Wait()
    close(ch)

    fmt.Println("Sorted array:", mergeSorted(ch))
    fmt.Println("Main goroutine is done")
}

func sortChunk(a []int, i int, ch chan any, wg *sync.WaitGroup) {
    fmt.Printf("Sorting chunk #%d in goroutine: %v\n", i, a)
    sort.Ints(a)
    fmt.Printf("Sorted chunk #%d in goroutine: %v\n", i, a)

    ch <- a
    wg.Done()
}

func mergeSorted(chunks chan any) []int {
    var merged []int

    for chunk := range chunks {
        merged = append(merged, chunk.([]int)...)
    }
    sort.Ints(merged)

    return merged
}

func chunkSlice(a []int, chunkSize int) [][]int {
    var chunks [][]int

    fmt.Println("Chunk size is :", chunkSize)
    for {
        if len(a) == 0 {
            break
        }

        if len(a) < chunkSize {
            chunkSize = len(a)
        }

        chunks = append(chunks, a[0:chunkSize])
        a = a[chunkSize:]
    }

    return chunks
}

func getMainArray() []int {
    var x int

    fmt.Print("Enter the number of integers: ")
    for {
        _, err := fmt.Scan(&x)
        if nil != err || x < 1 {
            fmt.Println("Please enter a positive integer")
            continue
        }

        break
    }

    arr := make([]int, x)
    fmt.Printf("Enter %d integers separated by spaces: ", x)
    for i := 0; i < x; i++ {
        _, err := fmt.Scan(&arr[i])
        if nil != err {
            fmt.Printf("Please enter valid %d integers separated by spaces: ", x)
        }
    }
    fmt.Println("Unsorted array:", arr)

    return arr
}
