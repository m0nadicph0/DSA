# Backtracking

Backtracking is a general algorithmic technique that is used to solve problems by incrementally building a solution 
and testing whether the current solution satisfies the problem constraints. If the current solution does not satisfy 
the constraints, the algorithm backtracks to the previous step and tries a different approach. This process is repeated
 until a solution that satisfies the problem constraints is found or all possible solutions have been tried.

Backtracking is often used for problems that involve searching through a large set of possible solutions, such as the 
classic "n-queens" problem in which you must place n chess queens on an n x n board so that no two queens threaten each other.

The advantage of backtracking is that it can often be used to solve problems that are otherwise difficult or impossible 
to solve with a straightforward algorithm. However, it can be inefficient for some problems, especially those with large 
search spaces or many constraints.

# General format of backtracking algorithms

The general format of a backtracking algorithm is as follows:

```go
func backtrack(result *[]Solution, choices []Choice, solution Solution) {
    if isComplete(solution) {
        *result = append(*result, solution)
        return
    }

    for _, c := range choices {
        if isValid(c, solution) {
            solution = applyChoice(c, solution)
            backtrack(result, choices, solution)
            solution = undoChoice(c, solution)
        }
    }
}

func solveProblem() []Solution {
    var result []Solution
    backtrack(&result, initialChoices, initialSolution)
    return result
}
```

Here, result is a pointer to an array of Solution objects that will contain all of the solutions found by the algorithm.
 choices is an array of Choice objects that represent all of the possible choices that can be made at the current step
 of the algorithm. solution is the current partial solution that the algorithm is building.

The backtrack function recursively explores all possible choices by iterating over the choices array and making each 
valid choice. If the resulting solution is complete (i.e., satisfies all constraints), it is added to the result array. 
Otherwise, the function calls itself recursively with the updated solution.

isComplete is a function that returns true if the current solution is complete and false otherwise. isValid is a function
 that returns true if a given choice is valid in the current solution and false otherwise. applyChoice is a function that 
applies a given choice to the current solution and returns the updated solution. undoChoice is a function that undoes the 
effect of a given choice on the current solution and returns the updated solution.

Finally, the solveProblem function initializes the result array and calls backtrack with the initial choices and solution.
 It returns the result array containing all of the solutions found by the algorithm.

## Easy problems


1. Generate all permutations of a string
1. Generate all subsets of a set
1. Generate all possible combinations of a string of digits that map to letters on a telephone keypad
1. Generate all possible solutions to the n-queens problem
1. Generate all possible solutions to the sudoku puzzle
1. Generate all possible valid IP addresses from a given string of digits
1. Generate all possible solutions to the knight's tour problem on a chessboard
1. Generate all possible solutions to the rat in a maze problem
1. Generate all possible solutions to the sum of subsets problem
1. Generate all possible solutions to the coin change problem
1. Generate all possible solutions to the subset sum problem
1. Generate all possible solutions to the partition problem
1. Generate all possible solutions to the subset sum problem with repeated elements
1. Generate all possible solutions to the largest sum contiguous subarray problem
1. Generate all possible solutions to the longest common subsequence problem
1. Generate all possible solutions to the longest increasing subsequence problem
1. Generate all possible solutions to the maximum subarray problem
1. Generate all possible solutions to the longest palindromic subsequence problem
1. Generate all possible solutions to the longest repeated substring problem
1. Generate all possible solutions to the longest substring without repeating characters problem
1. Generate all possible solutions to the longest valid parentheses problem
1. Generate all possible solutions to the word break problem
1. Generate all possible solutions to the boolean expression evaluation problem
1. Generate all possible solutions to the n-rooks problem
1. Generate all possible solutions to the all paths in a matrix problem
1. Generate all possible solutions to the Hamiltonian cycle problem
1. Generate all possible solutions to the graph coloring problem
1. Generate all possible solutions to the k-coloring problem
1. Generate all possible solutions to the topological sorting problem
1. Generate all possible solutions to the traveling salesman problem


## Intermediate Problems


1. Generate all possible solutions to the regular expression matching problem
1. Generate all possible solutions to the longest common substring problem
1. Generate all possible solutions to the longest zigzag subsequence problem
1. Generate all possible solutions to the word ladder problem
1. Generate all possible solutions to the longest word in a dictionary problem
1. Generate all possible solutions to the maximum number of non-overlapping substrings problem
1. Generate all possible solutions to the maximum sum path in a triangle problem
1. Generate all possible solutions to the longest increasing subarray problem with k swaps allowed
1. Generate all possible solutions to the longest common subsequence of three strings problem
1. Generate all possible solutions to the longest increasing subsequence of three sequences problem
1. Generate all possible solutions to the largest rectangle in a binary matrix problem
1. Generate all possible solutions to the longest palindromic substring problem
1. Generate all possible solutions to the longest common prefix problem
1. Generate all possible solutions to the largest sum subsequence of non-adjacent elements problem
1. Generate all possible solutions to the longest subarray with absolute difference less than or equal to limit problem
1. Generate all possible solutions to the count of palindromic substrings problem
1. Generate all possible solutions to the find k-th permutation problem
1. Generate all possible solutions to the find k-th smallest element in a sorted matrix problem
1. Generate all possible solutions to the find k-th smallest sum of a matrix with sorted rows problem
1. Generate all possible solutions to the find k-th smallest sum of a matrix with sorted columns problem
1. Generate all possible solutions to the find k-th smallest sum of a matrix with sorted diagonals problem
1. Generate all possible solutions to the find k-th largest sum of a matrix with sorted rows problem
1. Generate all possible solutions to the find k-th largest sum of a matrix with sorted columns problem
1. Generate all possible solutions to the find k-th largest sum of a matrix with sorted diagonals problem
1. Generate all possible solutions to the find k-th smallest element in a BST problem
1. Generate all possible solutions to the find k-th smallest element in a sorted array problem
1. Generate all possible solutions to the find k-th smallest element in an unsorted array problem
1. Generate all possible solutions to the find k-th smallest element in a max-heap problem
1. Generate all possible solutions to the find k-th largest element in a min-heap problem
1. Generate all possible solutions to the find k-th largest element in an unsorted array problem

# Advanced problems

1. Generate all possible solutions to the knapsack problem
1. Generate all possible solutions to the cutting stock problem
1. Generate all possible solutions to the multi-dimensional knapsack problem
1. Generate all possible solutions to the traveling salesman problem with time windows
1. Generate all possible solutions to the capacitated vehicle routing problem
1. Generate all possible solutions to the job shop scheduling problem
1. Generate all possible solutions to the bin packing problem
1. Generate all possible solutions to the set cover problem
1. Generate all possible solutions to the maximum independent set problem
1. Generate all possible solutions to the maximum clique problem
1. Generate all possible solutions to the maximum matching problem
1. Generate all possible solutions to the maximum flow problem
1. Generate all possible solutions to the minimum cut problem
1. Generate all possible solutions to the minimum vertex cover problem
1. Generate all possible solutions to the minimum dominating set problem
1. Generate all possible solutions to the minimum spanning tree problem
1. Generate all possible solutions to the maximum cut problem
1. Generate all possible solutions to the maximum independent domination problem
1. Generate all possible solutions to the maximum 2-satisfiability problem
1. Generate all possible solutions to the maximum 3-satisfiability problem
1. Generate all possible solutions to the maximum satisfiability problem
1. Generate all possible solutions to the maximum common subgraph problem
1. Generate all possible solutions to the maximum common subsequence problem
1. Generate all possible solutions to the maximum common subtree problem
1. Generate all possible solutions to the maximum bipartite matching problem
1. Generate all possible solutions to the minimum weighted vertex cover problem
1. Generate all possible solutions to the minimum weighted dominating set problem
1. Generate all possible solutions to the minimum weighted independent set problem
1. Generate all possible solutions to the minimum weighted cut problem
1. Generate all possible solutions to the minimum weighted edge cover problem

## Examples of backtracking

### Subsets

> Problem: Given a set of integers, find all subsets of the set.

To solve this problem using backtracking, we can start with an empty subset and recursively generate all subsets by either including or excluding the next element in the set. The recursive function will backtrack by undoing the inclusion or exclusion of an element in a subset and continue with the remaining elements until all subsets have been generated.

Here are the steps to solve this problem using backtracking:

1. Start with an empty subset.
1. For each element in the set, generate all subsets that include the element by adding it to the current subset and recursively calling the function with the remaining elements in the set.
1. Generate all subsets that exclude the element by not adding it to the current subset and recursively calling the function with the remaining elements in the set.
1. When there are no more elements in the set, add the current subset to the list of subsets.
1. Backtrack by removing the last element from the current subset and continue with the remaining elements until all subsets have been generated.


```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    backtrack(&result, []int{}, nums)
    return result
}

func backtrack(result *[][]int, temp []int, nums []int) {
    *result = append(*result, temp)
    for i := 0; i < len(nums); i++ {
        backtrack(result, append(append([]int{}, temp...), nums[i]), nums[i+1:])
    }
}
```

In this code, the `subsets` function takes in a set of integers nums and returns 
a list of all possible subsets. The `backtrack` function recursively generates all
 subsets by taking in a reference to the list of subsets result, the current subset 
temp, and the remaining elements in the set nums.

In the` backtrack` function, the current subset is added to the list of subsets by 
appending it to `resul`t. Then, for each element in the remaining set `nums`, the function 
is recursively called with the element included in the subset by appending it to `temp`.
 Finally, the function backtracks by removing the last element from `temp` and continuing 
with the remaining elements until all subsets have been generated.

### Permutations

> Problem: Given a set of numbers, find all unique permutations of the set.

To solve this problem using backtracking, we can start with an empty permutation and recursively generate all permutations by choosing the next unused number from the set to be added to the permutation. We can keep track of the used numbers and backtrack by removing the last added number from the permutation and trying another unused number until all permutations have been generated.

Here are the steps to solve this problem using backtracking:

1. Start with an empty permutation and a set of unused numbers.
1. For each unused number in the set, add the number to the permutation and mark it as used.
1. Recursively generate all permutations with the remaining unused numbers and the updated permutation.
1. When there are no more unused numbers, add the permutation to the list of unique permutations.
1. Backtrack by removing the last added number from the permutation and unmarking it as used.
1. Continue with the remaining unused numbers until all permutations have been generated.

```go

func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    used := make([]bool, len(nums))
    results := [][]int{}
    backtrack(nums, used, []int{}, &results)
    return results
}

func backtrack(nums []int, used []bool, temp []int, results *[][]int) {
    if len(temp) == len(nums) {
        *results = append(*results, append([]int{}, temp...))
    } else {
        for i := 0; i < len(nums); i++ {
            if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
                continue
            }
            used[i] = true
            temp = append(temp, nums[i])
            backtrack(nums, used, temp, results)
            used[i] = false
            temp = temp[:len(temp)-1]
        }
    }
}
```

In this code, the permuteUnique function takes in a set of numbers nums and returns 
a list of all unique permutations. The function first sorts the set to handle duplicate 
numbers and initializes a list of boolean flags used to keep track of used numbers. 
The backtrack function recursively generates all unique permutations by taking in the 
set of numbers nums, the list of boolean flags used, the current permutation temp, and 
a reference to the list of unique permutations `results`.

In the backtrack function, the base case is when the length of the current permutation 
is equal to the length of the set, in which case the permutation is added to the list 
of unique permutations. Otherwise, for each unused number in the set, the function adds 
the number to the permutation if it has not been used before and recursively generates 
all permutations with the remaining unused numbers and the updated permutation. 
The function backtracks by removing the last added number from the permutation and 
unmarking it as used, and continues with the remaining unused numbers until all permutations have been generated.

### Find words

> Problem: Given a matrix of lowercase letters and a dictionary of words, find all words in the matrix that are present in the dictionary. A word can start from any character in the matrix and can move to any of the 8 adjacent cells.

To solve this problem using backtracking, we can start with each cell in the matrix and recursively explore all possible paths from that cell that form a word in the dictionary. We can keep track of the current word and backtrack by removing the last added character from the word and exploring other paths until all possible words have been found.

Here are the steps to solve this problem using backtracking:

1. Start with each cell in the matrix and a dictionary of words.
1. For each cell, recursively explore all possible paths that form a word in the dictionary.
1. For each path, add the current cell's character to the current word and check if the word is in the dictionary.
1. If the word is in the dictionary, add it to the list of found words and continue exploring other paths.
1. Backtrack by removing the last added character from the word and continue exploring other paths.
1. Continue with the remaining cells until all possible words have been found.

Here's the code in Golang:

```go

func findWords(board [][]byte, words []string) []string {
    m, n := len(board), len(board[0])
    res := []string{}
    dict := make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            backtrack(board, dict, i, j, []byte{}, &res)
        }
    }
    return res
}

func backtrack(board [][]byte, dict map[string]bool, i, j int, word []byte, res *[]string) {
    m, n := len(board), len(board[0])
    if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == '#' {
        return
    }
    c := board[i][j]
    word = append(word, c)
    if dict[string(word)] {
        delete(dict, string(word))
        *res = append(*res, string(word))
    }
    board[i][j] = '#'
    backtrack(board, dict, i+1, j, word, res)
    backtrack(board, dict, i-1, j, word, res)
    backtrack(board, dict, i, j+1, word, res)
    backtrack(board, dict, i, j-1, word, res)
    backtrack(board, dict, i+1, j+1, word, res)
    backtrack(board, dict, i+1, j-1, word, res)
    backtrack(board, dict, i-1, j+1, word, res)
    backtrack(board, dict, i-1, j-1, word, res)
    board[i][j] = c
    word = word[:len(word)-1]
}
```

In this code, the `findWords` function takes in a matrix of characters board and a dictionary 
of words `words` and returns a list of words in the matrix that are present in the dictionary. 
The function first initializes a map dict to store the words in the dictionary and a list res 
to store the found words. The function then iterates over each cell in the matrix and calls the 
backtrack function to explore all possible paths that form a word in the dictionary.

In the backtrack function, we first check if the current cell is within the matrix bounds and 
if the cell has not already been visited by checking if it is marked with the # character. 
If the cell is valid, we add its character to the current word and check if the word is in the dictionary. 
If the word is in the dictionary, we add it to the list of found words and delete it from the dictionary to avoid duplicates.

We then mark the current cell as visited by changing its character to # and explore all possible paths 
by recursively calling the backtrack function on its neighboring cells. After exploring all possible 
paths, we unmark the current cell by changing its character back to its original value and remove the 
last added character from the current word to backtrack to the previous cell.

Finally, the findWords function returns the list of found words.


### Palindrome partitioning

> Problem: Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

Example:

```
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Explanation:
The palindrome partitions of "aab" are:
- "a", "a", "b" => all substrings are palindrome
- "aa", "b" => all substrings are palindrome
```

Solution: This problem can be solved using backtracking. The idea is to explore all possible ways to partition the string into palindromic substrings. For each possible partition, we check if the current substring is a palindrome using a separate helper function. If it is a palindrome, we add it to the current partition and call the backtrack function recursively with the remaining part of the string and the updated partition. If it is not a palindrome, we backtrack to the previous partition and explore other possibilities.

Here are the steps to solve the problem in English:

1. Define a recursive function called backtrack that takes the input string, the current partition, the current index, and the result slice as arguments.
1. If the current index is equal to the length of the input string, add the current partition to the result slice and return.
1. For each index from the current index to the end of the string, check if the substring from the current index to the current index plus the current length is a palindrome. If it is a palindrome, add the substring to the current partition and call the backtrack function recursively with the remaining part of the string (from the next index) and the updated partition.
1. After exploring all possible paths from the current index, remove the last added substring from the current partition to backtrack to the previous index.

```go
func partition(s string) [][]string {
    res := [][]string{}
    backtrack(s, []string{}, 0, &res)
    return res
}

func backtrack(s string, partition []string, index int, res *[][]string) {
    if index == len(s) {
        temp := make([]string, len(partition))
        copy(temp, partition)
        *res = append(*res, temp)
        return
    }
    for i := index; i < len(s); i++ {
        if isPalindrome(s[index : i+1]) {
            partition = append(partition, s[index:i+1])
            backtrack(s, partition, i+1, res)
            partition = partition[:len(partition)-1]
        }
    }
}

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```

In this code, we define the `partition` function that initializes an empty result slice and calls the backtrack function with the input string, an empty partition slice, the starting index 0, and a pointer to the result slice.

The backtrack function checks if the current index is equal to the length of the input string, and if so, adds the current partition to the result slice and returns. Otherwise, it iterates over the string from the current index to the end, checking if each substring is a palindrome using the isPalindrome helper function. If a substring is a palindrome, it adds it to the current partition and calls the backtrack function recursively with the remaining part of the string (from the next index) and the updated partition. After exploring all possible paths from the current index, it removes the last added substring from the current partition to backtrack to the previous index.

### Combination sum

> Problem: Given a set of integers and a target sum, find all unique combinations of integers in the set that add up to the target sum. The same number may be used multiple times in a combination.

Example:

lua
Copy code
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2+2+3=7 and 7=7 are the only two combinations that add up to the target sum 7.
Solution: This problem can be solved using backtracking. The idea is to explore all possible combinations of integers in the set, keeping track of the current sum and the current combination. We can stop exploring a path if the current sum exceeds the target sum.

Here are the steps to solve the problem in English:

1. Sort the input set in non-descending order to avoid duplicate combinations.
1. Define a recursive function called backtrack that takes the current combination, the current sum, and the current index as arguments.
1. If the current sum equals the target sum, add the current combination to the list of found combinations.
1. If the current sum exceeds the target sum or the current index is greater than the length of the set, stop exploring this path.
1. Otherwise, for each index from the current index to the end of the set, add the number at the index to the current combination and call the backtrack function recursively with the updated combination, the updated sum (sum + number at index), and the same index.
1. After exploring all possible paths from the current index, remove the last added number from the current combination to backtrack to the previous index.



```go
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := [][]int{}
    backtrack(candidates, target, []int{}, 0, &res)
    return res
}

func backtrack(candidates []int, target int, combination []int, index int, res *[][]int) {
    if target == 0 {
        *res = append(*res, combination)
        return
    }
    for i := index; i < len(candidates); i++ {
        if candidates[i] > target {
            break
        }
        newComb := make([]int, len(combination)+1)
        copy(newComb, combination)
        newComb[len(newComb)-1] = candidates[i]
        backtrack(candidates, target-candidates[i], newComb, i, res)
    }
}
```

In this code, we first sort the input set in non-descending order using sort.Ints(candidates) to avoid duplicate combinations. We then define the combinationSum function that initializes an empty result slice and calls the backtrack function with the input set, target sum, an empty combination slice, and the starting index 0.

The backtrack function checks if the current sum equals the target sum and appends the current combination to the result slice if so. It then iterates over the set from the current index to the end, adding the number at each index to the current combination and calling the backtrack function recursively with the updated combination, updated sum (target - number at index), and the same index. Finally, after exploring all possible paths from the current index, the last added number is removed from the current combination to backtrack to the previous index.

### Unique paths

Problem: Given a matrix of integers, matrix, find all unique paths from the top-left corner to the bottom-right corner, where you can only move down or right.

**Example Input:**

```
matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
]
```

**Example Output:**

```
[[1, 4, 7, 8, 9], [1, 4, 5, 8, 9], [1, 2, 5, 8, 9], [1, 2, 5, 6, 9], [1, 2, 3, 6, 9]]
```
Solution:

The idea behind the solution is to use backtracking to explore all possible paths from the top-left corner to the bottom-right corner. We can start at the top-left corner and recursively explore all possible paths by either moving down or right until we reach the bottom-right corner.

We can represent each path as a list of integers, where each integer represents the value of the corresponding cell in the matrix. We can use a helper function dfs to perform the backtracking. The function takes the current position (i, j) in the matrix, the current path path, and the set of all unique paths result.

At each step, we check if we have reached the bottom-right corner of the matrix. If so, we add the current path to the set of all unique paths. Otherwise, we explore all possible paths by moving down and right, and recursively call the dfs function with the updated position and path.


```go
func uniquePaths(matrix [][]int) [][]int {
    result := [][]int{}
    path := []int{}
    dfs(matrix, 0, 0, path, &result)
    return result
}

func dfs(matrix [][]int, i, j int, path []int, result *[][]int) {
    if i == len(matrix)-1 && j == len(matrix[0])-1 {
        // reached the bottom-right corner
        path = append(path, matrix[i][j])
        *result = append(*result, append([]int{}, path...))
        return
    }
    
    if i < len(matrix)-1 {
        // move down
        path = append(path, matrix[i][j])
        dfs(matrix, i+1, j, path, result)
        path = path[:len(path)-1]
    }
    
    if j < len(matrix[0])-1 {
        // move right
        path = append(path, matrix[i][j])
        dfs(matrix, i, j+1, path, result)
        path = path[:len(path)-1]
    }
}
```

We first create an empty slice result to store all the unique paths. We then call the dfs function with the starting position (0, 0), an empty path, and a reference to the result slice.

In the dfs function, we first check if we have reached the bottom-right corner of the matrix. If so, we append the current cell value to the path, make a copy of the path using append([]int{}, path...), and append it to the result slice.

Otherwise, we explore all possible paths by moving down and right, and recursively calling the dfs function with the updated position and path. We append the current cell value to the path before each recursive call, and remove it after each call using `
