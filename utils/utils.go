package utils

import (
	term "github.com/buger/goterm"
  "gotris/model"
)

func draw(chr string, pos model.Coord, offset model.Coord){
  term.MoveCursor(pos.X + offset.X,pos.Y + offset.Y)
  term.Print(chr)
}

func DrawBox(size model.Coord, offset model.Coord){
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
  if(false){
    draw(ver_side,model.Coord{0,0}, offset)
    draw(hor_side,model.Coord{0,0}, offset)
  }

  var x,y int = 1,1
  //Draw corners
  draw(tl_corner,model.Coord{x,y}, offset)
  draw(tr_corner,model.Coord{size.X,y}, offset)
  draw(bl_corner,model.Coord{x,size.Y}, offset)
  draw(br_corner,model.Coord{size.X,size.Y}, offset)

  // Draw sides
  for (y < size.Y -1){
    y += 1
    draw(hor_side,model.Coord{x,y},offset)
    draw(hor_side,model.Coord{size.X,y},offset)
  }
  y = 1

  for(x < size.X - 1){
    x += 1
    draw(ver_side,model.Coord{x,y}, offset)
    draw(ver_side,model.Coord{x,size.Y}, offset)
  }
}

func DrawText(text string, Pos model.Coord, offset model.Coord){
  draw(text,Pos,offset)
}

func DrawBlock(b *model.Block, offset model.Coord){
  for i := 3; i>=0;i--{
    for j:= 0;j<3;j++{
      obj := b.Elements[i][j]
      if obj.Chr != "" {
          draw(obj.Chr, obj.Pos, offset)
      }
    }
  }
}