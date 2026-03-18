/* Problem: Word Search (Backtracking on a Grid)

Description
Given a 2D grid of characters and a target word, write a recursive function to determine if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where:
- Adjacent cells are horizontally or vertically neighboring
- The same cell cannot be used more than once in a single path

Input:
board = [
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
word = "ABCCED"

Output:
true

Explanation:
The word can be formed by the path:
A → B → C → C → E → D

Constraints
- Use recursion (DFS + backtracking)
- You may not reuse the same cell in a path
- Modify the board in-place OR use a visited structure
- Return a boolean

Hints
- Loop through every cell as a starting point
- Create a recursive helper(row, col, index)
- Base case:
  If index == len(word), return true
- Check boundaries and character match
- Mark cell as visited before recursion
- Backtrack (unmark it) after exploring

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a word search problem using recursion (Backtracking)")
	var boardSize, inputsPerArray int

	fmt.Println("What is the board size?")
	fmt.Scan(&boardSize)
	fmt.Println("How many letters per row?")
	fmt.Scan(&inputsPerArray)
	board := make([][]string, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]string, inputsPerArray)
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < inputsPerArray; j++ {
			fmt.Scan(&board[i][j])
		}
	}
	var wordToFind string
	fmt.Println("Enter the word that you want to find")
	fmt.Scan(&wordToFind)
	shortestPath := findAnswer(board, wordToFind)
	fmt.Println(shortestPath)
}

func findAnswer(board [][]string, wordToFind string) bool {
	rows := len(board)
	cols := len(board[0])

	var helper func(int, int, int) bool
	helper = func(r, c, index int) bool {
		if index == len(wordToFind) {
			return true
		}
		// boundary + mismart check
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != string(wordToFind[index]) {
			return false
		}
		//mark visited
		temp := board[r][c]
		board[r][c] = "*"

		//explore the board
		found := helper(r+1, c, index+1) ||
			helper(r-1, c, index+1) ||
			helper(r, c+1, index+1) ||
			helper(r, c-1, index+1)
		//backtrack
		board[r][c] = temp
		return found
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if helper(i, j, 0) {
				return true
			}
		}
	}

	return false
}
