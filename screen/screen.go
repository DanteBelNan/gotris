package screen

import (
  "gotris/utils"
  "gotris/model"
)


func DrawInterface(){
  drawPlayground(model.Coord{24,28},model.Coord{1,0})
  drawScoreBoard(model.Coord{20,10},model.Coord{26,0})
}

func drawPlayground(pos model.Coord,offset model.Coord){
  utils.DrawBox(pos,offset)
}

func drawScoreBoard(pos model.Coord,offset model.Coord){
  utils.DrawBox(pos,offset)
  utils.DrawText("STATS: ", model.Coord{2,2},offset)
  utils.DrawText("NEXT ITEM: ", model.Coord{2,3},offset)

}