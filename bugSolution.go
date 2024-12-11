```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel to prevent race condition

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        ch <- i
                }(i)
        }

        wg.Wait() // Wait for all goroutines to finish sending
        close(ch) // Close the channel after all values are sent

        for i := range ch {
                fmt.Println(i)
        }
}
```