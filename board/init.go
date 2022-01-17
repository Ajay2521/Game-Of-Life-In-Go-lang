package board

// InitBoard - sets the values of the board with zero as value
func (b *board) InitBoard(boardSize int, board [][]int) [][]int {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = 0
		}
	}
	return board
}
