# Top Down examples

### 1.  Fibonacci series
The Fibonacci sequence is a sequence of numbers where each number is the sum of the previous two. The first two numbers in the sequence are 0 and 1.

The recursive relation for the Fibonacci sequence can be defined as follows:

```go
fib(n) = fib(n-1) + fib(n-2)
```
where `fib(n)` represents the `n`-th number in the Fibonacci sequence.

Using top-down dynamic programming, we can optimize the recursive algorithm by storing the results of previously computed subproblems in an array or a hash table. This technique is called memoization.

Here is how the top-down dynamic programming approach for the Fibonacci sequence would look like:

1.  Check if the result for `fib(n)` is already computed and stored in the memoization table. If so, return the result.
2.  If `n` is 0 or 1, return `n`.
3.  Otherwise, compute `fib(n-1)` and `fib(n-2)` recursively.
4.  Store the result of `fib(n)` in the memoization table.
5.  Return `fib(n)`.

By using memoization, we avoid recomputing the same subproblems multiple times, leading to a significant improvement in time complexity compared to the naive recursive approach.



### 2.  Longest common subsequence
The Longest Common Subsequence (LCS) problem is a classic dynamic programming problem. Given two sequences `X` and `Y`, the LCS is the longest subsequence that is present in both `X` and `Y`. A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

For example, consider the following sequences:
```
X: AGGTAB
Y: GXTXAYB
```
The longest common subsequence of `X` and `Y` is `GTAB`, which has a length of 4.

To solve the LCS problem, we can use dynamic programming. We can build a table where the `i`-th row and `j`-th column of the table represents the length of the LCS of the first `i` characters of `X` and the first `j` characters of `Y`. The table can be filled using the recursive relation that we discussed earlier:

```
LCS(X, Y, i, j) = {
    0                           if i == 0 or j == 0
    LCS(X, Y, i-1, j-1) + 1     if X[i-1] == Y[j-1]
    max(LCS(X, Y, i, j-1), LCS(X, Y, i-1, j))  if X[i-1] != Y[j-1]
}
```
Using this recursive relation, we can fill the table row by row, starting from the first row and the first column, until we fill the last row and the last column. The value in the last row and last column of the table gives us the length of the LCS of the entire sequences `X` and `Y`.

Here is what the table would look like for the example sequences:

```
  Y  G X T X A Y B
X   0 0 0 0 0 0 0 0
A   0 0 0 0 0 1 1 1
G   1 1 1 1 1 1 1 1
G   1 1 1 1 1 1 1 1
T   1 1 1 2 2 2 2 2
A   1 1 1 2 2 3 3 3
B   1 1 1 2 2 3 3 4
```
The value in the last row and last column of the table is 4, which is the length of the LCS of `X` and `Y`.

To find the actual LCS, we can backtrack from the last cell of the table to the first cell, using the values in the table to guide our decisions. We can start at the last cell of the table and follow the path with the maximum values, until we reach the first row or column of the table. The characters on this path form the LCS. For the example sequences, the path with the maximum values is:

```
X: AGGTAB
Y: GXTXAYB
     ↑ ↑  ↑
LCS: G T A B
```
Thus, the LCS of `X` and `Y` is `GTAB`.

Using top-down dynamic programming, we can optimize the recursive algorithm by storing the results of previously computed subproblems in an array or a hash table. This technique is called memoization.

Here is how the top-down dynamic programming approach for the LCS problem would look like:

1.  Check if the result for `LCS(X, Y, i, j)` is already computed and stored in the memoization table. If so, return the result.
2.  If `i` or `j` is 0, return 0.
3.  If the last characters of `X` and `Y` are the same, compute `LCS(X, Y, i-1, j-1) + 1`.
4.  If the last characters of `X` and `Y` are different, compute `max(LCS(X, Y, i, j-1), LCS(X, Y, i-1, j))`.
5.  Store the result of `LCS(X, Y, i, j)` in the memoization table.
6.  Return `LCS(X, Y, i, j)`.

By using memoization, we avoid recomputing the same subproblems multiple times, leading to a significant improvement in time complexity compared to the naive recursive approach.


### 3.  Knapsack problem

The Knapsack problem is a classic optimization problem in which we are given a set of items, each with a weight and a value, and a knapsack with a capacity. The goal is to choose a subset of the items that maximizes the total value while keeping the total weight within the capacity of the knapsack.

For example, suppose we have the following items:

```
Item 1: weight = 2, value = 6
Item 2: weight = 2, value = 10
Item 3: weight = 3, value = 12
Item 4: weight = 4, value = 13

```
And a knapsack with a capacity of 7. We can choose the items in the following way to maximize the total value while keeping the total weight within the capacity of the knapsack:

```
Item 2: weight = 2, value = 10
Item 3: weight = 3, value = 12
```
The total weight of these items is 5, which is within the capacity of the knapsack, and the total value is 22, which is the maximum possible value that can be obtained.

To solve the Knapsack problem, we can use dynamic programming. We can build a table where the `i`-th row and `j`-th column of the table represents the maximum value that can be obtained by choosing a subset of the first `i` items, with a total weight not exceeding `j`. The table can be filled using the recursive relation that we discussed earlier:

```
knapsack(i, j) = {
    0                                            if i == 0 or j == 0
    knapsack(i-1, j)                              if weight[i-1] > j
    max(knapsack(i-1, j), value[i-1] + knapsack(i-1, j-weight[i-1]))   otherwise
}

```
Using this recursive relation, we can fill the table row by row, starting from the first row and the first column, until we fill the last row and the last column. The value in the last row and last column of the table gives us the maximum value that can be obtained by choosing a subset of all the items, with a total weight not exceeding the capacity of the knapsack.

Here is what the table would look like for the example items:

```
Capacity  0 1 2  3  4  5  6  7
Item 0    0 0 0  0  0  0  0  0
Item 1    0 0 6  6  6  6  6  6
Item 2    0 0 10 10 16 16 16 16
Item 3    0 0 10 12 16 22 22 28
Item 4    0 0 10 12 16 22 25 28

```
The value in the last row and last column of the table is 28, which is the maximum possible value that can be obtained.

To find the actual items that are included in the optimal solution, we can backtrack from the last cell of the table to the first cell, using the values in the table to guide our decisions. Starting from the last cell of the table, we can check whether the value in the current cell is equal to the value in the cell above it. If it is, it means that the item in the current row was not included in the optimal solution. If the value in the current cell is greater than the value in the cell above.


### 4.  Coin change problem
The coin change problem is a classic algorithmic problem that involves finding the minimum number of coins required to make a given amount of money. Here's an explanation of the problem with examples, followed by the recursive relation for the top-down dynamic programming approach.

**Problem Statement** Suppose you are given a set of coins of different denominations and a total amount of money. You need to find the minimum number of coins required to make up that amount. Assume that you have an infinite supply of each coin denomination.

For example, let's say you have the following coins: {1, 5, 10, 25} cents. If you need to make up an amount of 37 cents, you can do it using one quarter (25 cents) and one dime (10 cents) and two pennies (1 cent). This requires a total of four coins. But if you need to make up an amount of 30 cents, you can do it using three dimes (10 cents) and no pennies. This requires a total of three coins.

The goal is to find the minimum number of coins required to make up a given amount of money.

**Recursive Relation** To solve the coin change problem using top-down dynamic programming, we can start by defining a recursive function `minCoins(n)` that takes the amount of money `n` as input and returns the minimum number of coins required to make up that amount.

We can define this function as follows:
```
minCoins(n) = 0 if n = 0
minCoins(n) = INF if n < 0
minCoins(n) = 1 + min(minCoins(n - coin[i])) for i in {0, 1, 2, ..., k-1}
```

Here, `INF` is a large value that represents an impossible number of coins (e.g., if `n` is negative and there are no coins to subtract from it). `k` is the number of different coin denominations available, and `coin[i]` is the value of the `i`-th coin denomination.

The base case of the recursion is when the amount of money `n` is zero. In this case, we need zero coins, so we return 0.

If `n` is negative, then we cannot make up that amount with the available coins, so we return `INF`.

Otherwise, we need to consider all the possible ways of making up the amount `n` using the available coins. We do this by subtracting each coin denomination from `n` and recursively computing the minimum number of coins required to make up the remaining amount. We add 1 to the result to account for the coin that we subtracted.

Finally, we take the minimum of all these possible solutions and return it.

Note that this recursive function has overlapping subproblems. For example, if we compute `minCoins(10)` and `minCoins(15)`, both of these computations will require computing `minCoins(5)`. To avoid redundant computations, we can store the results of each computation in a memoization table and look them up instead of recomputing them.



### 5.  Maximum subarray problem

The maximum subarray problem is a classic algorithmic problem that involves finding the contiguous subarray with the largest sum within a given array of integers. Here's an explanation of the problem with examples, followed by the recursive relation for the top-down dynamic programming approach.

**Problem Statement** Suppose you are given an array of integers. You need to find the contiguous subarray with the largest sum.

For example, let's say you have the following array: \[-2, 1, -3, 4, -1, 2, 1, -5, 4\]. The contiguous subarray with the largest sum is \[4, -1, 2, 1\], which has a sum of 6.

The goal is to find the contiguous subarray with the largest sum.

**Recursive Relation** To solve the maximum subarray problem using top-down dynamic programming, we can start by defining a recursive function `maxSubArray(nums, i)` that takes the array `nums` and an index `i` as input and returns the maximum sum of any contiguous subarray starting at index `i`.

We can define this function as follows:

```
maxSubArray(nums, i) = 0 if i >= len(nums)
maxSubArray(nums, i) = nums[i] if i == len(nums) - 1
maxSubArray(nums, i) = max(nums[i], nums[i] + maxSubArray(nums, i+1))
```
Here, `len(nums)` is the length of the input array.

The base case of the recursion is when the index `i` is equal to or greater than the length of the input array. In this case, we cannot form any subarray, so we return 0.

If `i` is equal to `len(nums) - 1`, then we have only one element left in the array, so the maximum sum of any contiguous subarray starting at index `i` is simply that element itself. We return `nums[i]` in this case.

Otherwise, we need to consider two possibilities: either we include the current element `nums[i]` in the subarray, or we don't. If we don't include it, then the maximum sum of any contiguous subarray starting at index `i` is simply the maximum sum of any contiguous subarray starting at index `i+1`. If we do include it, then the maximum sum of any contiguous subarray starting at index `i` is `nums[i]` plus the maximum sum of any contiguous subarray starting at index `i+1`.

Finally, we take the maximum of these two possibilities and return it.

Note that this recursive function has overlapping subproblems. For example, if we compute `maxSubArray(nums, 0)` and `maxSubArray(nums, 1)`, both of these computations will require computing `maxSubArray(nums, 2)`. To avoid redundant computations, we can store the results of each computation in a memoization table and look them up instead of recomputing them.


### 6.  Longest increasing subsequence
The longest increasing subsequence problem is a classic algorithmic problem that involves finding the length of the longest increasing subsequence of a given array of integers. Here's an explanation of the problem with examples, followed by the recursive relation for the top-down dynamic programming approach.

**Problem Statement** Suppose you are given an array of integers. You need to find the length of the longest increasing subsequence.

For example, let's say you have the following array: \[10, 9, 2, 5, 3, 7, 101, 18\]. The longest increasing subsequence in this array is \[2, 3, 7, 101\], which has a length of 4.

The goal is to find the length of the longest increasing subsequence.

**Recursive Relation** To solve the longest increasing subsequence problem using top-down dynamic programming, we can start by defining a recursive function `lengthOfLIS(nums, i, prev)` that takes the array `nums`, an index `i`, and the previous element `prev` as input and returns the length of the longest increasing subsequence that can be formed starting from index `i` with `prev` as the previous element.

We can define this function as follows:
```
lengthOfLIS(nums, i, prev) = 0 if i >= len(nums)
lengthOfLIS(nums, i, prev) = max(lengthOfLIS(nums, i+1, prev), 1 + lengthOfLIS(nums, i+1, nums[i])) if nums[i] > prev
lengthOfLIS(nums, i, prev) = lengthOfLIS(nums, i+1, prev) otherwise
```
Here, `len(nums)` is the length of the input array.

The base case of the recursion is when the index `i` is equal to or greater than the length of the input array. In this case, we cannot form any subsequence, so we return 0.

If the current element `nums[i]` is greater than the previous element `prev`, then we have two options: either we include `nums[i]` in the subsequence, or we don't. If we include it, then the length of the longest increasing subsequence that can be formed starting from index `i` with `prev` as the previous element is `1 + lengthOfLIS(nums, i+1, nums[i])`. If we don't include it, then the length of the longest increasing subsequence that can be formed starting from index `i` with `prev` as the previous element is simply `lengthOfLIS(nums, i+1, prev)`.

If `nums[i]` is not greater than `prev`, then we cannot include it in the subsequence, so we simply move on to the next element by calling `lengthOfLIS(nums, i+1, prev)`.

Finally, we take the maximum of these three possibilities and return it.

> **Note** that this recursive function has overlapping subproblems. For example, if we compute `lengthOfLIS(nums, 0, -inf)` and `lengthOfLIS(nums, 1, -inf)`, both of these computations will require computing `lengthOfLIS(nums, 2, 2)` if `nums[2]` is greater than `nums[1]`. To avoid redundant computations, we can store the results of each computation in a memoization table and look them up instead of recomputing them.


### 7.  Matrix chain multiplication
The matrix chain multiplication problem is a classic algorithmic problem that involves finding the most efficient way to multiply a given sequence of matrices. Here's an explanation of the problem with examples, followed by the recursive relation for the top-down dynamic programming approach.

**Problem Statement** Suppose you are given a sequence of matrices to multiply. The goal is to find the most efficient way to multiply these matrices, i.e., the order in which to multiply them such that the total number of scalar multiplications is minimized.

For example, let's say you have the following sequence of matrices:
```
A1: 10x20
A2: 20x30
A3: 30x40
```
One way to multiply these matrices is to first multiply `A1` and `A2`, and then multiply the result with `A3`. The total number of scalar multiplications required for this approach is `10*20*30 + 10*30*40 = 18000`.

Another way to multiply these matrices is to first multiply `A2` and `A3`, and then multiply the result with `A1`. The total number of scalar multiplications required for this approach is `20*30*40 + 10*20*40 = 24000`.

The goal is to find the order of multiplication that requires the minimum number of scalar multiplications.

**Recursive Relation** To solve the matrix chain multiplication problem using top-down dynamic programming, we can start by defining a recursive function `matrixChainOrder(p, i, j)` that takes the array `p`, and the indices `i` and `j` as input and returns the minimum number of scalar multiplications required to multiply the matrices from index `i` to index `j`.

We can define this function as follows:

```
matrixChainOrder(p, i, j) = 0 if i == j
matrixChainOrder(p, i, j) = min(matrixChainOrder(p, i, k) + matrixChainOrder(p, k+1, j) + p[i-1]*p[k]*p[j]) for k in range(i, j)
```
Here, `p` is an array of matrix dimensions such that the `i`-th matrix has dimensions `p[i-1] x p[i]`, and `i` and `j` are the indices of the subsequence of matrices that we want to multiply.

The base case of the recursion is when `i` is equal to `j`. In this case, we cannot multiply any matrices, so we return 0.

If `i` is not equal to `j`, then we need to consider all possible ways to split the subsequence of matrices from index `i` to index `j` into two subsequences and recursively compute the minimum number of scalar multiplications required for each split. We can do this by iterating over all possible values of `k` such that `i <= k < j`. For each value of `k`, we compute the minimum number of scalar multiplications required to multiply the matrices from index `i` to index `k`, and the minimum number of scalar multiplications required to multiply the matrices from index `k+1` to index `j`, and add these two quantities to the number of scalar multiplications required to multiply the resulting matrices. The resulting matrices have dimensions `p[i-1] x p[k]` and `p[k] x p[j]`. Therefore, the total number of scalar multiplications required to multiply these matrices is `p[i-1]*p[k]*p[j]`.

Finally, we take the minimum of all these possibilities and return it.

Note that this recursive function has overlapping subproblems. Let's consider the example of the following sequence of matrices:
```
A1: 10x20
A2: 20x30
A3: 30x40
A4: 40x30
```

To compute the minimum number of scalar multiplications required to multiply all these matrices, we can use the recursive function `matrixChainOrder(p, i, j)` as defined in the previous answer. Let's call this function for `i=1` and `j=4`, i.e., we want to multiply all the matrices in the sequence.

```
matrixChainOrder([10, 20, 30, 40, 30], 1, 4)
```
The first step of the recursion is to compute `matrixChainOrder([10, 20, 30, 40, 30], 1, 2)` and `matrixChainOrder([10, 20, 30, 40, 30], 2, 4)`. We can visualize this as follows:

```
            [A1] x [A2] x [A3] x [A4]
           /                   \
[A1] x [A2] x [A3]     [A2] x [A3] x [A4]
```
Note that both of these subproblems involve computing the minimum number of scalar multiplications required to multiply a subsequence of matrices, namely, the subsequence from index 1 to index 2, and the subsequence from index 2 to index 4, respectively.

The second step of the recursion is to compute `matrixChainOrder([10, 20, 30, 40, 30], 1, 3)`, `matrixChainOrder([10, 20, 30, 40, 30], 2, 3)`, `matrixChainOrder([10, 20, 30, 40, 30], 2, 4)`, and `matrixChainOrder([10, 20, 30, 40, 30], 3, 4)`. We can visualize this as follows:

```
            [A1] x [A2] x [A3] x [A4]
           /                   \
[A1] x [A2] x [A3]     [A2] x [A3] x [A4]
     /            \           /           \
[A1] x [A2]     [A2] x [A3] [A3] x [A4]    [A1] x [A2] x [A3] x [A4]
```
Note that both `matrixChainOrder([10, 20, 30, 40, 30], 1, 2)` and `matrixChainOrder([10, 20, 30, 40, 30], 2, 4)` are reused in the computation of these subproblems.

`A1`, `A2`, `A3`, and `A4` are matrices with the given dimensions.

For example, `A1` is a 10 x 20 matrix, `A2` is a 20 x 30 matrix, `A3` is a 30 x 40 matrix, and `A4` is a 40 x 30 matrix.

The problem is to compute the minimum number of scalar multiplications required to multiply these matrices in the given order. The solution to this problem involves finding the optimal sequence of matrix multiplications to minimize the number of scalar multiplications.

This illustrates the overlapping subproblems in the matrix chain multiplication problem. The same subproblem is computed multiple times in the recursion, leading to redundant computations. Dynamic programming can be used to avoid these redundant computations and speed up the algorithm.


### 8.  Edit distance problem
The edit distance problem is a classic dynamic programming problem that is used to determine the minimum number of operations required to transform one string into another. The three allowed operations are:

1.  Insert a character
2.  Delete a character
3.  Replace a character

Given two strings `s1` and `s2`, the goal is to determine the minimum number of operations required to transform `s1` into `s2`.

Recursive Relation: To solve this problem using top-down dynamic programming, we can use the recursive function `editDistance(s1, s2, i, j)`, where `s1` and `s2` are the input strings and `i` and `j` are the indices of the current characters being compared. The recursive relation is as follows:

```
editDistance(s1, s2, i, j) = 
    if i == 0: j                 // If s1 is empty, insert all characters of s2
    if j == 0: i                 // If s2 is empty, delete all characters of s1
    if s1[i-1] == s2[j-1]:       // If the characters match, no operation is required
        editDistance(s1, s2, i-1, j-1)
    else:
        1 + min(
            editDistance(s1, s2, i-1, j),       // Delete a character from s1
            editDistance(s1, s2, i, j-1),       // Insert a character into s1
            editDistance(s1, s2, i-1, j-1)      // Replace a character in s1
        )
```
The base cases of the recursion are when either `i` or `j` is equal to 0, in which case we need to insert or delete all the remaining characters in the other string.

Example: Let's say we want to find the edit distance between the strings "kitten" and "sitting". We can call the `editDistance` function with `i` and `j` set to the lengths of the two strings:

```
editDistance("kitten", "sitting", 6, 7)
```
The first step of the recursion is to compute `editDistance("kitten", "sitting", 5, 6)` and `editDistance("kitten", "sitting", 6, 5)`. We can visualize this as follows:

```
k i t t e n
      |   |
      |   s
   delete
```

```
k i t t e n
        |
        i
     insert
```

Note that both of these subproblems involve computing the edit distance between a prefix of one of the strings and a prefix of the other string.

The second step of the recursion is to compute `editDistance("kitten", "sitting", 5, 5)`, `editDistance("kitten", "sitting", 4, 6)`, and `editDistance("kitten", "sitting", 5, 7)`. We can visualize this as follows:


```
k i t t e n
      |   |
      |   i
   replace

```

```
k i t t e n
        |
        n
     insert

```

```
k i t t e n
      |   |
      s   s
   replace

```
Note that the subproblem `editDistance("kitten", "sitting", 5, 6)` is reused in the computation of the subproblems `editDistance("kitten", "sitting", 4, 6)`


### 9.  Longest palindromic subsequence
The longest palindromic subsequence problem is a classic dynamic programming problem that is used to find the length of the longest subsequence of a string that is also a palindrome.

A subsequence of a string is a sequence of characters that can be obtained by deleting some characters from the string without changing the order of the remaining characters. A palindrome is a string that reads the same backward as forward.

For example, given the input string "BBABCBCAB", the longest palindromic subsequence is "BABCBAB", which has a length of 7.

Recursive Relation: To solve this problem using top-down dynamic programming, we can use the recursive function `longestPalindromeSubsequence(s, i, j)`, where `s` is the input string and `i` and `j` are the indices of the current characters being compared. The recursive relation is as follows:

```
longestPalindromeSubsequence(s, i, j) = 
    if i > j: 0                             // Base case: empty string
    if i == j: 1                            // Base case: single-character string
    if s[i] == s[j]:                        // Characters match
        2 + longestPalindromeSubsequence(s, i+1, j-1)
    else:
        max(
            longestPalindromeSubsequence(s, i+1, j),    // Delete the first character
            longestPalindromeSubsequence(s, i, j-1)     // Delete the last character
        )

```
The base cases of the recursion are when `i` is greater than `j`, in which case we have an empty string, or when `i` is equal to `j`, in which case we have a single-character string.

**Example**: Let's say we want to find the length of the longest palindromic subsequence of the string "BBABCBCAB". We can call the `longestPalindromeSubsequence` function with `i` set to 0 and `j` set to the length of the string minus one:

```
longestPalindromeSubsequence("BBABCBCAB", 0, 8)
```
The first step of the recursion is to compute `longestPalindromeSubsequence("BBABCBCAB", 1, 7)` and `longestPalindromeSubsequence("BBABCBCAB", 0, 7)`. We can visualize this as follows:

```
B B A B C B C A B
  x         x
  |---------|
    subproblem
```

```
B B A B C B C A B
x               x
|---------------|
   subproblem
```
Note that both of these subproblems involve computing the length of the longest palindromic subsequence of a prefix of the string and a suffix of the string.

The second step of the recursion is to compute `longestPalindromeSubsequence("BBABCBCAB", 2, 6)` and `longestPalindromeSubsequence("BBABCBCAB", 1, 6)`. We can visualize this as follows:

```
B B A B C B C A B
    x     x
    |-----|
   subproblem

```

```
B B A B C B C A B
  x         x
  |---------|
     subproblem
```
Note that the subproblem `longestPalindromeSubsequence("BBABCBCAB", 1, 7)` is reused in the computation of the subproblem `longestPalindromeSubsequence("BBABCBCAB", 1, 6)`.

The final step of the recursion is to compute `longestPalindromeSubsequence("BBABCBCAB", 2, 5)` and `longestPalindromeSubsequence("BBABCBCAB", 1, 5)`. We can visualize this as follows:

```
B B A B C B C A B
      x x
      |-|
 subproblem
```

```
B B A B C B C A B
    x     x
    |-----|
   subproblem
```
Note that the subproblem `longestPalindromeSubsequence("BBABCBCAB", 1, 6)` is reused in the computation of the subproblem `longestPalindromeSubsequence("BBABCBCAB", 1, 5)`.

At this point, we have reached the base case of the recursion when `i` is greater than `j` in the subproblem `longestPalindromeSubsequence("BBABCBCAB", 3, 4)`. We return 0 from this subproblem, and the recursion unwinds.

Finally, we get the maximum of the two subproblems `longestPalindromeSubsequence("BBABCBCAB", 1, 7)` and `longestPalindromeSubsequence("BBABCBCAB", 0, 7)` that were computed in the first step of the recursion. These correspond to the two options of deleting the first character or deleting the last character to obtain a palindromic subsequence. The maximum of these two options is the length of the longest palindromic subsequence of the original string "BBABCBCAB", which is 7.

Thus, the solution to the problem is 7.


### 10. Rod cutting problem

The rod cutting problem is a classic optimization problem in computer science and mathematics that involves determining the maximum revenue that can be obtained by cutting a rod of a given length into pieces of specific lengths and selling them at given prices. The problem can be formulated as follows:

Given a rod of length n inches and a list of prices pi for i = 1, 2, ..., n, determine the maximum revenue that can be obtained by cutting the rod into pieces of lengths i = 1, 2, ..., n and selling them at the corresponding prices.

For example, suppose we have a rod of length 4 inches and the following price list:

| Length (i) | 1   | 2   | 3   | 4   |
| --- | --- | --- | --- | --- |
| Price (pi) | 1   | 5   | 8   | 9   |

To maximize revenue, we need to determine how to cut the rod into pieces of length 1, 2, 3, and 4 inches such that the total revenue is maximized. Here are some possible ways to cut the rod:

* Cut the rod into four pieces of length 1 inch. The total revenue would be 4 x 1 = 4.
* Cut the rod into two pieces of length 2 inches. The total revenue would be 2 x 5 = 10.
* Cut the rod into one piece of length 1 inch and one piece of length 3 inches. The total revenue would be 1 x 1 + 1 x 8 = 9.
* Cut the rod into one piece of length 2 inches and one piece of length 2 inches. The total revenue would be 2 x 5 = 10.
* Cut the rod into one piece of length 4 inches. The total revenue would be 1 x 9 = 9.

In this case, cutting the rod into two pieces of length 2 inches would maximize revenue with a total revenue of 10.

Here's another example. Suppose we have a rod of length 8 inches and the following price list:

| Length (i) | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Price (pi) | 1   | 5   | 8   | 9   | 10  | 17  | 17  | 20  |

To maximize revenue, we need to determine how to cut the rod into pieces of length 1, 2, 3, ..., 8 inches such that the total revenue is maximized. Here are some possible ways to cut the rod:

* Cut the rod into eight pieces of length 1 inch. The total revenue would be 8 x 1 = 8.
* Cut the rod into four pieces of length 2 inches. The total revenue would be 4 x 5 = 20.
* Cut the rod into two pieces of length 4 inches. The total revenue would be 2 x 9 = 18.
* Cut the rod into one piece of length 2 inches, one piece of length 3 inches, and one piece of length 4 inches. The total revenue would be 1 x 5 + 1 x 8 + 1 x 9 = 22.
* Cut the rod into one piece of length 8 inches. The total revenue would be 1 x 20 = 20.

In this case, cutting the rod into one piece of length 2 inches, one piece of length 3 inches, and one piece of length 4 inches would maximize revenue with a total revenue of 22.

here are the steps to solve the rod cutting problem using top-down dynamic programming:

1.  Define a memoization table to store the maximum revenue for each subproblem. The table will have n + 1 rows (corresponding to rod lengths 0 to n) and n + 1 columns (corresponding to subproblems of cutting a rod of length 0 to n into pieces of length 0 to n).
    
2.  Implement a recursive function that takes the current length of the rod and the memoization table as input, and returns the maximum revenue that can be obtained by cutting the rod.
    
3.  Inside the recursive function, check if the maximum revenue for the current subproblem has already been computed and stored in the memoization table. If so, return the stored value.
    
4.  If the maximum revenue has not been computed yet, consider all possible ways of cutting the rod into pieces and compute the revenue for each possible cut. Choose the cut that maximizes revenue and store the maximum revenue in the memoization table.
    
5.  Return the maximum revenue for the current subproblem.
    

Here's the code in Go:

```go
func rodCutting(n int, prices []int) int {
    memo := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        memo[i] = make([]int, n+1)
        for j := 0; j <= n; j++ {
            memo[i][j] = -1
        }
    }
    return rodCuttingHelper(n, prices, memo)
}

func rodCuttingHelper(n int, prices []int, memo [][]int) int {
    if n == 0 {
        return 0
    }
    if memo[n][n] != -1 {
        return memo[n][n]
    }
    maxRevenue := 0
    for i := 1; i <= n; i++ {
        revenue := prices[i-1] + rodCuttingHelper(n-i, prices, memo)
        if revenue > maxRevenue {
            maxRevenue = revenue
        }
    }
    memo[n][n] = maxRevenue
    return maxRevenue
}
```
In this implementation, the `rodCutting` function initializes the memoization table and calls the recursive `rodCuttingHelper` function to compute the maximum revenue. 

The `rodCuttingHelper` function checks if the maximum revenue for the current subproblem has already been computed and stored in the memoization table. If so, it returns the stored value. Otherwise, it considers all possible ways of cutting the rod into pieces and recursively computes the maximum revenue for each possible cut. The maximum revenue is stored in the memoization table and returned.



### 11. Subset sum problem

The Subset Sum problem is a classic problem in computer science and mathematics. It is a decision problem that asks whether a given set of integers contains a non-empty subset that adds up to a given target sum.

Formally, given a set of n integers S={a1,a2,...,an}, and a target sum T, the problem is to determine whether there exists a subset S'⊆S such that the sum of the elements in S' equals T.

For example, suppose we have the set S={2, 5, 8, 3, 1} and the target sum T=9. We need to determine whether there exists a subset of S that adds up to 9.

We can see that there are two subsets of S that add up to 9: {2, 5, 1} and {3, 1, 5}. Therefore, the answer to this instance of the Subset Sum problem is "yes".

Another example is if we have the set S={-7, -3, -2, 5, 8} and the target sum T=0. We need to determine whether there exists a subset of S that adds up to 0.

We can see that there is one subset of S that adds up to 0: {-3, -2, 5}. Therefore, the answer to this instance of the Subset Sum problem is "yes".

However, if we have the set S={1, 3, 5, 7, 9} and the target sum T=2, we need to determine whether there exists a subset of S that adds up to 2.

We can see that there is no subset of S that adds up to 2. Therefore, the answer to this instance of the Subset Sum problem is "no".

The Subset Sum problem is a well-known NP-complete problem, which means that it is believed to be computationally difficult to find an algorithm that solves all instances of the problem efficiently.


Here are the steps to solve the Subset Sum problem using top-down dynamic programming:

1.  Define a memoization table that will store the results of the subproblems. Initialize the table with a special value, such as -1, to indicate that the subproblem has not been solved yet.
2.  Define a recursive function that will solve the subproblems. The function should take the set of integers S, the current index i, and the current target sum t as input.
3.  Check if the current subproblem has already been solved by looking up the memoization table. If the solution is already known, return the cached result.
4.  If the current target sum is 0, return true, since the empty subset always adds up to 0.
5.  If the current index is less than 0, return false, since we have exhausted all the integers in the set without finding a subset that adds up to the target sum.
6.  If the current integer a\[i\] is greater than the target sum t, skip it and move on to the next integer by calling the recursive function with the same index i-1.
7.  If none of the previous cases apply, we have two options: either include the current integer a\[i\] in the subset, or skip it and move on to the next integer. We can solve both options recursively and combine their results using the logical OR operator.
8.  Store the result of the subproblem in the memoization table and return it.

Here is the Go code that implements the above steps:

```go
func subsetSumDP(S []int, T int, memo [][]int, i int) bool {
    if T == 0 {
        return true
    }
    if i < 0 {
        return false
    }
    if memo[i][T] != -1 {
        return memo[i][T] == 1
    }
    if S[i] > T {
        memo[i][T] = subsetSumDP(S, T, memo, i-1)
    } else {
        memo[i][T] = subsetSumDP(S, T, memo, i-1) || subsetSumDP(S, T-S[i], memo, i-1)
    }
    return memo[i][T] == 1
}

func subsetSum(S []int, T int) bool {
    memo := make([][]int, len(S))
    for i := range memo {
        memo[i] = make([]int, T+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return subsetSumDP(S, T, memo, len(S)-1)
}

```
In this implementation, the memoization table is a two-dimensional slice of integers. The value `memo[i][T]` represents the solution to the subproblem of finding a subset that adds up to T using the first i+1 integers of the set S. The value -1 in the memoization table indicates that the subproblem has not been solved yet, 0 indicates that the subproblem has been solved but the answer is false, and 1 indicates that the subproblem has been solved and the answer is true.

The main function subsetSum initializes the memoization table and calls the recursive function subsetSumDP with the initial parameters. The function subsetSumDP implements the steps described above, and the result is returned as a boolean value indicating whether there exists a subset of S that adds up to T.

### 12. Minimum edit distance

The Minimum Edit Distance (MED) problem is a way to measure the similarity between two strings. It calculates the minimum number of operations needed to transform one string into the other. The operations are typically insertion, deletion, and substitution of individual characters.

For example, suppose we have two strings "kitten" and "sitting". The minimum edit distance between them is 3, because we can transform "kitten" into "sitting" by performing the following three operations:

1.  Replace "k" with "s"
2.  Replace "e" with "i"
3.  Insert "g" at the end

Therefore, the minimum edit distance between "kitten" and "sitting" is 3.

Another example is "book" and "back". The minimum edit distance between them is 2, because we can transform "book" into "back" by performing the following two operations:

1.  Replace "o" with "a"
2.  Delete "k"

Therefore, the minimum edit distance between "book" and "back" is 2.

Here's one more example: "abcdef" and "azced". The minimum edit distance between them is 3, because we can transform "abcdef" into "azced" by performing the following three operations:

1.  Replace "b" with "z"
2.  Delete "d"
3.  Replace "f" with "d"

Therefore, the minimum edit distance between "abcdef" and "azced" is 3.

The MED problem is used in a variety of applications, such as spell checking, machine translation, and DNA sequencing, where the goal is to find the most likely match between two sequences of text or genetic code.

Here are the steps to solve the Minimum Edit Distance problem using top-down dynamic programming:

1.  Define a recursive function `MED(i, j)` that takes two indices, `i` and `j`, representing the current position in the two strings being compared.
2.  Define a 2D memoization table `memo` of size `m x n`, where `m` and `n` are the lengths of the two strings being compared.
3.  In the recursive function `MED`, check if the solution for the current position `(i, j)` has already been computed and stored in the memoization table. If yes, return the precomputed solution. Otherwise, compute the solution recursively by considering three possible operations: insertion, deletion, and substitution. Choose the operation with the minimum cost and return its value.
4.  To compute the cost of each operation, use the following rules:
    * Insertion: add 1 to the result of `MED(i, j+1)`
    * Deletion: add 1 to the result of `MED(i+1, j)`
    * Substitution: add 1 to the result of `MED(i+1, j+1)` if the characters at positions `i` and `j` are different, or 0 otherwise.
5.  Finally, call the recursive function `MED(0, 0)` to solve the problem.

Here is the code snippet in Go:

```go
func MED(i, j int, s1, s2 string, memo [][]int) int {
    if i == len(s1) {
        return len(s2) - j
    }
    if j == len(s2) {
        return len(s1) - i
    }
    if memo[i][j] != 0 {
        return memo[i][j]
    }
    cost := 0
    if s1[i] != s2[j] {
        cost = 1
    }
    insert := MED(i, j+1, s1, s2, memo) + 1
    delete := MED(i+1, j, s1, s2, memo) + 1
    substitute := MED(i+1, j+1, s1, s2, memo) + cost
    minCost := insert
    if delete < minCost {
        minCost = delete
    }
    if substitute < minCost {
        minCost = substitute
    }
    memo[i][j] = minCost
    return minCost
}

func minEditDistance(s1, s2 string) int {
    memo := make([][]int, len(s1))
    for i := range memo {
        memo[i] = make([]int, len(s2))
    }
    return MED(0, 0, s1, s2, memo)
}
```
To use the code, simply call the `minEditDistance` function with two strings as arguments, like this:

```go
distance := minEditDistance("kitten", "sitting")
fmt.Println(distance) // Output: 3
```
Note that the code uses memoization to avoid recomputing the same subproblems multiple times, which improves the time complexity of the algorithm from exponential to polynomial.


### 13. Optimal binary search tree
The Optimal Binary Search Tree (OBST) problem is a classic algorithmic problem in computer science. Given a set of keys with their respective frequencies, the goal is to construct a binary search tree that minimizes the expected search cost.

A binary search tree is a tree data structure where each node has at most two children, and the keys in the left subtree of a node are smaller than the key in the node, while the keys in the right subtree are greater. The search cost of a node is defined as the depth of the node in the tree, i.e., the number of edges from the root to the node.

Here's an example to illustrate the OBST problem:

Suppose we have the following set of keys with their respective frequencies:

```
keys = [1, 2, 3, 4, 5]
freqs = [0.15, 0.10, 0.05, 0.10, 0.20]
```
The frequencies represent the probability of each key being searched. For example, `freqs[0]` is the probability of searching for the key `1`.

To construct an optimal binary search tree, we first need to sort the keys in ascending order:

```
keys_sorted = [1, 2, 3, 4, 5]
```
Then, we need to build the tree recursively by dividing the keys into two subtrees, and choosing the root node that minimizes the expected search cost.

Let `OBST(i, j)` denote the expected search cost of a subtree containing the keys from `keys[i]` to `keys[j]`. The base case is when `i > j`, in which case the subtree is empty and the expected search cost is 0.

The recursive formula for `OBST(i, j)` is:
```
OBST(i, j) = min { OBST(i, r-1) + OBST(r+1, j) + sum(freqs[i:j+1]) }
```
where `r` is the index of the root node chosen from the subtree containing the keys from `keys[i]` to `keys[j]`, and `sum(freqs[i:j+1])` is the sum of the frequencies of the keys in the subtree.

We need to try all possible values of `r` between `i` and `j` to find the root node that minimizes the expected search cost.

The final solution to the OBST problem is `OBST(0, n-1)`, where `n` is the number of keys in the set.

In the example above, the optimal binary search tree is:

```
     3
    / \
   2   5
  /   / \
 1   4   -
```
The expected search cost is:

```
0.15*3 + 0.10*2 + 0.05*3 + 0.10*3 + 0.20*2 = 0.80
```

which is the minimum expected search cost for any binary search tree that can be constructed from the set of keys and frequencies.

To solve the OBST problem using top-down dynamic programming, we can use memoization to avoid recomputing the same subproblems multiple times. Here are the steps to implement the algorithm:

1.  Define a two-dimensional memoization table `dp` to store the expected search cost of a subtree containing the keys from `keys[i]` to `keys[j]`. Initialize all values in the table to `-1` to indicate that the subproblems have not been solved yet.
    
2.  Implement a recursive function `obst(i, j)` that takes the indices of the first and last keys in the subtree, and returns the expected search cost of the subtree.
    
3.  In the `obst` function, if `dp[i][j]` is not equal to `-1`, return `dp[i][j]` as the result of the subproblem.
    
4.  If `i > j`, return `0` as the expected search cost of an empty subtree.
    
5.  Otherwise, initialize a variable `min_cost` to a large value, and loop through all possible root nodes `r` between `i` and `j`.
    
6.  Compute the expected search cost of the left subtree by calling `obst(i, r-1)`, and the expected search cost of the right subtree by calling `obst(r+1, j)`.
    
7.  Compute the sum of the frequencies of the keys in the current subtree by calling `sum(freqs[i:j+1])`.
    
8.  Compute the total expected search cost of the current subtree as `left_cost + right_cost + subtree_freq`, where `left_cost` is the expected search cost of the left subtree, `right_cost` is the expected search cost of the right subtree, and `subtree_freq` is the sum of the frequencies of the keys in the current subtree.
    
9.  If the total expected search cost is less than `min_cost`, update `min_cost` with the new value.
    
10. Store the result of the subproblem in `dp[i][j]` and return `min_cost` as the expected search cost of the subtree.
    

Here's the Go code that implements the above algorithm:

```go
func optimalBST(keys []int, freqs []float64) float64 {
    n := len(keys)
    dp := make([][]float64, n)
    for i := range dp {
        dp[i] = make([]float64, n)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    return obst(keys, freqs, dp, 0, n-1)
}

func obst(keys []int, freqs []float64, dp [][]float64, i, j int) float64 {
    if i > j {
        return 0
    }
    if dp[i][j] != -1 {
        return dp[i][j]
    }
    minCost := math.Inf(1)
    for r := i; r <= j; r++ {
        leftCost := obst(keys, freqs, dp, i, r-1)
        rightCost := obst(keys, freqs, dp, r+1, j)
        subtreeFreq := sum(freqs[i:j+1])
        totalCost := leftCost + rightCost + subtreeFreq
        if totalCost < minCost {
            minCost = totalCost
        }
    }
    dp[i][j] = minCost
    return minCost
}

func sum(arr []float64) float64 {
    var s float64
    for _, v := range arr {
        s += v
    }
    return s
}
```
The `optimalBST` function initializes the memoization table and calls the `obst` function to solve the problem. The `obst` function checks if the subproblem has already been solved by checking the memoization table `dp`, and returns the result if it exists. Otherwise, it computes the expected search cost of each possible subtree by recursively calling `obst` on its left and right subtrees, and updates the memoization table with the minimum expected search cost. Finally, it returns the minimum expected search cost of the subtree.

Note that the time complexity of this algorithm is O(n^3), since we need to compute the expected search cost of O(n^2) subtrees, and each computation takes O(n) time. However, the use of memoization reduces the number of redundant computations, making the algorithm practical for small to medium-sized inputs.

Also note that in the code above, we use `math.Inf(1)` to represent infinity, since Go does not have a built-in infinity value for floats.

### 14. Longest bitonic subsequence
The Longest Bitonic Subsequence (LBS) problem is a classic problem in computer science that involves finding the length of the longest subsequence of a given sequence that first increases and then decreases.

To understand the problem better, let's start with an example sequence:

```
2 5 3 9 7 8 6
```
In this sequence, the longest bitonic subsequence is `2 5 9 7 6`, which has a length of 5. The subsequence first increases (`2 5 9`) and then decreases (`9 7 6`). Note that there are other bitonic subsequences in the sequence, such as `2 3 9 8 6` or `2 3 7 8 6`, but these are not as long as `2 5 9 7 6`.

Another example:
```
10 22 9 33 49 50 31 60
```
In this sequence, the longest bitonic subsequence is `10 22 33 50 60 31 9`, which has a length of 7. The subsequence first increases (`10 22 33 50 60`) and then decreases (`60 31 9`).

One more example:
```
4 5 6 3 2 1
```
In this sequence, the longest bitonic subsequence is `4 5 6`, which has a length of 3. The subsequence first increases (`4 5 6`) and then decreases (there are no elements left to consider).

The LBS problem is of great importance in computer science and has many real-world applications, such as in data compression, signal processing, and image recognition. The problem can be solved using dynamic programming, which involves breaking down the problem into smaller subproblems and solving them in a bottom-up manner.

here are the steps to solve the Longest Bitonic Subsequence problem using top-down dynamic programming approach:

1.  Define a recursive function `LBS` that takes the sequence `arr`, the current index `i`, the previous index `j`, and a boolean `isIncreasing` to represent if the sequence is currently increasing or decreasing. The function should return the length of the longest bitonic subsequence starting at index `i` with the previous index `j` and the current direction `isIncreasing`.
    
2.  In the `LBS` function, first check if the result for the current index, previous index, and direction is already computed and stored in a memo table. If so, return the memoized result.
    
3.  If the current index is equal to the length of the sequence, return 0 since there are no more elements left to consider.
    
4.  If the previous index is equal to -1, set the longest bitonic subsequence length to be the maximum of the result of two recursive calls - one where the current element is included in the increasing subsequence and one where the current element is included in the decreasing subsequence.
    
5.  If `isIncreasing` is true, recursively call `LBS` twice - once with the current index included in the increasing subsequence and once with the current index excluded from the increasing subsequence. The result should be the maximum of the two recursive calls.
    
6.  If `isIncreasing` is false, recursively call `LBS` twice - once with the current index included in the decreasing subsequence and once with the current index excluded from the decreasing subsequence. The result should be the maximum of the two recursive calls.
    
7.  Store the result for the current index, previous index, and direction in the memo table and return the result.
    
8.  Call the `LBS` function with the sequence `arr`, current index 0, previous index -1, and direction set to true (since we start with an increasing subsequence).
    
9.  The returned value is the length of the longest bitonic subsequence.
    

Here's the implementation of the above steps in Go:

```go
package main

import (
    "fmt"
)

func main() {
    arr := []int{2, 5, 3, 9, 7, 8, 6}
    length := LBS(arr)
    fmt.Println(length) // Output: 5
}

func LBS(arr []int) int {
    memo := make(map[string]int)
    return max(LBSHelper(arr, 0, -1, true, memo), LBSHelper(arr, 0, -1, false, memo))
}

func LBSHelper(arr []int, i, j int, isIncreasing bool, memo map[string]int) int {
    if j >= 0 {
        key := fmt.Sprintf("%d-%d-%t", i, j, isIncreasing)
        if val, ok := memo[key]; ok {
            return val
        }
    }
    if i == len(arr) {
        return 0
    }
    if j == -1 {
        return max(LBSHelper(arr, i+1, i, true, memo)+1, LBSHelper(arr, i+1, i, false, memo)+1)
    }
    if isIncreasing {
        res := max(LBSHelper(arr, i+1, j, true, memo), LBSHelper(arr, i+1, j, false, memo))
        if arr[i] > arr[j] {
            res = max(res, LBSHelper(arr, i+1, i, false, memo)+1)
        }
        if j == -1 {
            res = max(res, LBSHelper(arr, i+1, i, true, memo)+1)
        }
        if j == -1 || arr[i] > arr[j] {
            memo[fmt.Sprintf("%d-%d-%t", i, j, isIncreasing)] = res
        }
        return res
    } else {
        res := LBSHelper(arr, i+1, j, false, memo)
        if arr[i] < arr[j] {
            res = max(res, LBSHelper(arr, i+1, i, true, memo)+1)
        }
        if j == -1 || arr[i] < arr[j] {
            memo[fmt.Sprintf("%d-%d-%t", i, j, isIncreasing)] = res
        }
        return res
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
In this implementation, the `LBS` function initializes the memo table and calls the `LBSHelper` function with the sequence, current index 0, previous index -1, and direction set to true (since we start with an increasing subsequence).

The `LBSHelper` function is the recursive function that computes the length of the longest bitonic subsequence for a given starting index, previous index, and direction. The function uses memoization to avoid recomputing results for previously solved subproblems.

The `max` function is a helper function that returns the maximum of two integers.


### 15. Longest common substring

The Longest Common Substring (LCS) problem is a classic computer science problem that involves finding the longest common substring that is present in two or more given strings. A substring is a contiguous sequence of characters within a string.

For example, consider the two strings "ABCDEF" and "BCDEFG". The longest common substring between the two strings is "BCDEF" which has a length of 5.

Let's take another example to explain the LCS problem in more detail. Consider the following two strings:

String 1: "ABCDGH" String 2: "AEDFHR"

The longest common substring between these two strings is "ADH", which has a length of 3. Note that "A", "D", and "H" are all contiguous in both strings, which is what makes them a common substring.

Here's another example:

String 1: "Old MacDonald had a farm" String 2: "E-I-E-I-O"

The longest common substring between these two strings is "O", which appears twice in both strings. Note that the substrings "I", "E", and "A" are also common between the two strings, but they are not as long as "O".

The LCS problem can also be extended to more than two strings. For example:

String 1: "ABCD" String 2: "ACDF" String 3: "ACEF"

The longest common substring between these three strings is "A", which appears in all three strings.

In summary, the LCS problem involves finding the longest common substring that is present in two or more given strings. This problem has many applications in areas such as bioinformatics, text processing, and data compression.

To solve the Longest Common Substring (LCS) problem using top-down dynamic programming (DP), we can follow the following steps:

1.  Define the base case: If either of the strings is empty, the LCS length is zero.
    
2.  Define the state: We can define the state as a pair of indices `(i, j)` representing the position in the two strings.
    
3.  Define the recursive function: The recursive function will take two indices as input, and it will return the length of the LCS for the given indices. The function should check if the current characters at the two indices match, and recursively call itself for the next indices with one added to the LCS length.
    
4.  Add memoization: To avoid recomputing the same subproblems, we can use memoization to store the LCS length for each pair of indices.
    
5.  Find the LCS: After computing the memo table, we can backtrack from the maximum value in the table to find the LCS itself.
    

Here's an implementation of the above algorithm in Go:
```go
func lcsTopDown(s1, s2 string) string {
    memo := make(map[[2]int]int) // memo table

    // recursive function to compute LCS length
    var lcsLength func(int, int) int
    lcsLength = func(i, j int) int {
        if i == len(s1) || j == len(s2) {
            return 0 // base case
        }
        if val, ok := memo[[2]int{i, j}]; ok {
            return val // memoization
        }
        if s1[i] == s2[j] {
            memo[[2]int{i, j}] = 1 + lcsLength(i+1, j+1)
        } else {
            memo[[2]int{i, j}] = max(lcsLength(i+1, j), lcsLength(i, j+1))
        }
        return memo[[2]int{i, j}]
    }

    // compute LCS length
    length := lcsLength(0, 0)

    // backtrack to find LCS
    lcs := make([]byte, length)
    i, j := 0, 0
    for length > 0 && i < len(s1) && j < len(s2) {
        if s1[i] == s2[j] {
            lcs[length-1] = s1[i]
            i++
            j++
            length--
        } else if memo[[2]int{i+1, j}] > memo[[2]int{i, j+1}] {
            i++
        } else {
            j++
        }
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
To use this function, simply call `lcsTopDown(s1, s2)` with the two strings as input. The function will return the LCS as a string.


### 16. Maximum sum increasing subsequence
The Maximum Sum Increasing Subsequence (MSIS) problem is a classic computer science problem that involves finding the increasing subsequence of a given sequence with the maximum possible sum. A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

For example, consider the following sequence of numbers:

Sequence: 2, 4, 6, 1, 3, 8, 5

The increasing subsequences of this sequence are:

* 2, 4, 6, 8
* 1, 3, 5

The maximum sum increasing subsequence of this sequence is 2 + 4 + 6 + 8 = 20. Note that we have selected the first increasing subsequence as it has the maximum sum.

Let's take another example to explain the MSIS problem in more detail. Consider the following sequence of numbers:

Sequence: 4, 6, 1, 3, 8, 4, 5, 7, 9

The increasing subsequences of this sequence are:

* 4, 6, 8, 9
* 1, 3, 4, 5, 7, 9

The maximum sum increasing subsequence of this sequence is 1 + 3 + 4 + 5 + 7 + 9 = 29. Note that we have selected the second increasing subsequence as it has the maximum sum.

Here's another example:

Sequence: 1, 2, 3, 4, 5

The increasing subsequences of this sequence are:

* 1, 2, 3, 4, 5

The maximum sum increasing subsequence of this sequence is the sequence itself, i.e., 1 + 2 + 3 + 4 + 5 = 15.

In summary, the MSIS problem involves finding the increasing subsequence of a given sequence with the maximum possible sum. This problem has many applications in areas such as finance, scheduling, and data analysis.
To solve the Maximum Sum Increasing Subsequence (MSIS) problem using top-down dynamic programming (DP), we can follow the following steps:

1.  Define the base case: If the current index is out of bounds, return 0.
    
2.  Define the state: We can define the state as a pair of indices `(i, prev)` representing the current index and the index of the previous element in the increasing subsequence.
    
3.  Define the recursive function: The recursive function will take the current index and the index of the previous element as input, and it will return the maximum sum of an increasing subsequence that starts from the current index. The function should check if the current element is greater than the previous element, and recursively call itself for the next index with the current element added to the maximum sum.
    
4.  Add memoization: To avoid recomputing the same subproblems, we can use memoization to store the maximum sum for each pair of indices.
    
5.  Find the MSIS: After computing the memo table, we can backtrack from the maximum value in the table to find the MSIS itself.
    

Here's an implementation of the above algorithm in Go:

```go
func msisTopDown(arr []int) []int {
    memo := make(map[[2]int]int) // memo table

    // recursive function to compute MSIS sum
    var msisSum func(int, int) int
    msisSum = func(i, prev int) int {
        if i == len(arr) {
            return 0 // base case
        }
        if val, ok := memo[[2]int{i, prev}]; ok {
            return val // memoization
        }
        sumWithout := msisSum(i+1, prev)
        sumWith := 0
        if arr[i] > arr[prev] {
            sumWith = arr[i] + msisSum(i+1, i)
        }
        memo[[2]int{i, prev}] = max(sumWithout, sumWith)
        return memo[[2]int{i, prev}]
    }

    // compute MSIS sum
    sum := msisSum(0, 0)

    // backtrack to find MSIS
    msis := []int{}
    i, prev := 0, 0
    for i < len(arr) {
        if arr[i] > arr[prev] && msisSum(i, prev) == sum {
            msis = append(msis, arr[i])
            sum -= arr[i]
            prev = i
        }
        i++
    }
    return msis
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
To use this function, simply call `msisTopDown(arr)` with the array of integers as input. The function will return the MSIS as a slice of integers.

### 17. Painting fence problem
The Painting Fence problem is a classic problem in computer science that involves dividing a long fence into sections and painting each section a specific color. There are many variations of this problem, each with different constraints and requirements.

Here are some examples of the Painting Fence problem with different constraints:

1.  Simple Painting Fence Problem:

Suppose you have a fence with N sections, and you want to paint each section a different color. You have K colors available to use. You can use each color any number of times, but no two adjacent sections can be the same color. In how many ways can you paint the fence?

For example, if N = 3 and K = 2, the possible color combinations are:

* Red, Blue, Red
* Red, Blue, Green
* Blue, Red, Blue
* Blue, Red, Green
* Green, Blue, Red
* Green, Blue, Green
* Red, Green, Blue
* Green, Red, Blue
* Blue, Green, Red
* Blue, Green, Blue
* Green, Red, Green
* Red, Green, Red

So there are 12 possible ways to paint the fence.

2.  Painting Fence with Maximum Repetitions:

Suppose you have a fence with N sections, and you want to paint each section a specific color. You have K colors available to use, and each color can be used any number of times. However, you want to maximize the number of consecutive sections that are painted the same color. In how many ways can you paint the fence?

For example, if N = 5 and K = 3, the possible color combinations are:

* Red, Red, Red, Red, Red
* Red, Red, Red, Red, Blue
* Red, Red, Red, Blue, Red
* Red, Red, Blue, Red, Red
* Red, Blue, Red, Red, Red
* Blue, Red, Red, Red, Red
* Red, Red, Blue, Blue, Blue
* Red, Blue, Blue, Blue, Red
* Blue, Blue, Blue, Red, Red
* Blue, Blue, Red, Red, Blue

So there are 10 possible ways to paint the fence with the maximum number of consecutive sections painted the same color.

3.  Painting Fence with Minimum Cost:

Suppose you have a fence with N sections, and you want to paint each section a specific color. Each color has a cost associated with it. You have a budget B to spend on painting the fence, and you want to minimize the total cost while still painting each section a specific color. In how many ways can you paint the fence?

For example, if N = 4, K = 3, and the cost of each color is:

* Red: 3
* Blue: 2
* Green: 4

And the budget is B = 8, the possible color combinations are:

* Red, Blue, Red, Blue (cost: 7)
* Red, Blue, Green, Blue (cost: 9)
* Red, Green, Blue, Green (cost: 13)
* Blue, Red, Blue, Green (cost: 9)
* Blue, Green, Blue, Red (cost: 11)
* Green, Red, Blue, Red (cost: 12)

So the optimal way to paint the fence with a budget of 8 is to paint it Red, Blue, Red, Blue, with a total cost of 7.

here are the steps to solve the Painting Fence problem using top-down dynamic programming, along with code snippets in Go:

**Step 1**: Define the recursive function

The first step is to define a recursive function that will solve the problem. The function should take as input the current section of the fence to be painted, the previous color used (if any), and the number of colors remaining to use. The function should return the number of ways to paint the remaining sections of the fence.

Here's the code for the recursive function:

```go
func paintFence(n, k int, prevColor int, remainingColors int, memo map[[3]int]int) int {
    // Base case: we've painted all the sections of the fence
    if n == 0 {
        return 1
    }
    
    // Check if we've already computed this value
    key := [3]int{n, prevColor, remainingColors}
    if val, ok := memo[key]; ok {
        return val
    }
    
    // Recursive case: paint the current section of the fence
    count := 0
    for i := 0; i < k; i++ {
        if i != prevColor && remainingColors > 0 {
            count += paintFence(n-1, k, i, remainingColors-1, memo)
        }
    }
    
    // Store the computed value in the memoization table
    memo[key] = count
    
    return count
}
```
**Step 2**: Initialize the memoization table

Next, we need to initialize the memoization table with default values. The memoization table will store the results of previous function calls so that we can avoid recomputing the same values multiple times.

Here's the code to initialize the memoization table:

```go
func numWaysToPaintFence(n, k int) int {
    memo := make(map[[3]int]int)
    return paintFence(n, k, -1, k, memo)
}
```
**Step 3**: Call the recursive function

Finally, we need to call the recursive function with the appropriate parameters to solve the problem.

Here's the complete code to solve the Painting Fence problem using top-down dynamic programming:

```go
func paintFence(n, k int, prevColor int, remainingColors int, memo map[[3]int]int) int {
    // Base case: we've painted all the sections of the fence
    if n == 0 {
        return 1
    }
    
    // Check if we've already computed this value
    key := [3]int{n, prevColor, remainingColors}
    if val, ok := memo[key]; ok {
        return val
    }
    
    // Recursive case: paint the current section of the fence
    count := 0
    for i := 0; i < k; i++ {
        if i != prevColor && remainingColors > 0 {
            count += paintFence(n-1, k, i, remainingColors-1, memo)
        }
    }
    
    // Store the computed value in the memoization table
    memo[key] = count
    
    return count
}

func numWaysToPaintFence(n, k int) int {
    memo := make(map[[3]int]int)
    return paintFence(n, k, -1, k, memo)
}
```
To use the code, simply call the `numWaysToPaintFence` function with the number of sections `n` and the number of colors `k`. The function will return the number of ways to paint the fence.


### 18. Maximum sum subsequence
The Maximum sum subsequence problem is a classic computer science problem that involves finding a subsequence of an array or list of integers that has the largest possible sum.

In other words, given an array of integers, the goal is to find the contiguous subarray with the maximum sum. This problem has various applications, such as in financial analysis, where it can be used to calculate the maximum profit that can be obtained from a series of stock prices.

Here are some examples to illustrate the problem:

Example 1: Consider the array \[−2, 1, −3, 4, −1, 2, 1, −5, 4\]. The maximum sum subsequence is \[4, −1, 2, 1\], which has a sum of 6.

Example 2: Consider the array \[1, 2, 3, -4, 5\]. The maximum sum subsequence is \[1, 2, 3, -4, 5\], which has a sum of 7.

Example 3: Consider the array \[-2, -3, 4, -1, -2, 1, 5, -3\]. The maximum sum subsequence is \[4, -1, -2, 1, 5\], which has a sum of 7.

Example 4: Consider the array \[2, -1, 2, 3, 4, -5\]. The maximum sum subsequence is \[2, -1, 2, 3, 4\], which has a sum of 10.

Example 5: Consider the array \[−1, −2, −3\]. The maximum sum subsequence is the empty subsequence \[\], which has a sum of 0.

To solve the Maximum sum subsequence problem, we can use various algorithms, such as the Kadane's algorithm or the dynamic programming approach. These algorithms have different time and space complexities and are suitable for different problem sizes and constraints.

Here are the steps to solve the Maximum sum subsequence problem using the top-down dynamic programming approach:

1.  Define the recursive function: The first step is to define a recursive function that takes an array and the index of the last element included in the subsequence as parameters. The function should return the maximum sum subsequence that can be obtained from the subarray starting from index 0 up to the index of the last element included.
    
2.  Define the base case: If the index is less than 0, return 0 as the sum of an empty subsequence is 0.
    
3.  Check if the subproblem has already been solved: If the sum for the current subsequence has already been computed, return the stored value.
    
4.  Define the recursive relation: The recursive function should consider two cases: include the current element in the subsequence or exclude it. If the current element is included, add its value to the maximum sum subsequence from the previous subarray. If the current element is excluded, consider the maximum sum subsequence from the previous subarray.
    
5.  Store the result: Store the result of the current subproblem in a memoization table.
    
6.  Return the maximum sum subsequence: Once all subproblems have been solved, return the maximum sum subsequence obtained from the entire array.
    

Here's the Go code that implements the above steps:
```go
func maxSumSubsequence(nums []int) int {
    memo := make(map[int]int)
    return maxSumSubsequenceRec(nums, len(nums)-1, memo)
}

func maxSumSubsequenceRec(nums []int, lastIdx int, memo map[int]int) int {
    if lastIdx < 0 {
        return 0
    }
    
    if val, ok := memo[lastIdx]; ok {
        return val
    }
    
    includeLast := nums[lastIdx] + maxSumSubsequenceRec(nums, lastIdx-1, memo)
    excludeLast := maxSumSubsequenceRec(nums, lastIdx-1, memo)
    memo[lastIdx] = max(includeLast, excludeLast)
    
    return memo[lastIdx]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
The above code defines the `maxSumSubsequence` function that initializes the memoization table and calls the recursive function. The `maxSumSubsequenceRec` function implements the recursive function that computes the maximum sum subsequence using memoization. Finally, the `max` function returns the maximum value between two integers.

### 19. 0/1 Knapsack problem
The 0/1 Knapsack problem is a classic problem in computer science and optimization. It is a combinatorial optimization problem, which involves selecting a set of items with maximum value, while staying within a fixed weight capacity. The name "0/1" refers to the fact that each item is either included in the knapsack (assigned a value of 1), or not (assigned a value of 0).

To better understand the problem, let's consider an example. Suppose you are a thief who has just broken into a jewelry store. You have a knapsack with a maximum weight capacity of 10 kg, and you have to select the most valuable items to steal. The items available are:

| Item | Value ($) | Weight (kg) |
| --- | --- | --- |
| Diamond | 60  | 5   |
| Gold | 30  | 3   |
| Silver | 20  | 4   |
| Platinum | 40  | 6   |
| Ruby | 50  | 7   |

You want to maximize the value of the items you can steal, while ensuring that the weight of the items you select does not exceed the knapsack's capacity of 10 kg.

The problem can be formulated mathematically as follows:

Maximize: Value = 60x1 + 30x2 + 20x3 + 40x4 + 50x5

Subject to: 5x1 + 3x2 + 4x3 + 6x4 + 7x5 <= 10 (Knapsack capacity) x1, x2, x3, x4, x5 ∈ {0, 1} (Item selection)

where x1, x2, x3, x4, x5 are decision variables that indicate whether each item is selected (x = 1) or not (x = 0).

The solution to this problem is to select the Diamond, Gold, and Platinum items, which have a total value of 60 + 30 + 40 = $130 and a total weight of 5 + 3 + 6 = 14 kg. The Silver and Ruby items are not selected because their combined weight exceeds the capacity of the knapsack.

This is just one example of the 0/1 Knapsack problem. There are many real-world applications of this problem, such as resource allocation in project management, selection of investment portfolios, and packing of containers in logistics.

The 0/1 Knapsack problem can be solved using dynamic programming. The top-down approach is also known as memoization, which involves storing the results of subproblems to avoid recomputation.

Here are the steps to solve the 0/1 Knapsack problem using top-down DP:

Step 1: Define a recursive function to compute the maximum value that can be obtained for a given weight capacity and set of items.

Step 2: Check if the result for this combination of weight capacity and set of items has already been computed and stored in a memoization table. If so, return the stored result.

Step 3: If the weight capacity is 0 or the set of items is empty, return 0.

Step 4: If the weight of the last item in the set is greater than the weight capacity, exclude it from the set and compute the maximum value for the remaining items and weight capacity.

Step 5: Otherwise, compute the maximum value by either excluding the last item and computing the maximum value for the remaining items and weight capacity, or by including the last item and computing the maximum value for the remaining items and the remaining weight capacity after including the last item. Take the maximum of these two values.

Step 6: Store the computed result in the memoization table and return the result.

Step 7: Call the recursive function with the initial weight capacity and set of items, and return the maximum value.

Here is the code in Go:

```go
func knapsack(w []int, v []int, n int, capacity int, memo map[[2]int]int) int {
    if res, ok := memo[[2]int{n, capacity}]; ok {
        return res
    }
    if capacity == 0 || n == 0 {
        memo[[2]int{n, capacity}] = 0
        return 0
    }
    if w[n-1] > capacity {
        res := knapsack(w, v, n-1, capacity, memo)
        memo[[2]int{n, capacity}] = res
        return res
    }
    res := max(knapsack(w, v, n-1, capacity, memo), v[n-1]+knapsack(w, v, n-1, capacity-w[n-1], memo))
    memo[[2]int{n, capacity}] = res
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    w := []int{5, 3, 4, 6, 7}
    v := []int{60, 30, 20, 40, 50}
    capacity := 10
    memo := make(map[[2]int]int)
    fmt.Println(knapsack(w, v, len(w), capacity, memo))
}
```
In this code, `w` and `v` are slices containing the weights and values of the items, `n` is the number of items, `capacity` is the maximum weight capacity of the knapsack, and `memo` is a memoization table that stores the results of subproblems.

The `knapsack` function computes the maximum value that can be obtained for a given weight capacity and set of items using the recursive approach described above. The `max` function returns the maximum of two integers.

The `main` function calls the `knapsack` function with the initial weight capacity and set of items, and prints the maximum value.

### 20. Egg dropping problem

The Egg dropping problem is a classic problem in computer science and mathematics that involves finding the highest floor from which an egg can be dropped without breaking, given a certain number of eggs and a certain number of floors.

The problem is often stated as follows:

Suppose you have k eggs and you want to determine the highest floor from which you can drop an egg without it breaking. You have access to a building with n floors, and you can drop the eggs from any floor. If an egg breaks when dropped from a certain floor, it will also break when dropped from any higher floor. However, if an egg does not break when dropped from a certain floor, it will also not break when dropped from any lower floor.

The goal is to determine the minimum number of drops you need to make in order to find the highest floor from which the egg can be dropped without breaking.

Here are some examples to illustrate the problem:

**Example 1:** Suppose you have 2 eggs and 10 floors. You need to find the highest floor from which an egg can be dropped without breaking. You can drop the eggs from any floor, and once an egg breaks, you cannot use it again. What is the minimum number of drops you need to make?

Solution: One possible strategy is to start by dropping one egg from the first floor, then dropping it from the second floor, then the third floor, and so on, until it breaks. This would require a maximum of 10 drops in the worst case.

However, a more efficient strategy is to start by dropping one egg from the fifth floor. If it breaks, you can use the second egg to test the floors below the fifth floor. If it doesn't break, you can use the second egg to test the floors above the fifth floor. This strategy would require a maximum of 6 drops in the worst case.

**Example 2**: Suppose you have 3 eggs and 100 floors. You need to find the highest floor from which an egg can be dropped without breaking. What is the minimum number of drops you need to make?

Solution: One possible strategy is to start by dropping an egg from the 50th floor. If it breaks, you can use the remaining two eggs to test the floors below the 50th floor. If it doesn't break, you can use the remaining two eggs to test the floors above the 50th floor. If the second egg breaks on the 75th floor, you can use the third egg to test the floors between the 51st and 74th floors. This strategy would require a maximum of 14 drops in the worst case.

Example 3: Suppose you have 1 egg and 100 floors. You need to find the highest floor from which an egg can be dropped without breaking. What is the minimum number of drops you need to make?

Solution: In this case, you need to drop the egg from each floor one by one until it breaks, starting from the first floor. This strategy would require a maximum of 100 drops in the worst case.

These examples demonstrate the importance of developing an efficient strategy for solving the egg dropping problem, especially when dealing with a large number of floors or a limited number of eggs.

here are the steps to solve the Egg Dropping Problem using Top-Down Dynamic Programming (Memoization) along with code snippets in Go:

**Step 1**: Define a 2D memoization array dp, where dp\[i\]\[j\] represents the minimum number of drops needed to determine the critical floor with i eggs and j floors.

**Step 2**: Initialize the memoization array dp with a maximum value or -1, to indicate that the corresponding value has not yet been computed.

**Step 3**: Define a recursive function dropEggs(nEggs, nFloors int) int that takes the number of eggs and number of floors and returns the minimum number of drops needed to determine the critical floor.

**Step 4**: Check if the value of dp\[nEggs\]\[nFloors\] has already been computed. If it has, return the value.

**Step 5**: If nEggs is 1, then the problem reduces to finding the critical floor using a linear search. So, return nFloors.

**Step 6**: If nFloors is 0, then no drops are needed. So, return 0.

**Step 7**: Initialize minDrops variable to a large number.

**Step 8**: Use a loop to try dropping an egg from each floor from 1 to nFloors.

**Step 9**: If the egg breaks, then recurse with (nEggs - 1, floor - 1) as we have one less egg and we need to check the floors below the current floor.

**Step 10**: If the egg does not break, then recurse with (nEggs, nFloors - floor) as we still have the same number of eggs and we need to check the floors above the current floor.

**Step 11**: Compute the minimum number of drops needed by taking the maximum of the drops needed in the two cases and adding 1 for the current drop.

**Step 12**: Update the value of dp\[nEggs\]\[nFloors\] with the computed minimum number of drops.

**Step 13**: Return the computed minimum number of drops.

Here's the code snippet in Go:

```go
package main

import "fmt"

const MAX_VALUE = 1000000

var dp [][]int

func dropEggs(nEggs, nFloors int) int {
    if nEggs == 1 {
        return nFloors
    }
    if nFloors == 0 {
        return 0
    }
    if dp[nEggs][nFloors] != -1 {
        return dp[nEggs][nFloors]
    }
    minDrops := MAX_VALUE
    for floor := 1; floor <= nFloors; floor++ {
        drops := 1 + max(dropEggs(nEggs-1, floor-1), dropEggs(nEggs, nFloors-floor))
        if drops < minDrops {
            minDrops = drops
        }
    }
    dp[nEggs][nFloors] = minDrops
    return minDrops
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func eggDrop(nEggs, nFloors int) int {
    dp = make([][]int, nEggs+1)
    for i := 0; i <= nEggs; i++ {
        dp[i] = make([]int, nFloors+1)
        for j := 0; j <= nFloors; j++ {
            dp[i][j] = -1
        }
    }
    return dropEggs(nEggs, nFloors)
}

func main() {
    nEggs := 2
    nFloors := 10
    minDrops := eggDrop(nEggs, nFloors)
    fmt.Printf("Minimum number of drops needed for %d eggs and %d floors is %d\n", nEggs, nFloors, minDrops)
}
```
In this example, the eggDrop function is called with nEggs=2 and nFloors=10. The minimum number of drops needed to determine the critical floor with 2 eggs and 10 floors is then printed to the console. You can change the values of nEggs and nFloors to test the function with different input values.



### 21. Partition problem
The Partition problem is a classical problem in computer science and mathematics, which involves partitioning a given set of integers into two subsets, such that the sum of elements in each subset is equal.

Here are some examples that illustrate the Partition problem:

**Example 1**: Consider the set of integers {1, 5, 11, 5}. The task is to partition this set into two subsets such that the sum of elements in each subset is equal. One possible solution is to divide the set into {1, 5, 5} and {11}, which gives two subsets with a sum of 11 each.

**Example 2**: Let's take another example with the set {3, 1, 5, 9, 12}. We need to find if it is possible to divide this set into two subsets with the same sum. There are different ways to partition this set, but one possible solution is {3, 5, 9} and {1, 12}. Both subsets have a sum of 17.

**Example 3**: Consider a set of integers {7, 3, 1, 5, 4, 8}. We need to determine if it is possible to partition this set into two subsets with the same sum. One possible solution is {7, 3, 5} and {1, 4, 8}. Both subsets have a sum of 15.

**Example 4**: Let's consider another set of integers {1, 2, 3, 4, 5, 6, 7}. In this case, it is not possible to partition the set into two subsets with the same sum. This can be proven mathematically, but intuitively, we can see that the sum of all the elements in the set is 28, which is an even number. Since we need to divide the set into two subsets with the same sum, each subset should have a sum of 14. However, it is not possible to choose a subset of integers with a sum of 14 from the given set.

The Partition problem is an **NP-complete** problem, which means that there is no known polynomial-time algorithm that can solve it for all possible inputs. However, there are heuristic approaches and approximation algorithms that can provide good solutions in practice for many instances of the problem.

The Partition problem can be solved using dynamic programming (DP) in a top-down approach. Here are the steps to solve the problem using top-down DP:

Step 1: Define a recursive function that takes the following arguments:

* An array of integers
* The length of the array
* The current index (i.e., the index of the element that we are currently considering)
* The sum of the first subset that we are trying to create
* A memoization table to store the results of previous computations

The function should return true if it is possible to partition the array into two subsets with equal sums, and false otherwise.

Step 2: Inside the recursive function, check if we have already computed the result for the current state (i.e., the current index and the sum of the first subset). If yes, return the stored result from the memoization table.

Step 3: If the current index is equal to the length of the array, check if the sum of the first subset is equal to half the sum of all the elements in the array. If yes, return true; otherwise, return false.

Step 4: Otherwise, we have two choices: we can either include the current element in the first subset or exclude it. If we include the current element, the sum of the first subset increases by the value of the current element. If we exclude it, the sum of the first subset remains the same.

Step 5: Recursively call the function twice for each choice (include or exclude the current element), with the updated index and sum of the first subset.

Step 6: Store the result of the current state (i.e., the current index and sum of the first subset) in the memoization table and return the final result.

Here's the code snippet in Go:

```go
func canPartition(nums []int) bool {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    if totalSum%2 != 0 {
        return false
    }
    memo := make([][]int8, len(nums))
    for i := range memo {
        memo[i] = make([]int8, totalSum/2+1)
    }
    return canPartitionHelper(nums, 0, 0, totalSum/2, memo)
}

func canPartitionHelper(nums []int, index int, sum int, target int, memo [][]int8) bool {
    if index == len(nums) {
        return sum == target
    }
    if memo[index][sum] != 0 {
        return memo[index][sum] == 1
    }
    included := false
    if sum+nums[index] <= target {
        included = canPartitionHelper(nums, index+1, sum+nums[index], target, memo)
    }
    excluded := canPartitionHelper(nums, index+1, sum, target, memo)
    result := included || excluded
    if result {
        memo[index][sum] = 1
    } else {
        memo[index][sum] = -1
    }
    return result
}
```
In this code, we first check if the total sum of the elements in the array is even. If not, we return false because it is not possible to partition the array into two subsets with equal sums. We then create a memoization table to store the results of previous computations.

The `canPartitionHelper` function is the recursive function that takes the array of integers, the current index, the sum of the first subset, the target sum (i.e., half the sum of all the elements in the array), and the memoization table as arguments. The function returns true if it is possible to partition the array into two subsets with equal sums, and false otherwise.

Inside the `canPartitionHelper` function, we first check if we have already computed the result for the current state (i.e., the current index and the sum of the first subset) by checking the memoization table. If yes, we return the stored result. If not, we proceed with the recursive computation.

We have two choices at each step: we can either include the current element in the first subset or exclude it. If we include the current element, we check if the sum of the first subset exceeds the target sum. If it does not, we recursively call the function with the updated index and sum of the first subset. If we exclude the current element, we recursively call the function with the updated index and the same sum of the first subset.

After the recursive calls, we store the result of the current state (i.e., the current index and sum of the first subset) in the memoization table and return the final result.

The `canPartition` function is a wrapper function that first checks if it is possible to partition the array into two subsets with equal sums by computing the total sum of the elements in the array and checking if it is even. If not, it returns false. Otherwise, it calls the `canPartitionHelper` function with the array of integers, the initial index (i.e., 0), the initial sum of the first subset (i.e., 0), the target sum (i.e., half the sum of all the elements in the array), and the memoization table. It returns the result of the `canPartitionHelper` function.


### 22. Longest zigzag subsequence
The Longest Zigzag Subsequence problem is a classic computer science problem that involves finding the length of the longest subsequence of a given sequence such that the subsequence alternates between increasing and decreasing values.

For example, consider the following sequence of integers:
```
1 7 4 9 2 5
```
One possible zigzag subsequence of this sequence is:
```
1 7 4 9 2 5
```
This subsequence alternates between increasing and decreasing values, with 1< 7 > 4 < 9 > 2 < 5.

Another possible zigzag subsequence of the same sequence is:
```
1 7 4 5
```
This subsequence also alternates between increasing and decreasing values, with 1< 7 > 4 < 5.

The longest zigzag subsequence of this sequence is:
```
1 7 2 5
```
This subsequence has a length of 4 and alternates between increasing and decreasing values, with 1 < 7 > 2 < 5.

Here is another example sequence:
```
10 22 9 33 49 50 31 60
```
One possible zigzag subsequence of this sequence is:
```
10 22 9 33 31 60
```
This subsequence alternates between increasing and decreasing values, with 10 < 22 > 9 < 33 > 31 < 60.

Another possible zigzag subsequence of the same sequence is:
```
10 22 9 50 31
```
This subsequence also alternates between increasing and decreasing values, with 10 <22 > 9 < 50 > 31.

The longest zigzag subsequence of this sequence is:
```
10 22 9 49 31 60
```
This subsequence has a length of 6 and alternates between increasing and decreasing values, with 10 < 22 > 9 < 49 > 31 < 60.

The Longest Zigzag Subsequence problem is to find the length of the longest zigzag subsequence of a given sequence. It is a challenging problem, but it can be solved using dynamic programming techniques.

Here are the steps to solve the Longest Zigzag Subsequence problem using top-down dynamic programming:

**Step 1**: Define a memoization array to store the results of previously computed subproblems. The size of the array should be equal to the length of the input sequence.

**Step 2**: Define a recursive function that takes in the following parameters:

* The input sequence
* The current index in the sequence
* The previous value in the subsequence
* A boolean flag indicating whether the current value is greater than the previous value or not
* The memoization array

The function should do the following:

* If the current index is greater than or equal to the length of the sequence, return 0
* If the result for the current subproblem is already computed and stored in the memoization array, return the stored result
* Otherwise, compute the result for the current subproblem by considering two cases:
    * If the current value is greater than the previous value, find the length of the longest zigzag subsequence starting from the next index where the current value is smaller than the previous value
    * If the current value is smaller than the previous value, find the length of the longest zigzag subsequence starting from the next index where the current value is greater than the previous value
* Store the computed result in the memoization array and return it.

**Step 3**: Initialize the memoization array with a value of -1 to indicate that the results for all subproblems are yet to be computed.

**Step 4**: Call the recursive function with the following parameters:

* The input sequence
* The current index (0)
* The previous value (the first value in the sequence)
* A boolean flag indicating that the current value is greater than the previous value (true)
* The memoization array

**Step 5**: Return the result obtained from the recursive function call.

Here's the code snippet in Go:
```go
func longestZigzagSubsequence(seq []int, index int, prev int, isIncreasing bool, memo []int) int {
    if index >= len(seq) {
        return 0
    }

    if memo[index] != -1 {
        return memo[index]
    }

    var maxLength int
    if isIncreasing {
        for i := index + 1; i < len(seq); i++ {
            if seq[i] < prev {
                length := 1 + longestZigzagSubsequence(seq, i, seq[i], !isIncreasing, memo)
                if length > maxLength {
                    maxLength = length
                }
            }
        }
    } else {
        for i := index + 1; i < len(seq); i++ {
            if seq[i] > prev {
                length := 1 + longestZigzagSubsequence(seq, i, seq[i], !isIncreasing, memo)
                if length > maxLength {
                    maxLength = length
                }
            }
        }
    }

    memo[index] = maxLength
    return maxLength
}

func longestZigzagSubsequenceLength(seq []int) int {
    memo := make([]int, len(seq))
    for i := range memo {
        memo[i] = -1
    }

    return 1 + longestZigzagSubsequence(seq, 1, seq[0], true, memo)
}
```
In this implementation, we add 1 to the final result since the first element of the zigzag subsequence is always included.


### 23. String matching problems
The String matching problem is a classic algorithmic problem in computer science that involves finding the occurrence of a pattern within a given text string. It is commonly solved using dynamic programming (DP) techniques.

The problem can be stated as follows: Given a text string `T` and a pattern string `P`, find all occurrences of `P` within `T`. For example, consider the following text string and pattern:

```
T = "ababababa"
P = "aba"
```
There are four occurrences of the pattern `aba` within the text `ababababa`: `ab[aba]baba`, `aba[bab]aba`, `abab[aba]ba`, and `ababab[aba]`.

One approach to solving this problem using dynamic programming is to construct a table of boolean values `dp[i][j]` that indicates whether the substring `T[1...i]` contains the pattern `P[1...j]`. The table can be filled in using the following recursive formula:
```
dp[i][j] = true if dp[i-1][j-1] and T[i] == P[j]
dp[i][j] = false otherwise
```
The base cases are `dp[0][0] = true` (an empty pattern matches an empty text), `dp[0][j] = false` for `j > 0` (an empty text cannot match a non-empty pattern), and `dp[i][0] = true` for `i > 0` (any non-empty text matches an empty pattern).

Once the table is filled in, the occurrences of the pattern within the text can be found by scanning the last row of the table and identifying all `j` such that `dp[n][j]` is true, where `n` is the length of the text.

Here are some more examples:

**Example 1:**
```
T = "hello world"
P = "world"
```
The table would be as follows:
```
T   h   e   l   l   o       w   o   r   l   d
P   w   o   r   l   d
    F   F   F   F   F   F   T   F   F   F   F
```
The pattern `world` occurs once in the text `hello world`.
**Example 2:**
```
T = "abacabadabacaba"
P = "aba"
```
The table would be as follows:

```
T   a   b   a   c   a   b   a   d   a   b   a   c   a   b   a
P   a   b   a
    T   F   T   F   T   F   T   F   T   F   T   F   T   F   T
```
The pattern `aba` occurs four times in the text `abacabadabacaba`.
**Example 3:**
```
T = "abcd"
P = "efg"
```
The table would be as follows:

```
T   a   b   c   d
P   e   f   g
    F   F   F   F
```
The pattern `efg` does not occur in the text `abcd`.

here are the steps to solve the string matching problem using top-down dynamic programming (also known as memoization) with code snippets in Go:

1.  Define a function `match(i, j)` that returns true if the substring `T[1...i]` contains the pattern `P[1...j]`, and false otherwise. This function will be called recursively to fill in the DP table.

```go
func match(i, j int) bool {
    // base cases
    if j == 0 {
        return true
    }
    if i == 0 {
        return false
    }

    // memoization
    if dp[i][j] != -1 {
        return dp[i][j] == 1
    }

    // recursive cases
    if T[i] == P[j] {
        dp[i][j] = boolToInt(match(i-1, j-1))
    } else {
        dp[i][j] = 0
    }
    return dp[i][j] == 1
}
```
2.  Create a two-dimensional array `dp` of size `(n+1) x (m+1)` (where `n` is the length of the text `T` and `m` is the length of the pattern `P`) to store the DP table. Initialize all entries to -1 to indicate that they have not been computed yet.

```go 
n := len(T)
m := len(P)
dp := make([][]int, n+1)
for i := 0; i <= n; i++ {
    dp[i] = make([]int, m+1)
    for j := 0; j <= m; j++ {
        dp[i][j] = -1
    }
}
```
3.  Call the `match` function for each pair of indices `(i, j)` to fill in the DP table.
```go
for i := 1; i <= n; i++ {
    for j := 1; j <= m; j++ {
        match(i, j)
    }
}
```
4.  Scan the last row of the DP table to find all indices `j` such that `dp[n][j]` is true, where `n` is the length of the text `T`. These indices correspond to the positions in the text where the pattern `P` occurs.
```go
matches := []int{}
for j := 1; j <= m; j++ {
    if dp[n][j] == 1 {
        matches = append(matches, n-j+1)
    }
}
```
Putting it all together, the complete code to solve the string matching problem using top-down DP in Go would look like this:

```go
func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

func match(i, j int, T, P string, dp [][]int) bool {
    // base cases
    if j == 0 {
        return true
    }
    if i == 0 {
        return false
    }

    // memoization
    if dp[i][j] != -1 {
        return dp[i][j] == 1
    }

    // recursive cases
    if T[i-1] == P[j-1] {
        dp[i][j] = boolToInt(match(i-1, j-1, T, P, dp))
    } else {
        dp[i][j] = 0
    }
    return dp[i][j] == 1
}

func StringMatching(T, P string) []int {
    n := len(T)
    m := len(P)

    // create DP table
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j <= m; j++ {
            dp[i][j] = -1
        }
    }

    // fill in DP table
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            match(i, j, T, P, dp)
        }
    }

    // find matching positions
    matches := []int{}
    for j := 1; j <= m; j++ {
        if dp[n][j] == 1 {
            matches = append(matches, n-j+1)
        }
    }
    return matches
}
```
To use this function, you can call `StringMatching` with the text and pattern strings as arguments, like this:

```go
T := "AABAACAADAABAABA"
P := "AABA"
matches := StringMatching(T, P)
fmt.Println(matches)  // prints [13 10 7 4]
```

### 24. Maximum sum subarray with at most k elements
The Maximum sum subarray with at most k elements problem is a classic computer science problem that involves finding the contiguous subarray of length at most k within a given array, such that the sum of its elements is maximum. Here are a few examples to help illustrate the problem:

Example 1: Consider the array \[1, 2, 3, -4, 5\] and let k=2. The maximum sum subarray of at most k elements is \[2, 3\], since it has the largest sum of any subarray of length at most 2.

Example 2: Consider the array \[-2, -3, 4, -1, -2, 1, 5, -3\] and let k=3. The maximum sum subarray of at most k elements is \[4, -1, -2\], since it has the largest sum of any subarray of length at most 3.

Example 3: Consider the array \[1, 2, -5, 4, 5, 6\] and let k=1. The maximum sum subarray of at most k elements is \[6\], since it has the largest sum of any subarray of length at most 1.

Example 4: Consider the array \[3, -1, 2, 1, -5, 4, 2\] and let k=4. The maximum sum subarray of at most k elements is \[3, -1, 2, 1, -5, 4\], since it has the largest sum of any subarray of length at most 4.


here are the steps to solve the Maximum sum subarray with at most k elements problem using top-down dynamic programming (DP) approach:

1.  Define a 2D DP array `dp[i][j]` to store the maximum sum subarray of at most `j` elements ending at index `i` of the given array.
2.  Initialize `dp[0][j]` to be the first `j` elements of the array.
3.  For `i` from 1 to the length of the array, and `j` from 1 to `k`, update `dp[i][j]` as the maximum of the following:
    * `dp[i-1][j] + arr[i]`: Include the `i`th element in the subarray.
    * `dp[i-1][j-1] + arr[i]`: Start a new subarray of length 1 at `i`th position.
    * `dp[i-1][j]`: Don't include the `i`th element in the subarray.
4.  The maximum sum subarray of at most `k` elements in the array is the maximum value of `dp[i][k]` for `i` from `k` to the length of the array.

Here's the Go code snippet that implements the above approach:

```go
func maxSumSubarrayAtMostK(arr []int, k int) int {
    n := len(arr)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, k+1)
        dp[i][1] = arr[i]
    }
    for i := 1; i < n; i++ {
        for j := 2; j <= k; j++ {
            dp[i][j] = dp[i-1][j] // don't include the ith element
            if dp[i-1][j-1] > 0 {
                dp[i][j] = max(dp[i][j], dp[i-1][j-1]+arr[i]) // start a new subarray of length 1 at i
            }
            dp[i][j] = max(dp[i][j], dp[i-1][j]+arr[i]) // include the ith element in the subarray
        }
    }
    maxSum := dp[k-1][k]
    for i := k; i < n; i++ {
        maxSum = max(maxSum, dp[i][k])
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
You can use the `maxSumSubarrayAtMostK` function to find the maximum sum subarray of at most `k` elements in a given array. Here's an example of how to use the function:
```go
arr := []int{1, 2, 3, -4, 5}
k := 2
maxSum := maxSumSubarrayAtMostK(arr, k)
fmt.Println(maxSum) // Output: 5
```
In this example, we have an array `arr` with values `[1, 2, 3, -4, 5]` and we want to find the maximum sum subarray of at most 2 elements. The `maxSumSubarrayAtMostK` function is called with `arr` and `k` as parameters and it returns the maximum sum subarray which is `[2, 3]` with a sum of 5.

You can change the values of `arr` and `k` to test the function on different inputs.

### 25. Longest alternating subsequence
The Longest Alternating Subsequence (LAS) problem is a classical problem in computer science and mathematics that asks for the length of the longest subsequence in a given sequence such that the elements in the subsequence alternate in sign (positive/negative) and are strictly increasing or decreasing.

More formally, given a sequence of numbers `A[1], A[2], ..., A[n]`, we want to find the length of the longest subsequence `B` such that:

1.  `B[1]` and `B[2]` have different signs.
2.  For all `i`, `B[i]` and `B[i+1]` have different signs.
3.  `B` is strictly increasing or decreasing.

For example, consider the sequence `A = [1, 3, -2, -4, 5, -6, 7, 8]`. The longest alternating subsequence in this sequence is `[1, -2, 5, -6, 7, 8]`, which has length 6. Note that this subsequence is both strictly increasing and alternating in sign.

Here are a few more examples to illustrate the problem:

1.  `A = [1, 2, 3, 4, 5]`. The longest alternating subsequence is `[1, 2, 1, 2, 1]`, which has length 5.
2.  `A = [5, 4, 3, 2, 1]`. The longest alternating subsequence is `[5, 4, 5, 4, 5]`, which has length 5.
3.  `A = [1, 2, 3, 2, 1, 0, 1, 2, 3, 4, 5]`. The longest alternating subsequence is `[1, 2, 1, 0, 1, 2, 3, 4, 5]`, which has length 9.

There are various algorithms to solve the LAS problem with different time complexities, including dynamic programming and binary search. 

The Longest Alternating Subsequence (LAS) problem has several uses in different fields, some of which are:

1.  Bioinformatics: In bioinformatics, the LAS problem is used to determine the longest alternating subsequence of DNA or protein sequences, which can provide insights into the structure and function of these molecules.
    
2.  Data compression: The LAS problem is used in data compression to find patterns in a sequence that can be exploited to reduce its size. For example, if a sequence contains a long alternating subsequence, this can be compressed by representing it with fewer bits.
    
3.  Finance: In finance, the LAS problem can be used to identify trading signals in financial time series data. For example, a long alternating subsequence of positive and negative returns could be interpreted as a signal to buy and sell stocks.
    
4.  Image processing: The LAS problem can be used in image processing to identify patterns in grayscale or color images. For example, a long alternating subsequence of pixel values could indicate the presence of an edge or boundary in the image.
    
5.  Game theory: The LAS problem has applications in game theory, where it is used to model alternating strategies between two players in a game. For example, in a game of rock-paper-scissors, the longest alternating subsequence could represent the optimal strategy for a player to use against their opponent.


Here are the steps to solve the Longest Alternating Subsequence (LAS) problem using top-down dynamic programming in Go:

1.  Define the state: The state of the problem can be defined by two indices `i` and `j`, where `i` is the index of the current element and `j` is the index of the last element added to the subsequence. We also need to keep track of the last sign added to the subsequence.
    
2.  Define the base case: The base case is when `i` is equal to the length of the sequence. In this case, the length of the longest alternating subsequence is 0.
    
3.  Define the recursive case: In the recursive case, we have two options: either add the current element to the subsequence or skip it. If we add the current element, we need to check if the last sign added to the subsequence is different from the sign of the current element. If so, we can continue recursively with the next index and update the length of the subsequence accordingly. If we skip the current element, we can also continue recursively with the next index without changing the length of the subsequence.
    
4.  Memoize the results: To avoid redundant computations, we can memoize the results of the recursive calls using a 2D array.
    

Here is the Go code that implements the above steps:

```go
func lasTopDown(A []int) int {
    n := len(A)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return lasHelper(A, n, 0, -1, memo)
}

func lasHelper(A []int, n, i, j, sign int, memo [][]int) int {
    if i == n {
        return 0
    }
    if memo[i][j+1] != -1 {
        return memo[i][j+1]
    }
    add := 0
    if j == -1 || (A[i] > 0 && sign < 0) || (A[i] < 0 && sign > 0) {
        add = 1 + lasHelper(A, n, i+1, j+1, -sign, memo)
    }
    skip := lasHelper(A, n, i+1, j, sign, memo)
    memo[i][j+1] = max(add, skip)
    return memo[i][j+1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
The `lasTopDown` function initializes the memoization array and calls the `lasHelper` function with the appropriate arguments. The `lasHelper` function checks if the result has already been memoized and returns it if so. It then computes the length of the subsequence recursively by considering both the cases of adding and skipping the current element. Finally, it memoizes the result and returns it.

Note that the `max` function is used to return the maximum of two integers.


### 26. Tower of Hanoi problem
The Tower of Hanoi is a classic puzzle where you have three pegs and a number of discs of different sizes. The objective is to move the entire stack to another peg, following three rules:

1.  Only one disc can be moved at a time.
2.  Each move consists of taking the upper disc from one of the stacks and placing it on top of another stack or an empty peg.
3.  No disc may be placed on top of a smaller disc.

The puzzle can be solved recursively, solving for smaller numbers of discs first and using that solution to solve for larger numbers of discs.

For example, if we start with three discs, the starting position has a neat stack in ascending order of size on one peg. The smallest disc is at the top, forming a conical shape. The target position has the entire stack moved to another peg, with the same neat stack arrangement.

To solve the puzzle, we take the smallest disc from the first peg and move it to the second peg. Next, we move the middle-sized disc from the first peg to the third peg. We then move the smallest disc from the second peg to the third peg. After that, we move the largest disc from the first peg to the second peg. We then move the smallest disc from the third peg to the first peg. Next, we move the middle-sized disc from the third peg to the second peg. Finally, we move the smallest disc from the first peg to the second peg.

The process of moving the discs is done recursively. To solve for four discs, we use the same steps, but with an extra disc. We first move the smallest disc from the first peg to the second peg, followed by moving the second-smallest disc from the first peg to the third peg. We then move the smallest disc from the second peg to the third peg, followed by moving the third-smallest disc from the first peg to the second peg. We continue until we have moved the entire stack to the target peg.

This recursive process can be continued for any number of discs, and the solution will always follow the same set of steps.

Here are the steps to solve the Tower of Hanoi problem:

1.  Identify the number of discs you have and label the pegs as A, B, and C.
2.  If you have only one disc, move it from peg A to peg C.
3.  If you have two discs, move the smallest disc from peg A to peg B. Then move the largest disc from peg A to peg C. Finally, move the smallest disc from peg B to peg C.
4.  If you have more than two discs, follow these steps:
        * Move the top n-1 discs from peg A to peg B using peg C as the temporary peg.
        * Move the largest disc from peg A to peg C.
        * Move the n-1 discs from peg B to peg C using peg A as the temporary peg.

These steps can be repeated recursively for any number of discs. The key is to move the smaller discs first and to always move the largest disc last. By following these rules, you can solve the puzzle and move the entire stack of discs from one peg to another.


Here are the steps to solve the Tower of Hanoi problem using top-down dynamic programming, along with code snippets in Go:

1.  Define the base case: If there is only one disc, it can be moved directly to the target peg. If there are no discs, there are no moves required.

```go
func solveTowerOfHanoi(numDiscs, fromPeg, toPeg int) int {
    if numDiscs == 1 {
        // If there is only one disc, it can be moved directly to the target peg.
        return 1
    }
    if numDiscs == 0 {
        // If there are no discs, there are no moves required.
        return 0
    }
    // ...
}
```
2.  Check if the optimal solution for `numDiscs` has already been calculated. If yes, return the result.

```go
func solveTowerOfHanoi(numDiscs, fromPeg, toPeg int) int {
    // Check if the optimal solution for numDiscs has already been calculated. If yes, return the result.
    if memo[numDiscs][fromPeg][toPeg] != 0 {
        return memo[numDiscs][fromPeg][toPeg]
    }
    // ...
}
```
3.  For `numDiscs > 1`, recursively solve the problem for `numDiscs - 1` discs, moving them from the source peg to a spare peg. Then move the largest disc from the source peg to the target peg. Finally, recursively solve the problem for `numDiscs - 1` discs, moving them from the spare peg to the target peg.
```go
func solveTowerOfHanoi(numDiscs, fromPeg, toPeg int) int {
    // Check if the optimal solution for numDiscs has already been calculated. If yes, return the result.
    if memo[numDiscs][fromPeg][toPeg] != 0 {
        return memo[numDiscs][fromPeg][toPeg]
    }

    // Solve the problem recursively.
    sparePeg := 6 - fromPeg - toPeg // The spare peg is the one not currently used as source or target.
    numMoves := 0

    // Move n-1 discs from the source peg to the spare peg.
    numMoves += solveTowerOfHanoi(numDiscs-1, fromPeg, sparePeg)

    // Move the largest disc from the source peg to the target peg.
    numMoves++
    
    // Move n-1 discs from the spare peg to the target peg.
    numMoves += solveTowerOfHanoi(numDiscs-1, sparePeg, toPeg)

    // Memoize the result for future use.
    memo[numDiscs][fromPeg][toPeg] = numMoves

    // Return the number of moves required to solve the problem.
    return numMoves
}
```

4.  Memoize the result and return it.
```go
package main

import "fmt"

var memo [21][4][4]int // Memoization array to store pre-calculated solutions.

func solveTowerOfHanoi(numDiscs, fromPeg, toPeg int) int {
    // Check if the optimal solution for numDiscs has already been calculated. If yes, return the result.
    if memo[numDiscs][fromPeg][toPeg] != 0 {
        return memo[numDiscs][fromPeg][toPeg]
    }

    // Solve the problem recursively.
    sparePeg := 6 - fromPeg - toPeg // The spare peg is the one not currently used as source or target.
    numMoves := 0

    if numDiscs == 1 {
        // Base case: If there is only one disc, it can be moved directly to the target peg.
        numMoves = 1
    } else {
        // Move n-1 discs from the source peg to the spare peg.
        numMoves += solveTowerOfHanoi(numDiscs-1, fromPeg, sparePeg)

        // Move the largest disc from the source peg to the target peg.
        numMoves++

        // Move n-1 discs from the spare peg to the target peg.
        numMoves += solveTowerOfHanoi(numDiscs-1, sparePeg, toPeg)
    }

    // Memoize the result for future use.
    memo[numDiscs][fromPeg][toPeg] = numMoves

    // Return the number of moves required to solve the problem.
    return numMoves
}

func main() {
    numDiscs := 3
    fromPeg := 1
    toPeg := 3

    numMoves := solveTowerOfHanoi(numDiscs, fromPeg, toPeg)

    fmt.Printf("To solve the Tower of Hanoi problem with %d discs from peg %d to peg %d, %d moves are required.\n", numDiscs, fromPeg, toPeg, numMoves)
}
```
In this implementation, the `memo` array is used to store pre-calculated solutions to subproblems. The `solveTowerOfHanoi` function takes three arguments: the number of discs, the source peg, and the target peg. It returns the minimum number of moves required to solve the problem.

The function first checks if the optimal solution for the given input has already been calculated. If yes, it returns the result. If not, it solves the problem recursively using the top-down approach, memoizing the results of subproblems along the way to avoid redundant calculations.

In the `main` function, the `numDiscs`, `fromPeg`, and `toPeg` variables are initialized and used to call the `solveTowerOfHanoi` function. The minimum number of moves required to solve the problem is printed to the console.



### 27. Longest common prefix
The Longest Common Prefix (LCP) problem is a classic computer science problem that involves finding the longest common prefix string among a set of strings. In other words, given a set of strings, the goal is to find the longest string that is a prefix of all the strings.

For example, consider the following set of strings:

arduinoCopy code

```
{"apple", "app", "ape"}
```

The longest common prefix among these strings is "ap".

Another example:

```
{"car", "cat", "cab"}
```

The longest common prefix among these strings is "ca".

Here are a few more examples:

* {"flower", "flow", "flight"} -> "fl"
* {"dog", "doughnut", "door"} -> "do"
* {"house", "houston", "hotdog"} -> "ho"

If there is no common prefix among the strings, the LCP is an empty string.

Here's an example of a set of strings that have no common prefix:
```
{"hello", "world", "goodbye"}
```
In this case, the LCP is an empty string.

Here are the steps to solve the Longest Common Prefix (LCP) problem using top-down dynamic programming:

1.  Define a recursive function that takes in two parameters: the current index in the strings and the current LCP. The function should return the length of the LCP.
    
2.  Check if the current index is out of bounds for any of the strings. If it is, return the length of the LCP.
    
3.  Check if the current character at the current index is the same for all the strings. If it is, call the recursive function with the current index incremented and the LCP with the current character appended to it.
    
4.  If the current character is not the same for all the strings, return the length of the LCP.
    
5.  In the main function, call the recursive function with the initial index of 0 and an empty string as the initial LCP.
    
6.  Return the LCP length.
    

Here's the code snippet in Go:
```go
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    return lcpHelper(strs, 0, "")
}

func lcpHelper(strs []string, index int, lcp string) string {
    if index >= len(strs[0]) {
        return lcp
    }
    currentChar := strs[0][index]
    for _, str := range strs {
        if index >= len(str) || str[index] != currentChar {
            return lcp
        }
    }
    lcp += string(currentChar)
    return lcpHelper(strs, index+1, lcp)
}
```
In this code, we first check if the input array is empty. If it is, we return an empty string. Otherwise, we call the `lcpHelper` function with the initial index of 0 and an empty string as the initial LCP.

The `lcpHelper` function checks if the current index is out of bounds for any of the strings. If it is, it returns the current LCP. If the current character at the current index is the same for all the strings, it calls itself with the current index incremented and the LCP with the current character appended to it. If the current character is not the same for all the strings, it returns the current LCP.

Finally, in the `longestCommonPrefix` function, we return the result of the `lcpHelper` function.

### 28. Bellman-Ford algorithm

The Bellman-Ford algorithm is a dynamic programming algorithm used to find the shortest path between a source vertex and all other vertices in a weighted graph, even in the presence of negative weight edges. The algorithm works by computing the shortest path from the source vertex to every other vertex by relaxing the edges in the graph in a specific order.

Here's a step-by-step explanation of the algorithm, along with examples to illustrate its operation:

1.  Initialize the distance of the source vertex to 0, and the distance of all other vertices to infinity.

For example, consider the following weighted graph:
 ```
      5     3
   A ----> B ----> C
   | 1        | -2
   v          v
   D <---- E ----> F
      4      -1
```
Assume we are finding the shortest path from vertex A to all other vertices.

The initialization step sets the distance of vertex A to 0 and the distance of all other vertices to infinity:
 ```
  Vertex  | Distance
---------------------
    A    |    0
    B    |  Inf
    C    |  Inf
    D    |  Inf
    E    |  Inf
    F    |  Inf
```
2.  Relax all the edges in the graph V-1 times, where V is the number of vertices in the graph.

Relaxing an edge means updating the distance of the destination vertex if a shorter path is found through the source vertex. For each vertex, the algorithm considers all its outgoing edges and checks whether the path to the destination vertex can be improved by going through the source vertex.

For example, in the first iteration, we relax all the edges:
```
  Vertex  | Distance
---------------------
    A    |    0
    B    |    5
    C    |    8
    D    |    1
    E    |    5
    F    |   11
```
To relax the edge (A, B), we compare the distance of vertex B (infinity) with the distance of vertex A (0) plus the weight of the edge (5). Since 0 + 5 < infinity, we update the distance of vertex B to 5. We do the same for all the other edges and vertices.

3.  Check for negative-weight cycles.

If there are negative-weight cycles in the graph, the algorithm will not be able to find the shortest path, as the distance of the vertices will keep decreasing with each iteration. To check for negative-weight cycles, we relax all the edges one more time. If any distance can still be improved, it means there is a negative-weight cycle.

For example, in the second iteration, we relax all the edges again:

```
  Vertex  | Distance
---------------------
    A    |    0
    B    |    5
    C    |    3
    D    |    1
    E    |    5
    F    |   -2
```

The distance of vertex F can still be improved, even after V-1 iterations. This means there is a negative-weight cycle in the graph.

4.  Handle negative-weight cycles.

If there is a negative-weight cycle in the graph, the algorithm cannot find the shortest path, as the distance of the vertices will keep decreasing with each iteration. However, we can still detect the negative-weight cycle using the previous step.

For example, in the previous step, we found a negative-weight cycle in the graph. The cycle is formed by the edges (B, C), (C, E), and (E, B). To handle negative-weight cycles, we can either

The Bellman-Ford algorithm is typically implemented using a bottom-up approach rather than top-down dynamic programming. Therefore, it is not recommended to use a top-down approach for this algorithm. However, I can provide steps for implementing the bottom-up approach in Go:

1.  Initialize the distance of the source vertex to 0, and the distance of all other vertices to infinity.
    
2.  Create an array to store the distance to each vertex and set the distance of the source vertex to 0.
    
3.  Iterate V-1 times (where V is the number of vertices in the graph), and for each iteration, iterate over all the edges in the graph and update the distance to the destination vertex if a shorter path is found through the current vertex.
    
4.  If, after V-1 iterations, any distance can still be improved, it means there is a negative-weight cycle in the graph.
    

Here's an example implementation of the Bellman-Ford algorithm in Go:
```go
package main

import "fmt"

type edge struct {
    from int
    to int
    weight int
}

func main() {
    n := 6  // number of vertices
    m := 7  // number of edges
    source := 0  // source vertex
    edges := []edge{{0, 1, 5}, {0, 3, 1}, {1, 2, 3}, {2, 5, -1}, {3, 1, 1}, {3, 4, 4}, {4, 5, -2}}

    distances := make([]int, n)
    for i := range distances {
        distances[i] = int(^uint(0) >> 1)  // set all distances to infinity
    }
    distances[source] = 0  // set distance of source vertex to 0

    for i := 0; i < n-1; i++ {
        for _, e := range edges {
            if distances[e.from]+e.weight < distances[e.to] {
                distances[e.to] = distances[e.from] + e.weight
            }
        }
    }

    // check for negative-weight cycle
    for _, e := range edges {
        if distances[e.from]+e.weight < distances[e.to] {
            fmt.Println("Negative-weight cycle found")
            return
        }
    }

    // print distances to all vertices
    for i, d := range distances {
        fmt.Printf("Distance from %d to %d: %d\n", source, i, d)
    }
}
```

### 29. Shortest path in a DAG (Directed Acyclic Graph)

The shortest path in a DAG problem involves finding the shortest path from a source vertex to a destination vertex in a directed acyclic graph (DAG). A DAG is a graph with directed edges and no cycles. The shortest path problem is a classic problem in graph theory and has many practical applications.

Here are some examples of the shortest path in a DAG problem:

Example 1: Suppose you have a DAG with four vertices and five edges, as shown below. You want to find the shortest path from vertex A to vertex D.

```
A --> B --> C --> D
 \        ^
  \      /
   --> E -->
```

The shortest path from A to D is A → B → C → D, with a length of 3.

Example 2: Suppose you have a DAG with six vertices and seven edges, as shown below. You want to find the shortest path from vertex A to vertex F.

```
A --> B --> C --> D
 \        ^      ^
  \      /      /
   --> E --> F--
```

The shortest path from A to F is A → B → C → D → F, with a length of 4.

Example 3: Suppose you have a DAG with five vertices and six edges, as shown below. You want to find the shortest path from vertex A to vertex E.

```
A --> B --> C --> D
         \      ^
          \    /
           --> E
```

The shortest path from A to E is A → B → C → D → E, with a length of 4.

Example 4: Suppose you have a DAG with five vertices and six edges, as shown below. You want to find the shortest path from vertex A to vertex E.

```
A --> B --> C --> D
 \        ^      ^
  \      /      /
   --> E <-----
```

In this example, there is no path from A to E, because there is a cycle in the graph. This means that there is no shortest path from A to E.
Here are the steps to solve the shortest path in a DAG problem using top-down dynamic programming:

1.  Define the problem: Define the problem as finding the shortest path from a source vertex to a destination vertex in a DAG.
    
2.  Identify the subproblems: Break the problem down into smaller subproblems, where each subproblem is finding the shortest path from the source vertex to a specific intermediate vertex.
    
3.  Define the recurrence relation: Define the recurrence relation that relates the solution to a subproblem to the solutions of its smaller subproblems.
    
4.  Implement top-down dynamic programming: Implement the solution to the problem using top-down dynamic programming, which involves solving the subproblems recursively and storing the solutions in a memoization table to avoid recomputing the same subproblems multiple times.
    

Here is the code snippet in Go:
```go
type graph struct {
    vertices int
    edges    [][]int
}

func shortestPathDAG(g graph, src, dest int, memo []int) int {
    if src == dest {
        return 0
    }
    if memo[src] != -1 {
        return memo[src]
    }
    shortest := math.MaxInt32
    for _, neighbor := range g.edges[src] {
        pathLen := shortestPathDAG(g, neighbor, dest, memo)
        shortest = min(shortest, pathLen+1)
    }
    memo[src] = shortest
    return shortest
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func main() {
    g := graph{
        vertices: 4,
        edges: [][]int{
            {1, 2},
            {2},
            {3},
            {},
        },
    }
    memo := make([]int, g.vertices)
    for i := range memo {
        memo[i] = -1
    }
    src, dest := 0, 3
    shortest := shortestPathDAG(g, src, dest, memo)
    fmt.Println("Shortest path from", src, "to", dest, ":", shortest)
}
```
In this code, the `graph` struct represents the DAG as an adjacency list. The `shortestPathDAG` function uses memoization to store the solutions to subproblems and return the shortest path from the source vertex to the destination vertex. The `min` function returns the minimum of two integers. The `main` function creates a DAG and calls the `shortestPathDAG` function to find the shortest path from the source vertex to the destination vertex.
