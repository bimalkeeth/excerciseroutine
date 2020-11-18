package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	work:=make(chan string,1024)
	numWorker:=50000
	var wg sync.WaitGroup
	for i:=0;i<numWorker;i++ {
		go WebWorker(work,&wg)
	}
	urls:=[3]string{"http://example.com","http://reddit.com","http://twitter.com"}
	for i:=0;i<numWorker;i++{
		for _,url:=range urls{
			wg.Add(1)
			work<-url
		}

	}
	wg.Wait()
}


func WebWorker(in <-chan string,wg *sync.WaitGroup)  {
    for{
    	url :=<- in
    	res,err:=http.Get(url)
    	if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Printf("Get %s:%d\n",url,res.StatusCode)
		}
		wg.Done()
	}
}