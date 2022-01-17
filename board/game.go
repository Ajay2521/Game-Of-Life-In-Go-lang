package board

import "fmt"

const birthCondition = 3 // if a cell is surrounded by 3 living cell, then new cell born
const surCond1 = 3       // includes the cell itself, hence +1 to the rule condition
const surCond2 = 4       // includes the cell itself, hence +1 to the rule condition

/* perform live or death rules */
/*
1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/
func (b *board) LiveOrDeath( boardSize int, newBoard [][]int, nextBoard [][]int ) [][]int {

	var neighbourCount int

	// function call for Initializing a new board(next generation board)
	b.InitBoard(boardSize, nextBoard)

	for x := 0; x < boardSize; x++ {

		for y := 0; y < boardSize; y++ {

			var row, col, value int

			value = 0

			neighbourCount = 0 /* reset count for new cell */

			for row = -1; row <= 1; row++ {
				for col = -1; col <= 1; col++ {
					if (x+row >= 0) &&
						(y+col >= 0) &&
						(x+row < boardSize) &&
						(y+col < boardSize) {

						// assign the cell value (0 or 1) to variable value
						value = newBoard[x+row][y+col]

						/* value is determined base on the cell.
						if there is a neighbour cell,value =1, else value =0 */
						if value == 1 {
							neighbourCount++
						}
					}
				}
			}

			if newBoard[x][y] == 0 {
				// cell is uninhabited, check if conditions for birth are met
				if neighbourCount == birthCondition {
					nextBoard[x][y] = 1 // birth
				} else {
					nextBoard[x][y] = 0
				}
			} else if newBoard[x][y] == 1 {
				// cell exists, check if it survies or dies
				if neighbourCount == surCond1 || neighbourCount == surCond2 {
					// survive
					nextBoard[x][y] = newBoard[x][y]
				} else {
					// death
					nextBoard[x][y] = 0
				}
			}
		}
	}
	fmt.Println()
	return nextBoard
}
