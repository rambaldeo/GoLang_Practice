/* Problem: Generate All Valid Parentheses

Description
Given an integer n, write a recursive function to generate all combinations of well-formed parentheses consisting of n pairs.

A combination is considered valid if:
- Every opening parenthesis "(" has a corresponding closing parenthesis ")"
- Parentheses are properly nested

Input: 3

Output (order doesn’t matter):
((()))
(()())
(())()
()(())
()()()

Constraints
- Use recursion (backtracking) to build the combinations
- You cannot generate all permutations and then filter — build only valid sequences
- At no point should closing parentheses exceed opening ones
- Return a slice of strings (in Go: []string)

Hints
- Think of this as making two choices at each step:
  Add "(" if you still have opening parentheses left
  Add ")" if it won’t make the sequence invalid
- Track:
  Number of "(" used
  Number of ")" used
- Base case: when the length of the string is 2 * n, add it to the result
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a recursion problem to generate all valid parentheses")
	var userInput int
	fmt.Println("Please enter the integer")
	fmt.Scan(&userInput)

	buildString := building(userInput)
	for _, parentheses := range buildString {
		fmt.Println(parentheses)
	}
}

func building(userInput int) []string {
	result := []string{}
	var helper func(current string, open int, close int)

	helper = func(arr string, open, close int) {
		//base case
		if len(arr) == 2*userInput {
			result = append(result, arr)
			return
		}
		//add (
		if open < userInput {
			helper(arr+"(", open+1, close)
		}
		//add )
		if close < open {
			helper(arr+")", open, close+1)
		}
	}
	helper("", 0, 0)
	return result
}
