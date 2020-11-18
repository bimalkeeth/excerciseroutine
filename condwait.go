package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	Action *sync.Cond
}

func main() {

	obj:=Object{Action: sync.NewCond(&sync.Mutex{})}

	var attachListener func(cd *sync.Cond,fn func())

	attachListener = func(cd *sync.Cond,fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func(){

			wg.Done()
			cd.L.Lock()
			defer cd.L.Unlock()
			cd.Wait()
			fn()
            go attachListener(cd,fn)

		}()
		wg.Wait()
	}
	attachListener(obj.Action, func() {
		fmt.Println("No i feel like a javascript thing:fire one")
	})
	attachListener(obj.Action, func() {
		fmt.Println("No i feel like a javascript thing:fire two")
	})
	attachListener(obj.Action, func() {
		fmt.Println("No i feel like a javascript thing:fire three")
	})

	for range time.Tick(time.Second * 2){
		obj.Action.Signal()
	}
}
