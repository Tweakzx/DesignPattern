package flyweight

type ChessPieceUnit struct {
	ID        int
	Name      string
	Color     string
	default_X int
	default_Y int
}

var units = map[int]*ChessPieceUnit{
	1: {ID: 1, Name: "車", Color: "紅", default_X: 0, default_Y: 0},
	2: {ID: 2, Name: "炮", Color: "紅", default_X: 0, default_Y: 0},
	3: {ID: 3, Name: "馬", Color: "紅", default_X: 0, default_Y: 0},
	4: {ID: 4, Name: "相", Color: "紅", default_X: 0, default_Y: 0},
	5: {ID: 5, Name: "仕", Color: "紅", default_X: 0, default_Y: 0},
	6: {ID: 6, Name: "帥", Color: "紅", default_X: 0, default_Y: 0},
	//省略其他棋子
	7:  {ID: 1, Name: "車", Color: "黑", default_X: 0, default_Y: 0},
	8:  {ID: 2, Name: "炮", Color: "黑", default_X: 0, default_Y: 0},
	9:  {ID: 3, Name: "馬", Color: "黑", default_X: 0, default_Y: 0},
	10: {ID: 4, Name: "象", Color: "黑", default_X: 0, default_Y: 0},
	11: {ID: 5, Name: "士", Color: "黑", default_X: 0, default_Y: 0},
	12: {ID: 6, Name: "將", Color: "黑", default_X: 0, default_Y: 0},
	//省略其他棋子
}

type ChessPiece struct {
	Unit  *ChessPieceUnit
	X     int
	Y     int
	Eaten bool
}

func ChessPieceUnitFactory(id int) *ChessPieceUnit {
	return units[id]
}

type ChessBoard struct {
	ChessPieces map[int]*ChessPiece
}

func (b *ChessBoard) Init() {
	for id := range units {
		p := ChessPieceUnitFactory(id)
		b.ChessPieces[id] = &ChessPiece{
			Unit:  p,
			X:     p.default_X,
			Y:     p.default_Y,
			Eaten: false,
		}
	}
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{
		ChessPieces: map[int]*ChessPiece{},
	}
	board.Init()
	return board
}

func checkrules(id, x, y int) {
	//省略
}
func eat(id, x, y int) {
	//省略
}
func (c *ChessBoard) Move(id, x, y int) {
	checkrules(id, x, y)
	eat(id, x, y)
	c.ChessPieces[id].X = x
	c.ChessPieces[id].Y = y
}
