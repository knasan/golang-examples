package main

import "fmt"

func myPanic() {
  panic("test panic")
}

func myRecover() {
  if r := recover(); r != nil {
    fmt.Println("Panic From:", r)
  }
}


func main() {
  defer myRecover()
  myPanic()
}
