package screen

import (
  "gotris/utils"
  "gotris/model"
  "fmt"
  "strings"
  "errors"
)


type Matrix [22][26]string

type Map struct {
  matrix [][]string
  scoreMap [][]string
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
    rowsScore: rows -2 ,
    colsScore: cols - 16,
    score: 0,
  }
  m.matrix = make([][]string, m.rowsMap)
  for i := range m.matrix{
    m.matrix[i] = make([]string,m.colsMap)
  }
  m.scoreMap = make([][]string, m.rowsScore)
  for i := range m.scoreMap{
    m.scoreMap[i] = make([]string,m.colsScore)
  }

  m.populateMap()
  m.populateScore()
  return m
}

func (m *Map) Init(){

  m.DrawMap()
  // m.drawScore()
  
  drawScoreBoard(model.Coord{m.rowsScore,m.colsScore},model.Coord{m.colsMap,0},m.score)
  //So much faster, maybe the scorebard can be rendered this ugly way, since it works much better
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

func (m *Map) populateScore(){
  for  i:=0; i<m.rowsScore;i++ {
    for j:=0;j<m.colsScore;j++ {
      if (i == 0 && j == 0){
        m.scoreMap[i][j] = "┌"
      } else if (i == 0 && j == m.colsScore-1){
        m.scoreMap[i][j] = "┐"
      } else if (i == m.rowsScore -1 && j == 0){
         m.scoreMap[i][j] = "└"
      } else if (i == m.rowsScore-1 && j == m.colsScore-1){
        m.scoreMap[i][j] = "┘"
      } else if (i == 0 || i == m.rowsScore -1){
        m.scoreMap[i][j] = "-"
      } else if (j == 0 || j == m.colsScore -1){
        m.scoreMap[i][j] = "|"
      } else {
        m.scoreMap[i][j] = " "
      }
    }

  }
}

func WriteInMatrix(line *[]string, text string) error {
	if len(text) >= len(*line) {
		return errors.New("text can not have more characters than line")
	}

	for i := 0; i < len(text); i++ {
		char := string(text[i])
		(*line)[i] = char
	}

	return nil
}


func (m *Map) DrawMap(){
  utils.Debug(fmt.Sprintf("%d", m.rowsMap),5)
  for i:=0; i<m.rowsMap;i++ {
    utils.DrawLine(strings.Join(m.matrix[i],""),model.Coord{0,i},model.Coord{1,1})
  }
}

func (m *Map) drawScore(){
  for i:=0; i<m.rowsScore;i++ {
    utils.DrawLine(strings.Join(m.scoreMap[i],""),model.Coord{0,i},model.Coord{m.rowsMap+6,1})
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