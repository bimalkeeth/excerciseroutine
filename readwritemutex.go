package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	Mutex *sync.RWMutex
	round map[string]int
}

func newData()*data{
	d:=make(map[string]int)
	return &data{
		Mutex: &sync.RWMutex{},
		round: d,
	}
}
func(d *data)update(wid string){
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	count,ok:=d.round[wid]
	if !ok{
		fmt.Println("error occurred")
		return
	}
	d.round[wid]=count+1
}
func dowork(wid string,d *data,wg *sync.WaitGroup){
	for range time.Tick(time.Second*2){
		d.update(wid)
	}
	wg.Done()
}

func getdata(d *data){
	for range time.Tick(time.Second*2){
		d.Mutex.RLock()
		fmt.Println(d)
		d.Mutex.RUnlock()
	}
}

func main() {
   var wg sync.WaitGroup
   d:=newData()
   d.round["one"]=0
   d.round["two"]=0
   d.round["three"]=0

   go dowork("one",d,&wg);wg.Add(1)
   go dowork("two",d,&wg);wg.Add(1)
   go dowork("three",d,&wg);wg.Add(1)

   go getdata(d)
   wg.Wait()

}
