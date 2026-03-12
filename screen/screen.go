package screen

import (
  "gotris/utils"
  "gotris/model"
  "fmt"
)

type Matrix [22][26]string

func DrawInterface(){
  const mapSizeX = 22
  const mapSizeY = 26
  var mapBlock [mapSizeX][mapSizeY]string

  scoreSizeX,scoreSizeY := 20, 10
  score := 0

  drawPlayground(model.Coord{mapSizeX + 2,mapSizeY + 2},model.Coord{1,0})
  drawScoreBoard(model.Coord{scoreSizeX,scoreSizeY},model.Coord{mapSizeY,0},score)
  moveItems(model.Coord{2,1},mapBlock)
}

func drawPlayground(pos model.Coord,offset model.Coord){
  utils.DrawBox(pos,offset)
}

func drawScoreBoard(pos model.Coord,offset model.Coord, score int){
  utils.DrawBox(pos,offset)
  utils.DrawText("STATS: ", model.Coord{2,2},offset)
  utils.DrawText("NEXT ITEM: ", model.Coord{2,3},offset)
  block := model.NewDashBoardBlock(model.Coord{pos.X+offset.X,1})
  utils.DrawBlock(block,model.Coord{-4,2})
  scoreText := fmt.Sprintf("Score: %v", score)
  utils.DrawText(scoreText, model.Coord{2,9},offset)
}

func moveItems(offset model.Coord, mapBlock Matrix){
  // block := model.NewBlock()
  // block.DrawBlock(offset)
}