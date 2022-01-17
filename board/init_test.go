package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InitTestSuite struct {
	suite.Suite
	b Board
}

func (suite *InitTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *InitTestSuite) TestInit1() {
	var boardSize1 int = 5
	var boardIn1 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}
	var expectedOut1 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}

	var boardSize2 int = 2
	var boardIn2 = [][]int{{0, 0}, {0, 0}}
	var expectedOut2 = [][]int{{0, 0}, {0, 0}}

	var testDataSet = []struct {
		boardSize     int
		board         [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn1, expectedOut1},
		{boardSize2, boardIn2, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.InitBoard(testData.boardSize, testData.board))
	}
}

func (suite *InitTestSuite) TestInit2() {
	var boardSize1 int = 8
	var boardIn1 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	var expectedOut1 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}

	var boardSize2 int = 0
	var boardIn2 = [][]int{{}}
	var expectedOut2 = [][]int{{}}

	var testDataSet = []struct {
		boardSize     int
		board         [][]int
		expectedBoard [][]int
	}{
		{boardSize1, boardIn1, expectedOut1},
		{boardSize2, boardIn2, expectedOut2},
	}

	for _, testData := range testDataSet {
		suite.Equal(testData.expectedBoard, suite.b.InitBoard(testData.boardSize, testData.board))
	}

}

func TestInitTestSuite(t *testing.T) {
	suite.Run(t, new(InitTestSuite))
}
