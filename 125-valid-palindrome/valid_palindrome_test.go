package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	s    string
	want bool
}

func TestExample1(t *testing.T) {
	example1 := Example{
		s:    "A man, a plan, a canal: Panama",
		want: true,
	}
	got := isPalindrome(example1.s)
	assert.Equal(t, example1.want, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		s:    "race a car",
		want: false,
	}
	got := isPalindrome(example2.s)
	assert.Equal(t, example2.want, got)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		s:    " ",
		want: true,
	}
	got := isPalindrome(example3.s)
	assert.Equal(t, example3.want, got)
}
