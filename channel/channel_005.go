package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	data := make(chan bool)
	die := time.After(time.Second * 3)

	go func(b chan bool) {
		for {
			select {
			case d := <-data:
				fmt.Println(d, "got data")
			case <-die:
				fmt.Println("killing go routine")
				return
			default:
				fmt.Println("chilling")
				time.Sleep(time.Millisecond * 500)
			}
		}

	}(data)

	go func(b chan bool) {
		time.Sleep(time.Millisecond * 2)
		data <- true
		close(data)
	}(data)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	for {
		<-c
		fmt.Println("killing")
		break
	}

}
