# Fast and Slow pointers

Fast and slow pointers are commonly used in computer science algorithms, particularly in linked lists and other data structures.

A fast pointer is a pointer that moves through the data structure more quickly than a slow pointer. 
This can be used to achieve a number of different outcomes, depending on the specific algorithm being 
used. For example, in cycle detection algorithms for linked lists, a fast pointer is used to detect if 
there is a loop in the list by moving through it at a faster rate than the slow pointer.

On the other hand, a slow pointer is a pointer that moves through the data structure at a slower rate 
than the fast pointer. This can be used to achieve a number of different outcomes as well, depending 
on the specific algorithm being used. For example, in algorithms that require finding the middle element 
of a linked list, a slow pointer can be used to find this element by moving through the list at a slower 
rate than the fast pointer, which moves through the list twice as quickly.

Overall, fast and slow pointers are useful tools for optimizing algorithms that involve traversing data structures, 
and can be used in a variety of different ways to achieve specific outcomes.

## Easy problems

1. Remove the Nth node from the end of a linked list.
1. Determine if a linked list has a cycle.
1. Find the middle node of a linked list.
1. Reverse a linked list.
1. Merge two sorted linked lists.
1. Remove duplicates from a sorted linked list.
1. Find the kth element from the end of a linked list.
1. Remove all elements from a linked list that are greater than a given value.
1. Determine if a string is a palindrome using a linked list.
1. Find the intersection of two linked lists.
1. Determine if a linked list is a palindrome.
1. Remove all elements from a linked list that are less than a given value.
1. Find the length of a linked list.
1. Check if a linked list is sorted.
1. Rotate a linked list by K nodes.
1. Remove all elements from a linked list that occur more than once.
1. Swap every two adjacent nodes in a linked list.
1. Find the maximum element in a linked list.
1. Reverse a linked list in groups of K.
1. Find the sum of two linked lists that represent numbers.
1. Remove all even elements from a linked list.
1. Remove all odd elements from a linked list.
1. Find the median of a sorted linked list.
1. Remove duplicates from an unsorted linked list.
1. Sort a linked list using the merge sort algorithm.
1. Find the minimum element in a linked list.
1. Find the second largest element in a linked list.
1. Find the nth node from the beginning of a linked list.
1. Find the sum of all elements in a linked list.
1. Remove the middle element from a linked list.
 
## Intermediate problems
1. Given a linked list, rotate the list to the right by k places.
1. Given a linked list with a cycle, find the length of the cycle.
1. Given two linked lists that represent numbers, find the product of the numbers.
1. Given a linked list and a value x, partition the list such that all nodes less than x come before nodes greater than or equal to x.
1. Given a linked list, remove all nodes that have duplicate values, leaving only distinct values.
1. Given a linked list and a positive integer k, reverse the nodes of the list in groups of k.
1. Given a linked list with a cycle, find the starting node of the cycle.
1. Given a linked list with a cycle, remove the cycle.
1. Given a linked list with a cycle, find the length of the list.
1. Given two sorted linked lists, merge them into a single sorted list.
1. Given a linked list, remove all nodes that are duplicates of previous nodes.
1. Given a linked list, find the kth node from the end of the list using only one pass.
1. Given a linked list, remove all nodes with odd indices.
1. Given a linked list, remove all nodes with even indices.
1. Given a linked list, check if it is a palindrome in O(n) time and O(1) space.
1. Given a linked list, determine if it has an even or odd length.
1. Given a linked list and a value x, partition the list into two parts such that nodes less than x come before nodes greater than or equal to x, while preserving the original order of the nodes.
1. Given a linked list, rearrange the list such that all even-indexed nodes come before all odd-indexed nodes.
1. Given a linked list with a cycle, find the length of the longest non-cyclic path.
1. Given a linked list, remove all nodes that are greater than their next node.
1. Given a linked list, remove all nodes that are smaller than their previous node.
1. Given a linked list, swap every k nodes (where k is a positive integer) with their adjacent nodes.
1. Given a linked list, sort the list using the quicksort algorithm.
1. Given a linked list, remove all nodes with values less than the median value.
1. Given a linked list, remove all nodes with values greater than the median value.
1. Given a linked list, rotate the list to the left by k places.
1. Given a linked list, remove all nodes that have values that appear more than once in the list.
1. Given a linked list, find the node where the list divides into two equal halves.
1. Given a linked list and a value x, remove all nodes that have value x while preserving the original order of the nodes.
1. Given a linked list, remove all nodes that are greater than or equal to a given value.

# Illustrative examples

### Finding the middle node of a linked list:

```go
type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```
In this example, we use two pointers, slow and fast, to traverse the linked list. 
The slow pointer moves one step at a time, while the fast pointer moves two steps 
at a time. When the fast pointer reaches the end of the list, the slow pointer 
will be at the middle node.

### Removing the Nth node from the end of a linked list:

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    slow, fast := dummy, head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

```
In this example, we use two pointers, slow and fast, to traverse the linked list. 
The fast pointer moves n nodes ahead of the slow pointer. Then, we move both pointers 
at the same time until the fast pointer reaches the end of the list. At this point, 
the slow pointer will be pointing at the node to be removed, so we simply remove it 
by updating the Next pointer of the previous node.

### Checking if a linked list has a cycle:

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}

```
In this example, we use two pointers, slow and fast, to traverse the linked list. 
The slow pointer moves one step at a time, while the fast pointer moves two steps 
at a time. If there is a cycle in the list, the fast pointer will eventually catch 
up to the slow pointer. If there is no cycle, the fast pointer will reach the end of the list.

### Detecting the start of a cycle in a linked list:

```go
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    if fast == nil || fast.Next == nil {
        return nil
    }
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
```
In this example, we use two pointers, slow and fast, to detect the start of a cycle in a linked list. We start by setting slow and fast to head and iterating over the list. If there is a cycle in the list, the fast pointer will eventually catch up to the slow pointer. We then set slow back to head and iterate over the list again, this time moving both pointers at the same time. The point at which they meet is the start of the cycle.

### Finding the kth to last element in a linked list:

```go
func kthToLast(head *ListNode, k int) int {
    slow, fast := head, head
    for i := 0; i < k; i++ {
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    return slow.Val
}

```

In this example, we use two pointers, slow and fast, to find the kth to last element in a linked list. 
We start by setting fast to head and iterating over the list k times. Then, we move both pointers at 
the same time until fast reaches the end of the list. At this point, slow will be pointing at the kth 
to last element.



