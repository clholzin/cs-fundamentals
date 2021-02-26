



# 02/26/2021

## 1512. Number of Good Pairs

### Given an array of integers nums.

A pair (i,j) is called good if nums[i] == nums[j] and i < j.

Return the number of good pairs.


### Example 1:

Input: nums = [1,2,3,1,1,3]

Output: 4

Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

### Example 2:

Input: nums = [1,1,1,1]

Output: 6

Explanation: Each pair in the array are good.

### Example 3:

Input: nums = [1,2,3]
Output: 0
 

### Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100



## 3. Longest Substring Without Repeating Characters

### Given a string s, find the length of the longest substring without repeating characters.

 

### Example 1:

Input: s = "abcabcbb"

Output: 3

Explanation: The answer is "abc", with the length of 3.

### Example 2:

Input: s = "bbbbb"

Output: 1

Explanation: The answer is "b", with the length of 1.

### Example 3:

Input: s = "pwwkew"

Output: 3

Explanation: The answer is "wke", with the length of 3.

Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

### Example 4:

Input: s = ""

Output: 0
 

### Constraints:

0 <= s.length <= 5 * 104

s consists of English letters, digits, symbols and spaces.


