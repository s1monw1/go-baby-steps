package main

import (
	"fmt"
	"log"
	"time"
)

// delays for the given duration to then print the printable to the console
// once finished, true is put on the callback chan
func delayAndPrint(delay time.Duration, printable string, callback chan bool) {
	time.Sleep(delay * time.Millisecond)
	fmt.Println(printable)
	callback <- true
}
func main() {
	// create a callback channel to communicate each goroutines state
	finished := make(chan bool)

	start := time.Now()
	go delayAndPrint(0, "Hello ", finished)
	go delayAndPrint(200, "world!", finished)
	// await both goroutines to finish by consuming both expected messages from the channel
	<-finished
	<-finished
	elapsed := time.Since(start)

	log.Printf("goroutines took %s", elapsed)
}
