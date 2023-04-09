
# Dynamic Programming

Dynamic programming is a problem-solving technique that breaks down complex problems into smaller, more manageable subproblems and stores the solutions to these subproblems in memory for future use. By solving these subproblems only once and reusing their solutions, dynamic programming can greatly reduce the time and computational resources needed to solve the overall problem. It's like solving a puzzle piece by piece and reusing the solutions to the smaller pieces to solve the bigger puzzle faster.


To break down a complex problem into simpler subproblems, you can follow these steps:

1.  Identify the optimal substructure: This means that the solution to the larger problem can be constructed from the solutions of its subproblems.
    
2.  Define the recursive relationship: This means defining a function that can be called recursively to solve the subproblems.
    
3.  Store the solutions to subproblems: By storing the solutions, you can avoid having to solve the same subproblem multiple times, which can greatly improve efficiency.
    

Here are five examples of how dynamic programming can be used to solve complex problems:

1.  Fibonacci Sequence: The Fibonacci sequence is a series of numbers in which each number is the sum of the two preceding ones. To find the nth number in the sequence, you can use dynamic programming to store the solutions to smaller subproblems and avoid redundant calculations.
    
2.  Shortest Path Problem: Given a graph with weighted edges, the shortest path problem involves finding the shortest path between two vertices. Dynamic programming can be used to solve this problem by breaking it down into smaller subproblems and storing the solutions.
    
3.  Longest Common Subsequence: Given two strings, the longest common subsequence problem involves finding the longest subsequence that is common to both strings. Dynamic programming can be used to solve this problem by breaking it down into smaller subproblems and storing the solutions.
    
4.  Knapsack Problem: The knapsack problem involves finding the most valuable combination of items that can be placed in a knapsack of limited size. Dynamic programming can be used to solve this problem by breaking it down into smaller subproblems and storing the solutions.
    
5.  Edit Distance: Edit distance is a measure of similarity between two strings, calculated as the minimum number of edits (insertions, deletions, or substitutions) required to transform one string into the other. Dynamic programming can be used to solve this problem by breaking it down into smaller subproblems and storing the solutions.

The following are the steps to identify the optimal substructure of a problem:

1.  Identify the larger problem: First, you need to understand what the larger problem is that you are trying to solve.
    
2.  Break down the problem into smaller subproblems: Once you have identified the larger problem, try to break it down into smaller subproblems that are easier to solve.
    
3.  Look for patterns: Examine the smaller subproblems to see if there are any patterns or repeated subproblems.
    
4.  Identify the optimal substructure: If there is a pattern, it is likely that the problem has an optimal substructure. This means that the solution to the larger problem can be constructed from the solutions of its subproblems.
    
5.  Define the recursive relationship: Once you have identified the optimal substructure, you need to define a recursive function that can be called repeatedly to solve the subproblems.

Here are some examples of problems with optimal substructure:

1.  Fibonacci sequence: The Fibonacci sequence is a series of numbers in which each number is the sum of the two preceding ones. The optimal substructure of this problem is that the nth number in the sequence can be constructed from the solutions of its two preceding subproblems (n-1 and n-2).
    
2.  Longest common subsequence: Given two strings, the longest common subsequence problem involves finding the longest subsequence that is common to both strings. The optimal substructure of this problem is that the length of the longest common subsequence can be constructed from the solutions of its smaller subproblems.
    
3.  Shortest path problem: Given a graph with weighted edges, the shortest path problem involves finding the shortest path between two vertices. The optimal substructure of this problem is that the shortest path between two vertices can be constructed from the solutions of its smaller subproblems.
    
4.  Knapsack problem: Given a set of items with weight and value, and a knapsack with a capacity, the knapsack problem involves finding the maximum value that can be obtained by filling the knapsack with items. The optimal substructure of this problem is that the maximum value that can be obtained for a given capacity can be constructed from the solutions of its smaller subproblems.
    
5.  Edit distance problem: Given two strings, the edit distance problem involves finding the minimum number of operations (insertion, deletion, or substitution) required to transform one string into another. The optimal substructure of this problem is that the edit distance between two strings can be constructed from the solutions of its smaller subproblems.

### The Fibonacci series
The Fibonacci series is a sequence of numbers where each number is the sum of the two preceding ones, starting from 0 and 1. The first few numbers in the series are:

`0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...` 

The optimal substructure of the Fibonacci series can be illustrated using a tree diagram.

At each level of the tree, we have two subproblems that are the solutions to the two preceding numbers in the sequence. We start with the base case of f(0) = 0 and f(1) = 1, and build up the tree by computing the solutions to each subproblem until we reach the desired value of f(n).

Here is an example of the optimal substructure of the Fibonacci series for n=5:

```text
             f(5)
           /     \
        f(4)     f(3)
        /   \    /   \
     f(3)  f(2) f(2) f(1)
     /  \
  f(2) f(1)

```


To compute f(5), we need to first compute the solutions to the subproblems f(4) and f(3). We can see from the diagram that f(4) is the sum of the solutions to the subproblems f(3) and f(2), and f(3) is the sum of the solutions to the subproblems f(2) and f(1).

We can continue to compute the solutions to the subproblems until we reach the base cases of f(0) and f(1), at which point we can compute the solution to f(2), then f(3), and so on until we reach the desired value of f(n).

This tree diagram illustrates the optimal substructure of the Fibonacci series, which is the property that the solution to a larger problem can be constructed from the solutions to its smaller subproblems.

### The Longest Common Subsequence

The longest common subsequence (LCS) problem involves finding the longest subsequence that is common to two given strings. The optimal substructure of this problem can be illustrated using a table diagram.

Consider the two strings "ABCDGH" and "AEDFHR". We can construct a table where the rows represent the characters in the first string and the columns represent the characters in the second string. We fill in the table by comparing each pair of characters and checking if they match. If they do, we increment the count of the LCS by 1 and move diagonally to the next cell. Otherwise, we take the maximum of the counts in the cell to the left or the cell above.

Here is an example of the optimal substructure of the LCS problem for the strings "ABCDGH" and "AEDFHR":

```
        A   B   C   D   G   H
    --------------------------
    | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
  A | 1 | 1 | 1 | 1 | 1 | 1 | 1 |
  E | 1 | 1 | 1 | 1 | 1 | 1 | 1 |
  D | 1 | 1 | 1 | 2 | 2 | 2 | 2 |
  F | 1 | 1 | 1 | 2 | 2 | 2 | 2 |
  H | 1 | 1 | 1 | 2 | 2 | 3 | 3 |
  R | 1 | 1 | 1 | 2 | 2 | 3 | 3 |

```
We start by initializing the first row and column to 0. Then, for each cell (i,j) in the table, we compare the characters at indices i-1 and j-1 in the two strings. If they match, we add 1 to the count of the LCS and move diagonally to the next cell. Otherwise, we take the maximum of the counts in the cell to the left or the cell above.

The optimal substructure of the LCS problem can be seen in the table, where the count of the LCS for the entire string pair is the value in the bottom-right corner of the table. The solution to the LCS problem for a pair of strings of length n and m can be constructed from the solutions to its smaller subproblems for pairs of strings of length n-1 and m-1.

This table diagram illustrates the optimal substructure of the LCS problem, which is the property that the solution to a larger problem can be constructed from the solutions to its smaller subproblems.

### Shortest path problem

The shortest path problem involves finding the shortest path between two nodes in a weighted graph. The optimal substructure of this problem can be illustrated using a graph diagram.

Consider the following directed graph with weights on the edges:

```
     5      2
 A -----> B -----> C
 ^       / \       |
 |      /   \      |
 4    1      3    4
 |  /         \   |
 | /           \  |
 V/             \ V
 D <----- E <----- F
    2        6
```
Suppose we want to find the shortest path between nodes A and C. We can use Dijkstra's algorithm to compute the shortest paths from node A to every other node in the graph. We start by initializing the distances from node A to all other nodes as infinity, except for the distance from node A to itself, which is 0. We then visit each node in the graph and update the distances to its neighboring nodes if a shorter path is found.

Here is an example of the optimal substructure of the shortest path problem for the graph above, where we compute the shortest paths from node A to all other nodes:

```
Initial state:
   A: 0
   B: inf
   C: inf
   D: inf
   E: inf
   F: inf

First iteration:
   A: 0
   B: 5
   C: inf
   D: 4
   E: inf
   F: inf

Second iteration:
   A: 0
   B: 5
   C: 7
   D: 4
   E: 9
   F: inf

Third iteration:
   A: 0
   B: 5
   C: 7
   D: 4
   E: 8
   F: 13

Final distances:
   A: 0
   B: 5
   C: 7
   D: 4
   E: 8
   F: 13

```
### Knapsack problem
The knapsack problem involves finding the maximum value that can be obtained by filling a knapsack of capacity `W` with a set of `n` items, each with a weight `w` and a value `v`. The optimal substructure of this problem can be illustrated using a table diagram.

Consider the following example of a knapsack problem with a capacity of `W=5` and a set of items with weights and values as shown:

```
      item 1   item 2   item 3   item 4   item 5
weight   2        1        3        2        1
value    12       10       20       15       5
```
We can use dynamic programming to solve this problem by building a table that stores the maximum value that can be obtained for each possible capacity and subset of items. We start with a base case of no items and a knapsack of zero capacity, which has a maximum value of zero. We then fill in the table for each additional item, considering two options: either include the item in the knapsack, or exclude it. We choose the option that yields the maximum value.

Here is an example of the optimal substructure of the knapsack problem for the example above, where we compute the maximum value that can be obtained for each possible capacity and subset of items:

```
         0   1   2   3   4   5
item 0   0   0   0   0   0   0
item 1   0   0   12  12  12  12
item 2   0   10  12  22  22  22
item 3   0   10  15  25  27  37
item 4   0   10  15  30  32  42
item 5   0   5   15  30  35  45
```
In the table above, each cell `(i,j)` represents the maximum value that can be obtained using the first `i` items and a knapsack of capacity `j`. The base case is represented by the first row and column with zeros. The other cells are filled in by considering two options: either include the `i`th item in the knapsack, which has weight `w_i` and value `v_i`, and add its value to the maximum value that can be obtained using the remaining capacity `j - w_i` and the first `i-1` items; or exclude the `i`th item, and use the maximum value that can be obtained using the same capacity and the first `i-1` items.

We can see from the example that the optimal substructure of the knapsack problem is that the maximum value that can be obtained using a set of items and a knapsack of capacity `W` can be constructed from the maximum values that can be obtained using the same capacity and a subset of the items. We can start with the base case of no items and a knapsack of zero capacity, and build up the maximum values for larger subsets of items and larger capacities by considering the two options of including or excluding the next item. The final solution to the problem is the maximum value that can be obtained using all `n` items and the full capacity `W`.

This table diagram illustrates the optimal substructure of the knapsack problem, which is the property that the solution to a larger problem can be constructed from the solutions to its smaller subproblems.

### Edit distance problem

The edit distance problem involves finding the minimum number of operations required to transform one string into another. The possible operations are insertion, deletion, and substitution of a single character. The optimal substructure of this problem can be illustrated using a table diagram.

Consider the following example of two strings, `s1="kitten"` and `s2="sitting"`. We can use dynamic programming to solve this problem by building a table that stores the minimum number of operations required to transform a prefix of `s1` into a prefix of `s2`. We start with a base case of empty strings, which requires zero operations. We then fill in the table for each additional character of `s1` and `s2`, considering three options: insert the next character of `s2` to match the current prefix of `s1`, delete the current character of `s1` to match the prefix of `s2`, or substitute the current character of `s1` with the corresponding character of `s2`.

Here is an example of the optimal substructure of the edit distance problem for the example above, where we compute the minimum number of operations required to transform each prefix of `s1` into a prefix of `s2`:

```
       ''  s   i   t   t   i   n   g
   ------------------------------------
'' |  0   1   2   3   4   5   6   7
k  |  1   1   2   3   4   5   6   7
i  |  2   2   1   2   3   4   5   6
t  |  3   3   2   1   2   3   4   5
t  |  4   4   3   2   1   2   3   4
e  |  5   5   4   3   2   2   3   4
n  |  6   6   5   4   3   3   2   3

```

In the table above, each cell `(i,j)` represents the minimum number of operations required to transform the first `i` characters of `s1` into the first `j` characters of `s2`. The base case is represented by the first row and column with increasing values. The other cells are filled in by considering three options: insert the next character of `s2` to match the current prefix of `s1`, which adds one operation to the minimum required for the previous prefix of `s1` and `s2`; delete the current character of `s1` to match the prefix of `s2`, which adds one operation to the minimum required for the current prefix of `s1` and the previous prefix of `s2`; or substitute the current character of `s1` with the corresponding character of `s2`, which adds one operation to the minimum required for the previous prefix of `s1` and `s2` if the characters are different, and zero otherwise.

We can see from the example that the optimal substructure of the edit distance problem is that the minimum number of operations required to transform a larger prefix of `s1` into a prefix of `s2` can be constructed from the minimum numbers of operations required to transform smaller prefixes of `s1` and `s2`. We can start with the base case of empty strings, and build up the minimum numbers of operations for larger prefixes of `s` 

## Defining recursive relationships in DP

Defining the recursive relationship in DP involves expressing the solution of a larger problem in terms of the solutions of smaller subproblems. This can often be done using a recursive function that calls itself on smaller subproblems until the base case is reached.

Here are some examples of how to define the recursive relationship in DP:

### 1. Fibonacci sequence:

The Fibonacci sequence is a classic example of DP. The sequence is defined as the sum of the two preceding terms, with the base cases of `F(0)=0` and `F(1)=1`. We can define the recursive relationship as follows:

```
F(n) = F(n-1) + F(n-2), for n > 1
```
This means that the value of the `n`th Fibonacci number can be calculated as the sum of the `n-1`th and `n-2`th Fibonacci numbers. We can use memoization to avoid redundant calculations of the same subproblems.

### 3. Binomial coefficient:

The binomial coefficient `C(n,k)` represents the number of ways to choose `k` objects from `n` distinct objects. We can define the recursive relationship using the following formula:

```
C(n,k) = C(n-1,k-1) + C(n-1,k), for 0 < k < n
C(n,n) = C(n,0) = 1
```
This means that the number of ways to choose `k` objects from `n` objects can be calculated as the sum of the number of ways to choose `k-1` objects from the first `n-1` objects and the number of ways to choose `k` objects from the first `n-1` objects. The base cases are when `k` is equal to 0 or `n`, in which case there is only one way to choose the objects.

### 3. Longest common subsequence:

The longest common subsequence (LCS) of two strings `s1` and `s2` is the longest subsequence that is common to both strings. We can define the recursive relationship as follows:

```
LCS(s1, s2) = LCS(s1-1, s2-1) + s1[-1], if s1[-1] == s2[-1]
LCS(s1, s2) = max(LCS(s1-1, s2), LCS(s1, s2-1)), otherwise
LCS("", s2) = LCS(s1, ""), where "" represents the empty string
```
This means that the LCS of two strings can be calculated by comparing the last characters of each string. If the last characters match, then the LCS is the LCS of the two strings with the last characters removed, plus the last character. If the last characters do not match, then the LCS is the maximum of the LCS of the first string with the last character removed and the LCS of the second string with the last character removed. The base cases are when one of the strings is empty.

### 4. Knapsack problem:

The knapsack problem involves finding the maximum value that can be obtained by filling a knapsack with a capacity `W` with a subset of `n` items, each with a weight `wi` and value `vi`. We can define the recursive relationship using the following formula:

```
Knapsack(W, i) = 0, if i == 0 or W == 0
Knapsack(W, i) = Knapsack(W, i-1), if wi > W
Knapsack(W, i) = max(Knapsack(W, i-1), vi + Knapsack(W-wi, i-1)), otherwise
```
This means that the maximum value that can be obtained

### 5. Rod cutting problem

The rod cutting problem involves finding the maximum revenue that can be obtained by cutting a rod of length `n` into pieces of different lengths, each with a price `pi`. We can define the recursive relationship using the following formula:

```
R(n) = max(pi + R(n-i)), for i in [1, n]
R(0) = 0
```
This means that the maximum revenue that can be obtained from a rod of length `n` can be calculated by trying all possible lengths `i` to cut the rod, calculating the revenue obtained by cutting the rod at length `i` and recursively calculating the revenue obtained from the remaining piece of length `n-i`. The base case is when the length of the rod is 0, in which case the revenue obtained is also 0. We can again use memoization or tabulation to avoid redundant calculations of the same subproblems.

### 6. Coin change problem
The coin change problem involves finding the minimum number of coins required to make a given sum `S` using a set of coins with different denominations. We can define the recursive relationship using the following formula:

```
C(S) = min(C(S-c)+1), for all c in the set of denominations and S-c >= 0
C(0) = 0
```
This means that the minimum number of coins required to make a sum `S` can be calculated by trying all possible denominations `c` and recursively calculating the minimum number of coins required to make the remaining sum `S-c`. The base case is when the sum is 0, in which case the minimum number of coins required is also 0. Again, we can use memoization or tabulation to avoid redundant calculations of the same subproblems.

—

## Types of Dynamic Programming

There are two main types of dynamic programming:

1.  **Top-down DP (Memoization):** This approach involves breaking down the problem into subproblems and recursively solving them. In order to avoid redundant calculations, we can store the results of the subproblems in a lookup table (memoization). When we encounter a subproblem that has already been solved, we can simply look up the result in the table rather than recalculating it. This approach is also called memoization because we are "memorizing" the results of the subproblems.
    
2.  **Bottom-up DP (Tabulation):** This approach involves solving the subproblems in a bottom-up manner, starting with the smallest subproblems and building up to the larger ones. We store the results of each subproblem in a table (tabulation) and use them to solve larger subproblems. This approach is also called tabulation because we are filling in a table with the results of the subproblems.
    

Both approaches are used in dynamic programming, and which one to use depends on the specific problem and the programmer's preference.

### Differences between the two approaches

Top-down and bottom-up are two common approaches to dynamic programming, and they differ in their implementation and usage.

The main differences between the two approaches are as follows:

1.  Memoization vs. Tabulation: Top-down DP uses memoization, which involves storing the results of previous computations in a memo table and looking them up when needed. Bottom-up DP uses tabulation, which involves building a table of results iteratively from smaller subproblems to larger ones.
    
2.  Recursion vs. Iteration: Top-down DP uses recursion to divide the problem into subproblems, while bottom-up DP uses iteration to build up the solution from the smallest subproblems to the largest.
    
3.  Memory usage: Top-down DP requires additional memory for the memo table, which can grow very large for some problems. Bottom-up DP typically uses less memory since the table is built up iteratively and only needs to store the results of smaller subproblems.
    
4.  Time complexity: Top-down DP may have additional overhead due to the recursive calls and memo table lookups, while bottom-up DP may have better time complexity due to the elimination of these overheads.
    
5.  Initialization: Top-down DP may require additional initialization steps to set up the memo table and recursive calls, while bottom-up DP may require initialization of the table with base cases.
    

Choosing between top-down and bottom-up DP depends on several factors such as problem size, available memory, and personal preference. In general, top-down DP is preferred when the problem has a recursive structure that can be easily represented and memoized, and when only a small portion of the memo table is actually used. Bottom-up DP is preferred when the problem can be represented in a tabular form and when the entire memo table needs to be filled in to obtain the final solution. In some cases, a hybrid approach combining both techniques may also be used.

#### Top-Down

Pros:  
1. It is a natural transformation from the normal Complete Search recursion  
2. Computes the sub-problems only when necessary (sometimes this is faster)  
Cons:  
1. Slower if many sub-problems are revis- ited due to function call overhead (this is not usually penalized in programming contests) 
2. If there are M states, an O(M) table size is required, which can lead to MLE for some harder problems (except if we use the trick in Section 8.3.4)


#### Bottom-Up

**Pros**:  
1. Faster if many sub-problems are revisited as there is no overhead from recursive calls 
2. Can save memory space with the ‘space saving trick’ technique  
**Cons**:  
1. For programmers who are inclined to recursion, this style may not be intuitive
2. If there are M states, bottom-up DP visits and fills the value of all these M states


### Examples of top-down DP

#### 1. Fibonacci sequence

The Fibonacci sequence is a classic example of dynamic programming. We can use top-down dynamic programming (memoization) to calculate the Fibonacci numbers recursively.

```go
func fib(n int, memo []int) int {
    if memo[n] != -1 {
        return memo[n]
    }
    if n == 0 {
        memo[n] = 0
    } else if n == 1 {
        memo[n] = 1
    } else {
        memo[n] = fib(n-1, memo) + fib(n-2, memo)
    }
    return memo[n]
}

func main() {
    n := 6
    memo := make([]int, n+1)
    for i := range memo {
        memo[i] = -1
    }
    fmt.Println(fib(n, memo))
}

```

Illustration:

```
n = 6

                 fib(6)
                /     \
           fib(5)     fib(4)
           /    \     /    \
       fib(4)  fib(3) fib(3) fib(2)
       /    \
   fib(3) fib(2)
   

```

Memoization Table:

```
n    |   0  |   1  |   2  |   3  |   4  |   5  |   6  |
-------------------------------------------------------
memo |   0  |   1  |  -1  |  -1  |  -1  |  -1  |  -1  |
```

#### 2. Longest common subsequence (LCS)


Here's an example of solving the Longest Common Subsequence (LCS) problem using top-down dynamic programming approach in Go language, along with a recursion tree diagram and memo table.

Problem statement: Given two strings `s1` and `s2`, find the length of their longest common subsequence.

Algorithm: We can solve this problem using a top-down dynamic programming approach with memoization. We start by defining a recursive function `LCS(i, j)` that finds the LCS between `s1[:i]` and `s2[:j]`. If the characters at the ends of both strings match, then the LCS is the LCS of the two strings without their last characters, plus the last character. If they don't match, then the LCS is the maximum of the LCS of `s1[:i-1]` and `s2[:j]` and the LCS of `s1[:i]` and `s2[:j-1]`. We memoize the results of this function to avoid recomputing them.

Recursion Tree Diagram:

```
            LCS(5, 4)
           /         \
    LCS(4, 3)         LCS(5, 3)
    /      \           /       \
LCS(3, 2) LCS(4, 2) LCS(4, 2) LCS(5, 2)
```
Memo Table:

```
     0  1  2  3  4  5
   -------------------
0 |  0  0  0  0  0  0
1 |  0  1  1  1  1  1
2 |  0  1  1  2  2  2
3 |  0  1  2  2  2  2
4 |  0  1  2  2  3  3
5 |  0  1  2  2  3  4
```
Code:

```go
func LCS(s1, s2 string) int {
    memo := make([][]int, len(s1)+1)
    for i := range memo {
        memo[i] = make([]int, len(s2)+1)
    }
    return lcsHelper(s1, s2, len(s1), len(s2), memo)
}

func lcsHelper(s1, s2 string, i, j int, memo [][]int) int {
    if memo[i][j] != 0 {
        return memo[i][j]
    }
    if i == 0 || j == 0 {
        return 0
    }
    if s1[i-1] == s2[j-1] {
        memo[i][j] = 1 + lcsHelper(s1, s2, i-1, j-1, memo)
        return memo[i][j]
    }
    memo[i][j] = max(lcsHelper(s1, s2, i-1, j, memo), lcsHelper(s1, s2, i, j-1, memo))
    return memo[i][j]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

Note: The `max` function is used to find the maximum of two integers. You can call the `LCS` function with two strings as arguments to get the length of their LCS.

#### Longest increasing subsequence 

Here's an example of the Longest Increasing Subsequence (LIS) problem using top-down dynamic programming with memoization, along with a recursion tree diagram and memo table, in Go:

Problem Statement: Given an unsorted array of integers, find the length of the longest increasing subsequence.

Solution: We can solve this problem using dynamic programming with memoization. We define the function `LIS` to return the length of the longest increasing subsequence of the subarray `nums[i:]`, where `i` is the index of the first element of the subarray. The base case is when the subarray is empty, in which case the LIS is 0. For each index `i` in the subarray, we check if `nums[i]` can be included in the LIS by comparing it with the previous element. If it can be included, we recursively call `LIS(i+1)` and add 1 to the result. We store the result in a memo table to avoid redundant computations.

Recursion Tree Diagram:

```
   LIS(0)
  /   |   \
LIS(1) LIS(2) LIS(3)
  |      |   /  \
LIS(2) LIS(3) ... LIS(n-1)
  |
LIS(3)
  |
  ...

```

Memo Table:

```
  i  0 1 2 3 4 5
LIS  0 3 2 1 1 1
```

Code snippet

```go
func lengthOfLIS(nums []int) int {
    memo := make([]int, len(nums))
    for i := range memo {
        memo[i] = -1 // initialize memo table with -1
    }
    return LIS(nums, 0, memo)
}

func LIS(nums []int, i int, memo []int) int {
    if i == len(nums) {
        return 0 // base case: empty subarray
    }
    if memo[i] != -1 {
        return memo[i] // return memoized result
    }
    maxLIS := 1 // minimum LIS is 1 (the element itself)
    for j := i+1; j < len(nums); j++ {
        if nums[j] > nums[i] {
            maxLIS = max(maxLIS, 1+LIS(nums, j, memo))
        }
    }
    memo[i] = maxLIS // memoize result
    return maxLIS
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
In the code above, we first initialize the memo table with -1. Then, we call the `LIS` function with the input array `nums`, starting from index 0, and the memo table. In the `LIS` function, we first check if the result is already memoized, and return it if it is. Otherwise, we loop through the remaining subarray to find the maximum LIS that can be obtained by including the current element. We recursively call `LIS` with the next index and add 1 to the result if the current element can be included. Finally, we memoize the result and return it. The `max` function is used to find the maximum of two integers.

#### Edit Distance

Here's an example of the Edit Distance problem using top-down dynamic programming with memoization, along with a recursion tree diagram and memo table, in Go:

Problem Statement: Given two strings `s1` and `s2`, find the minimum number of operations required to convert `s1` to `s2`. The allowed operations are: insert a character, delete a character, or replace a character.

Solution: We can solve this problem using dynamic programming with memoization. We define the function `editDistance` to return the minimum number of operations required to convert the substring `s1[i:]` to the substring `s2[j:]`, where `i` and `j` are the indices of the first characters of the substrings. The base case is when one of the substrings is empty, in which case the edit distance is the length of the other substring. For each character in the substrings, we check if they are equal. If they are, we don't need to perform any operation and we recursively call `editDistance(i+1, j+1)` with the same substrings. If they are not equal, we have three options: insert a character, delete a character, or replace a character. We recursively call `editDistance(i+1, j)`, `editDistance(i, j+1)`, and `editDistance(i+1, j+1)`, and take the minimum of the three results. We store the result in a memo table to avoid redundant computations.

Recursion Tree Diagram:

```
                             (horses, rose, 0, 0)
                             /               |                 \
         (horses, rose, 0, 1)            (horses, rose, 1, 0)           (horses, rose, 1, 1)
           /      |      \               /      |      \               /      |      \
(horses, rose, 0, 2)  .     .     .    (horses, rose, 2, 0)  .     .    .     (horses, rose, 2, 1)
        /     |    \                    /     |    \                   /     |      \
(horses, rose, 0, 3) .  .  .    .   .  (horses, rose, 3, 0) .   .   .    (horses, rose, 3, 1)
     /   |  \                            /   |  \                        /    |   \
.   .   .   .   .                  (horses, rose, 4, 0)   .   .   .  (horses, rose, 4, 1)
                              /        |         \
                            .   .   .   .   (horses, rose, 5, 0)
```

Memo Table:

```
   j  0 1 2 3 4 5 6
i  |   h o r s e s
0  | 0 1 2 3 4 5 6
1  r 1 1 1 2 3 4 5
2  o 2 2 2 2 3 4 5
3  s 3 3 3 3 3 4 5
4  e 4 4 4 4 4 4 5
```

Code 

```
func minDistance(word1 string, word2 string) int {
    memo := make([][]int, len(word1)+1)
    for i := range memo {
        memo[i] = make([]int, len(word2)+1)
        for j := range memo[i] {
            memo[i][j] = -1 // initialize memo table with -1
        }
    }
    return editDistance(word1, word2, 0, 0, memo)
}

func editDistance(s1, s2 string, i, j int, memo [][]int) int {
    if memo[i][j] != -1 {
        return memo[i][j] // return memoized result
    }
    if i == len(s1) {
        return len(s2) - j // base case: s1 is empty
    }
    if j == len(s2) {
        return len(s1) - i // base case: s2 is empty
    }
    if s1[i] == s2[j] {
        memo[i][j] = editDistance(s1, s2, i+1, j+1, memo) // no operation needed
    } else {
        insert := 1 + editDistance(s1, s2, i, j+1, memo)    // insert a character
        delete := 1 + editDistance(s1, s2, i+1, j, memo)    // delete a character
        replace := 1 + editDistance(s1, s2, i+1, j+1, memo) // replace a character
        memo[i][j] = min(insert, min(delete, replace))
    }
    return memo[i][j]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

```

### Examples of bottom-up dp

#### Fibonacci sequence
Here's another example of how to solve the Fibonacci sequence problem using the bottom-up approach

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    
    // create a memo table to store the computed values
    memo := make([]int, n+1)
    memo[0], memo[1] = 0, 1
    
    // compute the values bottom-up
    for i := 2; i <= n; i++ {
        memo[i] = memo[i-1] + memo[i-2]
    }
    
    // return the nth Fibonacci number
    return memo[n]
}

```
Tabulation
```
 n | memo[n]
---|--------
 0 |   0
 1 |   1
 2 |   1
 3 |   2
 4 |   3
 5 |   5
 6 |   8
 7 |  13
 8 |  21
 9 |  34
10 |  55
```
In this example, the bottom-up approach is used to compute the nth Fibonacci number. The memo table is used to store the intermediate results of the subproblems, which are then used to compute the final result. The table example shows the values stored in the memo table for different values of n.

#### Coin change problem
here's an example of how to solve the coin change problem using the bottom-up approach with a table example in Go:

```go
func coinChange(coins []int, amount int) int {
    // initialize a memo table to store the minimum number of coins needed for each amount
    memo := make([]int, amount+1)
    for i := range memo {
        memo[i] = amount + 1
    }
    memo[0] = 0
    
    // compute the minimum number of coins needed for each amount bottom-up
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if i >= coin {
                memo[i] = min(memo[i], memo[i-coin]+1)
            }
        }
    }
    
    // return -1 if it's not possible to make the amount with the given coins, otherwise return the minimum number of coins needed
    if memo[amount] > amount {
        return -1
    }
    return memo[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

```
Tabulation
```
  amount | memo[amount]
---------|-------------
       0 |            0
       1 |            1
       2 |            1
       3 |            2
       4 |            2
       5 |            1
       6 |            2
       7 |            2
       8 |            2
       9 |            3
      10 |            2

```
In this example, the bottom-up approach is used to solve the coin change problem, which involves finding the minimum number of coins needed to make a given amount using a given set of coin denominations. The memo table is used to store the minimum number of coins needed for each amount. The table example shows the values stored in the memo table for different amounts.

#### Maximum subarray problem

```go
func maxSubArray(nums []int) int {
    // initialize a memo table to store the maximum sum of subarrays ending at each index
    memo := make([]int, len(nums))
    memo[0] = nums[0]
    maxSum := memo[0]
    
    // compute the maximum sum of subarrays ending at each index bottom-up
    for i := 1; i < len(nums); i++ {
        memo[i] = max(memo[i-1]+nums[i], nums[i])
        maxSum = max(maxSum, memo[i])
    }
    
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```
Tabulation

```
 i | nums[i] | memo[i] | maxSum
---|---------|---------|-------
 0 |      -2 |      -2 |    -2
 1 |       1 |       1 |     1
 2 |      -3 |      -2 |     1
 3 |       4 |       4 |     4
 4 |      -1 |       3 |     4
 5 |       2 |       5 |     5
 6 |       1 |       6 |     6
 7 |      -5 |       1 |     6

```
In this example, the bottom-up approach is used to solve the **Maximum Subarray problem**, which involves finding the `maximum sum of a contiguous subarray within a given array of integers`. The memo table is used to store the maximum sum of subarrays ending at each index, which is then used to compute the final result. The table example shows the values stored in the memo table and the maximum sum found so far for each index.

#### Longest common subsequence problem
here's an example of how to solve the Longest Common Subsequence (LCS) problem using the bottom-up approach.

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    // initialize a memo table to store the length of LCS between the two texts
    m, n := len(text1), len(text2)
    memo := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        memo[i] = make([]int, n+1)
    }
    
    // compute the length of LCS bottom-up
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                memo[i][j] = memo[i-1][j-1] + 1
            } else {
                memo[i][j] = max(memo[i-1][j], memo[i][j-1])
            }
        }
    }
    
    return memo[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

Tabulation

```
     | "" | r | o | s | e
-----|---|---|---|---|---
   ""|  0|  0|  0|  0|  0
   h |  0|  0|  0|  0|  0
   o |  0|  0|  1|  1|  1
   r |  0|  1|  1|  1|  1
   s |  0|  1|  1|  2|  2
   e |  0|  1|  1|  2|  3

```
n this example, the bottom-up approach is used to solve the LCS problem, which involves finding the length of the longest common subsequence between two given strings. The memo table is used to store the length of LCS between substrings of the two input strings, and is computed bottom-up by comparing each character in the two strings. The table example shows the values stored in the memo table for the example input strings "horse" and "rose". The final result is found in the bottom-right corner of the table, which indicates that the length of the LCS between "horse" and "rose" is 3.


#### Knapsack problem

```go
func knapsack(weights, values []int, capacity int) int {
    n := len(weights)
    memo := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        memo[i] = make([]int, capacity+1)
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= capacity; j++ {
            if weights[i-1] > j {
                memo[i][j] = memo[i-1][j]
            } else {
                memo[i][j] = max(memo[i-1][j], memo[i-1][j-weights[i-1]]+values[i-1])
            }
        }
    }
    
    return memo[n][capacity]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

Tabulation:
```
         | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10
---------|---|---|---|---|---|---|---|---|---|---|---
 0 items | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0
 1 item  | 0 | 0 | 0 | 0 | 0 | 6 | 6 | 6 | 6 | 6 | 6
 2 items | 0 | 0 | 0 | 0 | 7 | 7 | 7 | 7 | 7 | 13| 13
 3 items | 0 | 0 | 0 | 10| 10| 10| 10| 17| 17| 17| 17

``` 
In this example, the bottom-up approach is used to solve the Knapsack problem, which involves finding the maximum value that can be obtained from a set of items with given weights and values, subject to a weight limit. The memo table is used to store the maximum value that can be obtained using a subset of the input items and a given capacity. The table example shows the values stored in the memo table for the example input weights {3, 1, 5} and values {4, 6, 10}, with a capacity of 10. The final result is found in the bottom-right corner of the table, which indicates that the maximum value that can be obtained using the input items and capacity is 17.

####  Partition problem

Here's an example of the Partition problem using bottom-up DP, with table example and code snippets in Go.

The Partition problem asks if it is possible to divide a set of positive integers into two subsets such that the sum of elements in both subsets is equal.

Let's say we have the set `nums = [1, 5, 11, 5]`. We want to check if we can divide it into two subsets with equal sum.

We can approach this problem using bottom-up DP:

1.  We start by initializing a 2D boolean memo table with size `n+1` by `sum/2 + 1`, where `n` is the length of the input array `nums`, and `sum` is the total sum of the input array.
    
2.  We set the base cases, where if the sum is 0, then it is possible to divide the array into two equal subsets, so we set the memo table to `true`. If the input array is empty or the sum is negative, then it is not possible to divide the array into two equal subsets, so we set the memo table to `false`.
    
3.  For each element in the input array `nums` and each possible sum `s` from 1 to `sum/2`, we calculate if it is possible to achieve the sum `s` using the current element. If it is possible to achieve the sum `s` using the current element, then we set the corresponding entry in the memo table to `true`.
    
4.  After filling in the entire memo table, the answer is the value at the bottom-right corner of the memo table. If it is `true`, then it is possible to divide the input array into two subsets with equal sum. Otherwise, it is not possible.
    

Here's the code snippet in Go:

```go
func canPartition(nums []int) bool {
    n := len(nums)
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum%2 != 0 {
        return false
    }
    target := sum / 2
    
    memo := make([][]bool, n+1)
    for i := range memo {
        memo[i] = make([]bool, target+1)
    }
    
    for i := 0; i <= n; i++ {
        memo[i][0] = true
    }
    for j := 1; j <= target; j++ {
        memo[0][j] = false
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= target; j++ {
            memo[i][j] = memo[i-1][j]
            if j >= nums[i-1] {
                memo[i][j] = memo[i][j] || memo[i-1][j-nums[i-1]]
            }
        }
    }
    
    return memo[n][target]
}

```

Here's the memo table for the example input `nums = [1, 5, 11, 5]`:

|     | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 0   | T   | F   | F   | F   | F   | F   | F   | F   | F   | F   | F   | F   |
| 1   | T   | T   | F   | F   | F   | F   | F   | F   | F   |     |     |     |


## Examples of top-down approach
1.  Fibonacci series
2.  Longest common subsequence
3.  Knapsack problem
4.  Coin change problem
5.  Maximum subarray problem
6.  Longest increasing subsequence
7.  Matrix chain multiplication
8.  Edit distance problem
9.  Longest palindromic subsequence
10. Rod cutting problem
11. Subset sum problem
12. Minimum edit distance
13. Optimal binary search tree
14. Longest bitonic subsequence
15. Counting problems
16. Longest common substring
17. Maximum sum increasing subsequence
18. Painting fence problem
19. Maximum sum subsequence
20. 0/1 Knapsack problem
21. Egg dropping problem
22. Partition problem
23. Longest zigzag subsequence
24. String matching problems
25. Maximum sum subarray with at most k elements
26. Longest alternating subsequence
27. Tower of Hanoi problem
28. Longest common prefix
29. Bellman-Ford algorithm
30. Shortest path in a DAG (Directed Acyclic Graph)

## Examples of bottom-up approach

1.  Fibonacci series
2.  Longest common subsequence
3.  Knapsack problem
4.  Coin change problem
5.  Maximum subarray problem
6.  Longest increasing subsequence
7.  Matrix chain multiplication
8.  Edit distance problem
9.  Longest palindromic subsequence
10. Rod cutting problem
11. Subset sum problem
12. Minimum edit distance
13. Optimal binary search tree
14. Longest bitonic subsequence
15. Counting problems (such as counting the number of ways to make change)
16. Longest common substring
17. Maximum sum increasing subsequence
18. Painting fence problem
19. Maximum sum subsequence without adjacent elements
20. 0/1 Knapsack problem
21. Egg dropping problem
22. Partition problem
23. Longest zigzag subsequence
24. String matching problems
25. Maximum sum subarray with at most k elements
26. Longest alternating subsequence
27. Tower of Hanoi problem
28. Longest common prefix
29. Shortest common supersequence
30. Maximum size subset with given sum.


