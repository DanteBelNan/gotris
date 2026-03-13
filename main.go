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
    Map := screen.NewMap(22,26)
    Map.Init()
    // screen.DrawInterface()

    time.Sleep(time.Second)
  }
}