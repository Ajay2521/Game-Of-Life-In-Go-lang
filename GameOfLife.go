/*
 *   Author - Ajay M
 *   Project - Game of Life
 *   Language - Go Lang
 *
 *   This is an implementation of Conway's Game of Life,
 *   https://en.wikipedia.org/wiki/Conway's_Game_of_Life.
 */

// creating and defining the main package of the program
package main

// importing necessary function needed in our program
// fmt - input and output operation
// math/rand - generating random numbers
// time - making programing to execute based on time interval / sleep
import (
	"fmt"
	"math/rand"
	"time"
)

// Global variables
const birthCondition = 3 // if a cell is surrounded by 3 living cell, then new cell born
const surCond1 = 3       // includes the cell itself, hence +1 to the rule condition
const surCond2 = 4       // includes the cell itself, hence +1 to the rule condition
var boardSize int        // size of the board which will be given by user

var iterate int // number of generation are needed to be generated

// main function of the program
func main() {

	// x and y considered for iteration of cells through x and y axis
	var x int = 0
	var y int = 0

	var blockChoice int

	// getting board size from user
	fmt.Print("\nEnter the Board size needed : ")
	fmt.Scan(&boardSize)

	// defining board using 2D-slice data structure
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	// defining new board using 2D-slice data structure
	newBoard := make([][]int, boardSize)
	for i := range newBoard {
		newBoard[i] = make([]int, boardSize)
	}

	// getting no. of iteration needed from user
	fmt.Print("\nEnter the Maximum Iteration needed : ")
	fmt.Scan(&iterate)

	// function call for Initializing a initial board
	initBoard(board)

	// getting user choice for user generated initial block or auto generation of initial block
	fmt.Print("\n1. Press 1 for User Generated Initial Board.\n2. Press 2 for auto Generated Initial Board")
	fmt.Print("\nEnter Your choice : ")
	fmt.Scan(&blockChoice)

	// switch block to call the block generation function based on user choice
	switch blockChoice {
	case 1:
		userBoard(board)
	case 2:
		randBoard(board)
	default:
		fmt.Println("\nInvalid Input")
	}

	// function call for Initializing a new board which is considered as next generation block
	initBoard(newBoard)

	fmt.Println("\n########## Initial board ##########")

	// function call to display the board value
	displayBoard(board)

	fmt.Print("\nThe Game of Life begins...")

	// loop runs for user given iteration time
	for i := 1; i <= iterate; i++ {

		// function call to find the living cell and death cell in the board
		liveOrDeath(board, newBoard, x, y)

		fmt.Printf("########## Iteration : %d ##########\n\n", i)

		// function call to display the board value
		displayBoard(newBoard)

		// update the current block with the new block(next generation block)
		updateBoard(board, newBoard, x, y)
	}

	fmt.Println("\nThe Game of Life has came to an end.\n")
}

// initBoard - sets the values of the board with zero as value
func initBoard(board [][]int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = 0
		}
	}
}

// userBoard - sets 1's to cells which was given by the user
func userBoard(board [][]int) {

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

		board[row][col] = 1
	}
}

// randBoard - sets 1's in random order/index on the board
func randBoard(board [][]int) {
	for i := 0; i < boardSize*boardSize; i++ {
		x := rand.Int() % boardSize
		y := rand.Int() % boardSize
		board[x][y] = 1
	}
}

// displayBoard - print the values of board as symbol
// * -> denotes living cell/alive cell
// - -> denotes death cell
func displayBoard(board [][]int) {
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
}

/* perform live or death rules */
/*
1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/
func liveOrDeath(board [][]int, newBoard [][]int, x int, y int) {

	var neighbourCount int

	// function call for Initializing a new board(next generation board)
	initBoard(newBoard)

	for x = 0; x < boardSize; x++ {

		for y = 0; y < boardSize; y++ {

			neighbourCount = countNeighbours(board, x, y)

			if board[x][y] == 0 {
				// cell is uninhabited, check if conditions for birth are met
				if neighbourCount == birthCondition {
					newBoard[x][y] = 1 // birth
				} else {
					newBoard[x][y] = 0
				}
			} else if board[x][y] == 1 {
				// cell exists, check if it survies or dies
				if neighbourCount == surCond1 || neighbourCount == surCond2 {
					// survive
					newBoard[x][y] = board[x][y]
				} else {
					// death
					newBoard[x][y] = 0
				}
			}
		}
	}
	fmt.Println()
}

// countNeighbours - counts the neighbours of a given cell
func countNeighbours(board [][]int, x, y int) int {

	var row, col, value, neighbourCount int

	value = 0

	neighbourCount = 0 /* reset count for new cell */

	for row = -1; row <= 1; row++ {
		for col = -1; col <= 1; col++ {
			if (x+row >= 0) &&
				(y+col >= 0) &&
				(x+row < boardSize) &&
				(y+col < boardSize) {

				// assign the cell value (0 or 1) to variable value
				value = board[x+row][y+col]

				/* value is determined base on the cell.
				if there is a neighbour cell,value =1, else value =0 */
				if value == 1 {
					neighbourCount++
				}
			}
		}
	}
	return neighbourCount
}

// updateBoard - updates the board by newBoard(next generation board) for next iteration
func updateBoard(board [][]int, newBoard [][]int, x, y int) {
	for x = 0; x < boardSize; x++ {
		for y = 0; y < boardSize; y++ {
			board[x][y] = newBoard[x][y]
		}
	}
}
