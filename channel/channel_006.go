package main

import "fmt"

func main() {
	alphas := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	fmt.Println(alphas)
	ctrlVal := 5
	bufChanData := make(chan rune, ctrlVal)
	clearBuf := make(chan bool)
	go func(bcd chan rune) {
		for {
			select {
			case cb := <-clearBuf:
				for i := 0; i <= ctrlVal; i++ {
					fmt.Println(cb, <-bcd)
				}
			}
		}
	}(bufChanData)
	for i := 0; i <= ctrlVal; i++ {
		if i == ctrlVal {
			clearBuf <- true
			break
		}
		bufChanData <- alphas[i]
		fmt.Println("Added rune ", string(alphas[i]))
	}

}
