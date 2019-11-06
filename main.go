package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Start processing on", hostname, "host...")
	} else {
		fmt.Println("Cannot parse hostname:", err.Error())
		fmt.Println("Start processing...")
	}
	for {
		randomNumber := rand.Uint32()
		duration := time.Duration(randomNumber)
		fmt.Println("Waiting for", duration)
		time.Sleep(duration)
		if time.Millisecond * 500 < duration && duration < 2 * time.Second {
			fmt.Println("Found proper random number:", randomNumber)
			fmt.Println("Exiting with duration", duration, "which is in between 500ms and 1s")
			break
		}
	}
}
