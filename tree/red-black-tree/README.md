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

Red-Black Trees support a variety of operations, including some more complex ones. Here are some of the more complex operations that can be performed on a Red-Black Tree:

1. **Augmented Tree Queries**: In addition to storing key-value pairs in the nodes, Red-Black Trees can also store additional information (such as counts, sums, or other aggregate values) about the subtree rooted at each node. This allows for efficient queries on the tree, such as finding the sum of all values between two given keys.
1. **Range Queries**: Red-Black Trees can efficiently find all nodes whose keys lie within a given range. This can be useful for a variety of applications, such as finding all data items within a certain time interval.
1. **Tree Traversal**: Red-Black Trees can be traversed in a variety of ways, such as inorder, preorder, and postorder. These traversals can be used to perform various operations on the tree, such as printing out its contents or building a balanced binary search tree from an unbalanced one.
1. **Tree Rebalancing**: Occasionally, Red-Black Trees may become unbalanced due to insertions or deletions. Rebalancing the tree involves a series of rotations and recolorings that restore the Red-Black properties of the tree. This can be a complex operation, but it is essential for maintaining the performance guarantees of the tree.
1. **Iterators**: Red-Black Trees can be implemented with iterators that allow efficient traversal of the tree in sorted order. This can be useful for iterating through large datasets or for building other data structures that rely on sorted data.

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

1. **Databases and Search Engines**: Red-Black Trees are commonly used in databases and search engines for indexing and searching data. They can efficiently store and retrieve data based on a key value, making them useful for fast lookups.
1. **Implementing Sets and Maps**: Red-Black Trees can be used to implement sets and maps with efficient insertion, deletion, and search operations. The balanced nature of Red-Black Trees ensures that these operations have logarithmic time complexity.
1. **File Systems**: Red-Black Trees can be used in file systems to keep track of directory structures and to efficiently search for files based on their names.
1. **Network Routing**: Red-Black Trees can be used in network routing protocols to determine the best path for data packets to travel through a network. They can be used to efficiently store and search for routing information.
1. **Graph Algorithms**: Red-Black Trees can be used in graph algorithms, such as Dijkstra's shortest path algorithm and Prim's minimum spanning tree algorithm. They can be used to efficiently store and retrieve graph edges based on their weights.
1. **Computational Geometry**: Red-Black Trees can be used in computational geometry for various applications such as nearest neighbor search, range queries, and intersection detection.

## Difference between AVL trees and Red-Black trees

AVL trees and Red-Black trees are both balanced binary search trees, but there are some key differences between them:
1. **Balancing Criteria**: AVL trees are more strictly balanced than Red-Black trees. In an AVL tree, the heights of the left and right subtrees of each node differ by at most 1. In contrast, Red-Black trees relax this balance criterion to allow for some imbalance, as long as it can be corrected through rotations and color changes.
1. **Height**: AVL trees tend to be more height-balanced than Red-Black trees. This means that AVL trees can have a shorter average path length from the root to a node, making them potentially faster for certain operations. However, this comes at the cost of more expensive balancing operations.
1. **Insertion and Deletion**: Red-Black trees tend to be faster than AVL trees for insertion and deletion operations. This is because Red-Black trees have a more flexible balancing criteria, which allows for fewer rotations and color changes during insertion and deletion operations.
1. **Memory Overhead**: AVL trees require more memory overhead than Red-Black trees. This is because AVL trees require a balance factor or height value to be stored in each node, whereas Red-Black trees only require a single bit to indicate the color of each node.
1. **Applications**: AVL trees are typically used in applications where a high degree of balance is required, such as in databases or real-time systems. Red-Black trees are more commonly used in general-purpose applications, such as in C++'s STL library.

These are some of the key differences between AVL trees and Red-Black trees. Ultimately, the choice between the two will depend on the specific requirements of the application in question.

## Easy problems


1. Insert a node into a Red-Black Tree.
1. Delete a node from a Red-Black Tree.
1. Check if a given value is present in a Red-Black Tree.
1. Find the minimum value in a Red-Black Tree.
1. Find the maximum value in a Red-Black Tree.
1. Find the successor of a given value in a Red-Black Tree.
1. Find the predecessor of a given value in a Red-Black Tree.
1. Check if a Red-Black Tree is empty.
1. Count the number of nodes in a Red-Black Tree.
1. Find the height of a Red-Black Tree.
1. Perform an inorder traversal of a Red-Black Tree.
1. Perform a preorder traversal of a Red-Black Tree.
1. Perform a postorder traversal of a Red-Black Tree.
1. Count the number of black nodes in a Red-Black Tree.
1. Count the number of red nodes in a Red-Black Tree.
1. Check if a given Red-Black Tree is balanced.
1. Check if a given Red-Black Tree is a valid Binary Search Tree.
1. Find the kth smallest element in a Red-Black Tree.
1. Find the kth largest element in a Red-Black Tree.
1. Check if a given Red-Black Tree is a valid Red-Black Tree.
1. Count the number of leaf nodes in a Red-Black Tree.
1. Check if a given Red-Black Tree is a full binary tree.
1. Check if a given Red-Black Tree is a complete binary tree.
1. Mirror a given Red-Black Tree.
1. Convert a given Red-Black Tree into a binary search tree.
1. Print all the paths from the root to the leaf nodes of a Red-Black Tree.
1. Find the depth of a given node in a Red-Black Tree.
1. Find the width of a Red-Black Tree.
1. Check if two Red-Black Trees are structurally identical.
1. Find the diameter of a Red-Black Tree.

## Intermediate problems


1. Verify if a Red-Black Tree has a balanced number of black nodes.
1. Find the maximum sum of elements in a path from the root to a leaf node in a Red-Black Tree.
1. Find the maximum sum of elements in any path in a Red-Black Tree.
1. Find the lowest common ancestor of two nodes in a Red-Black Tree.
1. Check if a Red-Black Tree is a valid AVL Tree.
1. Count the number of paths in a Red-Black Tree with a given sum.
1. Count the number of paths in a Red-Black Tree that start and end at the same node.
1. Check if a Red-Black Tree is height-balanced.
1. Implement an iterator for a Red-Black Tree that returns nodes in ascending order.
1. Implement an iterator for a Red-Black Tree that returns nodes in descending order.
1. Check if a Red-Black Tree is symmetric around its center.
1. Rotate a Red-Black Tree to the left.
1. Rotate a Red-Black Tree to the right.
1. Convert a sorted array into a balanced Red-Black Tree.
1. Convert a Red-Black Tree into an AVL Tree.
1. Find the shortest path between two nodes in a Red-Black Tree.
1. Find the longest path between two nodes in a Red-Black Tree.
1. Find the height of the largest perfect subtree in a Red-Black Tree.
1. Count the number of subtrees in a Red-Black Tree with a given number of nodes.
1. Check if a Red-Black Tree is a complete k-ary tree.
1. Convert a Red-Black Tree into a 2-3 Tree.
1. Find the number of nodes in the largest complete subtree of a Red-Black Tree.
1. Implement a Red-Black Tree using a treap.
1. Count the number of nodes in the largest perfect binary subtree of a Red-Black Tree.
1. Find the diameter of a Red-Black Tree using dynamic programming.
1. Find the height of a Red-Black Tree using dynamic programming.
1. Count the number of paths in a Red-Black Tree that have a given length.
1. Check if a given binary tree can be converted into a Red-Black Tree.
1. Find the closest value in a Red-Black Tree to a given value.
1. Remove all nodes from a Red-Black Tree that lie in a given range.

