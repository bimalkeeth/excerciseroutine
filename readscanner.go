package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read(r io.Reader)<-chan string{
	lines:=make(chan string)
	go func(){
		defer close(lines)
		scan:=bufio.NewScanner(r)
		for scan.Scan(){
		 lines<-scan.Text()
		}
	}()
	return lines
}

func main() {
	mes:=read(os.Stdin)
	for anu:=range mes{
		fmt.Println("Message is out",anu)
	}

}
