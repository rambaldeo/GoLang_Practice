/*Problem: Generate All Subsets of a Set (Power Set)
Description
Given a set of integers, write a recursive function to generate all possible subsets (the power set).
Input: [1, 2, 3]
Output (order doesn’t matter):
[]
[1]
[2]
[3]
[1,2]
[1,3]
[2,3]
[1,2,3]
Constraints
You cannot use loops to combine subsets in the main recursion—only recursion itself.
You can use helper functions or pass current subset as an argument.
Return a slice of slices (in Go: [][]int) containing all subsets.

Hints
Think of the recursion as two choices for each element:
Include the current element
Exclude the current element
Base case: when you have considered all elements, add the subset to the result. */

package main

import (
	"fmt"
)

func main() {
	var arrayLength int

	//Take user input
	fmt.Println("How many number do you want to enter:")
	fmt.Scanln(&arrayLength)
	fmt.Println("Please input the intergers")

	//Creating the array to hold the variables
	array := make([]int, arrayLength)
	for i := 0; i < arrayLength; i++ {
		fmt.Scanln(&array[i])
	}
	fmt.Println("Thank you, the integers that you entered is: ")
	for index, a := range array {
		fmt.Printf("Index:%d is value:%d", a, index)
	}
	fmt.Println("\nThe subsets will be displayed before: ")
	//Get the subsets
	generateSubsets(array, 0)

	//Getting the power sets
	//results := []int{}
	//generatePowerSets(array, 0, 0, &results)
	PowerSets := PowerSetsAnswer(array)
	for _, s := range PowerSets {
		fmt.Println(s)
	}

}

func generateSubsets(array []int, n int) {
	length := len(array)
	var s []int = array[0:n]
	if length == n {
		fmt.Printf("%d \n", s)
		fmt.Println("This is the end, goodbye")
		return
	}
	if n == 0 {
		generateSubsets(array, n+1)
	} else {
		fmt.Printf("%d \n", s)
		generateSubsets(array, n+1)
	}

}

func PowerSetsAnswer(array []int) [][]int {
	var results [][]int
	var current []int

	var helper func(index int)
	helper = func(index int) {
		if index == len(array) {
			subset := make([]int, len(current))
			copy(subset, current)
			results = append(results, subset)
			return
		}

		//Including the element
		current = append(current, array[index])
		helper(index + 1)

		//Backtrack (removing the element)
		current = current[:len(current)-1]
	}
	helper(0)
	return results
}

func generatePowerSets(array []int, n int, m int, result *[]int) {
	length := len(array)
	var sets []int = array[m:n]
	//start at 0 index and get the subArrays with that index

	//If you reached the end of n, increase the m value
	if n == length {
		fmt.Printf("%d \n", sets)
		generatePowerSets(array, 0, m+1, result)
	}

	if m == length && n == length {
		fmt.Printf("%d \n", sets)
		return
	}

}
