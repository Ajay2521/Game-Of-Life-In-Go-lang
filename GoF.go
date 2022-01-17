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
import (
	"fmt"
	"GoF/board"
)

// Global variables
var boardSize, iteration, choice int
var newBoard, nextBoard [][]int

// main function of the program
func main() {

	userBoard := board.NewBoard()

	// function call for geting user input
	boardSize, iteration, choice = userBoard.UserInput(boardSize, iteration, choice)

	// function call for create board
	newBoard, nextBoard = userBoard.CreateBoard(boardSize)

	// function call for Initializing a initial board
	newBoard = userBoard.InitBoard(boardSize, newBoard)

	// switch block to call the block generation function based on user choice
	switch choice {
	case 1:
		newBoard = userBoard.UserBoard(newBoard)
	case 2:
		newBoard = userBoard.RandBoard(boardSize,newBoard)
	default:
		fmt.Println("\nInvalid Input")
	}

	fmt.Println("\n########## Initial board ##########")

	// function call to display the board value
	userBoard.DisplayBoard(boardSize, newBoard)

	fmt.Print("\nThe Game of Life begins...")

	// loop runs for user given iteration time
	for i := 1; i <= iteration; i++ {

		// function call to find the living cell and death cell in the board
		nextBoard = userBoard.LiveOrDeath(boardSize, newBoard, nextBoard)

		fmt.Printf("########## Iteration : %d ##########\n\n", i)

		// function call to display the board value
		userBoard.DisplayBoard(boardSize, nextBoard)

		// update the current block with the new block(next generation block)
		userBoard.UpdateBoard(boardSize, newBoard, nextBoard)
	}

	fmt.Println("\nThe Game of Life has came to an end.")
}



