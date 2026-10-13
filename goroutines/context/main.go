package main

import (
	"context"
	"fmt"
	"time"
)

// родительский контекст
func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo завершилась :(")
			return
		default:
			fmt.Println("foo прододжает выполнение")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// дочерний контекст
func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo завершилась :(")
			return
		default:
			fmt.Println("Boo прододжает выполнение")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
