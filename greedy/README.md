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

## Defining the problem

Defining a greedy problem involves identifying a problem where the greedy strategy of making locally optimal choices leads to a globally optimal solution. A greedy problem can be defined by specifying the following:

1. Objective function: The objective function defines the quantity that we want to optimize. For example, in the coin change problem, the objective is to minimize the number of coins used to make change for a given amount.

2. Set of constraints: The constraints define the conditions that must be satisfied by any valid solution. For example, in the coin change problem, the constraints are that the total value of the coins used must equal the given amount, and that only valid coin denominations can be used.

3. Candidate solutions: The candidate solutions are the set of choices that can be made at each step of the problem. For a greedy problem, the candidate solutions are those that appear to be the best locally, based on some heuristic or rule.

Once we have defined a greedy problem, we can generate a set of candidate solutions by applying the greedy strategy to the problem. This involves making locally optimal choices at each step of the problem, based on some heuristic or rule. The set of candidate solutions generated by the greedy strategy can then be evaluated to determine which one is globally optimal, according to the objective function and constraints of the problem.

For example, in the coin change problem, the candidate solutions are the set of valid coin denominations that are less than or equal to the given amount. The greedy strategy is to always choose the largest denomination that is less than or equal to the remaining amount, until the remaining amount is zero. This generates a set of candidate solutions that corresponds to the sequence of coin denominations used to make change for the given amount. The globally optimal solution is the one that uses the minimum number of coins, according to the objective function and constraints of the problem.

here are some more examples of greedy problems and candidate solutions:

**Activity selection problem**: Given a set of activities, each with a start time and end time, the objective is to select the maximum number of non-overlapping activities. The candidate solutions are the activities sorted by their end times, and the greedy strategy is to always choose the activity with the earliest end time that does not overlap with any previously chosen activity.

**Huffman coding**: Given a set of characters and their frequencies, the objective is to encode the characters with variable-length codes such that the total length of the encoded message is minimized. The candidate solutions are the pairs of characters with their frequencies, and the greedy strategy is to repeatedly merge the two least frequent characters into a single node, until all nodes are merged into a single tree.

**Fractional knapsack problem**: Given a set of items, each with a weight and value, and a knapsack with a maximum weight capacity, the objective is to maximize the total value of items that can be put into the knapsack. The candidate solutions are the items sorted by their value-to-weight ratios, and the greedy strategy is to always choose the item with the highest value-to-weight ratio that can fit into the remaining capacity of the knapsack.

**Dijkstra's shortest path algorithm**: Given a weighted graph with non-negative edge weights and a source vertex, the objective is to find the shortest path to all other vertices. The candidate solutions are the vertices that have not yet been visited, and the greedy strategy is to always choose the vertex with the smallest distance from the source that has not yet been visited, and update the distances to its neighbors.

In each of these examples, the candidate solutions are generated based on some heuristic or rule that appears to be the best locally, and the globally optimal solution is determined by evaluating the set of candidate solutions according to the objective function and constraints of the problem.

##  Activity selection problem

here is an explanatory solution for the Activity selection problem based on the general format of greedy algorithms:

1.  _Define the problem and candidate solutions_: The Activity selection problem involves selecting a maximum number of non-overlapping activities from a set of activities, each with a start time and end time. The candidate solutions are the activities themselves, and the objective is to select the maximum number of non-overlapping activities.
    
2.  _Identify a criterion for evaluating the candidate solutions_: The criterion for evaluating the candidate solutions is the end time of each activity. The objective is to select the maximum number of non-overlapping activities with the earliest possible end times.
    
3.  _Define a rule for selecting the best candidate solution:_ The rule for selecting the best candidate solution is to always choose the activity with the earliest end time that does not overlap with any previously chosen activity. This is based on the assumption that selecting activities with earlier end times will leave more opportunities for selecting additional activities later.
    
4.  _Repeat the selection rule until a satisfactory solution is found:_ We start by sorting the activities in ascending order of their end times. We then iterate through the sorted list of activities, and at each step, we select the activity with the earliest end time that does not overlap with any previously chosen activity. We continue this process until all activities have been considered or no further activities can be selected.
    

By following this greedy algorithm, we can select a maximum number of non-overlapping activities with the earliest possible end times, which is the globally optimal solution to the Activity selection problem.

```go
package main

import (
    "fmt"
    "sort"
)

type Activity struct {
    start int
    end   int
}

// A function to sort activities by their end times in ascending order
func sortByEndTime(activities []Activity) {
    sort.Slice(activities, func(i, j int) bool {
        return activities[i].end < activities[j].end
    })
}

// A function to select the maximum number of non-overlapping activities
func selectActivities(activities []Activity) []Activity {
    var selected []Activity

    // Sort activities by their end times
    sortByEndTime(activities)

    // Select the first activity
    selected = append(selected, activities[0])
    lastSelected := 0

    // Iterate through the remaining activities
    for i := 1; i < len(activities); i++ {
        // Check if the current activity overlaps with the last selected activity
        if activities[i].start >= activities[lastSelected].end {
            // If not, select the current activity and update the last selected activity
            selected = append(selected, activities[i])
            lastSelected = i
        }
    }

    return selected
}

func main() {
    // Define a set of activities
    activities := []Activity{
        {1, 3},
        {2, 5},
        {3, 6},
        {4, 7},
        {5, 9},
        {6, 8},
        {7, 10},
        {8, 11},
        {9, 12},
        {10, 14},
    }

    // Select the maximum number of non-overlapping activities
    selected := selectActivities(activities)

    // Print the selected activities
    fmt.Println("Selected activities:")
    for _, activity := range selected {
        fmt.Printf("Start time: %d, End time: %d\n", activity.start, activity.end)
    }
}
```

