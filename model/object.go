package model

type Coord struct{
  X int
  Y int
}

type Object struct{
	Chr string
	Pos Coord
}



type Block struct{
	Elements [4][3]Object
}
//There should be multiple constructors, one for each kind of block
func NewBlock(Pos Coord) *Block {
    return &Block{
        Elements: [4][3]Object{
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -3}}, {},{}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -2}}, {}, {}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -1}}, {}, {}, },
			{ {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + 0}}, {}, {}, },
        },
    }
}

func NewDashBoardBlock(Pos Coord) *Block {
    return &Block{
        Elements: [4][3]Object{
            { {Chr: "*", Pos: Coord{Pos.X + 0, Pos.Y + 0}}, {},{}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + 1}}, {}, {}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + 2}}, {}, {}, },
			{ {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + 3}}, {}, {}, },
        },
    }
}
