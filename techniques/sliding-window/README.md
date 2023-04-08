# Sliding Window

The sliding window technique is a commonly used algorithmic technique used to solve problems that involve arrays or lists. The technique involves dividing the array into smaller contiguous sub-arrays, called windows, and moving these windows over the original array to perform some specific operation.

The idea is to maintain a "window" of elements in the array, and move that window one element at a time to process each possible sub-array of size k (where k is a fixed size). This technique is often used when the problem requires finding some pattern or sub-array in the given array.

The sliding window technique has a time complexity of O(n), where n is the size of the array. It is often used to solve problems related to string, arrays, and lists.

Here is an example of how the sliding window technique can be used to solve a problem:

**Problem statement:** Given an array of integers, find the maximum sum of any contiguous subarray of size k.

Solution using the sliding window technique:

1. Initialize a variable 'sum' to 0.
1. Compute the sum of the first k elements and assign it to 'sum'.
1. Create a loop that starts at k and runs until the end of the array.
1. At each iteration, subtract the element that is k steps behind from the sum and add the current element to the sum.
1. Compare the current sum with the maximum sum seen so far and update the maximum sum if necessary.
1. Return the maximum sum seen so far.

This algorithm runs in `O(n)` time since it only needs to iterate over the array once.

## Easy Problems

1. Maximum sum subarray of size K
1. Smallest subarray with a given sum
1. Longest substring with at most K distinct characters
1. Maximum of all subarrays of size K
1. Minimum window substring
1. Subarray product less than K
1. Check if a string contains all anagrams of another string
1. Longest Repeating Character Replacement
1. Minimum size subarray sum
1. Find all anagrams in a string
1. Longest Substring Without Repeating Characters
1. Subarrays with distinct integers
1. Longest Continuous Increasing Subsequence
1. Longest Mountain in Array
1. Minimum Adjacent Swaps for K Consecutive Ones
1. Find the Duplicate Number
1. Fruit into Baskets
1. Maximum Erasure Value
1. Longest Turbulent Subarray
1. Maximum Length of Repeated Subarray
1. Grumpy Bookstore Owner
1. Minimum Operations to Reduce X to Zero
1. Binary Subarrays With Sum
1. Degree of an Array
1. Partition Array Into Three Parts With Equal Sum
1. Longest Chunked Palindrome Decomposition
1. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
1. Find All Numbers Disappeared in an Array
1. Minimum Number of Operations to Move All Balls to Each Box
1. Maximum Number of Vowels in a Substring of Given Length

## Intermediate Problems

1. Minimum window substring
1. Longest substring with at most K distinct characters
1. Longest substring without repeating characters
1. Maximum length of a concatenated string with unique characters
1. Minimum size subarray sum
1. Longest subarray with absolute difference less than or equal to limit
1. Number of subarrays with bounded maximum value
1. Longest subarray with sum divisible by K
1. Minimum window subsequence
1. Minimum operations to make a subsequence
1. Minimum window containing subsequences of all strings
1. Longest subarray with equal number of 0s and 1s
1. Subarrays with K different integers
1. Minimum swaps to group all 1's together
1. Find the longest substring with k unique characters in a given string
1. Maximum sum of two non-overlapping subarrays
1. Maximum sum of subsequence with no adjacent elements
1. Maximum product of two non-overlapping subarrays
1. Maximum sum of subsequence with no adjacent elements
1. Maximum sum of subsequence with at least K elements
1. Maximum length of subsequence with difference between adjacent elements equal to 1
1. Maximum sum subsequence with at least one element picked from each half
1. Smallest subsequence with sum greater than or equal to target
1. Maximum average subarray of size K
1. Largest sum of averages
1. Longest contiguous subarray with all ones
1. Longest substring with unique characters after k replacements
1. Maximum product subarray of size k
1. Maximum sum of subsequence with no 3 adjacent elements
1. Minimum cost to connect sticks

## Illustrative examples

Here are some examples of applying the sliding window technique:

### Maximum sum subarray of size K

```go
func maxSumSubarrayOfSizeK(k int, arr []int) int {
    windowStart, windowSum, maxSum := 0, 0, 0
    for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
        windowSum += arr[windowEnd]
        if windowEnd >= k-1 {
            maxSum = max(maxSum, windowSum)
            windowSum -= arr[windowStart]
            windowStart++
        }
    }
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### Smallest subarray with a given sum
```go
func smallestSubarrayWithSum(s int, arr []int) int {
    windowStart, windowSum, minLen := 0, 0, math.MaxInt32
    for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
        windowSum += arr[windowEnd]
        for windowSum >= s {
            minLen = min(minLen, windowEnd-windowStart+1)
            windowSum -= arr[windowStart]
            windowStart++
        }
    }
    if minLen == math.MaxInt32 {
        return 0
    }
    return minLen
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### Longest substring with at most K distinct characters
```go
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    freq := make(map[byte]int)
    windowStart, maxLen := 0, 0
    for windowEnd := 0; windowEnd < len(s); windowEnd++ {
        freq[s[windowEnd]]++
        for len(freq) > k {
            freq[s[windowStart]]--
            if freq[s[windowStart]] == 0 {
                delete(freq, s[windowStart])
            }
            windowStart++
        }
        maxLen = max(maxLen, windowEnd-windowStart+1)
    }
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### Maximum length of a concatenated string with unique characters
```go
func maxLength(arr []string) int {
    ans := 0
    dfs(0, []byte{}, &ans, arr)
    return ans
}

func dfs(idx int, cur []byte, ans *int, arr []string) {
    if idx == len(arr) {
        if len(cur) > *ans {
            *ans = len(cur)
        }
        return
    }
    dfs(idx+1, cur, ans, arr)
    for i := range arr[idx] {
        if isUnique(cur, arr[idx][i]) {
            cur = append(cur, arr[idx][i])
            dfs(idx+1, cur, ans, arr)
            cur = cur[:len(cur)-1]
        }
    }
}

func isUnique(cur []byte, c byte) bool {
    for i := range cur {
        if cur[i] == c {
            return false
        }
    }
    return true
}
```

### Subarrays with K different integers

```go
func subarraysWithKDistinct(A []int, K int) int {
    count1, count2 := make(map[int]int), make(map[int]int)
    windowStart1, window
    windowStart2, res := 0, 0
    for windowEnd := 0; windowEnd < len(A); windowEnd++ {
        count1[A[windowEnd]]++
        count2[A[windowEnd]]++
        for len(count1) > K {
            count1[A[windowStart1]]--
            if count1[A[windowStart1]] == 0 {
                delete(count1, A[windowStart1])
            }
            windowStart1++
        }
        for len(count2) >= K {
            count2[A[windowStart2]]--
            if count2[A[windowStart2]] == 0 {
                delete(count2, A[windowStart2])
            }
            windowStart2++
        }
        res += windowStart2 - windowStart1
    }
    return res
}
```









