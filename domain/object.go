package domain

type Coord struct{
  X int
  Y int
}

type Object struct{
	Chr string
	Pos Coord
	Offset Coord
}
