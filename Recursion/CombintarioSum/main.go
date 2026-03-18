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