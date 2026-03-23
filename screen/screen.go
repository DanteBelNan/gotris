package screen

import (
  "gotris/utils"
  "gotris/domain"
  "fmt"
  "strings"
)

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
    colsScore: cols - 10,
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
  m.DrawScore()
  
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
  utils.WriteInMatrix(m.scoreMap[1],"NEXT:",2)
}

func (m *Map) DrawMap(){
  utils.Debug(fmt.Sprintf("%d", m.rowsMap),5)
  for i:=0; i<m.rowsMap;i++ {
    utils.DrawLine(strings.Join(m.matrix[i],""),domain.Coord{0,i},domain.Coord{1,1})
  }
}

func (m *Map) DrawScore(){
  for i:=0; i<m.rowsScore;i++ {
    utils.DrawLine(strings.Join(m.scoreMap[i],""),domain.Coord{0,i},domain.Coord{m.rowsMap+6,1})
  }
}




