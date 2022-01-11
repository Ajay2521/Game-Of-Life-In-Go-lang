package main

import (
	"fmt"
	"reflect"
	"testing"
)

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

	testNewBoard := make([][]int, testBoardSize)
	for i := range testNewBoard {
		testNewBoard[i] = make([]int, testBoardSize)
	}

	testWantBoard := [][]int{{1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1}}

	testNewBoard  = liveOrDeath(testBoard, testNewBoard, x, y, testBoardSize)

	result := reflect.DeepEqual(testNewBoard,testWantBoard)
	fmt.Println(result)
	if result == false {
		t.Errorf("got %v, wanted %v", testNewBoard, testWantBoard)
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

	testNewBoard := make([][]int, testBoardSize)
	for i := range testNewBoard {
		testNewBoard[i] = make([]int, testBoardSize)
	}

	testWantBoard := [][]int{{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}}

	testNewBoard  = liveOrDeath(testBoard, testNewBoard, x, y, testBoardSize)

	result := reflect.DeepEqual(testNewBoard,testWantBoard)
	fmt.Println(result)
	if result == false {
		t.Errorf("got %v, wanted %v", testNewBoard, testWantBoard)
	}
}
