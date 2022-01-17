package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
	b Board
}

func (suite *GameTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *GameTestSuite) TestGame1() {
	var boardSize1 int = 5
	var boardIn11 = [][]int{{0, 1, 1, 1, 0},
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 1, 1, 1, 0}}
	var boardIn12 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}

	var expectedOut1 = [][]int{{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 0, 1}}

	var boardSize2 int = 2
	var boardIn21 = [][]int{{1, 1}, {1, 1}}
	var boardIn22 = [][]int{{0, 0}, {0, 0}}
	var expectedOut2 = [][]int{{1, 1}, {1, 1}}

	var testDataSet = []struct {
		boardSize     int
		newBoard      [][]int
		nextBoard     [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn11, boardIn12, expectedOut1},
		{boardSize2, boardIn21, boardIn22, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.LiveOrDeath(testData.boardSize, testData.newBoard, testData.nextBoard))
	}
}

func (suite *GameTestSuite) TestGame2() {
	var boardSize1 int = 8
	var boardIn11 = [][]int{{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 0, 1, 1, 1}}
	var boardIn12 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	var expectedOut1 = [][]int{{1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1}}

	var boardSize2 int = 1
	var boardIn21 = [][]int{{1}}
	var boardIn22 = [][]int{{0}}
	var expectedOut2 = [][]int{{0}}

	var testDataSet = []struct {
		boardSize     int
		newBoard      [][]int
		nextBoard     [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn11, boardIn12, expectedOut1},
		{boardSize2, boardIn21, boardIn22, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.LiveOrDeath(testData.boardSize, testData.newBoard, testData.nextBoard))
	}

}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}
