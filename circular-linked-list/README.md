# Circular Linked List

A circular linked list is a data structure that consists of a sequence of nodes, 
where each node contains a value and a pointer to the next node in the sequence.
In a circular linked list, the last node points back to the first node, forming a loop.

The basic operations of a circular linked list are:

1. `Insertion`: To insert a new node into the circular linked list, we need to create a new node with the given value and set its pointer to the next node in the sequence. Then, we update the pointer of the previous node to point to the new node, effectively inserting it into the sequence.
1. `Deletion`: To delete a node from the circular linked list, we need to update the pointer of the previous node to point to the next node in the sequence, effectively skipping over the node to be deleted. Then, we can free the memory occupied by the deleted node.
1. `Traversal`: To traverse a circular linked list, we start at any node and follow the pointers to the next nodes until we reach the starting node again. We can perform any desired operations on each node during the traversal.
1. `Searching`: To search for a specific value in the circular linked list, we can start at any node and follow the pointers to the next nodes until we find a node with the desired value or we reach the starting node again.
1. `Concatenation`: To concatenate two circular linked lists, we need to find the last node of the first list and set its pointer to the first node of the second list. Then, we update the pointer of the last node of the second list to point to the first node of the first list, effectively forming a new circular linked list that contains all the nodes of both lists.

These basic operations allow us to manipulate and process the data stored in a circular linked list.

## Applications

Circular linked lists have various applications in computer science and programming. Some of the common applications are:

1. **Implementation of data structures**: Circular linked lists can be used to implement various data structures such as a circular buffer, a circular queue, and a circular stack. These data structures have applications in computer networking, operating systems, and real-time systems, where data needs to be processed in a circular manner.
1. **Music and media players**: Circular linked lists can be used to implement music and media players, where a playlist is stored as a circular linked list. Each node in the list represents a song or a media file, and the player traverses the list to play the files in sequence.
1. **Navigation systems**: Circular linked lists can be used to implement navigation systems in GPS devices and maps, where the nodes represent the various locations or landmarks, and the list represents a route or a path between them.
1. **Game development**: Circular linked lists can be used to implement game loops and game states in video games, where the game world is represented as a circular linked list of objects and the game loop traverses the list to update and render the objects.
1. **Simulation and modeling**: Circular linked lists can be used to implement simulations and models in various fields such as physics, biology, and economics. The nodes in the list represent the various entities or agents, and the list represents the interactions and relationships between them.

These are just a few examples of the various applications of circular linked lists. In general, circular linked lists are useful whenever we need to process data in a circular or cyclic manner.


## Easy problems

1. Create an empty circular linked list.
1. Insert a new node at the beginning of the list.
1. Insert a new node at the end of the list.
1. Delete the first node of the list.
1. Delete the last node of the list.
1. Traverse the list and print the data of each node.
1. Find the length of the list.
1. Search for a node with a given value.
1. Delete all nodes with a given value.
1. Reverse the list.
1. Split the list into two equal halves.
1. Merge two sorted circular linked lists.
1. Remove duplicates from the list.
1. Find the maximum and minimum values in the list.
1. Find the sum of all values in the list.
1. Find the average of all values in the list.
1. Find the median of all values in the list.
1. Check if the list is palindrome.
1. Shift the list left by a given number of nodes.
1. Shift the list right by a given number of nodes.
1. Rotate the list by a given number of nodes.
1. Insert a new node at a given position in the list.
1. Delete a node at a given position in the list.
1. Sort the list in ascending order.
1. Sort the list in descending order.
1. Swap two nodes in the list.
1. Find the middle node of the list.
1. Check if the list contains a loop.
1. Find the starting node of a loop in the list.
1. Remove a loop from the list.

## Intermediate problems

Insert a new node in a sorted circular linked list.
1. Implement a circular queue using a circular linked list.
1. Implement a stack using a circular linked list.
1. Implement a deque (double-ended queue) using a circular linked list.
1. Implement a priority queue using a circular linked list.
1. Find the kth node from the end of the list.
1. Find the intersection point of two lists.
1. Find the union of two sorted lists.
1. Find the intersection of two sorted lists.
1. Find the median of two sorted lists.
1. Find the kth largest element in the list.
1. Find the sum of all subarrays of the list.
1. Find the product of all subarrays of the list.
1. Find the maximum subarray sum.
1. Find the maximum subarray product.
1. Find the maximum circular subarray sum.
1. Find the maximum circular subarray product.
1. Find the longest increasing subsequence.
1. Find the longest decreasing subsequence.
1. Find the longest bitonic subsequence.
1. Find the longest palindromic subsequence.
1. Find the longest common subsequence of two lists.
1. Find the longest repeating subsequence.
1. Find the shortest path in a graph represented by a circular linked list.
1. Convert a circular linked list to a binary tree.
1. Flatten a circular linked list that contains nested lists.
1. Implement a skip list using a circular linked list.
1. Implement a B-tree using a circular linked list.
1. Implement a red-black tree using a circular linked list.
1. Implement a AVL tree using a circular linked list.


