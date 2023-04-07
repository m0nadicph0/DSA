# Doubly Linked lists

A doubly linked list is a type of linked list where each node has pointers to both the previous and next nodes in the list. 
This allows for efficient traversal in both directions.

The basic operations for a doubly linked list include:

1. **Insertion**: Adding a new node to the list at a specified position.
    * To insert a node at the beginning of the list, set the new node's next pointer to the current head of the list and set the current head's prev pointer to the new node.
    * To insert a node at the end of the list, set the new node's prev pointer to the current tail of the list and set the current tail's next pointer to the new node.
    * To insert a node at a specific position, find the node at the position and set its prev pointer to the new node, set the new node's next pointer to the node at the position, set the new node's prev pointer to the node before the position, and set the node before the position's next pointer to the new node.
2. **Deletion**: Removing a node from the list at a specified position.
    * To delete the first node in the list, set the head of the list to the second node and set the second node's prev pointer to NULL.
    * To delete the last node in the list, set the tail of the list to the second-to-last node and set the second-to-last node's next pointer to NULL.
    * To delete a node at a specific position, find the node at the position, set the node before the position's next pointer to the node after the position, set the node after the position's prev pointer to the node before the position, and free the node at the position.
3. **Traversal**: Moving through the list to access or modify its nodes.
    * To traverse forward through the list, start at the head of the list and follow the next pointers until you reach the end.
    * To traverse backward through the list, start at the tail of the list and follow the prev pointers until you reach the beginning.
4. **Search**: Finding a node in the list with a specified value.
    * To search forward through the list, start at the head of the list and follow the next pointers until you find a node with the specified value or reach the end.
    * To search backward through the list, start at the tail of the list and follow the prev pointers until you find a node with the specified value or reach the beginning.

## API

```go
type DoublyLinkedList interface {
   InsertAtBeginning(value int)
   InsertAtEnd(value int)
   InsertAtPosition(value int, position int)
   DeleteAtBeginning() error
   DeleteAtEnd() error
   DeleteAtPosition(position int) error
   TraverseForward() error
   TraverseBackward() error
   Search(value int) (bool, error)
}
```

## Easy Problems

1. Implement a function to create a new doubly linked list.
1. Implement a function to insert a node at the beginning of the list.
1. Implement a function to insert a node at the end of the list.
1. Implement a function to insert a node at a specific position in the list.
1. Implement a function to delete the first node of the list.
1. Implement a function to delete the last node of the list.
1. Implement a function to delete a node at a specific position in the list.
1. Implement a function to traverse the list forward.
1. Implement a function to traverse the list backward.
1. Implement a function to search for a node with a specific value in the list.
1. Implement a function to find the length of the list.
1. Implement a function to check if the list is empty.
1. Implement a function to reverse the list.
1. Implement a function to sort the list in ascending order.
1. Implement a function to sort the list in descending order.
1. Implement a function to remove duplicate nodes from the list.
1. Implement a function to get the nth node from the list.
1. Implement a function to insert a node before a specific node in the list.
1. Implement a function to insert a node after a specific node in the list.
1. Implement a function to swap two nodes in the list.
1. Implement a function to find the maximum element in the list.
1. Implement a function to find the minimum element in the list.
1. Implement a function to find the sum of all elements in the list.
1. Implement a function to find the average of all elements in the list.
1. Implement a function to find the median of all elements in the list.
1. Implement a function to find the mode of all elements in the list.
1. Implement a function to find the range of all elements in the list.
1. Implement a function to find the variance of all elements in the list.
1. Implement a function to find the standard deviation of all elements in the list.
1. Implement a function to delete all nodes from the list.

## Intermediate Problems

1. Implement a function to rotate the list by k positions to the right.
1. Implement a function to delete all nodes with a specific value from the list.
1. Implement a function to merge two sorted doubly linked lists.
1. Implement a function to partition the list around a specific value.
1. Implement a function to split the list into two halves.
1. Implement a function to find the kth largest element in the list.
1. Implement a function to find the kth smallest element in the list.
1. Implement a function to find the nth node from the end of the list.
1. Implement a function to find the middle node of the list.
1. Implement a function to find the intersection of two doubly linked lists.
1. Implement a function to find the union of two doubly linked lists.
1. Implement a function to check if the list is a palindrome.
1. Implement a function to check if the list is cyclic.
1. Implement a function to remove the cycle from a cyclic list.
1. Implement a function to find the length of the cycle in a cyclic list.
1. Implement a function to swap every two nodes in the list.
1. Implement a function to reverse every k nodes in the list.
1. Implement a function to check if the list is a doubly palindrome.
1. Implement a function to find the longest increasing subsequence in the list.
1. Implement a function to find the longest decreasing subsequence in the list.
1. Implement a function to find the longest bitonic subsequence in the list.
1. Implement a function to find the longest common subsequence between two doubly linked lists.
1. Implement a function to find the longest palindrome subsequence in the list.
1. Implement a function to find the longest subarray with alternating sign in the list.
1. Implement a function to find the longest subarray with equal elements in the list.
1. Implement a function to find the longest subarray with consecutive elements in the list.
1. Implement a function to find the longest subarray with non-negative elements in the list.
1. Implement a function to find the longest subarray with non-positive elements in the list.
1. Implement a function to find the longest subarray with odd elements in the list.
1. Implement a function to find the longest subarray with even elements in the list.


