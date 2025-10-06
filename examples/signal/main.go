package main

import (
	"fmt"
	"time"

	"github.com/wxnacy/go-tools/signal"
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
