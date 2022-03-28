package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string, bufferLimit)
	out := make(chan string)
	go func() {
		for {
			if len(channel) == bufferLimit {
				for i := 0; i < bufferLimit; i++ {
					out <-<-channel
				}
				<-time.After(clearInterval)
			}
			channel <- data
		}
	}()
	return out
}

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}
