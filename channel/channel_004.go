package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan string)
	until := time.After(time.Second * 5)
	done := make(chan bool)
	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println(d)
			case <-until:
				done <- true
			}
		}
	}()
	for {
		select {
		case <-done:
			close(data)
			fmt.Println("channel closed")
			return
		default:
			data <- "lo"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
