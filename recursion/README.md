# Recursion

1.  **Understand the problem:** Read the problem statement and understand what the problem is asking. Identify the input and output formats and constraints of the problem.
2.  **Identify the base case:** A base case is a condition in which the function does not make any recursive calls and returns a value. Identify the simplest input for which the solution is already known, usually a case with the smallest input size or a trivial case.
3.  **Define the recursive case:** Identify how the problem can be broken down into smaller sub-problems. Express the solution to the problem in terms of a recursive function, where the function makes one or more recursive calls to itself with smaller input sizes.
4.  **Solve the problem recursively:** Call the function recursively with the smaller sub-problems until the base case is reached. At the base case, return the value directly, and use the returned values to compute the final answer.
5.  **Test the solution:** Test the solution with different input sizes and corner cases to ensure it works as expected.
6.  **Optimize the solution:** Once you have a working solution, think about how you can optimize the algorithm by reducing the number of recursive calls or avoiding unnecessary computations.
7.  **Analyze the time and space complexity:** Analyze the time and space complexity of the algorithm to understand its performance characteristics.
8.  **Refactor the code:** If necessary, refactor the code to improve its readability, maintainability, and modularity.

## Running Times

To express the running time of a recursive function in terms of the running time of its sub-problems, you need to define a recurrence relation. A recurrence relation expresses the running time of the function as a function of the running time of its sub-problems. Here's how you can define a recurrence relation for a recursive function:

1.  **Identify the base case**: The base case is the simplest version of the problem that can be solved without recursion. It is typically a small input size that can be solved directly without making any recursive calls.
    
2.  **Identify the recursive case**: The recursive case is the more complex version of the problem that requires one or more recursive calls to solve. It typically involves dividing the problem into smaller sub-problems, solving each sub-problem recursively, and then combining the results to solve the original problem.
    
3.  **Define the recurrence relation**: To define the recurrence relation, you need to express the running time of the function in terms of the running time of its sub-problems. This typically involves adding the time required to solve the base case to the time required to solve the recursive case.

### Master Theorem
The master theorem is a mathematical tool that helps us analyze the running time of some recursive algorithms. It gives us a formula that we can use to compute the running time in terms of the size of the input and the running time of the sub-problems.

The formula has three parts:

1.  The first part represents the time it takes to solve the problem on the top level.
2.  The second part represents the time it takes to solve the sub-problems.
3.  The third part represents any additional work done by the algorithm.

To use the master theorem, we need to know how the running time of the sub-problems changes as we divide the problem into smaller sub-problems. If we can express this running time in a specific way, we can use the master theorem to compute the overall running time.

For example, if the running time of the sub-problems is a constant fraction of the original problem size (i.e., it takes half as long to solve a problem that is half the size), then the master theorem tells us that the running time of the algorithm is O(log n).

In general, the master theorem is a useful tool for understanding the running time of recursive algorithms and can help us make better decisions about how to optimize them.

#### Mathematical Definition

The master theorem provides a way to calculate the asymptotic complexity of a divide-and-conquer algorithm by solving a recurrence relation of the form:

`T(n) = aT(n/b) + f(n)`

where:

* T(n) is the running time of the algorithm on an input of size n.
* a is the number of subproblems created in the recursion.
* n/b is the size of each subproblem.
* f(n) is the time required to combine the solutions of the subproblems.

The master theorem defines three cases, based on the relative growth rates of a and f(n)/n^k:

Case 1: `If a > b^k`, then `T(n) = O(n^log_b a)`.

Case 2: If `a = b^k`, then `T(n) = O(n^k log n)`.

Case 3: If `a < b^k`, then `T(n) = O(n^k)`.

Here, k is a constant that depends on the structure of the subproblems and f(n)/n^k is a measure of the work required to combine the subproblems. The master theorem provides a simple way to analyze the running time of many common recursive algorithms.

#### Examples
Here are some easy and intuitive examples of how the master theorem can be used:

1.  Suppose we have an algorithm that divides a problem into two sub-problems of size n/2 and then combines the results in constant time. In this case, the running time of the algorithm can be expressed as T(n) = 2T(n/2) + O(1). By applying the master theorem, we can see that the running time of this algorithm is O(n).
2.  Let's say we have an algorithm that divides a problem into four sub-problems of size n/4 and then combines the results in linear time. In this case, the running time of the algorithm can be expressed as T(n) = 4T(n/4) + O(n). By applying the master theorem, we can see that the running time of this algorithm is O(n log n).
3.  Suppose we have an algorithm that divides a problem into three sub-problems of size n/3 and then combines the results in quadratic time. In this case, the running time of the algorithm can be expressed as T(n) = 3T(n/3) + O(n^2). By applying the master theorem, we can see that the running time of this algorithm is O(n^2).
4.  Suppose we have an algorithm that divides a problem into two sub-problems of size n/2 and then combines the results in linear time. In this case, the running time of the algorithm can be expressed as T(n) = 2T(n/2) + O(n). By applying the master theorem, we can see that the running time of this algorithm is O(n log n).
5.  Let's say we have an algorithm that divides a problem into two sub-problems of size n/2 and then combines the results in quadratic time. In this case, the running time of the algorithm can be expressed as T(n) = 2T(n/2) + O(n^2). By applying the master theorem, we can see that the running time of this algorithm is O(n^2).
6.  Suppose we have an algorithm that divides a problem into three sub-problems of size n/3 and then combines the results in linear time. In this case, the running time of the algorithm can be expressed as T(n) = 3T(n/3) + O(n). By applying the master theorem, we can see that the running time of this algorithm is O(n).
7.  Let's say we have an algorithm that divides a problem into four sub-problems of size n/4 and then combines the results in constant time. In this case, the running time of the algorithm can be expressed as T(n) = 4T(n/4) + O(1). By applying the master theorem, we can see that the running time of this algorithm is O(n).
8.  Suppose we have an algorithm that divides a problem into two sub-problems of size n/2 and then combines the results in logarithmic time. In this case, the running time of the algorithm can be expressed as T(n) = 2T(n/2) + O(log n). By applying the master theorem, we can see that the running time of this algorithm is O(n log n).

Let us consider some examples:

#### Factorial
The factorial function is defined as follows:

```
n! = n * (n-1) * (n-2) * ... * 1
```

One way to compute the factorial of a number is to use a recursive algorithm:

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

To analyze the running time of this algorithm, we can use the **Master theorem**. Let's set `a = 1` and `b = 1`, since the algorithm makes one recursive call to compute the factorial of `n-1`, which has size `n-1`. The cost of multiplying n and the result of the recursive call is `O(1)`, since it takes constant time to perform multiplication.

Therefore, we have:

`T(n) = T(n-1) + O(1)`

This recurrence relation tells us that the running time of the algorithm depends on the running time of the subproblem of size `n-1`.

Using the Master theorem, we can classify this recurrence relation as follows:

* `a = 1, b = 1, and f(n) = O(1)`
* `log_b a = log_1 1 = 0`

Since` f(n) = O(1)` is equal to `n^(log_b a)`, which is a constant, we can apply case 1 of the Master theorem, which tells us that the running time of the algorithm is `Θ(n^log\_1 1 * log n) = Θ(log n)`.

In other words, the running time of the recursive factorial algorithm is logarithmic in the size of the input, which means that it grows slowly with the input size. Therefore, the recursive factorial algorithm is efficient and can be used to compute the factorial of large numbers, even though it uses recursion.

#### Fibonacci Sequence
Let's consider an example of the Fibonacci sequence, which is often used to illustrate the power of recursion.

The Fibonacci sequence is defined as follows:

F(0) = 0 F(1) = 1 F(n) = F(n-1) + F(n-2) for n > 1

One way to compute the nth Fibonacci number is to use a recursive algorithm:

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

To analyze the running time of this algorithm, we can use the Master theorem. Let's set a = 2 and b = 2, since the algorithm makes two recursive calls to compute F(n-1) and F(n-2), both of which have size n/2. The cost of combining the results is O(1), since it takes constant time to add two numbers. Therefore, we have:

T(n) = 2T(n/2) + O(1)

This is a classic example of the divide-and-conquer paradigm, where we divide the problem into two subproblems of size n/2, solve each subproblem recursively, and combine the results in linear time.

Using the Master theorem, we can classify this recurrence as follows:

* a = 2, b = 2, and f(n) = O(1)
* log_b a = log_2 2 = 1

Since f(n) = O(1) is polynomially smaller than n^(log_b a), which is linear, we can apply case 2 of the Master theorem, which tells us that the running time of the algorithm is Θ(n^log_2 2) = Θ(n). In other words, the running time of the recursive Fibonacci algorithm is linear in the size of the input, which means that it grows very quickly with the input size. As a result, this algorithm is not very efficient for large values of n, and an iterative solution or a more sophisticated algorithm is required to compute Fibonacci numbers efficiently.

#### Permutation
Permutation generation is the process of generating all possible permutations of a set of n elements. One way to generate permutations is to use a recursive algorithm that swaps elements in the set to generate new permutations.

Let's consider an example of a recursive algorithm for generating permutations:
```go
func permute(nums []int, start int) {
    if start == len(nums)-1 {
        fmt.Println(nums)
    } else {
        for i := start; i < len(nums); i++ {
            nums[start], nums[i] = nums[i], nums[start]
            permute(nums, start+1)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }
}
```
This algorithm uses a divide-and-conquer approach to generate permutations by swapping elements. The running time of the algorithm depends on the number of permutations that need to be generated, which is n! for a set of n elements.

To analyze the running time of this algorithm, we can use the Master theorem. Let's set a = 1 and b = 2, since the algorithm generates two subproblems of size n-1 by swapping elements and making recursive calls. The cost of swapping elements and making recursive calls is O(1), since it takes constant time to perform these operations.

Therefore, we have:

T(n) = 2 * T(n-1) + O(1)

This recurrence relation tells us that the running time of the algorithm depends on the running time of the two subproblems of size n-1.

Using the Master theorem, we can classify this recurrence relation as follows:

* a = 2, b = 2, and f(n) = O(1)
* log_b a = log_2 2 = 1

Since f(n) = O(1) is equal to n^(log_b a), which is a constant, we can apply case 2 of the Master theorem, which tells us that the running time of the algorithm is Θ(n^log_2 2 * log n) = Θ(n log n).

In other words, the running time of the recursive permutation generation algorithm is O(n log n), which is not as efficient as other non-recursive algorithms for generating permutations, but still manageable for small values of n.

## Easy problems
1.  Write a recursive function to find the factorial of a number. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/factorial.go#L3)]
2.  Write a recursive function to find the nth Fibonacci number.  [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/fibonacci.go#L3)]
3.  Write a recursive function to find the greatest common divisor of two numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/gcd.go#L7)]
4.  Write a recursive function to find the least common multiple of two numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/gcd.go#L21)]
5.  Write a recursive function to calculate the sum of an array of numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L4)]
6.  Write a recursive function to calculate the product of an array of numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L12)]
7.  Write a recursive function to reverse a string. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/strings.go#L4)]
8.  Write a recursive function to check if a string is a palindrome. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/strings.go#L9)]
9.  Write a recursive function to calculate the power of a number. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L4)]
10. Write a recursive function to find the sum of digits of a number. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L9)]
11. Write a recursive function to find the number of digits in a number. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L14)]
12. Write a recursive function to find the minimum element in an array. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L24)]
13. Write a recursive function to find the maximum element in an array.[[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L32)]
14. Write a recursive function to check if an array is sorted in ascending order. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L55)]
15. Write a recursive function to check if an array is sorted in descending order. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L67)]
16. Write a recursive function to search for an element in an array. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/collections.go#L79)]
17. Write a recursive function to calculate the nth triangular number.[[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L28)]
18. Write a recursive function to calculate the sum of the first n natural numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L36)]
19. Write a recursive function to calculate the sum of the first n even numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L44)]
20. Write a recursive function to calculate the sum of the first n odd numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L52)]
21. Write a recursive function to print the numbers from 1 to n. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L60)]
22. Write a recursive function to print the numbers from n to 1. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L69)]
24. Write a recursive function to calculate the product of the first n natural numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L78)]
25. Write a recursive function to calculate the product of the first n even numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L86)]
26. Write a recursive function to calculate the product of the first n odd numbers. [[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L94)]
27. Write a recursive function to calculate the sum of the digits of a binary number.[[code](https://github.com/m0nadicph0/DSA/blob/main/recursion/numerical.go#L102)]

## Intermediate Problems
1.  Write a recursive function to generate all permutations of a string.
2.  Write a recursive function to generate all combinations of a set of integers.
3.  Write a recursive function to generate all subsets of a set of integers.
4.  Write a recursive function to find the maximum element in an array.
5.  Write a recursive function to find the minimum element in an array.
6.  Write a recursive function to sort an array using the merge sort algorithm.
7.  Write a recursive function to sort an array using the quicksort algorithm.
8.  Write a recursive function to solve the Tower of Hanoi problem.
9.  Write a recursive function to solve the Josephus problem.
10. Write a recursive function to find the nth root of a number using the bisection method.
11. Write a recursive function to evaluate a postfix expression.
12. Write a recursive function to check if a given binary tree is a binary search tree.
13. Write a recursive function to calculate the height of a binary tree.
14. Write a recursive function to calculate the diameter of a binary tree.
15. Write a recursive function to find the lowest common ancestor of two nodes in a binary tree.
16. Write a recursive function to traverse a binary tree in preorder.
17. Write a recursive function to traverse a binary tree in inorder.
18. Write a recursive function to traverse a binary tree in postorder.
19. Write a recursive function to find the kth smallest element in a binary search tree.
20. Write a recursive function to find the kth largest element in a binary search tree.
21. Write a recursive function to find the nth prime number.
22. Write a recursive function to find the number of ways to climb n stairs, taking 1, 2, or 3 steps at a time.
24.  Write a recursive function to generate all subsets of a set.
25.  Write a recursive function to determine if a given sum can be obtained from a set of integers.
26.  Write a recursive function to solve the Tower of Hanoi problem.
27.  Write a recursive function to find the kth smallest element in a binary search tree.
28.  Write a recursive function to check if a binary tree is a binary search tree.
29.  Write a recursive function to print all nodes of a binary tree in level order.
30. Write a recursive function to find the diameter of a binary tree.
31. Write a recursive function to find the lowest common ancestor of two nodes in a binary tree.
32. Write a recursive function to check if a binary tree is symmetric.
33. Write a recursive function to find the shortest path in a maze.
34. Write a recursive function to solve the N-Queens problem.
35. Write a recursive function to generate all valid IP addresses from a given string of digits.
36. Write a recursive function to find the largest subarray with equal number of 0s and 1s.
37. Write a recursive function to check if a given graph is connected.
38. Write a recursive function to find the shortest path between two nodes in a graph.
39. Write a recursive function to find the articulation points in a graph.
40. Write a recursive function to find the strongly connected components of a directed graph.
41. Write a recursive function to find the maximum flow in a network flow graph.
42. Write a recursive function to find the minimum cut in a network flow graph.
43. Write a recursive function to find the Eulerian path in a graph.
44. Write a recursive function to find the Hamiltonian path in a graph.
45. Write a recursive function to solve the 0/1 Knapsack problem.
46. Write a recursive function to find the largest sum of a contiguous subarray.
47. Write a recursive function to find the longest common subsequence of two strings.
48. Write a recursive function to find the Edit distance between two strings.
49. Write a recursive function to find the Levenshtein distance between two strings.
50. Write a recursive function to find the number of ways to make change for a given amount of money.

## Advanced Problems
1.  Write a recursive function to generate all possible ways to parenthesize an expression.
2.  Write a recursive function to generate all possible ways to arrange a set of n queens on an nxn chessboard.
3.  Write a recursive function to solve the knapsack problem.
4.  Write a recursive function to solve the subset sum problem.
5.  Write a recursive function to solve the traveling salesman problem.
6.  Write a recursive function to solve the longest common subsequence problem.
7.  Write a recursive function to solve the longest increasing subsequence problem.
8.  Write a recursive function to find the maximum subarray sum.
9.  Write a recursive function to find the longest palindrome subsequence.
10. Write a recursive function to find the shortest path in a graph using Dijkstra's algorithm.
11. Write a recursive function to find the shortest path in a graph using Bellman-Ford algorithm.
12. Write a recursive function to solve the edit distance problem.
13. Write a recursive function to find the maximum flow in a flow network using Ford-Fulkerson algorithm.
14. Write a recursive function to solve the 0-1 knapsack problem.
15. Write a recursive function to solve the longest common substring problem.
16. Write a recursive function to solve the coin change problem.
17. Write a recursive function to solve the minimum coin change problem.
18. Write a recursive function to solve the longest palindrome substring problem.
19. Write a recursive function to solve the rod cutting problem.
20. Write a recursive function to solve the egg dropping problem.
21. Write a recursive function to solve the maximum sum increasing subsequence problem.
22. Write a recursive function to solve the longest zig-zag subsequence problem.
23.  Write a recursive function to solve the maximum length chain of pairs problem.
24. Write a recursive function to solve the longest bitonic subsequence problem.
25. Write a recursive function to solve the longest alternating subsequence problem.
26. Write a recursive function to solve the longest increasing subsequence in a matrix problem.
27. Write a recursive function to solve the longest palindromic subsequence in a matrix problem.
28. Write a recursive function to solve the longest increasing subsequence in a DAG (Directed Acyclic Graph) problem.
29. Write a recursive function to solve the all-pairs shortest path problem using Floyd-Warshall algorithm.
30. Write a recursive function to solve the maximum independent set problem in a tree.
31. Write a recursive function to solve the maximum independent set problem in a DAG (Directed Acyclic Graph).
32. Write a recursive function to solve the vertex cover problem in a tree.
33. Write a recursive function to solve the vertex cover problem in a DAG (Directed Acyclic Graph).
34. Write a recursive function to solve the Hamiltonian path problem in a graph.
35. Write a recursive function to solve the Eulerian path problem in a graph.
36. Write a recursive function to solve the minimum vertex cover problem in a graph.
37. Write a recursive function to solve the maximum clique problem in a graph.
38. Write a recursive function to solve the maximum cut problem in a graph.
39. Write a recursive function to solve the minimum cut problem in a graph.
40. Write a recursive function to solve the maximum independent set problem in a graph.
