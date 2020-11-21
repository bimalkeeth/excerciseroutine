package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second*5, cancel)

	done := ctx.Done()
	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("program is ending")
			return
		case <-time.After(time.Second):
			fmt.Println("tick", i)
		}
	}

}
