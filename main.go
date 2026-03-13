package main

import (
  "gotris/screen"
  term "github.com/buger/goterm"
  "time"
)


func main(){
  

  term.Clear()
  Map := screen.NewMap(22,26)
  Map.Init()

  for {
    term.Flush()
    Map.DrawMap()
    // screen.DrawInterface()

    time.Sleep(time.Second)
  }
}