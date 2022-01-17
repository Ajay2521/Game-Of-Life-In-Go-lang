package board

import "fmt"

// UserInput - function getting user input
func (b *board) UserInput(boardSize, iteration, choice int) (int, int, int) {

	// getting board size from user
	fmt.Print("\nEnter the Board size needed : ")
	fmt.Scan(&boardSize)

	// getting no. of iteration needed from user
	fmt.Print("\nEnter the Maximum Iteration needed : ")
	fmt.Scan(&iteration)

	// getting user choice for user generated initial block or auto generation of initial block
	fmt.Print("\n1. Press 1 for User Generated Initial Board.\n2. Press 2 for auto Generated Initial Board")
	fmt.Print("\nEnter Your choice : ")
	fmt.Scan(&choice)

	return boardSize, iteration, choice
}
