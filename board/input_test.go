package board

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InputTestSuite struct {
	suite.Suite
	b Board
}

func (suite *InputTestSuite) SetupSuite() {
	suite.b = NewBoard()
}

func (suite *InputTestSuite) TestInput1() {
	var testDataSet = []struct {
		boardSize         int
		iteration         int
		choice            int
		expectedBoardSize int
		expectedIteration int
		expectedChoice    int
	}{
		{3, 5, 1, 3, 5, 1},
		{1, 4, 2, 1, 4, 2},
	}

	for _, testData := range testDataSet {
		receivedBoardSize, receivedIteration, receivedChoice := suite.b.UserInput(testData.boardSize, testData.iteration, testData.choice)
		suite.Equal(testData.expectedBoardSize, receivedBoardSize)
		suite.Equal(testData.expectedIteration, receivedIteration)
		suite.Equal(testData.expectedChoice, receivedChoice)
	}
}

func (suite *InputTestSuite) TestInput2() {

	var testDataSet = []struct {
		boardSize         int
		iteration         int
		choice            int
		expectedBoardSize int
		expectedIteration int
		expectedChoice    int
	}{
		{3, 5, 1, 3, 5, 1},
		{1, 4, 2, 1, 4, 2},
	}

	for _, testData := range testDataSet {
		receivedBoardSize, receivedIteration, receivedChoice := suite.b.UserInput(testData.boardSize, testData.iteration, testData.choice)
		suite.Equal(testData.expectedBoardSize, receivedBoardSize)
		suite.Equal(testData.expectedIteration, receivedIteration)
		suite.Equal(testData.expectedChoice, receivedChoice)
	}
}

func TestInputTestSuite(t *testing.T) {
	suite.Run(t, new(InputTestSuite))
}
