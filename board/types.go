package board

type board struct{}

type Board interface {
	UserInput(boardSize, iteration, choice int) (int, int, int)
	CreateBoard(boardSize int) ([][]int, [][]int)
	InitBoard(boardSize int, board [][]int) [][]int
	DisplayBoard(boardSize int, board [][]int)
	UserBoard(newBoard [][]int) [][]int
	RandBoard(boardSize int, newBoard [][]int) [][]int
	LiveOrDeath(boardSize int, newBoard [][]int, nextBoard [][]int) [][]int
	UpdateBoard(boardSize int, newBoard [][]int, nextBoard [][]int) [][]int
}

func NewBoard() Board {
	return &board{}
}
