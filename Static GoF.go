package main

import (
	"fmt"
	"time"
)

const size = 20

var grid = [20][20]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

var neighbourCount [size][size]int

func main() {

	for steps := 0; steps < 10; steps++ {

		fmt.Println("\n")

		time.Sleep(2000 * time.Millisecond)

		fmt.Println("\nGneration -", steps)

		for i := 0; i < size; i++ {

			fmt.Print("\n")

			for j := 0; j < size; j++ {

				if grid[i][j] == 1 {

					fmt.Printf(" *")

				} else {

					fmt.Printf(" -")

				}

				neighbourCount[i][j] = countNegihbour(grid, i, j, size)
			}

		}

		for i := 0; i < size; i++ {

			for j := 0; j < size; j++ {

				if grid[i][j] >= 1 {

					if neighbourCount[i][j] <= 1 || neighbourCount[i][j] >= 4 {
						grid[i][j] = 0
					}
				} else {
					if neighbourCount[i][j] == 3 {
						grid[i][j] = 1
					}
				}
			}
		}
	}

}

func countNegihbour(grid [size][size]int, i, j, size int) int {

	count := 0

	if i-1 >= 0 && j-1 >= 0 {
		if grid[i-1][j-1] >= 1 {
			count++
		}
	}

	if i-1 >= 0 {
		if grid[i-1][j] >= 1 {
			count++
		}
	}

	if i-1 >= 0 && j+1 < size {
		if grid[i-1][j+1] >= 1 {
			count++
		}
	}

	if j-1 >= 0 {
		if grid[i][j-1] >= 1 {
			count++
		}
	}

	if j+1 < size {
		if grid[i][j+1] >= 1 {
			count++
		}
	}

	if i+1 < size && j-1 >= 0 {
		if grid[i+1][j-1] >= 1 {
			count++
		}
	}

	if i+1 < size {
		if grid[i+1][j] >= 1 {
			count++
		}
	}

	if i+1 < size && j+1 < size {
		if grid[i+1][j+1] >= 1 {
			count++
		}
	}

	return count
}
