package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UpdateTestSuite struct {
	suite.Suite
	b Board
}

func (suite *UpdateTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *UpdateTestSuite) TestUpdate1() {
	var boardSize1 int = 5
	var boardIn11 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}
	var boardIn12 = [][]int{{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0}}

	var expectedOut1 = [][]int{{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0}}

	var boardSize2 int = 2
	var boardIn21 = [][]int{{0, 0}, {0, 0}}
	var boardIn22 = [][]int{{0, 1}, {1, 0}}
	var expectedOut2 = [][]int{{0, 1}, {1, 0}}

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
		suite.Equal(testData.expectedBoard, suite.b.UpdateBoard(testData.boardSize, testData.newBoard, testData.nextBoard))
	}
}

func (suite *UpdateTestSuite) TestUpdate2() {
	var boardSize1 int = 8
	var boardIn11 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
	var boardIn12 = [][]int{{0, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0}}
	var expectedOut1 = [][]int{{0, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 1, 0}}

	var boardSize2 int = 1
	var boardIn21 = [][]int{{0}}
	var boardIn22 = [][]int{{1}}
	var expectedOut2 = [][]int{{1}}

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
		suite.Equal(testData.expectedBoard, suite.b.UpdateBoard(testData.boardSize, testData.newBoard, testData.nextBoard))
	}

}

func TestUpdateTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
