/* Problem: Combination Sum II

Description
Given a slice of integers (candidates) that may contain duplicates and a target integer, write a recursive function to find all unique combinations where the chosen numbers sum to the target.

Each number in candidates may be used **at most once** in each combination.

Return all possible unique combinations.

Input:
candidates = [10, 1, 2, 7, 6, 1, 5]
target = 8

Output (order doesn’t matter):
[1,1,6]
[1,2,5]
[1,7]
[2,6]

Constraints
- Use recursion (backtracking)
- Each number can only be used once per combination
- The input may contain duplicates
- Do not generate duplicate combinations
- Return a slice of slices (in Go: [][]int)

Hints
- Sort the candidates array first
- Use a loop to explore choices starting from the current index
- Skip duplicate elements at the same recursion level

- At each step:
  Choose a number and move to the next index

- Track:
  Current sum
  Current combination

- Base cases:
  If sum == target → add combination to result
  If sum > target → stop exploring

- Important:
  Use index progression (i + 1) to prevent reusing elements
  Skip duplicates using:
  if i > index && candidates[i] == candidates[i-1]

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is a problem where I am given a slice of integers (candidates) that may contain duplicates and a target integer, write a recursive function to find all unique combinations where the chosen numbers sum to the target.")
	//Given a slice of integers that contain duplicates and a target integer,
	//write a recursive function to find all unique combinations
	// where the chosen numbers sum to the target.

	var candiates, targetNumber int
	fmt.Println("Enter the number of candiates to be in the array")
	fmt.Scan(&candiates)
	fmt.Println("Please enter the candiates now")
	array := make([]int, candiates)
	for i := 0; i < candiates; i++ {
		fmt.Printf("%d:", i+1)
		fmt.Scan(&array[i])
	}
	fmt.Println("Please enter the target number")
	fmt.Scan(&targetNumber)
	UniqueCombinations := getCombo(array, targetNumber)
	fmt.Println(UniqueCombinations)
}

func getCombo(array []int, targetNumber int) [][]int {
	result := [][]int{}
	sort.Ints(array)
	var helper func([]int, int, int)

	helper = func(current []int, index int, currentSum int) {
		//base case where the sum is equal to the target number and have no duplicates
		if currentSum == targetNumber {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		//If the currentsum is larger or index out of range
		if currentSum > targetNumber || index >= len(array) {
			return
		}
		for i := index; i < len(array); i++ {
			//Skip duplicates
			if i > index && array[i] == array[i-1] {
				continue
			}
			//Pick the number and add it to the combination
			current = append(current, array[i])
			//Then move forward to the next index
			helper(current, i+1, currentSum+array[i])
			//Back track
			current = current[:len(current)-1]
		}

	}

	helper([]int{}, 0, 0)
	return result
}
