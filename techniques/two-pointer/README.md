# Two Pointer



The Two Pointer Technique is a popular algorithmic technique used to solve problems that involve iterating over two sequences or arrays simultaneously. The technique involves initializing two pointers at the beginning of the sequences and then manipulating the pointers until they meet or cross over, depending on the problem's requirements.

Here's a general outline of the two-pointer technique:

Initialize two pointers, one at the beginning and the other at the end (or some other appropriate location) of the sequence or array.
Iterate through the sequence or array, adjusting the pointers based on the problem's requirements, until the pointers meet or cross over.
Process the data or perform the necessary operations based on the pointers' current position.
The two-pointer technique is often used to solve problems that involve searching, sorting, or merging two arrays, finding pairs of elements with a certain property, or finding a subsequence of a certain length with a certain property. This technique can often result in a more efficient algorithm compared to other approaches because it avoids nested loops, reducing the overall time complexity of the solution.





Example 1: Finding a pair of elements in an array that sum to a target value

```go
func twoSum(nums []int, target int) []int {
    i, j := 0, len(nums)-1
    for i < j {
        sum := nums[i] + nums[j]
        if sum == target {
            return []int{i, j}
        } else if sum < target {
            i++
        } else {
            j--
        }
    }
    return []int{}
}
```

This function takes an array of integers and a target value, and returns the indices of the two elements in the array that sum to the target value. The two-pointer technique is used to iterate through the array from both ends simultaneously, adjusting the pointers based on the sum of the two elements. The time complexity of this algorithm is O(n), where n is the length of the array.

Example 2: Reversing a string

```go
func reverseString(s []byte) {
    i, j := 0, len(s)-1
    for i < j {
        s[i], s[j] = s[j], s[i]
        i++
        j--
    }
}
```

This function takes a string as a byte array and reverses it in place using the two-pointer technique. The two pointers start at the beginning and end of the array, and swap the elements at their positions until they meet in the middle. The time complexity of this algorithm is O(n/2), where n is the length of the string.

Example 3: Finding the longest palindrome in a string

```go
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)
        len2 := expandAroundCenter(s, i, i+1)
        maxLen := max(len1, len2)
        if maxLen > end-start {
            start = i - (maxLen-1)/2
            end = i + maxLen/2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left int, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

This function takes a string and finds the longest palindrome substring in it using the two-pointer technique. The function expandAroundCenter takes two pointers and expands them to both sides until they no longer form a palindrome. The function longestPalindrome iterates through each character in the string and calls expandAroundCenter with two different starting positions (one for odd-length palindromes and one for even-length palindromes). The function keeps track of the longest palindrome found so far and returns it at the end. The time complexity of this algorithm is O(n^2), where n is the length of the string.



## Problems

1. Given an array of integers, find a pair of elements that sum up to a given target value.
1. Given an array of integers, find the maximum subarray sum.
1. Given an array of integers, find the longest subarray with no more than K distinct elements.
1. Given a sorted array of integers, remove duplicates in place.
1. Given two sorted arrays of integers, merge them into a single sorted array.
1. Given an array of integers, find the longest increasing subsequence.
1. Given an array of integers, find all pairs of elements that differ by a given value.
1. Given an array of integers, find all triplets that sum up to zero.
1. Given an array of integers, find the smallest subarray with a sum greater than or equal to a given target value.
1. Given a linked list, reverse it in place.
1. Given an array of integers, find the median.
1. Given an array of integers, move all zeros to the end of the array while maintaining the relative order of non-zero elements.
1. Given an array of integers, remove all instances of a given value in place.
1. Given a linked list, remove all instances of a given value in place.
1. Given an array of integers, find the longest subarray with equal number of zeros and ones.
1. Given an array of integers, find the longest subarray with sum divisible by K.
1. Given a sorted array of integers, remove all instances of duplicates of length greater than or equal to three.
1. Given a sorted array of integers, find the closest pair to a given sum.
1. Given an array of integers, find all subarrays with a sum of zero.
1. Given an array of integers, find the subarray with the largest sum of absolute differences between adjacent elements.
1. Given an array of integers, find the maximum product of any three elements.
1. Given an array of integers, find the minimum distance between any two distinct elements.
1. Given an array of integers, find the longest subarray with at most two distinct elements.
1. Given a linked list, find the middle node.
1. Given a linked list, detect and remove a cycle, if one exists.
1. Given two sorted arrays of integers, find the median of their merged array.
1. Given a linked list, reverse every k nodes in place.
1. Given an array of integers, find the longest subarray with all elements distinct.
1. Given an array of integers, find the longest subarray with at most K distinct elements.
1. Given an array of integers, find the longest subarray with elements in non-decreasing order.
