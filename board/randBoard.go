package board

import "math/rand"

// RandBoard - sets 1's in random order/index on the board
func (b *board) RandBoard(boardSize int, newBoard [][]int) [][]int {
	for i := 0; i < boardSize*boardSize; i++ {
		x := rand.Int() % boardSize
		y := rand.Int() % boardSize
		newBoard[x][y] = 1
	}
	return newBoard
}
