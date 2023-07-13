package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChessBoard(t *testing.T) {
	board1 := NewChessBoard()
	board1.Move(1, 2, 3)

	board2 := NewChessBoard()
	board2.Move(1, 3, 4)

	//两个棋盘的棋子的享元是同一个对象
	assert.Equal(t, board1.ChessPieces[1].Unit, board2.ChessPieces[1].Unit)
}
