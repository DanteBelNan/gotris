package main

import (
  "gotris/screen"
  term "github.com/buger/goterm"
  "time"
)


func main(){
  

  term.Clear()

  for {
    term.Flush()
    screen.DrawInterface()

    time.Sleep(time.Second)
  }
}