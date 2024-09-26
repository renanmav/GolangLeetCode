package valid_parentheses

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false

Example 4:

Input: s = "([])"
Output: true

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			// if it's an opening bracket, push it to the stack
			stack = append(stack, char)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != char {
			// if it's a closing bracket and it doesn't match the last opening bracket, return false
			return false
		} else {
			// if it's a closing bracket and it matches the last opening bracket, pop the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
