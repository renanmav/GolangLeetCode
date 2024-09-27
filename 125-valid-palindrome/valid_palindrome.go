package valid_palindrome

import "strings"

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	alphanumeric := []rune{}
	for _, char := range lower {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			alphanumeric = append(alphanumeric, char)
		}
	}

	for i := 0; i < len(alphanumeric); i++ {
		if alphanumeric[i] != alphanumeric[len(alphanumeric)-i-1] {
			return false
		}
	}

	return true
}
