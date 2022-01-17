package board

import "fmt"

// userBoard - sets 1's to cells which was given by the user
func (b *board) UserBoard(newBoard [][]int) [][]int {

	// get no. of living cell from user
	var initialCell int

	fmt.Printf("\nEnter no. of Living cell in initial block : ")
	fmt.Scan(&initialCell)

	// get the row index and column index for which the value has to be set to 1's
	for i := 1; i <= initialCell; i++ {

		var row, col int

		fmt.Printf("\nEnter the row index for living cell %d : ", i)
		fmt.Scan(&row)
		fmt.Printf("\nEnter the colum index for living cell %d : ", i)
		fmt.Scan(&col)

		newBoard[row][col] = 1
	}
	return newBoard
}
