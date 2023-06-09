# Binary Tree

A binary tree is a tree data structure in which each node has at most two children, known as the left child and the right child. The left child node is typically smaller than the parent node, and the right child node is typically larger than the parent node in a binary search tree.

Binary trees are used in a wide range of applications, such as database indexing, sorting algorithms, and expression trees.

## Basic Operations
The basic operations that can be performed on a binary tree include:

1. **Insertion**: Inserting a new node into the binary tree in the appropriate location based on the key value.
2. **Deletion**: Removing a node from the binary tree while maintaining the structure of the tree.
3. **Traversal**: Visiting each node in the binary tree in a particular order. There are three common types of traversal:
    1. **In-order traversal**: Visit the left subtree, then the root, then the right subtree. [left->root->right]
    1. **Pre-order traversal**: Visit the root, then the left subtree, then the right subtree. [root->left->right]
    1. **Post-order traversal**: Visit the left subtree, then the right subtree, then the root.[left->right->root]
1. **Searching**: Searching the binary tree for a node with a specific key value.
1. **Height calculation**: Calculating the height of the binary tree, which is the number of edges on the longest path from the root to a leaf node.

## Other operations

There are several complex operations that can be performed on binary trees, which are built on top of the basic operations of insertion, deletion, traversal, searching, and height calculation. Here are some examples:

1. **Rotation**: Rotation is a technique used to balance a binary tree. In an unbalanced binary tree, nodes can be rotated to change the structure of the tree and balance it. There are two types of rotations: left rotation and right rotation.
1. **Serialization and deserialization**: Binary trees can be serialized to convert the tree data structure into a string, which can be easily stored or transmitted. Deserialization is the reverse process of creating a binary tree from its serialized string representation.
1. **Diameter calculation**: The diameter of a binary tree is the length of the longest path between any two nodes in the tree. This operation involves traversing the tree and keeping track of the maximum distance between any two nodes.
1. **Lowest common ancestor**: Given two nodes in a binary tree, the lowest common ancestor (LCA) is the lowest node that has both nodes as descendants. The LCA operation involves traversing the tree and keeping track of the path from the root to each node, and then finding the first common ancestor in the two paths.
1. **Morris traversal**: Morris traversal is a space-efficient algorithm for traversing a binary tree without using a stack or recursion. It modifies the structure of the tree temporarily to store information that can be used to backtrack to the parent node.

## API

```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type BinaryTree struct {
    Root *TreeNode
}

func (bt *BinaryTree) Insert(val int) {}
func (bt *BinaryTree) Delete(val int) {}
func (bt *BinaryTree) TraverseInOrder() []int {}
func (bt *BinaryTree) TraversePreOrder() []int {}
func (bt *BinaryTree) TraversePostOrder() []int {}
func (bt *BinaryTree) Search(val int) *TreeNode {}
func (bt *BinaryTree) Height() int {}
```

## Easy problems


1. Implement a function to calculate the size of a binary tree.
1. Implement a function to calculate the height of a binary tree.
1. Implement a function to check if a binary tree is balanced.
1. Implement a function to find the maximum value in a binary tree.
1. Implement a function to find the minimum value in a binary tree.
1. Implement a function to check if a binary tree is a binary search tree.
1. Implement a function to find the kth smallest value in a binary search tree.
1. Implement a function to find the path from the root to a given node in a binary tree.
1. Implement a function to check if two binary trees are equal.
1. Implement a function to check if a binary tree is a mirror of itself.
1. Implement a function to invert a binary tree.
1. Implement a function to convert a binary tree to its mirror.
1. Implement a function to check if a binary tree is symmetric.
1. Implement a function to check if a binary tree is a subtree of another binary tree.
1. Implement a function to find the lowest common ancestor of two nodes in a binary tree.
1. Implement a function to check if a binary tree is a complete binary tree.
1. Implement a function to calculate the diameter of a binary tree.
1. Implement a function to find the maximum sum of a path from the root to a leaf in a binary tree.
1. Implement a function to check if a binary tree is height-balanced.
1. Implement a function to check if a binary tree is full.
1. Implement a function to check if a binary tree is perfect.
1. Implement a function to print the level order traversal of a binary tree.
1. Implement a function to print the spiral order traversal of a binary tree.
1. Implement a function to print the diagonal order traversal of a binary tree.
1. Implement a function to count the number of leaf nodes in a binary tree.
1. Implement a function to count the number of nodes at a given level in a binary tree.
1. Implement a function to calculate the sum of all the nodes in a binary tree.
1. Implement a function to delete a node from a binary search tree.
1. Implement a function to convert a binary search tree to a sorted array.
1. Implement a function to find the inorder successor of a node in a binary search tree.

## Intermediate problems


1. Implement a function to check if a binary tree is a balanced binary search tree.
1. Implement a function to find the distance between two nodes in a binary tree.
1. Implement a function to serialize and deserialize a binary tree.
1. Implement a function to find the vertical sum of a binary tree.
1. Implement a function to print the boundary of a binary tree.
1. Implement a function to find the largest BST subtree in a binary tree.
1. Implement a function to find the maximum width of a binary tree.
1. Implement a function to check if a binary tree is a sum tree.
1. Implement a function to check if a binary tree is a subtree of another binary search tree.
1. Implement a function to convert a binary tree to a doubly linked list.
1. Implement a function to check if a binary tree is isomorphic to another binary tree.
1. Implement a function to check if a binary tree is a sum tree or not.
1. Implement a function to check if a binary tree is a heap.
1. Implement a function to find the lowest common ancestor of multiple nodes in a binary tree.
1. Implement a function to find the maximum sum of a path between any two nodes in a binary tree.
1. Implement a function to find the maximum difference between a node and its ancestor in a binary tree.
1. Implement a function to find the maximum sum of nodes at a given level in a binary tree.
1. Implement a function to find the diameter of a binary tree using dynamic programming.
1. Implement a function to check if a binary tree is a subtree of another binary tree without using extra space.
1. Implement a function to find the kth largest value in a binary search tree.
1. Implement a function to find the closest leaf node to a given node in a binary tree.
1. Implement a function to check if a binary tree is a foldable tree.
1. Implement a function to find the largest complete subtree in a binary tree.
1. Implement a function to find the lowest common ancestor of a binary tree in O(1) space.
1. Implement a function to check if a binary tree is a sum tree or not in O(n) time.
1. Implement a function to find the minimum depth of a binary tree using breadth-first search.
1. Implement a function to check if a binary tree is a subtree of another binary tree using Morris traversal.
1. Implement a function to convert a binary search tree to a balanced binary search tree.
1. Implement a function to find the maximum product of nodes in a binary tree.
1. Implement a function to check if a binary tree is a subtree of another binary tree using string matching.
