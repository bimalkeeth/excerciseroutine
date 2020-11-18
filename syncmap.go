package main

import "sync"

func main() {
  m:=sync.Map{}
  m.Store("","")

}
