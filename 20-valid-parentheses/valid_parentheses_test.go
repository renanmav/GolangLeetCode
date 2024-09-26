package valid_parentheses

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
		s:    "()",
		want: true,
	}
	got := isValid(example1.s)
	assert.Equal(t, example1.want, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		s:    "()[]{}",
		want: true,
	}
	got := isValid(example2.s)
	assert.Equal(t, example2.want, got)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		s:    "(]",
		want: false,
	}
	got := isValid(example3.s)
	assert.Equal(t, example3.want, got)
}

func TestExample4(t *testing.T) {
	example4 := Example{
		s:    "([])",
		want: true,
	}
	got := isValid(example4.s)
	assert.Equal(t, example4.want, got)
}

func TestExample5(t *testing.T) {
	example5 := Example{
		s:    "([)]",
		want: false,
	}
	got := isValid(example5.s)
	assert.Equal(t, example5.want, got)
}

func TestExample6(t *testing.T) {
	example6 := Example{
		s:    "]",
		want: false,
	}
	got := isValid(example6.s)
	assert.Equal(t, example6.want, got)
}

func TestExample7(t *testing.T) {
	example7 := Example{
		s:    "((",
		want: false,
	}
	got := isValid(example7.s)
	assert.Equal(t, example7.want, got)
}

func TestExample8(t *testing.T) {
	example8 := Example{
		s:    "){",
		want: false,
	}
	got := isValid(example8.s)
	assert.Equal(t, example8.want, got)
}
