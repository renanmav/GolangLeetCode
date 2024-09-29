package longest_palindromic_substring

/*
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func longestPalindrome(s string) string {
	return approach3(s)
}

// Approach 1: Check all substrings
// Time Complexity: O(n^3)
// Space Complexity: O(1)
func approach1(s string) string {
	panic("not implemented")
}

// Approach 2: Dynamic Programming
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func approach2(s string) string {
	panic("not implemented")
}

// Approach 3: Expand From Centers
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func approach3(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ { // O(nˆ2)
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Approach 4: Manacher's Algorithm
// Time Complexity: O(n)
// Space Complexity: O(n)
func approach4(s string) string {
	panic("not implemented")
}
