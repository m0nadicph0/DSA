# Heap

A Heap is a specialized tree-based data structure that satisfies the heap property, 
which can be either the `min-heap` property or the `max-heap` property. In a `min-heap`, 
the parent node is always smaller than or equal to its children, and the root node 
is the smallest element in the heap. In a `max-heap`, the parent node is always greater 
than or equal to its children, and the root node is the largest element in the heap.

A Heap can be implemented as an array, where the root element is stored at index 0 and 
the left and right children of a node at index i are stored at indices `2i+1` and `2i+2`, 
respectively. Because of this property, a heap can be stored efficiently in an array.

The main advantage of a heap data structure is its efficiency in finding the minimum or 
maximum element in the heap. In a min-heap, the minimum element is always at the root, 
and in a max-heap, the maximum element is always at the root. This operation takes constant 
time, `O(1)`, since the root element is always at the same index.

Heaps are commonly used to implement priority queues, which are data structures that 
maintain a set of elements with priorities and allow efficient access to the element 
with the highest or lowest priority. Priority queues are used in many algorithms, 
including Dijkstra's algorithm for finding the shortest path in a graph and Huffman 
coding for data compression.

Common operations on a Heap include insertion, deletion, and retrieval of the minimum 
or maximum element. The time complexity of these operations depends on the size of the 
heap and is typically O(log n), where n is the number of elements in the heap.

## Application of heap

A heap is a useful data structure that can be used in a variety of applications, such as:

1. **Sorting**: heaps are used to implement efficient sorting algorithms like heapsort.
2. **Priority queues**: heaps can be used to implement efficient priority queues where the elements with the highest or lowest priority are always at the root of the heap.
3. **Graph algorithms**: heaps can be used in graph algorithms like Dijkstra's shortest path algorithm and Prim's minimum spanning tree algorithm to efficiently find the next vertex to visit.
4. **Memory management**: heaps are used in dynamic memory allocation systems to keep track of allocated and deallocated memory blocks.
5. **Event-driven simulation**: heaps can be used to implement an event-driven simulation where events are sorted by their scheduled time and processed in the order of their occurrence.

Overall, heaps are useful for any application that requires efficient access to the minimum or maximum element in a collection, or efficient insertion and deletion of elements based on their priority or value.

## Basic operations

The basic operations of a Heap data structure include:

1. `Insertion`: adding a new element to the heap in its appropriate position based on the heap property (either min-heap or max-heap).
2. `Deletion`: removing the root element (either the smallest or largest element in the heap depending on whether it is a min-heap or max-heap) and adjusting the remaining elements to maintain the heap property.
3. `Peek`: retrieving the value of the root element without removing it from the heap.
4. `Heapify`: converting an array of elements into a heap structure by applying the heap property to each element.
5. `Merge`: combining two heaps into a single heap while maintaining the heap property.

The time complexity of these operations can vary depending on the implementation, but the most 
common time complexities are `O(log n)` for insertion, deletion, and peek, and `O(n)` for heapify 
and merge.

## Easy problems

1. Find the kth smallest element in an unsorted array of integers.
1. Merge k sorted arrays of integers into one sorted array.
1. Find the kth largest element in a binary search tree.
1. Sort an array using heapsort algorithm.
1. Find the median of a stream of integers.
1. Find the maximum sum of any subarray of size k in an array of integers.
1. Implement a priority queue using a heap.
1. Find the kth smallest element in a binary search tree.
1. Find the top k frequent elements in an array of integers.
1. Implement a stack using a heap.
1. Find the maximum sum of any subarray of size k in a stream of integers.
1. Implement a queue using a heap.
1. Find the smallest range that contains at least one element from each of k sorted lists.
1. Find the kth smallest element in a sorted matrix.
1. Merge k sorted linked lists into one sorted linked list.
1. Given a string, find the k most frequent characters in the string.
1. Given an array of integers, find the smallest range that contains all the elements of the array.
1. Given a binary tree, find the maximum element in the binary tree.
1. Given a binary tree, find the minimum element in the binary tree.
1. Given a binary tree, find the kth largest element in the binary tree.
1. Given a binary tree, find the kth smallest element in the binary tree.
1. Given a binary tree, check if it is a max heap or a min heap.
1. Given a binary tree, convert it into a max heap or a min heap.
1. Given a binary tree, print the elements in a level-order traversal using a heap.
1. Given a binary tree, print the elements in a depth-first traversal using a heap.
1. Given a binary tree, find the height of the binary tree using a heap.
1. Given a binary tree, find the diameter of the binary tree using a heap.
1. Given a binary tree, check if it is a complete binary tree using a heap.
1. Given a binary tree, check if it is a balanced binary tree using a heap.
1. Given a binary tree, check if it is a binary search tree using a heap.


## Intermediate problems

1. Given an array of integers, find the kth smallest element in the range [left, right].
1. Given an array of intervals, merge overlapping intervals.
1. Given a binary tree, find the kth largest element in the right subtree of a node.
1. Given a binary tree, find the kth smallest element in the left subtree of a node.
1. Implement a stack using two heaps.
1. Implement a queue using two heaps.
1. Given a matrix of integers, find the kth smallest element in the matrix.
1. Given a matrix of integers, find the kth largest element in the matrix.
1. Given a binary tree, find the sum of all the elements greater than or equal to k.
1. Given a binary tree, find the sum of all the elements less than or equal to k.
1. Given a binary tree, find the distance between two nodes.
1. Given a binary tree, find the vertical sum of the nodes.
1. Given a binary tree, find the level that has the maximum sum of the nodes.
1. Given a binary tree, find the level that has the minimum sum of the nodes.
1. Given a binary tree, find the deepest leaf node.
1. Given a binary tree, find the height of the deepest leaf node.
1. Given a binary tree, find the length of the longest path from the root to a leaf node.
1. Given a binary tree, find the length of the shortest path from the root to a leaf node.
1. Given a binary tree, find the length of the longest path between any two nodes.
1. Given a binary tree, find the length of the shortest path between any two nodes.
1. Given a binary tree, find the closest common ancestor of two nodes.
1. Given a binary tree, find the farthest node from a given node.
1. Given a binary tree, find the node with the maximum sum of the subtree rooted at that node.
1. Given a binary tree, find the node with the maximum average of the subtree rooted at that node.
1. Given a binary tree, find the node with the maximum number of children.
1. Given a binary tree, find the node with the maximum number of leaves.
1. Given a binary tree, find the node with the maximum depth.
1. Given a binary tree, find the node with the minimum depth.
1. Given a binary tree, find the node with the maximum width.
1. Given a binary tree, find the node with the maximum diameter.



