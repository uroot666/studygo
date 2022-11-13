package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(2000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Hello")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
