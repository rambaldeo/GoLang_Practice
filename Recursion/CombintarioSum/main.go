/* Problem: Combination Sum

Description
Given a slice of distinct integers (candidates) and a target integer, write a recursive function to find all unique combinations of candidates where the chosen numbers sum to the target.

You may use the same number **multiple times**.

Return all possible combinations.

Input:
candidates = [2, 3, 6, 7]
target = 7

Output (order doesn’t matter):
[2,2,3]
[7]

Constraints
- Use recursion (backtracking)
- You can reuse the same element multiple times
- Do not generate duplicate combinations
- Return a slice of slices (in Go: [][]int)

Hints
- At each step, you have a choice:
  Pick the current number (and stay at the same index)
  Skip the current number (move to next index)

- Track:
  Current sum
  Current combination

- Base cases:
  If sum == target → add combination to result
  If sum > target → stop exploring

- Important:
  Use an index to avoid going backwards and creating duplicates

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a problem to find all unique combinations of integers that add up to a number")
	var candiates, targetNumber int
	fmt.Println("Enter the number of candiates to be in the array")
	fmt.Scan(&candiates)
	fmt.Println("Please enter the canidates now")
	array := make([]int, candiates)
	for i := 0; i < candiates; i++ {
		fmt.Printf("%d:", i+1)
		fmt.Scanln(&array[i])
	}
	fmt.Println("Please enter the target number")
	fmt.Scan(&targetNumber)
	Combos := getcombo(array, targetNumber)
	fmt.Print(Combos)
}

func getcombo(array []int, targetNumber int) [][]int {
	result := [][]int{}
	var helper func([]int, int, int)

	helper = func(current []int, index int, currentSum int) {
		//base case
		//if the sum is equal to the target number
		if currentSum == targetNumber {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		if currentSum > targetNumber || index >= len(array) {
			return
		}
		// ✅ Choice 1: PICK the number
		current = append(current, array[index])
		helper(current, index, currentSum+array[index])
		// 🔙 Backtrack (undo the choice)
		current = current[:len(current)-1]

		// ❌ Choice 2: SKIP the number
		helper(current, index+1, currentSum)
	}

	helper(array, 0, 0)
	return result
}

func sumOfArray(tempArray []int, targetNum int) int {
	result := 0
	for i := 0; i < len(tempArray); i++ {
		result += tempArray[i]
	}
	return result
}
