package main

import (
	"fmt"
	"time"
)

func main() {

	for range time.Tick(time.Second *1){
		fmt.Println("Engine Roaring")
	}



}
