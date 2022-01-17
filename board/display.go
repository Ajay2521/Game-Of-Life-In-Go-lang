package board

import (
	"fmt"
	"time"
)

// displayBoard - print the values of board as symbol
// * -> denotes living cell/alive cell
// - -> denotes death cell
func (b *board) DisplayBoard(boardSize int, board [][]int) {
	time.Sleep(1000 * time.Millisecond)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == 1 {
				fmt.Printf(" *")
			} else {
				fmt.Printf(" -")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(board)
}
