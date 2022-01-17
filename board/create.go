package board

// CreateBoard - function Create a outer board.
func (b *board) CreateBoard(boardSize int) ([][]int, [][]int){

	// defining new board using 2D-slice data structure
	newBoard := make([][]int, boardSize)
	for i := range newBoard {
		newBoard[i] = make([]int, boardSize)
	}

	// defining new board using 2D-slice data structure
	nextBoard := make([][]int, boardSize)
	for i := range nextBoard {
		nextBoard[i] = make([]int, boardSize)
	}

	return newBoard, nextBoard
}
