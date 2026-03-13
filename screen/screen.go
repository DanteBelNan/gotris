package screen

import (
  "gotris/utils"
  "gotris/model"
  "fmt"
  "strings"
)


type Matrix [22][26]string

type Map struct {
  matrix [][]string
  rowsMap int
  colsMap int 
  rowsScore int
  colsScore int
  score int
}

func NewMap(rows,cols int) *Map{
  m := &Map{
    rowsMap: rows,
    colsMap: cols,
    rowsScore: cols - 2,
    colsScore: rows - 10,
    score: 0,
  }
  m.matrix = make([][]string, rows)
  for i := range m.matrix{
    m.matrix[i] = make([]string,cols)
  }

  return m
}

func (m *Map) Init(){
  m.populateMap()
  m.drawMap()
  drawScoreBoard(model.Coord{m.rowsScore,m.colsScore},model.Coord{m.colsMap,0},m.score)
}

func (m *Map) populateMap(){
  for  i:=0; i<m.rowsMap;i++ {
    for j:=0;j<m.colsMap;j++ {
      if (i == 0 && j == 0){
        m.matrix[i][j] = "┌"
      } else if (i == 0 && j == m.colsMap-1){
        m.matrix[i][j] = "┐"
      } else if (i == m.rowsMap -1 && j == 0){
         m.matrix[i][j] = "└"
      } else if (i == m.rowsMap-1 && j == m.colsMap-1){
        m.matrix[i][j] = "┘"
      } else if (i == 0 || i == m.rowsMap -1){
        m.matrix[i][j] = "-"
      } else if (j == 0 || j == m.colsMap -1){
        m.matrix[i][j] = "|"
      } else {
        m.matrix[i][j] = " "
      }
    }
  }
}



func (m *Map) drawMap(){
  utils.Debug(fmt.Sprintf("%d", m.rowsMap),5)
  for i:=0; i<m.rowsMap;i++ {
    utils.DrawLine(strings.Join(m.matrix[i],""),model.Coord{0,i},model.Coord{1,1})
  }
}

func DrawInterface(){
  const mapSizeX = 22
  const mapSizeY = 26
  var mapBlock [mapSizeX][mapSizeY]string

  scoreSizeX,scoreSizeY := 20, 10
  score := 0

  // drawPlayground(model.Coord{mapSizeX + 2,mapSizeY + 2},model.Coord{1,0})
  drawScoreBoard(model.Coord{scoreSizeX,scoreSizeY},model.Coord{mapSizeY,0},score)
  moveItems(model.Coord{2,1},mapBlock)
}

// func drawPlayground(pos model.Coord,offset model.Coord){
//   utils.DrawBox(pos,offset)
// }

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