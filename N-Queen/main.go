package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This program will solve the famous N-Queen problem. It will take N number of Queens and place them on the board where they do not interact with the other queens. The minimum board size is 4x4")
	fmt.Println("This solution will be using Backtracking")
	fmt.Println("Please enter the board size")

	var boardSize int
	fmt.Scan(&boardSize)

	if boardSize < 4 {
		fmt.Println("Please enter a larger board size")
		fmt.Scan(&boardSize)
	}

	fmt.Printf("The entered board size is %d\n", boardSize)

	// Get the solutions
	list := nQueen(boardSize)

	// Print solutions
	fmt.Printf("The total solution is %d", len(list))
	fmt.Println("If you would like to see the solutions, please select Yes")
	var userInput string
	fmt.Scan(&userInput)
	if strings.ToLower(userInput) == "yes" {
		for i, solution := range list {
			fmt.Printf("Solution %d: %v\n", i+1, solution)
		}
	}
}

func nQueen(n int) [][]int {

	// making the board
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	result := [][]int{}

	// determine where to place queens
	placeQueens(0, mat, &result)

	return result
}

// recursive function to place queens
func placeQueens(row int, mat [][]int, result *[][]int) {

	n := len(mat)

	// If all queens have been placed
	if row == n {

		ans := []int{}

		// Store the current solution
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mat[i][j] == 1 {
					ans = append(ans, j+1)
				}
			}
		}

		*result = append(*result, ans)
		return
	}

	// Backtracking logic will go here later
	//Consider the row and try placing
	//queen in all columns one by one
	for i := 0; i < n; i++ {
		//Check if queen can be placed
		if isSafe(row, mat, i) == 1 {
			mat[row][i] = 1
			placeQueens(row+1, mat, result)

			//backtrack
			mat[row][i] = 0
		}

	}
}

func isSafe(row int, mat [][]int, col int) int {
	n := len(mat)

	//Check this col on upper side
	for i := 0; i < row; i++ {
		if mat[i][col] == 1 {
			return 0
		}
	}
	//Check upper diagnoal on left side
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if mat[i][j] == 1 {
			return 0
		}
	}

	//Check upper diagnoal on right side
	for i, j := row-1, col+1; j < n && i >= 0; i, j = i-1, j+1 {
		if mat[i][j] == 1 {
			return 0
		}
	}

	return 1
}
