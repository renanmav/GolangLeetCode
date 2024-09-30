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
	return approach2(s)
}

// Approach 1: Check all substrings
// Time Complexity: O(n^3)
// Space Complexity: O(1)
func approach1(s string) string { // O(n^3)
	check := func(i, j int) bool {
		left := i
		right := j - 1
		for left < right { // O(n)
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	for length := len(s); length > 0; length-- { // O(n)
		for start := 0; start <= len(s)-length; start++ { // O(n)
			if check(start, start+length) {
				return s[start : start+length]
			}
		}
	}

	return ""
}

// Approach 2: Dynamic Programming
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func approach2(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// Create a dp table to store whether s[i:j+1] is a palindrome
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1

	// All substrings of length 1 are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Check for substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Check for substrings of length greater then 2
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1

			// If s[i] == s[j] and dp[i+1][j-1] is true, then s[i:j+1] is a palindrome
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}

	return s[start : start+maxLength]
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
