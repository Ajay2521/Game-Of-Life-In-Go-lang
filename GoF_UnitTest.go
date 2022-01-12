package main

import (
	"reflect"
	"testing"
)


type  TestCase struct{
	testBoard [][] int
	testReceivedBoard [][] int
	testExpectedBoard [][]int
}

func TestGeneration(t *testing.T) {

	var x int = 0
	var y int = 0
	var testBoardSize = 8

	testBoard := [][]int{{0, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 0, 1, 1, 1}}

	testReceivedBoard := make([][]int, testBoardSize)
	for i := range testReceivedBoard {
		testReceivedBoard[i] = make([]int, testBoardSize)
	}

	testExpectedBoard := [][]int{{1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1}}

	testCase := TestCase{
		testBoard: testBoard,
		testReceivedBoard: testReceivedBoard,
		testExpectedBoard: testExpectedBoard,
	}

	testCase.testReceivedBoard  = liveOrDeath(testCase.testBoard, testCase.testReceivedBoard, x, y, testBoardSize)

	result := reflect.DeepEqual(testCase.testReceivedBoard, testCase.testExpectedBoard)

	if result == false {
		t.Errorf("got %v, wanted %v", testCase.testReceivedBoard, testCase.testExpectedBoard)
	}
}

func TestGame(t *testing.T) {

	var x int = 0
	var y int = 0
	var testBoardSize = 8

	testBoard := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
	}

	testReceivedBoard := make([][]int, testBoardSize)
	for i := range testReceivedBoard {
		testReceivedBoard[i] = make([]int, testBoardSize)
	}

	testExpectedBoard := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}}

	testCase := TestCase{
		testBoard: testBoard,
		testReceivedBoard: testReceivedBoard,
		testExpectedBoard: testExpectedBoard,
	}

	testCase.testReceivedBoard  = liveOrDeath(testCase.testBoard, testCase.testReceivedBoard, x, y, testBoardSize)

	result := reflect.DeepEqual(testCase.testReceivedBoard, testCase.testExpectedBoard)

	if result == false {
		t.Errorf("got %v, wanted %v", testCase.testReceivedBoard, testCase.testExpectedBoard)
	}
}
