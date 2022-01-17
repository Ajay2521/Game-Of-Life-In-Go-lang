package board

// updateBoard - updates the board by newBoard(next generation board) for next iteration
func (b *board) UpdateBoard(boardSize int, newBoard [][]int, nextBoard [][]int) [][]int {
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			newBoard[x][y] = nextBoard[x][y]
		}
	}
	return newBoard
}
