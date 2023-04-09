# Red-Black Trees

A red-black tree is a self-balancing binary search tree, meaning that it automatically adjusts its structure to ensure that the tree remains balanced as nodes are added or removed. The balancing is achieved through the use of "coloring" on the nodes, which can be either red or black.

## Basic operations

The basic operations that can be performed on a red-black tree are:

1. **Insertion**: To insert a new node into the tree, the tree is first searched as in a regular binary search tree. Once the new node is inserted, the tree is rebalanced using a series of rotations and color changes to ensure that the tree remains balanced.
1. **Deletion**: To delete a node from the tree, the tree is first searched for the node to be deleted. Once the node is found, it is removed from the tree and the tree is rebalanced in the same way as during insertion.
1. **Search**: To search for a particular key in the tree, the tree is traversed starting from the root node and comparing each key with the search key until a match is found.
1. **Minimum/Maximum**: To find the minimum or maximum value in the tree, we simply follow the left (minimum) or right (maximum) child pointers from the root node until we reach the bottom of the tree.
1. **Successor/Predecessor**: To find the next or previous value in the tree, we first find the node containing the current value, and then follow a set of rules to find the next (successor) or previous (predecessor) node based on the order of the keys in the tree.

All of these operations can be performed in `O(log n)` time, making red-black trees a very efficient data structure for many applications.

## Complex operations

## Examples

```text
                8B
              /    \
            4R      12R
          /   \   /    \
        2B    6B  10B  14B
             / \ 
            5R 7R

```
In this tree, each node is either red or black, and the root node (8) is black. Each path from the root to a leaf node contains the same number of black nodes, which ensures that the tree is balanced. Note that the red nodes are not adjacent to each other, and that there are no two adjacent red nodes along the same path.

This particular tree represents the following set of keys: {2, 4, 5, 6, 7, 8, 10, 12, 14}. The black nodes are labeled with a "B" and the red nodes are labeled with an "R".

```text
                                15B
                          /            \
                        7R             18R
                     /      \        /       \
                   3B       12B    17B       22B
                  /  \      /   \          /    \
                2R   5R   10R  13R     20R     25R
               /           \              \
             1B             6B            21B

```
In this tree, each node is either red or black, and the root node (15) is black. Each path from the root to a leaf node contains the same number of black nodes, which ensures that the tree is balanced. Note that the red nodes are not adjacent to each other, and that there are no two adjacent red nodes along the same path.

This particular tree represents the following set of keys: {1, 2, 3, 5, 6, 7, 10, 12, 13, 15, 17, 18, 20, 21, 22, 25}. The black nodes are labeled with a "B" and the red nodes are labeled with an "R".


## API

```go
type Node struct {
    Key    int
    Value  interface{}
    Color  bool // true for black, false for red
    Parent *Node
    Left   *Node
    Right  *Node
}

type RedBlackTree struct {
    Root *Node
}

func NewRedBlackTree() *RedBlackTree {
    return &RedBlackTree{Root: nil}
}

func (t *RedBlackTree) Insert(key int, value interface{}) {
    // ... insert a new node with the given key and value into the tree ...
    // ... perform balancing to maintain Red-Black properties ...
}

func (t *RedBlackTree) Delete(key int) {
    // ... find and remove the node with the given key ...
    // ... perform balancing to maintain Red-Black properties ...
}

func (t *RedBlackTree) Search(key int) (interface{}, bool) {
    // ... traverse the tree to find the node with the given key ...
}

func (t *RedBlackTree) Minimum() (int, interface{}, bool) {
    // ... traverse the tree to find the node with the smallest key ...
}

func (t *RedBlackTree) Maximum() (int, interface{}, bool) {
    // ... traverse the tree to find the node with the largest key ...
}

func (t *RedBlackTree) Successor(key int) (int, interface{}, bool) {
    // ... find the node with the given key ...
    // ... traverse the tree to find the node with the next smallest key ...
}

func (t *RedBlackTree) Predecessor(key int) (int, interface{}, bool) {
    // ... find the node with the given key ...
    // ... traverse the tree to find the node with the next largest key ...
}

```

## Uses and applications

## Difference between AVL trees and Red-Black trees


