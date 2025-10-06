# Signal Package

The signal package provides utilities for handling OS signals in Go applications, particularly for graceful shutdowns.

## Installation

```go
import "github.com/wxnacy/gotool/signal"
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/wxnacy/gotool/signal"
)

func main() {
    // Create a new stop channel
    stopChan := signal.NewStopChan()
    defer stopChan.Stop()

    fmt.Println("Application started. Press CTRL+C to stop...")

    // Your main application logic
    for i := 0; i < 1000; i++ {
        // Check if a stop signal was received
        if stopChan.Check() {
            fmt.Println("Stop signal received, shutting down gracefully...")
            return
        }
        
        // Do some work
        fmt.Printf("Working... %d\n", i)
    }
}
```

### Advanced Usage with Goroutines

```go
package main

import (
    "fmt"
    "time"
    "github.com/wxnacy/gotool/signal"
)

func worker(id int, stopChan signal.StopChan, done chan bool) {
    for {
        // Check if a stop signal was received
        if stopChan.Check() {
            fmt.Printf("Worker %d received stop signal\n", id)
            done <- true
            return
        }
        
        // Do some work
        fmt.Printf("Worker %d is working...\n", id)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    // Create a new stop channel
    stopChan := signal.NewStopChan()
    defer stopChan.Stop()

    // Create done channel to wait for workers
    done := make(chan bool, 2)

    // Start workers
    go worker(1, stopChan, done)
    go worker(2, stopChan, done)

    fmt.Println("Workers started. Press CTRL+C to stop...")

    // Wait for all workers to finish
    for i := 0; i < 2; i++ {
        <-done
    }

    fmt.Println("All workers stopped.")
}
```

## API

### `NewStopChan() StopChan`

Creates and initializes a signal listening channel.

- Listens for SIGINT and SIGTERM signals
- Returns a StopChan that can be used to check for signals

### `StopChan.Stop()`

Stops listening for signals.

### `StopChan.Check() bool`

Checks if a stop signal has been received.

- Returns `true` if a signal was received, `false` otherwise
- Prints "收到终止信号，程序退出..." when a signal is received

## Examples

### Example 1: Simple Loop with Signal Handling

```go
package main

import (
    "fmt"
    "time"
    "github.com/wxnacy/gotool/signal"
)

func main() {
    stopChan := signal.NewStopChan()
    defer stopChan.Stop()

    fmt.Println("Starting process... Press CTRL+C to stop")

    i := 0
    for {
        if stopChan.Check() {
            fmt.Println("Process stopped by user")
            break
        }

        fmt.Printf("Processing item %d\n", i)
        i++
        time.Sleep(500 * time.Millisecond)
    }

    fmt.Println("Process finished")
}
```

### Example 2: Batch Processing with Signal Handling

```go
package main

import (
    "fmt"
    "time"
    "github.com/wxnacy/gotool/signal"
)

func processItem(item int) {
    fmt.Printf("Processing item %d...\n", item)
    time.Sleep(100 * time.Millisecond) // Simulate work
    fmt.Printf("Item %d processed\n", item)
}

func main() {
    stopChan := signal.NewStopChan()
    defer stopChan.Stop()

    items := make([]int, 100)
    for i := range items {
        items[i] = i
    }

    fmt.Println("Starting batch processing... Press CTRL+C to stop")

    for i, item := range items {
        // Check for stop signal before processing each item
        if stopChan.Check() {
            fmt.Printf("Stopped at item %d/%d\n", i, len(items))
            return
        }

        processItem(item)
    }

    fmt.Println("Batch processing completed")
}
```