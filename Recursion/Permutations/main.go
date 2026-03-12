/*
Problem: Generate All Permutations (Backtracking)
Description

Given a slice of integers, generate all possible permutations of those numbers.

A permutation is an ordering of all elements.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("This is a permutations that will be solved using recursion with backtracking")
	var arrayLength int
	fmt.Println("How many number do you want to enter:")
	fmt.Scanln(&arrayLength)
	//Creating the array to hold the variables
	fmt.Println("PLease enter the number")
	array := make([]int, arrayLength)
	for i := 0; i < arrayLength; i++ {
		fmt.Printf("%d.", i+1)
		fmt.Scanln(&array[i])
	}
	AllPermutations := perm(array)
	fmt.Printf("The total amount of permutations is:%d \n", len(AllPermutations))
	fmt.Println("Would you like to see all permutations? (yes/no)")
	var showAll string
	fmt.Scanln(&showAll)
	if strings.ToLower(showAll) == "yes" {
		for _, combinations := range AllPermutations {
			fmt.Println(combinations)
		}
	} else {
		fmt.Println("Have a good day, goodbye")
	}
}

func perm(TakenArray []int) [][]int {
	var helper func([]int, int)
	result := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			temp := make([]int, len(arr))
			copy(temp, arr)
			result = append(result, temp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					temp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = temp
				} else {
					temp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = temp
				}
			}
		}
	}
	helper(TakenArray, len(TakenArray))
	return result
}
