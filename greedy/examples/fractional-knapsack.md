# Fractional Knapsack

The Fractional Knapsack problem is a classic optimization problem that involves selecting a subset of items with given weights and values, subject to a maximum weight constraint, in such a way as to maximize the total value of the selected items. In the fractional version of the problem, items can be selected in fractions (i.e., partial amounts) rather than as discrete units.

Here's an example scenario to illustrate the problem:

Suppose we have a knapsack with a maximum weight capacity of 15 pounds, and we are given the following five items to choose from:

| Item | Weight (lbs) | Value ($) |
| --- | --- | --- |
| A   | 2   | 10  |
| B   | 3   | 20  |
| C   | 5   | 30  |
| D   | 7   | 40  |
| E   | 1   | 5   |

The candidate solutions in the Fractional Knapsack problem are subsets of the items, where each item can be chosen in fractions (i.e., a fraction of an item's weight can be chosen). For example, we could choose half of item A (i.e., 1 pound), all of item B (i.e., 3 pounds), and a quarter of item C (i.e., 1.25 pounds). We could also choose no items at all, or any other combination of items and fractions that meets the weight constraint.

The optimal solution to the Fractional Knapsack problem is the subset of items that maximizes the total value, subject to the weight constraint. In our example scenario, the optimal solution would be to choose all of items B and C (i.e., 8 pounds total), and a fraction of item A (i.e., 1 pound), for a total weight of 9 pounds and a total value of $60.

The key to solving the Fractional Knapsack problem is to use a greedy algorithm that repeatedly selects the item with the highest value-to-weight ratio, until either the knapsack is full or all items have been considered. This is a greedy approach because it always selects the item that offers the most value for its weight, without considering how that choice might affect future choices. However, this approach is optimal for the Fractional Knapsack problem because each item can be chosen in fractions, so it is always possible to achieve the optimal solution by selecting the items with the highest value-to-weight ratios.

```go
package main

import (
	"fmt"
	"sort"
)

type Item struct {
	weight float64
	value  float64
}

type ByRatio []Item

func (a ByRatio) Len() int           { return len(a) }
func (a ByRatio) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRatio) Less(i, j int) bool { return a[i].value/a[i].weight > a[j].value/a[j].weight }

func fractionalKnapsack(items []Item, capacity float64) float64 {
	sort.Sort(ByRatio(items)) // sort items by decreasing value-to-weight ratio
	var totalValue float64
	for _, item := range items {
		if capacity == 0 {
			break
		}
		if item.weight <= capacity { // select whole item
			totalValue += item.value
			capacity -= item.weight
		} else { // select fractional item
			totalValue += item.value * (capacity / item.weight)
			capacity = 0
		}
	}
	return totalValue
}

func main() {
	items := []Item{
		{2, 10},
		{3, 20},
		{5, 30},
		{7, 40},
		{1, 5},
	}
	capacity := 15
	optimalValue := fractionalKnapsack(items, float64(capacity))
	fmt.Println(optimalValue) // prints: 60
}
```

The Item struct represents an item in the knapsack, with a weight and a value. The ByRatio type is a custom type for sorting items by decreasing value-to-weight ratio. The fractionalKnapsack function takes a slice of items and a maximum capacity, and returns the total value of the items that can be included in the knapsack, according to the greedy algorithm.

The fractionalKnapsack function first sorts the items in decreasing order of value-to-weight ratio using the sort.Sort function and the ByRatio type. Then, it iterates over the items in the sorted order, selecting whole items if there is enough capacity remaining in the knapsack, and fractional items otherwise. The total value of the selected items is accumulated and returned at the end.

In the main function, we create a slice of items and a maximum capacity, and call the fractionalKnapsack function to compute the optimal value. Finally, we print the optimal value to the console.
