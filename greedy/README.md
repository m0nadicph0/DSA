# Greedy Algorithms

Greedy algorithms are problem-solving strategies that make **locally optimal choices** at each step 
in the hope of finding a **global optimum solution**. 

Here are some properties of greedy algorithms:

1. **Greedy Choice Property**: A greedy algorithm always makes the locally optimal choice at each step, hoping to find the global optimum solution.
1. **Optimal Substructure Property**: A problem has optimal substructure if an optimal solution to the problem contains optimal solutions to its subproblems. This property allows the greedy algorithm to make a series of locally optimal choices that ultimately lead to a global optimum solution.
1. **Often efficient**: Greedy algorithms often involve a simple and straightforward strategy that can be implemented quickly and efficiently.
1. **May not always find the globally optimal solution**: Greedy algorithms do not guarantee finding the globally optimal solution, as they make locally optimal choices at each step. In some cases, a greedy approach may lead to a suboptimal solution.
1. **Often used as a heuristic**: Greedy algorithms are often used as a heuristic to find a good solution quickly, which can then be refined using other techniques.
1. **Correctness may require proof**: In some cases, it may be necessary to prove the correctness of a greedy algorithm, particularly if the solution obtained is not guaranteed to be optimal.

## Examples

### Kruskal's algorithm for minimum spanning trees:
This algorithm selects the edge with the smallest weight and adds it to the spanning tree, 
as long as it doesn't create a cycle. This process is repeated until all vertices are connected. 
The properties of greedy algorithms that apply here are:
  - **Greedy choice property**: At each step, the algorithm chooses the edge with the minimum weight, which is the locally optimal choice.
  - **Optimal substructure**: The subproblem of choosing the minimum-weight edge from a set of edges leads to an optimal solution for the overall problem of finding the minimum spanning tree.

### Dijkstra's algorithm for shortest path: 
This algorithm selects the vertex with the smallest distance and adds it to the shortest path tree, updating the distances of its neighboring vertices. This process is repeated until the destination vertex is reached. The properties of greedy algorithms that apply here are:
  * **Greedy choice property**: At each step, the algorithm chooses the vertex with the smallest distance, which is the locally optimal choice.
  * **Optimal substructure**: The subproblem of finding the shortest path from a source vertex to a neighboring vertex leads to an optimal solution for the overall problem of finding the shortest path from the source vertex to the destination vertex.

### Huffman coding:
This algorithm constructs a binary tree that represents the optimal prefix code for a set of characters based on their frequency. The characters with the lowest frequency are assigned the shortest binary code. The properties of greedy algorithms that apply here are:
  * **Greedy choice property**: At each step, the algorithm chooses the two characters with the lowest frequency and merges them into a single node in the tree, which is the locally optimal choice.
  * **Optimal substructure**: The subproblem of merging two characters into a single node leads to an optimal solution for the overall problem of constructing the optimal prefix code.

### Activity selection problem:
This algorithm selects the maximum number of non-overlapping activities that can be scheduled in a given time frame. The activities are sorted by their finish times, and the algorithm selects the first activity that finishes, then iterates through the remaining activities, selecting the first one that finishes after the previous activity. The properties of greedy algorithms that apply here are:
  * **Greedy choice property**: At each step, the algorithm chooses the activity that finishes first, which is the locally optimal choice.
  * **Optimal substructure**: The subproblem of selecting the first activity that finishes after the previous activity leads to an optimal solution for the overall problem of scheduling the maximum number of non-overlapping activities.

### Coin change problem:
This algorithm finds the minimum number of coins needed to make change for a given amount. The algorithm selects the largest coin denomination that is less than or equal to the remaining amount, subtracts it from the remaining amount, and repeats until the remaining amount is zero. The properties of greedy algorithms that apply here are:
  * **Greedy choice property**: At each step, the algorithm chooses the largest coin denomination that is less than or equal to the remaining amount, which is the locally optimal choice.
  * **Optimal substructure**: The subproblem of finding the minimum number of coins needed to make change for a smaller amount leads to an optimal solution for the overall problem of finding the minimum number of coins needed to make change for the given amount.

## General format of Greedy Algorithms

The general format of greedy algorithms is as follows:

1. Define a problem and a set of candidate solutions.
1. Identify a criterion for evaluating the candidate solutions.
1. Define a rule for selecting the best candidate solution.
1. Repeat the selection rule until a satisfactory solution is found.

Here's an example of the greedy coin change problem in Go:

```go
func coinChangeGreedy(coins []int, amount int) int {
    // Base case: If the amount is zero, we don't need any coins
    if amount == 0 {
        return 0
    }
    
    // Greedy step: Find the largest coin denomination that is less than or equal to the given amount
    largestCoin := -1
    for _, coin := range coins {
        if coin <= amount {
            largestCoin = coin
        } else {
            break
        }
    }
    
    // Recursive case: Solve the subproblem of making change for the remaining amount after subtracting the largest coin denomination
    subproblemCoins := coinChangeGreedy(coins, amount-largestCoin)
    
    // Combined solution: Return the minimum number of coins needed, which is one plus the minimum number of coins needed to make change for the remaining amount
    if subproblemCoins == -1 {
        return -1
    }
    return subproblemCoins + 1
}

```
Here's how the code exhibits the greedy and optimal substructure properties:

**Greedy property:** The code implements a greedy strategy by finding the largest coin denomination that is less than or equal to the given amount. This ensures that the number of coins used is minimized, as larger denominations are always preferred over smaller ones. This greedy strategy works for certain coin systems, such as the US coin system, where the coin denominations are multiples of each other. However, it may not work for all coin systems, and may not always give the optimal solution.

**Base case:** The base case checks if the given amount is zero, and returns zero if it is. This defines the optimal solution to the smallest subproblem.

**Recursive case:** In the recursive case, we subtract the largest coin denomination from the given amount, and recursively solve the subproblem of making change for the remaining amount. This step reduces the given problem to a smaller subproblem that can be solved optimally.

**Combined solution:** We return the minimum number of coins needed, which is one plus the minimum number of coins needed to make change for the remaining amount. This combination of solutions exhibits the optimal substructure property, as the optimal solution to the overall problem can be obtained by recursively applying the optimal solution to the subproblems.

