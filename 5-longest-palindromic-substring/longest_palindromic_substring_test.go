package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	s    string
	want string
}

func TestExample1a(t *testing.T) {
	example1 := Example{
		s:    "babad",
		want: "bab",
	}
	got := longestPalindrome(example1.s)
	assert.Contains(t, []string{example1.want, "aba"}, got)
}

func TestExample1b(t *testing.T) {
	example1 := Example{
		s:    "babad",
		want: "aba",
	}
	got := longestPalindrome(example1.s)
	assert.Contains(t, []string{example1.want, "bab"}, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		s:    "cbbd",
		want: "bb",
	}
	got := longestPalindrome(example2.s)
	assert.Equal(t, example2.want, got)
}
