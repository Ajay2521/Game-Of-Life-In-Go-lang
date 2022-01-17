package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	b Board
}

func (suite *CreateTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *CreateTestSuite) TestCreate1() {
	var boardSize1 int = 5

	var expectedOut1 = [][]int{{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}

	var boardSize2 int = 2
	var expectedOut2 = [][]int{{0, 0}, {0, 0}}

	var testDataSet = []struct {
		boardSize         int
		expectedNewBoard  [][]int
		expectedNextBoard [][]int
	}{
		{boardSize1, expectedOut1, expectedOut1},
		{boardSize2, expectedOut2, expectedOut2},
	}

	for _, testData := range testDataSet {
		receivedNewBoard, receivedNextBoard := suite.b.CreateBoard(testData.boardSize)
		suite.Equal(testData.expectedNewBoard, receivedNewBoard)
		suite.Equal(testData.expectedNextBoard, receivedNextBoard)
	}
}

func (suite *CreateTestSuite) TestCreate2() {
	var boardSize1 int = 8
	var expectedOut1 = [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}

	var boardSize2 int = 0
	var expectedOut2 = [][]int{}

	var testDataSet = []struct {
		boardSize         int
		expectedNewBoard  [][]int
		expectedNextBoard [][]int
	}{
		{boardSize1, expectedOut1, expectedOut1},
		{boardSize2, expectedOut2, expectedOut2},
	}

	for _, testData := range testDataSet {
		receivedNewBoard, receivedNextBoard := suite.b.CreateBoard(testData.boardSize)
		suite.Equal(testData.expectedNewBoard, receivedNewBoard)
		suite.Equal(testData.expectedNextBoard, receivedNextBoard)
	}
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
