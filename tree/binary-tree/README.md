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
