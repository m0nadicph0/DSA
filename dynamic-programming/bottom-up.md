# Bottom-Up DP Examples

  * [1.  Fibonacci series](#1--fibonacci-series)
  * [2.  Longest common subsequence](#2--longest-common-subsequence)
  * [3.  Knapsack problem](#3--knapsack-problem)
  * [4.  Coin change problem](#4--coin-change-problem)
  * [5.  Maximum subarray problem](#5--maximum-subarray-problem)
  * [6.  Longest increasing subsequence](#6--longest-increasing-subsequence)
  * [7.  Matrix chain multiplication](#7--matrix-chain-multiplication)
  * [8.  Edit distance problem](#8--edit-distance-problem)
  * [9.  Longest palindromic subsequence](#9--longest-palindromic-subsequence)
  * [10. Rod cutting problem](#10-rod-cutting-problem)
  * [11. Subset sum problem](#11-subset-sum-problem)
  * [12. Minimum edit distance](#12-minimum-edit-distance)
  * [13. Optimal binary search tree](#13-optimal-binary-search-tree)
  * [14. Longest bitonic subsequence](#14-longest-bitonic-subsequence)
  * [15. Counting problems (such as counting the number of ways to make change)](#15-counting-problems-such-as-counting-the-number-of-ways-to-make-change)
  * [16. Longest common substring](#16-longest-common-substring)
  * [17. Maximum sum increasing subsequence](#17-maximum-sum-increasing-subsequence)
  * [18. Painting fence problem](#18-painting-fence-problem)
  * [19. Maximum sum subsequence without adjacent elements](#19-maximum-sum-subsequence-without-adjacent-elements)
  * [20. 0/1 Knapsack problem](#20-0-1-knapsack-problem)
  * [21. Egg dropping problem](#21-egg-dropping-problem)
  * [22. Partition problem](#22-partition-problem)
  * [23. Longest zigzag subsequence](#23-longest-zigzag-subsequence)
  * [24. String matching problems](#24-string-matching-problems)
  * [25. Maximum sum subarray with at most k elements](#25-maximum-sum-subarray-with-at-most-k-elements)
  * [26. Longest alternating subsequence](#26-longest-alternating-subsequence)
  * [27. Tower of Hanoi problem](#27-tower-of-hanoi-problem)
  * [28. Longest common prefix](#28-longest-common-prefix)
  * [29. Shortest common supersequence](#29-shortest-common-supersequence)
  * [30. Maximum size subset with given sum.](#30-maximum-size-subset-with-given-sum)




## 1.  Fibonacci series
The Fibonacci series is a famous mathematical sequence where each number is the sum of the two preceding ones, starting from 0 and 1. The first few numbers in the sequence are: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

The problem of finding the nth number in the Fibonacci series can be approached in various ways, but one of the most efficient methods is using dynamic programming. Specifically, we can use the bottom-up approach to avoid the computational overhead of recursion and redundant calculations.

The basic idea of the bottom-up DP approach is to start from the smallest subproblems (usually the base cases) and iteratively compute the solutions for larger subproblems until we reach the desired solution for the original problem.

Here are the steps to solve the Fibonacci series problem using bottom-up DP:

1.  Define an array or slice to store the computed values of the Fibonacci sequence. Initialize it with the base cases of the series, which are f\[0\]=0 and f\[1\]=1.
2.  Iterate over the remaining indices of the array, starting from 2 up to n (the desired index of the Fibonacci number to compute).
3.  For each index i, compute the ith Fibonacci number by summing the two preceding numbers f\[i-1\] and f\[i-2\], and store the result in f\[i\].
4.  After the iteration, the desired Fibonacci number is stored in f\[n\], so we can return it as the solution.

Here's an example implementation of the bottom-up DP approach in Go:
```go
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    f := make([]int, n+1)
    f[0], f[1] = 0, 1
    for i := 2; i <= n; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[n]
}
```
In this implementation, we first check if the input n is less than 2, in which case we can directly return n as the solution (since f\[0\]=0 and f\[1\]=1). Otherwise, we create a slice f of size n+1 to store the computed Fibonacci numbers. We initialize the first two values of f with the base cases. Then, we iterate over the remaining indices of f and compute each Fibonacci number by summing the two preceding ones. Finally, we return the nth Fibonacci number stored in f\[n\].

For example, if we call fibonacci(6), the output will be 8, since the 6th Fibonacci number is 8 (0, 1, 1, 2, 3, 5, 8).

##  2.  Longest common subsequence

The longest common subsequence (LCS) problem is a classic computer science problem that involves finding the longest sequence of characters that is common to two or more input strings. In other words, given two or more strings, the LCS problem seeks the longest subsequence that appears in all of them.

For example, consider the following two strings:


```
string 1: ABCDEF
string 2: ACDEG
```

The LCS of these two strings is "ACDE". This is the longest subsequence that appears in both strings.

Here is another example:


```
string 1: AGGTAB
string 2: GXTXAYB
```

The LCS of these two strings is "GTAB". Again, this is the longest subsequence that appears in both strings.

The LCS problem can also involve more than two strings. For example, consider the following three strings:


```
string 1: ABCBDAB
string 2: BDCABA
string 3: BADACB
```

The LCS of these three strings is "BDAB". This is the longest subsequence that appears in all three strings.

One way to solve the LCS problem is to use dynamic programming. This involves building a table that stores the LCS for all possible pairs of prefixes of the input strings. The final entry in the table will contain the LCS of the full input strings.

here are the steps to solve the LCS problem using bottom-up dynamic programming:

1.  Create a 2D table of size (m+1)x(n+1), where m and n are the lengths of the input strings. This table will store the LCS for all possible pairs of prefixes of the input strings.
    
2.  Initialize the first row and first column of the table to 0, as the LCS of a string with an empty string is 0.
    
3.  Traverse the table row by row and column by column, filling in each entry as follows:
        * If the characters at the current positions of the two strings match, set the value of the current entry in the table to 1 plus the value in the previous diagonal entry (i.e., one position up and to the left).
        * If the characters do not match, set the value of the current entry to the maximum of the value in the previous row entry and the value in the previous column entry.

4.  The final entry in the table will contain the LCS of the full input strings.
    
5.  To reconstruct the LCS, start at the final entry in the table and follow the path of maximum values (i.e., the path of diagonal entries) until you reach an entry with a value of 0. The characters corresponding to each diagonal entry on this path form the LCS.
    

Here is the code in Go:
```go
func lcs(s1, s2 string) string {
    m, n := len(s1), len(s2)
    table := make([][]int, m+1)
    for i := range table {
        table[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        table[i][0] = 0
    }
    for j := 0; j <= n; j++ {
        table[0][j] = 0
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s1[i-1] == s2[j-1] {
                table[i][j] = 1 + table[i-1][j-1]
            } else {
                table[i][j] = max(table[i-1][j], table[i][j-1])
            }
        }
    }
    lcs := make([]byte, 0, table[m][n])
    for i, j := m, n; i > 0 && j > 0; {
        if s1[i-1] == s2[j-1] {
            lcs = append(lcs, s1[i-1])
            i--
            j--
        } else if table[i-1][j] > table[i][j-1] {
            i--
        } else {
            j--
        }
    }
    // Reverse lcs and return as string
    for i, n := 0, len(lcs); i < n/2; i++ {
        lcs[i], lcs[n-1-i] = lcs[n-1-i], lcs[i]
    }
    return string(lcs)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
In the code above, the `lcs` function takes two input strings `s1` and `s2` and returns their LCS as a string. The function uses the bottom-up dynamic programming approach described earlier. The `max` function is a helper function that returns the maximum of two integers.


##  3.  Knapsack problem
The Knapsack problem is a well-known optimization problem in computer science and operations research. It involves selecting a set of items with given values and weights to fit into a knapsack with a limited weight capacity, such that the total value of the selected items is maximized.

For example, imagine that you are a thief who has broken into a jewelry store. You have a knapsack with a limited weight capacity, and you want to fill it with the most valuable items possible. The store has the following items:

| Item | Value ($) | Weight (kg) |
| --- | --- | --- |
| A   | 60  | 10  |
| B   | 100 | 20  |
| C   | 120 | 30  |
| D   | 40  | 5   |
| E   | 80  | 12  |

Your knapsack can only hold a maximum weight of 50 kg. Which items should you select to maximize the value of your loot?

To solve this problem, we can use a variety of algorithms, including dynamic programming and greedy algorithms. One possible approach is to use a dynamic programming algorithm that builds a table of subproblems and their solutions. In this case, the subproblems correspond to different subsets of items and different capacities of the knapsack.

Using this algorithm, we can construct a table where each cell represents the maximum value that can be obtained by filling a knapsack of a certain capacity with a certain subset of items. We can fill out this table by considering two cases for each item: either the item is included in the knapsack, or it is not.

After filling out the table, we can find the optimal subset of items by backtracking from the last cell of the table to the first, considering which items were included at each step. In this example, the optimal subset of items is {B, C, D}, with a total value of $260 and a total weight of 55 kg (which is just over the knapsack's weight capacity).

here are the steps to solve the Knapsack problem using a bottom-up dynamic programming (DP) approach:

1.  Create a two-dimensional slice (i.e., a table) with dimensions (number of items + 1) x (knapsack capacity + 1). Initialize all cells to 0, since the maximum value that can be obtained with a weight capacity of 0 or with no items is 0.
    
2.  Iterate over all possible combinations of items and knapsack capacities, and fill in the table by computing the maximum value that can be obtained using the items up to that point and the current knapsack capacity. For each item i and capacity j, there are two possible cases: a. The item is not included in the solution: In this case, the maximum value is the same as the value that can be obtained using the same capacity but with the previous items (i.e., table\[i-1\]\[j\]). b. The item is included in the solution: In this case, the maximum value is the sum of the item's value and the maximum value that can be obtained using the remaining capacity and the previous items (i.e., table\[i-1\]\[j-weight\[i\]\] + value\[i\]), where weight\[i\] is the weight of the current item and value\[i\] is its value.
    
3.  Once the table is filled, the maximum value that can be obtained using all items and a knapsack capacity of W is the value in the last cell of the table (i.e., table\[number of items\]\[W\]).
    
4.  To find the items that were selected, we can backtrack from the last cell of the table to the first, considering whether each item was included or not. If the value of the current cell is equal to the value of the previous cell (i.e., the current item was not included), we move to the cell above. Otherwise, we move diagonally up-left and include the current item in the solution.
    

Here's some Go code that implements this algorithm:
```go
func knapsack(items []item, W int) (int, []bool) {
    n := len(items)
    table := make([][]int, n+1)
    for i := range table {
        table[i] = make([]int, W+1)
    }

    for i := 1; i <= n; i++ {
        for j := 0; j <= W; j++ {
            if items[i-1].weight > j {
                table[i][j] = table[i-1][j]
            } else {
                table[i][j] = max(table[i-1][j], table[i-1][j-items[i-1].weight]+items[i-1].value)
            }
        }
    }

    value := table[n][W]
    selected := make([]bool, n)
    i, j := n, W
    for i > 0 && j > 0 {
        if table[i][j] != table[i-1][j] {
            selected[i-1] = true
            j -= items[i-1].weight
        }
        i--
    }

    return value, selected
}

type item struct {
    value, weight int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
Here, `items` is a slice of `item` structs, where each item has a value and a weight. The function `knapsack` takes this slice and a knapsack capacity `W` as input, and returns the maximum value that can be obtained and a slice of boolean values indicating which items were selected.

##  4.  Coin change problem
The Coin Change problem is a classic algorithmic problem in which we are given a set of coins of different denominations and a target amount. The task is to find the minimum number of coins needed to make the change for the target amount.

For example, suppose we have the following coins: \[1, 5, 10, 25\] and we want to make the change for the amount 32. The minimum number of coins needed to make the change for 32 is 2 (one 25 cent coin and one 7 cent coin).

Here are the steps to solve the Coin Change problem using bottom-up dynamic programming:

1.  Initialize a 1D array dp\[\] with size equal to the target amount + 1, and set all values to infinity except dp\[0\] which is set to 0.
2.  Loop through each coin denomination from the smallest to the largest, and for each denomination j, loop through each target amount i from j to the target amount.
3.  For each i, calculate dp\[i\] as the minimum of dp\[i\] and dp\[i-j\] + 1.
4.  After the loops, return dp\[target amount\].

Here's the Go code for the bottom-up DP solution:
```go
func coinChange(coins []int, target int) int {
    dp := make([]int, target+1)
    for i := range dp {
        dp[i] = math.MaxInt32
    }
    dp[0] = 0

    for _, j := range coins {
        for i := j; i <= target; i++ {
            dp[i] = min(dp[i], dp[i-j]+1)
        }
    }

    if dp[target] == math.MaxInt32 {
        return -1 // no solution
    }
    return dp[target]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```
In the code above, we first initialize the `dp[]` array with infinity except for`dp[0]`. We then loop through each coin denomination, and for each denomination j, we loop through each target amount i from j to the target amount. For each i, we calculate `dp[i] `as the minimum of `dp[i]` and `dp[i-j]+1.` Finally, we return `dp[target] `which gives us the minimum number of coins needed to make the change for the target amount.

##  5.  Maximum subarray problem
The Maximum subarray problem is a classic algorithmic problem that involves finding the contiguous subarray within a one-dimensional array of numbers that has the largest sum.

For example, consider the following array:

```
arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
```
The maximum subarray in this case would be `[4, -1, 2, 1]` with a sum of `6`.

To solve the Maximum subarray problem using bottom-up dynamic programming, you can follow these steps:

1.  Define a one-dimensional array `dp` with the same length as the input array `arr`.
2.  Initialize `dp[0]` with the first element of `arr`.
3.  Loop through the remaining elements of `arr` and for each element `arr[i]`, calculate `dp[i]` as the maximum value between `dp[i-1]+arr[i]` and `arr[i]`.
4.  After the loop is complete, the maximum sum subarray will be the contiguous subarray within `arr` that starts at the index `start` such that `dp[start]` is the maximum value within `dp`.

Here is the code snippet in Go to solve the Maximum subarray problem using bottom-up dynamic programming:
```go
func maxSubarraySum(arr []int) []int {
    n := len(arr)
    dp := make([]int, n)
    dp[0] = arr[0]
    maxSum := dp[0]
    start, end := 0, 0
    
    for i := 1; i < n; i++ {
        dp[i] = max(dp[i-1]+arr[i], arr[i])
        if dp[i] > maxSum {
            maxSum = dp[i]
            end = i
        }
    }
    
    // Find start index of max subarray
    for i := end; i >= 0; i-- {
        if dp[i] == arr[i] || dp[i-1]+arr[i] < dp[i] {
            start = i
            break
        }
    }
    
    return arr[start : end+1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```
You can test this function with the example array `[-2, 1, -3, 4, -1, 2, 1, -5, 4]` and it should return `[4, -1, 2, 1]` as the maximum subarray.

##  6.  Longest increasing subsequence
The Longest Increasing Subsequence (LIS) problem is a classic problem in computer science and mathematics that involves finding the length of the longest subsequence of a given sequence that is strictly increasing.

For example, consider the sequence {3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}. The longest increasing subsequence in this sequence is {1, 2, 5, 6}, which has length 4.

To solve the LIS problem using a bottom-up dynamic programming (DP) approach, we can use an array dp where dp\[i\] represents the length of the longest increasing subsequence that ends at index i. We start by initializing dp\[i\] to 1 for all i because the longest increasing subsequence that ends at i is just the element at index i itself.

We then iterate over the remaining indices j from 1 to n-1 (where n is the length of the sequence), and for each j, we iterate over all the previous indices i from 0 to j-1. For each i, we check if the element at index j is greater than the element at index i. If it is, we update dp\[j\] to be the maximum of dp\[j\] and dp\[i\] + 1, because the longest increasing subsequence that ends at j can be extended by adding the element at index j to the longest increasing subsequence that ends at i.

Finally, we iterate over the entire dp array to find the maximum value, which represents the length of the overall longest increasing subsequence.

Here is the Go code to implement this bottom-up DP approach:
```go
func LIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }

    for j := 1; j < n; j++ {
        for i := 0; i < j; i++ {
            if nums[j] > nums[i] {
                dp[j] = max(dp[j], dp[i]+1)
            }
        }
    }

    maxLIS := 0
    for _, v := range dp {
        if v > maxLIS {
            maxLIS = v
        }
    }
    return maxLIS
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
To use this function, we can simply call it with our input sequence:
```go
nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
fmt.Println(LIS(nums)) // Output: 4
```

##  7.  Matrix chain multiplication
The Matrix Chain Multiplication problem is a classic computer science problem that involves finding the most efficient way to multiply a chain of matrices. The goal is to minimize the total number of multiplications required to compute the product of all matrices.

For example, consider the following chain of matrices:
```
A = 10x30
B = 30x5
C = 5x60
D = 60x10
```
We can multiply the matrices in different orders. For instance, we can multiply `(AB)C` or `A(BC)` or `A((BC)D)` or `((AB)C)D`, and so on. Each way of multiplying the matrices has a different cost in terms of the total number of multiplications required. The goal is to find the order of multiplication that minimizes the total cost.

The solution to the Matrix Chain Multiplication problem can be obtained using dynamic programming, specifically the bottom-up approach. The idea is to start by solving subproblems of increasing size and combining their solutions to obtain the solution to the original problem.

Here are the steps to solve the Matrix Chain Multiplication problem using bottom-up dynamic programming:

1.  Define a 2D array `m` of size `n x n`, where `n` is the number of matrices in the chain. The element `m[i][j]` will store the minimum cost of multiplying the matrices from index `i` to index `j` inclusive.
    
2.  Initialize the diagonal elements of the array `m` to zero, since multiplying a single matrix has no cost.
    
3.  For each subchain length `l` from `2` to `n`, do the following:
    
    a. For each possible subchain starting index `i` (from `1` to `n-l+1`), do the following:
    
    i. Compute the minimum cost of multiplying the subchain `i` to `i+l-1` using the formula `m[i][i+l-1] = min(m[i][k] + m[k+1][i+l-1] + p[i-1]*p[k]*p[i+l-1])`, where `k` ranges from `i` to `i+l-2`, and `p` is an array of matrix dimensions such that the `i`-th matrix has dimensions `p[i-1] x p[i]`.
    
4.  The minimum cost of multiplying the entire chain is stored in `m[1][n]`.
    

Here's the Go code for solving the Matrix Chain Multiplication problem using bottom-up dynamic programming:

```go
func matrixChainOrder(p []int) int {
    n := len(p) - 1
    m := make([][]int, n+1)
    for i := 1; i <= n; i++ {
        m[i] = make([]int, n+1)
        m[i][i] = 0
    }
    for l := 2; l <= n; l++ {
        for i := 1; i <= n-l+1; i++ {
            j := i + l - 1
            m[i][j] = math.MaxInt32
            for k := i; k <= j-1; k++ {
                q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
                if q < m[i][j] {
                    m[i][j] = q
                }
            }
        }
    }
    return m[1][n]
}
```
The function `matrixChainOrder` takes an array `p` of matrix dimensions and returns the minimum cost of multiplying the entire chain of matrices.

##  8.  Edit distance problem

The Edit distance problem, also known as Levenshtein distance, is a measure of how different two strings are from each other. The goal of this problem is to determine the minimum number of operations required to transform one string into another, where an operation can be either insertion, deletion, or substitution of a single character.

For example, consider the two strings "kitten" and "sitting". To transform "kitten" into "sitting", we need to perform the following operations:

* Substitute "s" for "k"
* Substitute "i" for "e"
* Insert "t" after "i"
* Insert "t" after "n"
* Substitute "g" for "e"

Thus, the minimum number of operations required is 5.

To solve this problem using bottom-up dynamic programming, we can create a matrix where the rows represent the characters of the first string and the columns represent the characters of the second string. We then fill in the matrix by computing the minimum number of operations required to transform the substrings up to the current position.

The code snippet in Go for solving the Edit distance problem using bottom-up dynamic programming is as follows:

```go
func min(x, y, z int) int {
    if x < y {
        if x < z {
            return x
        } else {
            return z
        }
    } else {
        if y < z {
            return y
        } else {
            return z
        }
    }
}

func editDistance(s1, s2 string) int {
    m, n := len(s1), len(s2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
```
The `min` function takes three integers as input and returns the minimum of the three. The `editDistance` function takes two strings as input and returns the minimum number of operations required to transform the first string into the second string.

In the `editDistance` function, we first initialize the dynamic programming matrix `dp` with the base cases, where the distance between the empty string and any non-empty string is the length of the non-empty string. We then fill in the rest of the matrix by computing the minimum of the three possible operations (insertion, deletion, and substitution) and adding 1 to it.

Finally, we return the value in the bottom-right corner of the matrix, which represents the minimum number of operations required to transform the entire first string into the entire second string.

##  9.  Longest palindromic subsequence
The Longest Palindromic Subsequence (LPS) problem is a classic dynamic programming problem that involves finding the length of the longest subsequence that is also a palindrome. A subsequence is a sequence of characters that appear in the same order as they do in the original string, but not necessarily consecutively. A palindrome is a sequence of characters that reads the same backward as forward.

For example, the LPS of the string "character" is "carac" which is a palindrome of length 5.

To solve the LPS problem using bottom-up dynamic programming, we can use a 2D table where each cell (i,j) represents the length of the LPS of the subsequence starting at index i and ending at index j.

Here are the steps to solve the LPS problem using bottom-up DP:

1.  Create a 2D table lpsTable with dimensions n x n, where n is the length of the input string.
    
2.  Initialize all the cells in the table to 0.
    
3.  For each i, set lpsTable\[i\]\[i\] to 1 since a subsequence consisting of a single character is always a palindrome.
    
4.  Iterate over the table for all values of L starting from 2 up to the length of the input string.
    
5.  For each value of L, iterate over all possible starting indices i up to n-L+1.
    
6.  Compute the ending index j=i+L-1.
    
7.  If the characters at indices i and j are equal, set lpsTable\[i\]\[j\] to 2 + lpsTable\[i+1\]\[j-1\].
    
8.  Otherwise, set lpsTable\[i\]\[j\] to the maximum of lpsTable\[i+1\]\[j\] and lpsTable\[i\]\[j-1\].
    
9.  The final result will be the value in the cell lpsTable\[0\]\[n-1\].
    

Here is the Go code to implement the bottom-up DP solution for the LPS problem:
```go
func lps(input string) int {
    n := len(input)
    lpsTable := make([][]int, n)
    for i := range lpsTable {
        lpsTable[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        lpsTable[i][i] = 1
    }

    for L := 2; L <= n; L++ {
        for i := 0; i < n-L+1; i++ {
            j := i + L - 1
            if input[i] == input[j] {
                lpsTable[i][j] = 2 + lpsTable[i+1][j-1]
            } else {
                lpsTable[i][j] = max(lpsTable[i+1][j], lpsTable[i][j-1])
            }
        }
    }

    return lpsTable[0][n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
In this implementation, we create a 2D slice lpsTable to hold the values of the LPS subsequence lengths. We initialize the diagonal of the table to 1, as a subsequence of length 1 is always a palindrome. We then fill in the table from bottom-up, following the steps outlined above. The max function is used to find the maximum value of two integers. The function returns the length of the LPS of the input string.

##  10. Rod cutting problem
The Rod cutting problem is a classic problem in dynamic programming where we have a rod of length n and we need to cut it into smaller pieces to maximize its value. Each piece of the rod has a different value associated with it. The goal is to find the maximum value that can be obtained by cutting the rod into pieces.

For example, let's say we have a rod of length 6, and we have the following prices for different pieces of the rod:

| Length (in inches) | 1   | 2   | 3   | 4   | 5   | 6   |
| --- | --- | --- | --- | --- | --- | --- |
| Price ($) | 1   | 5   | 8   | 9   | 10  | 17  |
The optimal way to cut the rod in this case is to cut it into two pieces of length 2 and 4, yielding a total value of $5 + $9 = $14.

To solve this problem using bottom-up dynamic programming, we need to follow the following steps:

1.  Create a table to store the maximum value for each length of the rod.
    
2.  Initialize the first row of the table with zeros since we cannot cut a rod of length 0.
    
3.  For each length of the rod, consider all possible ways of cutting the rod and update the maximum value for that length using the formula:
    
    `max_value[length] = max(max_value[length], price[i] + max_value[length-i-1])`
    
    Here, `price[i]` is the value of the piece of length `i+1`, and `max_value[length-i-1]` is the maximum value that can be obtained from the remaining piece of length `length-i-1`.
    
4.  After considering all possible ways of cutting the rod for all lengths up to `n`, the maximum value for the entire rod will be stored in `max_value[n]`.
    

Here is the Go code to solve the Rod cutting problem using bottom-up dynamic programming:
```go
func rodCutting(lengths []int, prices []int, n int) int {
    max_value := make([]int, n+1)
    max_value[0] = 0

    for i := 1; i <= n; i++ {
        for j := 0; j < len(lengths); j++ {
            if lengths[j] <= i {
                max_value[i] = max(max_value[i], prices[j] + max_value[i-lengths[j]])
            }
        }
    }
    return max_value[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
Here, `lengths` and `prices` are arrays containing the lengths and prices of the different pieces of the rod, and `n` is the length of the rod that we want to cut. The function returns the maximum value that can be obtained by cutting the rod of length `n`.

##  11. Subset sum problem
**Subset Sum Problem:**

The Subset Sum Problem is a well-known computational problem in computer science, which can be stated as follows: Given a set of integers, is there a non-empty subset whose sum is zero? The problem is known to be NP-complete, which means that it is believed to be computationally intractable for large instances. However, dynamic programming can be used to solve the problem efficiently for small instances.

**Example:**

Suppose we have the set S = {3, 1, 5, 9, 12}, and we want to know whether there is a non-empty subset of S whose sum is zero. One possible solution is to take the subset {3, 9} and subtract their sum from the sum of the remaining elements, which is 1 + 5 + 12 = 18 - 12 = 6. Since 3 + 9 - 6 = 6, we can conclude that there is indeed a non-empty subset of S whose sum is zero.

**Steps to solve the problem using bottom-up DP:**

1.  Initialize a boolean table of size (n+1) x (sum+1), where n is the number of elements in the set and sum is the maximum possible sum of any subset.
    
2.  Initialize the first row of the table to False, except for the first cell which is set to True, since it is always possible to form a subset of sum zero from an empty set.
    
3.  For each element i in the set S, starting from the first element:
    
    a. For each possible sum j from 0 to sum, starting from j = i:
    
    i. Set the value of the table at position (i, j) to the OR of the values in the previous row, i.e. table\[i-1\]\[j\], and the value at position (i-1, j-i), i.e. table\[i-1\]\[j-i\]. This represents the possibility of forming a subset of sum j with or without the current element i.
    
4.  Return the value of the table at position (n, sum), which represents whether a non-empty subset of S can be formed with sum zero.
    

**Code snippets in Go:**
```go
func subsetSum(S []int, sum int) bool {
    n := len(S)
    table := make([][]bool, n+1)
    for i := range table {
        table[i] = make([]bool, sum+1)
    }

    for j := 0; j <= sum; j++ {
        table[0][j] = false
    }
    table[0][0] = true

    for i := 1; i <= n; i++ {
        for j := i; j <= sum; j++ {
            table[i][j] = table[i-1][j] || table[i-1][j-S[i-1]]
        }
    }

    return table[n][sum]
}

```
The function `subsetSum` takes a slice of integers `S` and an integer `sum` as input, and returns a boolean value indicating whether a non-empty subset of `S` can be formed with sum zero. The function implements the above steps using a two-dimensional boolean slice `table`. The function can be called as follows:

```go
S := []int{3, 1, 5, 9, 12}
sum := 0
fmt.Println(subsetSum(S, sum)) // Output: true
```
In this example, the function returns true, indicating that there is a non-empty subset of S whose sum is zero, as shown in the previous example.

##  12. Minimum edit distance
The Minimum Edit Distance problem is a classic dynamic programming problem that involves finding the minimum number of operations required to transform one string into another. The operations that can be performed include insertion, deletion, and substitution of a single character.

For example, consider the two strings "kitten" and "sitting". We want to find the minimum number of operations required to transform "kitten" into "sitting". One possible solution is as follows:

1.  Replace 'k' with 's'
2.  Replace 'e' with 'i'
3.  Insert 't' after 'i'
4.  Replace 't' with 't'
5.  Replace 'e' with 'i'
6.  Replace 'n' with 'g'

So the minimum edit distance between "kitten" and "sitting" is 6.

To solve this problem using bottom-up dynamic programming, we can create a 2D array where the value at position (i, j) represents the minimum edit distance required to transform the first i characters of string A into the first j characters of string B. We can fill up this array in a bottom-up manner by computing the minimum edit distance for each pair of substrings.

Here is the code snippet in Go:
```go
func minEditDist(A, B string) int {
    m, n := len(A), len(B)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 {
                dp[i][j] = j
            } else if j == 0 {
                dp[i][j] = i
            } else if A[i-1] == B[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
            }
        }
    }

    return dp[m][n]
}

func min(a, b, c int) int {
    if a < b && a < c {
        return a
    } else if b < a && b < c {
        return b
    } else {
        return c
    }
}
```
In the above code, we first initialize the 2D array with the base cases where one of the substrings is empty. Then, we iterate through each pair of substrings and fill up the array using the recurrence relation:
```go
dp[i][j] = min(
    dp[i-1][j] + 1,  // deletion
    dp[i][j-1] + 1,  // insertion
    dp[i-1][j-1] + 1 // substitution
)

```
Once we have filled up the entire array, the value at the bottom-right corner represents the minimum edit distance between the two strings.

##  13. Optimal binary search tree
Optimal binary search tree is a problem in computer science where given a set of keys and their respective probabilities, the task is to construct a binary search tree such that the expected search time is minimized.

An example of this problem would be given the following keys and their probabilities:
```
keys = { 10, 20, 30, 40 }
probabilities = { 0.1, 0.2, 0.3, 0.4 }

```
Our task is to construct a binary search tree that minimizes the expected search time. In this example, the optimal binary search tree would have the following structure:
```
     30
   /    \
10      40
  \
   20

```
The expected search time for this tree is 1.5.

Here are the steps to solve the Optimal binary search tree problem using bottom-up dynamic programming:

1.  Create an n+1 x n+1 table where n is the number of keys. The table will be used to store the minimum expected search time for a sub-tree containing keys i through j.
    
2.  Fill the diagonal of the table with the probability of the key at that index.
    
3.  Starting with sub-trees of size 2, iterate over all sub-trees of increasing size up to n.
    
4.  For each sub-tree, consider all possible root nodes and calculate the expected search time using the values stored in the table for its left and right subtrees.
    
5.  Store the minimum expected search time for the sub-tree in the table.
    
6.  Once all sub-trees have been considered, the optimal binary search tree can be constructed using the values in the table.
    

Here is an implementation of the bottom-up dynamic programming solution to the Optimal binary search tree problem in Go:

```go
func optimalBST(keys []int, probs []float64) float64 {
    n := len(keys)
    table := make([][]float64, n+1)
    for i := range table {
        table[i] = make([]float64, n+1)
    }

    // Fill the diagonal of the table with the probability of the key at that index.
    for i := 0; i < n; i++ {
        table[i][i] = probs[i]
    }
    table[n][n] = probs[n]

    // Fill the table bottom-up for increasing sub-tree sizes.
    for k := 2; k <= n; k++ {
        for i := 0; i <= n-k+1; i++ {
            j := i + k - 1
            table[i][j] = math.MaxFloat64
            sum := probs[i]
            for p := i; p <= j; p++ {
                // Calculate expected search time for sub-trees with root at p.
                cost := table[i][p-1] + table[p+1][j]
                if cost < table[i][j] {
                    table[i][j] = cost
                }
                if p < j {
                    sum += probs[p+1]
                }
            }
            table[i][j] += sum
        }
    }

    // Return the minimum expected search time for the entire tree.
    return table[0][n-1]
}
```
The time complexity of this implementation is O(n^3), where n is the number of keys.

##  14. Longest bitonic subsequence
The Longest bitonic subsequence problem is a variation of the Longest Increasing Subsequence (LIS) problem. In this problem, we are given a sequence of numbers, and we need to find the length of the longest subsequence that is first increasing and then decreasing.

For example, consider the sequence \[1, 3, 5, 4, 2\]. A bitonic subsequence of this sequence could be \[1, 3, 5, 4, 2\], which is first increasing and then decreasing. The length of this subsequence is 5.

To solve this problem using bottom-up dynamic programming, we can follow the following steps:

1.  Initialize an array dp of length n, where n is the length of the input sequence. Set each element of dp to 1, as each element in the sequence is a bitonic subsequence of length 1.
    
2.  Traverse the input sequence from left to right and for each index i, find the length of the longest increasing subsequence ending at index i. Store this value in an array lis.
    
3.  Traverse the input sequence from right to left and for each index i, find the length of the longest decreasing subsequence starting at index i. Store this value in an array lds.
    
4.  For each index i in the input sequence, compute the length of the longest bitonic subsequence ending at index i by adding the length of the longest increasing subsequence ending at index i and the length of the longest decreasing subsequence starting at index i. Subtract 1 from this value, as the element at index i is counted twice. Store this value in dp\[i\].
    
5.  Find the maximum value in the dp array and return it as the length of the longest bitonic subsequence.
    

Here's the Go code to implement the above approach:
```go
func LongestBitonicSubsequence(seq []int) int {
    n := len(seq)
    dp := make([]int, n)
    lis := make([]int, n)
    lds := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = 1
        lis[i] = 1
        lds[i] = 1
    }

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if seq[i] > seq[j] && lis[j]+1 > lis[i] {
                lis[i] = lis[j] + 1
            }
        }
    }

    for i := n - 2; i >= 0; i-- {
        for j := n - 1; j > i; j-- {
            if seq[i] > seq[j] && lds[j]+1 > lds[i] {
                lds[i] = lds[j] + 1
            }
        }
    }

    for i := 0; i < n; i++ {
        dp[i] = lis[i] + lds[i] - 1
    }

    max := dp[0]
    for i := 1; i < n; i++ {
        if dp[i] > max {
            max = dp[i]
        }
    }

    return max
}
```

The time complexity of this approach is O(n^2), where n is the length of the input sequence, as we need to traverse the input sequence twice and for each index, we need to find the length of the longest increasing and decreasing subsequences.

##  15. Counting problems (such as counting the number of ways to make change)

Counting problems refer to problems that involve counting the number of possible ways to do something, given a set of constraints or rules. One example of a counting problem is the "coin change" problem, which asks us to find the number of ways to make change for a given amount using a set of denominations.

For example, suppose we have the following denominations: {1, 5, 10, 25} and we want to make change for 50 cents. We want to find the number of ways to do this.

To solve this problem using dynamic programming, we can use a bottom-up approach. We start by defining a table, where the rows represent the different denominations, and the columns represent the different amounts we want to make change for. We initialize the first column to 1, since there is only one way to make change for 0 cents. Then, for each subsequent denomination, we fill in the table based on the following recurrence relation:

```
table[i][j] = table[i-1][j] + table[i][j-d[i]]
```

where i is the current denomination, j is the current amount, d[i] is the value of the i-th denomination, and table[i][j] represents the number of ways to make change for j cents using the first i denominations.

Here is the code snippet in Go:

```go
func coinChange(denominations []int, amount int) int {
    table := make([][]int, len(denominations)+1)
    for i := range table {
        table[i] = make([]int, amount+1)
        table[i][0] = 1
    }
    
    for i := 1; i <= len(denominations); i++ {
        for j := 1; j <= amount; j++ {
            if j < denominations[i-1] {
                table[i][j] = table[i-1][j]
            } else {
                table[i][j] = table[i-1][j] + table[i][j-denominations[i-1]]
            }
        }
    }
    
    return table[len(denominations)][amount]
}
```
We first initialize the table with the base case of one way to make change for 0 cents. Then, we iterate over each denomination and amount, filling in the table using the recurrence relation. Finally, we return the value in the last cell of the table, which represents the number of ways to make change for the given amount using all the denominations.

##  16. Longest common substring
The Longest Common Substring (LCS) problem is a classic computer science problem that involves finding the longest substring that two or more strings share in common. A substring is a contiguous sequence of characters within a string.

For example, consider the two strings "abcdxyz" and "xyzabcd". The longest common substring between these two strings is "abcd" with a length of 4.

To solve the LCS problem using the bottom-up dynamic programming (DP) approach, we can follow these steps:

1.  Initialize a 2D table of size (m+1) x (n+1), where m and n are the lengths of the two strings, and initialize all values to 0.
2.  Iterate over the table from the first row and first column to the last row and last column.
3.  If the characters in the two strings at the current indices match, set the value in the table at the current indices to the value in the table at the previous indices (i-1, j-1) plus 1.
4.  If the characters in the two strings at the current indices do not match, set the value in the table at the current indices to 0.
5.  Keep track of the maximum value in the table and its corresponding indices.
6.  Extract the longest common substring by starting at the maximum value in the table and tracing back through the table until a 0 is encountered.

Here is a sample implementation of the bottom-up DP approach in Go:
```go
func findLCS(s1, s2 string) string {
    m := len(s1)
    n := len(s2)
    table := make([][]int, m+1)
    for i := range table {
        table[i] = make([]int, n+1)
    }
    maxLength := 0
    maxI := 0
    maxJ := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s1[i-1] == s2[j-1] {
                table[i][j] = table[i-1][j-1] + 1
                if table[i][j] > maxLength {
                    maxLength = table[i][j]
                    maxI = i
                    maxJ = j
                }
            } else {
                table[i][j] = 0
            }
        }
    }
    lcs := ""
    for maxLength > 0 {
        lcs = string(s1[maxI-1]) + lcs
        maxI--
        maxJ--
        maxLength--
    }
    return lcs
}
```
This implementation takes two input strings `s1` and `s2`, initializes the table, iterates over the table and fills it out, and then extracts and returns the LCS. Note that this implementation only returns one of the possible longest common substrings.

##  17. Maximum sum increasing subsequence
subsequence is \[4, 6, 8\] with a total sum of 18.

To solve this problem using bottom-up dynamic programming, you can use a 1-dimensional array to store the maximum sum increasing subsequence ending at each position of the input sequence. At each position, you can compute the maximum sum increasing subsequence by iterating over all the previous positions that have a smaller value and adding the maximum sum increasing subsequence ending at that position if the value at the current position is greater than the value at the previous position. You then select the maximum sum increasing subsequence ending at all the previous positions and add the value at the current position to obtain the maximum sum increasing subsequence ending at the current position.

Here are the steps in English:

1.  Initialize an array called dp with the same length as the input sequence and set all values to 0.
2.  Set the first element of dp to the value of the first element in the input sequence.
3.  Iterate over the input sequence starting from the second element: a. For each element, iterate over all the previous elements in the input sequence that have a smaller value. b. If the value at the current position is greater than the value at the previous position, compute the maximum sum increasing subsequence ending at the previous position and add the value at the current position. c. Select the maximum of all the computed values as the maximum sum increasing subsequence ending at the current position and store it in dp.
4.  Return the maximum value in dp as the solution to the problem.

Here is the code snippet in Go:
```go
func maxSumIncreasingSubsequence(arr []int) int {
    n := len(arr)
    dp := make([]int, n)

    // Initialize the first element of dp.
    dp[0] = arr[0]

    // Iterate over the input sequence starting from the second element.
    for i := 1; i < n; i++ {
        maxSum := 0

        // Iterate over all the previous elements in the input sequence that have a smaller value.
        for j := 0; j < i; j++ {
            // If the value at the current position is greater than the value at the previous position,
            // compute the maximum sum increasing subsequence ending at the previous position and add the value at the current position.
            if arr[i] > arr[j] {
                maxSum = max(maxSum, dp[j])
            }
        }

        // Select the maximum of all the computed values as the maximum sum increasing subsequence ending at the current position.
        dp[i] = maxSum + arr[i]
    }

    // Return the maximum value in dp as the solution to the problem.
    return maxSlice(dp)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSlice(arr []int) int {
    maxVal := arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] > maxVal {
            maxVal = arr[i]
        }
    }
    return maxVal
}

```
Here, the `max` function returns the maximum value of two integers, and the `maxSlice` function returns the maximum value in a slice of integers.

The time complexity of this algorithm is O(n^2) since we iterate over all the previous positions for each element in the input sequence. However, we can optimize this algorithm further by using binary search to find the position of the previous element with the largest value that is less than the current element. This can reduce the time complexity to O(n log n).

Here is the optimized code snippet in Go:
```go
func maxSumIncreasingSubsequence(arr []int) int {
    n := len(arr)
    dp := make([]int, n)

    // Initialize the first element of dp.
    dp[0] = arr[0]

    // Initialize the length of the longest increasing subsequence to 1.
    lisLength := 1
    lis := make([]int, n)
    lis[0] = 0

    // Iterate over the input sequence starting from the second element.
    for i := 1; i < n; i++ {
        // Use binary search to find the position of the previous element with the largest value that is less than the current element.
        j := sort.Search(lisLength, func(k int) bool {
            return arr[i] <= arr[lis[k]]
        })

        // Update the length of the longest increasing subsequence if necessary.
        if j == lisLength {
            lisLength++
        }

        // Update the longest increasing subsequence.
        lis[j] = i

        // Compute the maximum sum increasing subsequence ending at the current position.
        dp[i] = dp[lis[j-1]] + arr[i]
    }

    // Return the maximum value in dp as the solution to the problem.
    return maxSlice(dp)
}
```
In this optimized code, we use a separate array called `lis` to store the positions of the elements in the longest increasing subsequence ending at each position of the input sequence. We also use a variable called `lisLength` to keep track of the length of the longest increasing subsequence seen so far. Finally, we use binary search to find the position of the previous element with the largest value that is less than the current element, and update the longest increasing subsequence and the maximum sum increasing subsequence accordingly.

##  18. Painting fence problem
The Painting fence problem is a classic dynamic programming problem that involves painting a fence with two colors such that no more than two adjacent fence posts have the same color. The problem is to find the number of ways in which the fence can be painted.

For example, let's say we have a fence with n posts. We need to paint the fence using two colors, say red and blue, such that no more than two adjacent posts have the same color. For example, if n = 3, then the fence can be painted in the following ways:

* RBR
* RBB
* BRB
* BBB
* BRR

We need to find the number of such ways.

To solve this problem using bottom-up dynamic programming, we can use the following steps:

1.  Define the subproblem: Let's say f(i) represents the number of ways to paint i posts such that no more than two adjacent posts have the same color.
    
2.  Define the base case: We can initialize f(1) and f(2) to 2, since there are only two ways to paint one post and two posts.
    
3.  Define the recurrence relation: For i > 2, we can paint the i-th post in either the same color as the (i-1)th post, or a different color. If we paint the i-th post in the same color as the (i-1)th post, then we cannot paint the (i-1)th post in the same color as the (i-2)th post. So, the number of ways to paint i posts in this case is f(i-2). If we paint the i-th post in a different color than the (i-1)th post, then the number of ways to paint i posts in this case is f(i-1) multiplied by the number of available colors, which is 2.
    
    Therefore, the recurrence relation is: f(i) = f(i-1) + f(i-2).
    
4.  Compute the solution: We can compute the solution by iterating from i=3 to n and computing f(i) using the recurrence relation. The final answer is f(n).
    

Here's the Go code snippet for the above algorithm:

```go
func numWays(n int) int {
    if n == 1 || n == 2 {
        return 2
    }
    f1 := 2
    f2 := 2
    for i := 3; i <= n; i++ {
        fi := f1 + f2
        f2 = f1
        f1 = fi
    }
    return f1
}

```
In the above code, f1 and f2 represent f(i-1) and f(i-2) respectively, and fi represents f(i). We start with f1 = f2 = 2 as the base case, and compute fi using the recurrence relation. Finally, we return f1 as the answer.

##  19. Maximum sum subsequence without adjacent elements
The Maximum Sum Subsequence Without Adjacent Elements problem involves finding the maximum sum that can be obtained by selecting a subsequence of non-adjacent elements from an array of integers.

For example, given the array `[3, 4, 5, 10, 11, 12]`, the maximum sum subsequence without adjacent elements is `[3, 5, 11]`, with a sum of 19.

To solve this problem using bottom-up dynamic programming, we can define a state function dp[i] that represents the maximum sum of a subsequence ending at index i. We can then use the following recurrence relation to calculate dp[i]:
```
dp[i] = max(dp[i-1], dp[i-2]+arr[i])

```
where arr is the input array and i >= 2.

The base cases are `dp[0] = arr[0] and dp[1] = max(arr[0], arr[1]).`

The final solution is the maximum value in the dp array.

Here's the Go code snippet that implements the bottom-up dynamic programming approach:

```go
func maxSumSubseq(arr []int) int {
    n := len(arr)
    dp := make([]int, n)
    dp[0] = arr[0]
    dp[1] = max(arr[0], arr[1])
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+arr[i])
    }
    return dp[n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
We start by initializing the dp array with the base cases. Then, we iterate over the input array and calculate the maximum sum for each index using the recurrence relation. Finally, we return the maximum value in the dp array, which represents the maximum sum subsequence without adjacent elements.

##  20. 0/1 Knapsack problem
The 0/1 Knapsack problem is a well-known optimization problem in computer science, where we have a set of items with weights and values, and a knapsack that can carry a maximum weight capacity. The goal is to maximize the total value of items that can be placed into the knapsack without exceeding its capacity.

For example, let's say we have a knapsack with a maximum capacity of 10 units, and we have the following items with their respective weights and values:

| Item | Weight | Value |
| --- | --- | --- |
| A   | 2   | 10  |
| B   | 5   | 25  |
| C   | 4   | 15  |
| D   | 3   | 12  |

We need to choose the items that can be placed into the knapsack in such a way that we maximize the total value while staying within the capacity limit of the knapsack.

To solve this problem using bottom-up dynamic programming, we can create a 2D array where the rows represent the items and the columns represent the weight capacity of the knapsack. The value at each cell (i, j) in the array represents the maximum value we can achieve using the first i items and a knapsack with a capacity of j.

We can use the following steps to solve the problem:

1.  Create a 2D array dp of size (n+1) x (W+1), where n is the number of items and W is the maximum weight capacity of the knapsack.
    
2.  Initialize the first row and column of the array to 0, as we can achieve zero value with zero capacity and with no items.
    
3.  For each item i from 1 to n, and for each capacity j from 1 to W, calculate the maximum value we can achieve using the first i-1 items and a knapsack with a capacity of j. If the weight of the i-th item is less than or equal to j, we can either include or exclude the i-th item from the knapsack. We choose the maximum value between the two cases.
    
4.  The final answer will be stored in the last cell of the dp array, dp\[n\]\[W\].
    

Here's the code snippet in Go:
```go
func knapsack(W int, wt []int, val []int, n int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, W+1)
    }

    for i := 0; i <= n; i++ {
        for j := 0; j <= W; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
            } else if wt[i-1] <= j {
                dp[i][j] = max(val[i-1]+dp[i-1][j-wt[i-1]], dp[i-1][j])
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }

    return dp[n][W]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
To use this function, you can call it with the maximum weight capacity of the knapsack, the weights and values of the items, and the number of items:

```go
W := 10
wt := []int{2, 5, 4, 3}
val := []int{10, 25, 15, 12}
n := len(val)

fmt.Println(knapsack(W, wt, val, n)) // Output: 37
```
In this example, the maximum value we can achieve is 37 by including items A, C, and D in the knapsack.

##  21. Egg dropping problem

The egg dropping problem is a classic problem in computer science and mathematics. It involves finding the minimum number of attempts required to determine the highest floor from which an egg can be dropped without breaking. The problem is commonly encountered in the field of building and construction, where engineers need to determine the maximum height from which they can drop a material without causing damage.

For example, suppose you have two eggs and a 10-story building. You need to find the highest floor from which you can drop an egg without breaking it. What is the minimum number of attempts required to find the answer?

To solve the egg dropping problem using bottom-up dynamic programming, you can follow these steps:

1.  Create a 2D array to store the results of subproblems. The array should have dimensions n+1 x k+1, where n is the number of eggs and k is the number of floors.
2.  Initialize the array with base cases. For example, if there is only one egg, then the minimum number of attempts required is equal to the number of floors. If there is only one floor, then the minimum number of attempts required is 1.
3.  Use a nested loop to fill in the array with the results of subproblems. Start with two eggs and two floors, and work your way up to n eggs and k floors.
4.  For each subproblem, try dropping the egg from each floor from 1 to j, where j is the current floor being tested. For each drop, calculate the number of attempts required to find the solution, taking into account both the case where the egg breaks and the case where it does not break. Take the maximum of these values, since you want to minimize the worst-case scenario.
5.  Store the result of the subproblem in the array.
6.  Once the array is filled in, return the value in the last cell, which represents the minimum number of attempts required to find the solution.

Here is the implementation of the above steps in Go:
```go
func eggDrop(n, k int) int {
    // Initialize the 2D array
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, k+1)
    }
    
    // Initialize base cases
    for j := 1; j <= k; j++ {
        dp[1][j] = j
    }
    for i := 1; i <= n; i++ {
        dp[i][1] = 1
    }
    
    // Fill in the array with results of subproblems
    for i := 2; i <= n; i++ {
        for j := 2; j <= k; j++ {
            dp[i][j] = math.MaxInt32
            for x := 1; x <= j; x++ {
                res := 1 + max(dp[i-1][x-1], dp[i][j-x])
                if res < dp[i][j] {
                    dp[i][j] = res
                }
            }
        }
    }
    
    // Return the final result
    return dp[n][k]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
In the above implementation, we use `math.MaxInt32` as the initial value for the results of subproblems that have not yet been calculated. We also define a `max` function to return the maximum of two integers. Finally, we return the value in the last cell of the 2D array as the final result.
##  22. Partition problem

The Partition problem is a classic dynamic programming problem that involves partitioning a set of integers into two subsets such that the difference between the sum of the elements in each subset is minimized.

For example, given the set `[1, 6, 11, 5]`, we can partition it into two subsets `[1, 5]`and `[6, 11] `such that the difference between their sums is minimized, which is 1 in this case.

To solve this problem using bottom-up dynamic programming, we can define a 2D table dp of size `(n+1 x sum+1)`, where n is the number of elements in the input set and sum is the sum of all elements in the input set. `dp[i][j] `represents whether there exists a subset of the first i elements of the input set such that its sum is j. If such a subset exists, then `dp[i][j]` is true, otherwise it is false.

To fill up the dp table, we can initialize `dp[0][0]` as true and `dp[i][0]` as true for all i, since the empty subset has a sum of 0. We can then iterate through the input set and for each element `nums[i]`, we can iterate through all possible sums j from sum/2 to 0 in descending order. If `dp[i-1][j]` is true, then we can also make `dp[i][j+nums[i]]` as true, since we can add `nums[i]` to a subset that has a sum of j to get a new subset that has a sum of `j+nums[i]`. If `dp[n][sum/2] `is true, then the minimum difference between the two subsets is` sum-2*(sum/2)`, which is simply the difference between the sum of all elements in the input set and the sum of one of the two subsets.

Here's the code in Go:

```go
func canPartition(nums []int) bool {
    n := len(nums)
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum%2 == 1 {
        return false
    }
    sum /= 2
    dp := make([][]bool, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]bool, sum+1)
        dp[i][0] = true
    }
    for i := 1; i <= n; i++ {
        for j := sum; j >= nums[i-1]; j-- {
            dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
        }
    }
    return dp[n][sum]
}

```
In this code, the function canPartition takes an input set nums as a parameter and returns a boolean indicating whether the set can be partitioned into two subsets with equal sums. The sum of all elements in the input set must be even for such a partition to exist. The dp table is filled up using a bottom-up approach, and the function returns dp\[n\]\[sum\], which indicates whether there exists a subset of the first n elements of the input set with a sum of sum.

##  23. Longest zigzag subsequence
The Longest Zigzag Subsequence problem is a classic dynamic programming problem that involves finding the longest subsequence in an array or sequence such that the elements alternate between being smaller and larger than their adjacent elements.

For example, given the sequence `[1, 7, 4, 9, 2, 5]`, the longest zigzag subsequence is `[1, 7, 4, 9, 2, 5]`, since the elements alternate between being smaller and larger than their adjacent elements.

To solve this problem using bottom-up dynamic programming, we can define a 2D table dp of size (n x 2), where n is the length of the input sequence. `dp[i][0]` stores the length of the longest zigzag subsequence that ends at index i, with the last element being smaller than the previous element, and `dp[i][1] `stores the length of the longest zigzag subsequence that ends at index i, with the last element being larger than the previous element.

To fill up the dp table, we can iterate through the input sequence and for each index i, check all previous indices j less than i. If `nums[i] `is greater than `nums[j]`, then we can update `dp[i][1]` as `max(dp[i][1], dp[j][0] + 1)`, since we can add `nums[i]` to the end of a zigzag subsequence that ends at index j and has the last element smaller than nums\[j\] to get a new zigzag subsequence that ends at index i and has the last element greater than nums\[j\]. Similarly, if `nums[i]` is less than `nums[j]`, then we can update `dp[i][0]` as `max(dp[i][0], dp[j][1] + 1)`, since we can add `nums[i]` to the end of a zigzag subsequence that ends at index j and has the last element greater than `nums[j]` to get a new zigzag subsequence that ends at index i and has the last element smaller than `nums[j]`.

The final answer to the problem is the maximum value in the last row of the dp table, which represents the length of the longest zigzag subsequence that ends at the last index of the input sequence.

Here's the code in Go:


```go
func longestZigzagSubsequence(nums []int) int {
    n := len(nums)
    dp := make([][2]int, n)
    for i := 0; i < n; i++ {
        dp[i][0] = 1
        dp[i][1] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i][1] = max(dp[i][1], dp[j][0]+1)
            } else if nums[i] < nums[j] {
                dp[i][0] = max(dp[i][0], dp[j][1]+1)
            }
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        ans = max(ans, max(dp[i][0], dp[i][1]))
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
In this code, the function longestZigzagSubsequence takes an input sequence nums as a parameter and returns the length of the longest zigzag subsequence in the sequence. The max function is a helper function to return the maximum of two integers.
##  24. String matching problems

The string matching problem is a classic computer science problem that involves finding all occurrences of a given pattern in a larger string. The problem arises in many applications such as text editors, DNA sequence analysis, and data compression.

For example, given a text string "ababababa" and a pattern string "aba", the goal is to find all occurrences of the pattern in the text. In this case, the pattern appears three times at positions 0, 2, and 4.

To solve the string matching problem using bottom-up dynamic programming (DP), we can use a 2D array to keep track of whether a given prefix of the text string matches a prefix of the pattern string. The DP table has dimensions (n+1) x (m+1), where n is the length of the text string and m is the length of the pattern string.

The algorithm works as follows:

1.  Initialize the DP table. Set dp\[0\]\[0\] to true and dp\[0\]\[j\] and dp\[i\]\[0\] to false for all i > 0 and j > 0.
2.  For each i from 1 to n and each j from 1 to m, set dp\[i\]\[j\] to true if: a. dp\[i-1\]\[j-1\] is true and text\[i-1\] matches pattern\[j-1\], or b. dp\[i\]\[j-1\] is true and pattern\[j-1\] is '_' (wildcard), or c. dp\[i-1\]\[j\] is true and pattern\[j-1\] is '_' (wildcard).
3.  After filling the DP table, iterate over the last row (dp\[n\]\[:\]) and return the indices where dp\[n\]\[j\] is true.

Here's the code in Go:

```go
func stringMatch(text string, pattern string) []int {
    n := len(text)
    m := len(pattern)
    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for j := 1; j <= m; j++ {
        if pattern[j-1] == '*' {
            dp[0][j] = dp[0][j-1]
        }
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if text[i-1] == pattern[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if pattern[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            }
        }
    }
    matches := make([]int, 0)
    for j := 1; j <= m; j++ {
        if dp[n][j] {
            matches = append(matches, j-1)
        }
    }
    return matches
}
```
In the code above, the function `stringMatch` takes two string arguments `text` and `pattern` and returns a slice of integers representing the indices of the pattern matches in the text. The function initializes the DP table, fills it using the bottom-up approach, and then returns the matches by iterating over the last row of the DP table.

##  25. Maximum sum subarray with at most k elements

The Maximum Sum Subarray with at most k elements problem is a classic dynamic programming problem. Given an array of integers, the task is to find the subarray with at most k elements that has the maximum sum.

For example, consider the array `arr = [2, 1, 4, 3, 5]` and `k = 3`. The subarray with at most 3 elements and maximum sum is `[4, 3, 5]` with a sum of 12.

To solve this problem using bottom-up dynamic programming, we will use a 2D array `dp` of size `(n+1)x(k+1)`, where `n` is the length of the input array. The value `dp[i][j]` represents the maximum sum subarray with at most `j` elements ending at index `i-1` of the input array. We will also use an auxiliary variable `ans` to keep track of the maximum sum subarray we have seen so far.

The steps to solve the problem are as follows:

1.  Initialize `dp[i][1]` to the value of `arr[i-1]` for all `i` from 1 to `n`. Initialize `dp[0][j]` to 0 for all `j` from 1 to `k+1`.
2.  For all `i` from 1 to `n`, and all `j` from 2 to `k+1`, compute `dp[i][j]` as follows:
    * Initialize `sum` to 0.
    * For all `p` from `i-1` to 0 (in reverse order), add `arr[p]` to `sum`.
    * Set `dp[i][j]` to the maximum of `dp[p][j-1] + sum` and `dp[i-1][j]`.
3.  Update `ans` to the maximum value in the last row of `dp`.
4.  Return `ans`.

Here is the Go code to implement the above algorithm:

```go
func maxSubArrayWithAtMostKElements(arr []int, k int) int {
    n := len(arr)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
    }

    ans := 0
    for i := 1; i <= n; i++ {
        dp[i][1] = arr[i-1]
        for j := 2; j <= k+1; j++ {
            sum := 0
            dp[i][j] = dp[i-1][j]
            for p := i-1; p >= 0 && i-p <= j; p-- {
                sum += arr[p]
                dp[i][j] = max(dp[i][j], dp[p][j-1]+sum)
            }
        }
        ans = max(ans, dp[i][k+1])
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
We initialize the `dp` array with size `(n+1)x(k+1)` and set the initial values of `dp[i][1]` to `arr[i-1]` and `dp[0][j]` to 0. Then, we iterate over all indices `i` from 1 to `n` and all indices `j` from 2 to `k+1` to compute `dp[i][j]`. We use a nested loop to compute the sum of the subarray ending at index `i-1` and containing at most `j` elements, by iterating over all indices `p` from `i-1` to 0 in reverse order and adding `arr[p]` to `sum`. We then set `dp[i][j]` to the maximum of `dp[p][j-1] + sum` and `dp[i-1][j]`, where `dp[p][j-1] + sum` represents the sum of the subarray ending at index `p-1` and containing `j-1` elements, plus the sum of the subarray from index `p` to index `i-1`. We update `ans` to the maximum value in the last row of `dp`.

Finally, we return `ans`, which represents the maximum sum subarray with at most `k` elements in the input array.

Note that this algorithm has a time complexity of O(n^2 * k) and a space complexity of O(n * k), which can be optimized using sliding window techniques to reduce the time complexity to O(n * k) and the space complexity to O(n). However, the above implementation provides a clear and simple illustration of the bottom-up dynamic programming approach to solve this problem.

##  26. Longest alternating subsequence
The Longest Alternating Subsequence (LAS) problem is a classic dynamic programming problem. Given an array of integers, the task is to find the length of the longest subsequence where the adjacent elements differ in sign.

For example, consider the array` [1, -5, 2, -3, 7, -8]`. The longest alternating subsequence in this array is `[1, -5, 2, -3, 7]`, which has a length of 5.

To solve this problem using bottom-up DP, we can define a 2D table where`dp[i][j]` represents the length of the longest alternating subsequence ending at index i with the last element having sign j, where j = 0 for negative and j = 1 for positive. We can initialize the table with 1 since every element is an alternating subsequence of length 1.

Then, for each pair of indices (i, j) where i > j, we can update the table as follows:

* If the element at index i has a different sign than the element at index j, we can update `dp[i][j]` to be `dp[j][1-j] + 1`, since we can add the element at index i to the alternating subsequence ending at index j.
* Otherwise, we can leave `dp[i][j]` as it is, since we cannot add the element at index i to the alternating subsequence ending at index j.

Finally, we can iterate over the last row of the table and return the maximum value as the length of the longest alternating subsequence.

Here's the code snippets in Go:

```go
func longestAlternatingSubsequence(arr []int) int {
    n := len(arr)
    dp := make([][2]int, n)
    for i := 0; i < n; i++ {
        dp[i][0] = 1
        dp[i][1] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if arr[i]*arr[j] < 0 {
                dp[i][arr[i]>0] = max(dp[i][arr[i]>0], dp[j][1-arr[i]>0]+1)
            } else {
                dp[i][arr[i]>0] = max(dp[i][arr[i]>0], dp[j][arr[j]>0])
            }
        }
    }
    res := 1
    for i := 0; i < n; i++ {
        res = max(res, max(dp[i][0], dp[i][1]))
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
In this implementation, we use the `max` function to compute the maximum value between two integers. We also use the `arr[i]>0` and `arr[j]>0` expressions to check the signs of the elements at indices `i` and `j`.

##  27. Tower of Hanoi problem

The Tower of Hanoi is a classic puzzle game that involves moving a stack of disks of different sizes from one peg to another while following certain rules. The goal is to move the entire stack to another peg, without placing a larger disk on top of a smaller one at any time.

Here's an example of the problem: suppose you have three pegs labeled A, B, and C, and you start with a stack of three disks of different sizes on peg A, with the largest disk on the bottom and the smallest on top. The goal is to move the entire stack to peg C, using peg B as an intermediary, while following the rules of the game.

The steps to solve the problem using bottom-up dynamic programming are:

1.  Define a function to calculate the optimal solution for moving a stack of n disks from peg A to peg C, using peg B as an intermediary.
    
2.  Create a table to store the optimal solutions for all subproblems of size 1 to n, where each entry in the table represents the minimum number of moves required to move a stack of that size.
    
3.  Fill in the base cases of the table, which are the solutions for moving a stack of size 1, 2, and 3 directly from A to C without using B.
    
4.  Use the bottom-up approach to fill in the remaining entries of the table by using the recursive formula:
    
    * Number of moves required to move a stack of n disks from A to C using B as an intermediary is equal to:
    * 2 times the number of moves required to move a stack of n-1 disks from A to B using C as an intermediary
    * plus 1 (to move the largest disk from A to C)
    * plus the number of moves required to move a stack of n-1 disks from B to C using A as an intermediary.
5.  Return the final solution, which is the entry in the table corresponding to moving a stack of size n from A to C.
    

Here's an implementation of the solution in Go:
```go
func towerOfHanoi(n int) int {
    // Create table to store optimal solutions
    table := make([]int, n+1)

    // Fill in base cases
    table[1] = 1
    table[2] = 3
    table[3] = 7

    // Fill in remaining entries of table using recursive formula
    for i := 4; i <= n; i++ {
        table[i] = 2*table[i-1] + 1
    }

    // Return final solution
    return table[n]
}

```
This function takes an integer `n` representing the number of disks in the initial stack, and returns an integer representing the minimum number of moves required to move the entire stack to the destination peg. The implementation uses a one-dimensional array to store the optimal solutions, and fills in the table in a bottom-up fashion using a loop. The time complexity of this solution is O(n), since we need to fill in the table for all subproblems of size 1 to n.
##  28. Longest common prefix
The Longest Common Prefix (LCP) problem is a classic problem in computer science where given a set of strings, you need to find the longest prefix common to all the strings.

For example, consider the following strings:

```
["flower", "flow", "flight"]
```
The longest common prefix for these strings is `"fl"`.

One way to solve this problem is by using a bottom-up dynamic programming (DP) approach. The basic idea is to build a 2D table where each cell `dp[i][j]` represents the length of the longest common prefix between the i-th and j-th strings.

The steps to solve the LCP problem using bottom-up DP are as follows:

1.  Initialize the DP table with `0` for all the cells `dp[i][j]` where `i != j`.
    
2.  For each cell `dp[i][j]`, if the i-th and j-th characters match, set `dp[i][j]` to `dp[i-1][j-1] + 1`. Otherwise, set `dp[i][j]` to `0`.
    
3.  Find the maximum value in the last row of the DP table. This value represents the length of the longest common prefix for all the strings.
    

Here is the Go code for the bottom-up DP approach to solve the LCP problem:

```go
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n == 0 {
        return ""
    }

    m := len(strs[0])
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            k := 0
            for ; k < m && k < len(strs[i]) && k < len(strs[j]); k++ {
                if strs[i][k] != strs[j][k] {
                    break
                }
            }
            dp[i][j] = k
            dp[j][i] = k
        }
    }

    ans := strs[0]
    for i := 1; i < n; i++ {
        for j := 0; j < m; j++ {
            if dp[i][i-1] < j+1 {
                break
            }
            if strs[i][j] != ans[j] {
                ans = ans[:j]
                break
            }
        }
    }

    return ans
}
```
In the code, we first initialize the DP table with 0 for all the cells `dp[i][j] `where` i != j`. Then, we iterate over all the cells in the DP table and fill in their values using the approach described above.

Finally, we iterate over the strings and the characters in the first string to find the longest common prefix. We do this by checking the value in the DP table for the i-th string and the `(i-1)-th` string, and then comparing the characters in the i-th string with the corresponding characters in the longest common prefix found so far. If there is a mismatch, we update the longest common prefix accordingly.
##  29. Shortest common supersequence

The Shortest Common Supersequence (SCS) problem is a classic computer science problem that requires finding the shortest possible string that contains all the given strings as subsequences.

For example, given two strings "ABCBDAB" and "BDCABA", the shortest common supersequence is "ABDCBDAB" since it contains both the given strings as subsequences and has the shortest possible length.

To solve the SCS problem using bottom-up dynamic programming (DP), we can follow the following steps:

1.  Define the DP table: We can define a 2D table of size `(n+1) x (m+1)`, where n and m are the lengths of the two given strings. The value at `dp[i][j]` will represent the length of the shortest common supersequence of the substrings` s1[0:i] and s2[0:j]`.
    
2.  Initialize the DP table: We can initialize the first row and column of the DP table to their respective indices, since the shortest common supersequence of a string with an empty string is the string itself.
    
3.  Fill the DP table: We can fill the DP table row by row, from left to right. For each cell `dp[i][j]`, we have two cases:
    
    a. If the last characters of` s1[0:i] and s2[0:j]` are equal, then the shortest common supersequence of these substrings will be the same as that of the substrings `s1[0:i-1] and s2[0:j-1]`, plus the common character. Therefore, `dp[i][j] = dp[i-1][j-1] + 1`.
    
    b. If the last characters of s1\[0:i\] and s2\[0:j\] are not equal, then we have two options: we can either include the last character of s1 or the last character of s2 in the supersequence. We choose the option that gives us the shorter supersequence. Therefore, `dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1`.
    
4.  Return the result: The length of the shortest common supersequence of the two given strings will be the value at the bottom-right corner of the DP table, i.e., `dp[n][m]`.
    

Here is the code snippet in Go for solving the SCS problem using bottom-up DP:
```go
func shortestCommonSupersequence(s1 string, s2 string) int {
    n, m := len(s1), len(s2)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }

    // initialize the first row and column
    for i := 0; i <= n; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j
    }

    // fill the DP table
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }

    // return the result
    return dp[n][m]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

```

##  30. Maximum size subset with given sum.
The Maximum size subset with given sum problem is a classic dynamic programming problem where we are given an array of integers and a target sum. The goal is to find the maximum size of a subset of the array such that the sum of elements in the subset is equal to the target sum.

For example, suppose we have an array `arr = [2, 3, 5, 8]` and a target sum of `10`. The maximum size subset with the given sum is `[2, 8]`, which has a size of 2.

To solve this problem using bottom-up dynamic programming, we can define a 2D table `dp` of size `(n+1) x (target+1)`, where `n` is the size of the array and `target` is the given target sum. The `dp[i][j]` entry in the table represents the maximum size subset of the first `i` elements of the array that sums up to `j`. The final solution will be `dp[n][target]`.

The recurrence relation for this problem is as follows:
```
if arr[i-1] > j:
    dp[i][j] = dp[i-1][j]
else:
    dp[i][j] = max(dp[i-1][j], 1 + dp[i-1][j-arr[i-1]])
```
Here, `arr[i-1]` represents the i-th element of the array since the table is 1-indexed.

The base case for the table is when `j=0`, which means we can always have an empty subset that sums up to 0. Hence, `dp[i][0] = 0` for all `i`.

To solve the problem, we need to fill the `dp` table in a bottom-up manner, starting from the base cases and filling the table row by row until we reach the final solution at `dp[n][target]`.

Here's the code snippet in Go:
```go
func maxSubsetSize(arr []int, target int) int {
    n := len(arr)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, target+1)
    }

    // base case
    for i := 0; i <= n; i++ {
        dp[i][0] = 0
    }

    // fill the table
    for i := 1; i <= n; i++ {
        for j := 1; j <= target; j++ {
            if arr[i-1] > j {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = max(dp[i-1][j], 1+dp[i-1][j-arr[i-1]])
            }
        }
    }

    return dp[n][target]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
To use the function, simply call `maxSubsetSize(arr, target)` with the array and target sum as inputs, and it will return the maximum size subset with the given sum.
