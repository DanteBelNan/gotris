package utils

import (
	term "github.com/buger/goterm"
  "gotris/domain"
  "errors"
)

func draw(chr string, pos domain.Coord, offset domain.Coord){
  term.MoveCursor(pos.X + offset.X,pos.Y + offset.Y)
  term.Print(chr)
}

func DrawLine(str string, pos domain.Coord, offset domain.Coord){
  term.MoveCursor(pos.X + offset.X,pos.Y+offset.Y)
  term.Print(str)
}

func Debug(err string, counter int){
  term.MoveCursor(0,40 + counter)
  term.Print(err)
}

func DrawBox(size domain.Coord, offset domain.Coord){
  //Corners
  var tl_corner,tr_corner,bl_corner,br_corner string
  tl_corner = "┌"
  tr_corner = "┐"
  bl_corner = "└"
  br_corner = "┘"

  //Sides
  var hor_side,ver_side string
  hor_side = "|"
  ver_side = "-"

  var x,y int = 1,1
  //Draw corners
  draw(tl_corner,domain.Coord{x,y}, offset)
  draw(tr_corner,domain.Coord{size.X,y}, offset)
  draw(bl_corner,domain.Coord{x,size.Y}, offset)
  draw(br_corner,domain.Coord{size.X,size.Y}, offset)

  // Draw sides
  for (y < size.Y -1){
    y += 1
    draw(hor_side,domain.Coord{x,y},offset)
    draw(hor_side,domain.Coord{size.X,y},offset)
  }
  y = 1

  for(x < size.X - 1){
    x += 1
    draw(ver_side,domain.Coord{x,y}, offset)
    draw(ver_side,domain.Coord{x,size.Y}, offset)
  }
}


func DrawBlock(b *domain.Block, offset domain.Coord){

}

func WriteInMatrix(line []string, text string, offset int) error {
  if len(text) >= len(line) {
    return errors.New("text can not have more characters than line")
  }

  for i := 0; i < len(text); i++ {
    char := string(text[i])
    line[i + offset] = char
  }

  return nil
}

