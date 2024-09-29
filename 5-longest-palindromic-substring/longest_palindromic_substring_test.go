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

func TestExample3(t *testing.T) {
	example3 := Example{
		s:    "a",
		want: "a",
	}
	got := longestPalindrome(example3.s)
	assert.Equal(t, example3.want, got)
}

func TestExample4(t *testing.T) {
	example4 := Example{
		s:    "aa",
		want: "aa",
	}
	got := longestPalindrome(example4.s)
	assert.Equal(t, example4.want, got)
}

func TestExample5(t *testing.T) {
	example5 := Example{
		s:    "abaxabaxabb",
		want: "baxabaxab",
	}
	got := longestPalindrome(example5.s)
	assert.Equal(t, example5.want, got)
}

func TestExample6(t *testing.T) {
	example6 := Example{
		s:    "aracecarsuperfast",
		want: "racecar",
	}
	got := longestPalindrome(example6.s)
	assert.Equal(t, example6.want, got)
}
