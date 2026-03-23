package domain

// import (
// 	"gotris/utils"
// )

type Block struct{
	Elements [4][3]Object
}

//There should be multiple constructors, one for each kind of block
func NewBlock(Pos Coord, Offset Coord) *Block {
    return &Block{
        Elements: [4][3]Object{
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -3}, Offset: Coord{Offset.X,Offset.Y}}, {},{}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -2}, Offset: Coord{Offset.X,Offset.Y}}, {}, {}, },
            { {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + -1}, Offset: Coord{Offset.X,Offset.Y}}, {}, {}, },
			{ {Chr: "*", Pos: Coord{Pos.X + 0,Pos.Y + 0}, Offset: Coord{Offset.X,Offset.Y}}, {}, {}, },
        },
    }
}

func (b *Block) DrawBlock(){
	for i := 3; i>=0;i--{
		for j:= 0;j<3;j++{
      obj := b.Elements[i][j]
      if obj.Chr != "" {
        // utils.DrawLine(obj.Chr, obj.Pos, obj.Offset)
      }
	}
  }
}
