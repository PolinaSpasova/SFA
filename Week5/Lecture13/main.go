package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	context.CancelFunc
	buffer     chan string
	bufferSize int
	/* Add other fields you might need */
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	var new BufferedContext
	new.Context = ctx
	new.buffer = make(chan string, bufferSize)
	new.bufferSize = bufferSize
	new.CancelFunc = cancel
	return &new
}
func (bc *BufferedContext) Done() <-chan struct{} {

	if bc.CancelFunc != nil {
		return bc.Context.Done()
	}
	if len(bc.buffer) == bc.bufferSize {
		return bc.Context.Done()
	}
	return nil
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fn(bc.Context, bc.buffer)
}

func main() {
	ctx := NewBufferedContext(time.Second, 10)
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 500) // try different values here
				fmt.Println("bar")
			}
		}
	})
}
