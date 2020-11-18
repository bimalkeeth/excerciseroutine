package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/",index)
	err:=http.ListenAndServe(":8088",nil)
	if err!=nil{
		log.Fatal("Error is in the system")
	}

}

func index(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Hello Cloud Native Go")
}
