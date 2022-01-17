package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RandBoardTestSuite struct {
	suite.Suite
	b Board
}

func (suite *RandBoardTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *RandBoardTestSuite) TestRandBoard1() {
	var boardSize1 int = 5
	var boardIn1 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}
	var expectedOut1 = [][]int{{0, 1, 1, 1, 0},
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 1, 1, 1, 0}}

	var boardSize2 int = 3
	var boardIn2 = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	var expectedOut2 = [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 0}}

	var testDataSet = []struct {
		boardSize     int
		newBoard      [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn1, expectedOut1},
		{boardSize2, boardIn2, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.RandBoard(testData.boardSize, testData.newBoard))
	}
}

func (suite *RandBoardTestSuite) TestRandBoard2() {
	var boardSize1 int = 8
	var boardIn1 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	var expectedOut1 = [][]int{{0, 1, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 1, 1, 0, 1},
		{0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1, 0, 1, 1},
		{1, 1, 0, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 1}}

	var boardSize2 int = 1
	var boardIn2 = [][]int{{0}}
	var expectedOut2 = [][]int{{1}}

	var testDataSet = []struct {
		boardSize     int
		newBoard      [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn1, expectedOut1},
		{boardSize2, boardIn2, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.RandBoard(testData.boardSize, testData.newBoard))
	}

}

func TestRandBoardTestSuite(t *testing.T) {
	suite.Run(t, new(RandBoardTestSuite))
}
