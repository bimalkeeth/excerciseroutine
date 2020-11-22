package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(time.Second * 10)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdin.Write(buf)
		case <-done:
			fmt.Println("time out")
			close(echo)
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)

		data = bytes.Trim(data, "\x00")
		s := "\n"
		data = append(data, s...)

		length := int64(len(data))
		fmt.Println(len)
		if length > 2 && l > 1 {
			out <- data
		}
	}
}
