package verifying_an_alien_dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	words  []string
	order  string
	wanted bool
}

func TestExample1(t *testing.T) {
	example1 := Example{
		words:  []string{"hello", "leetcode"},
		order:  "hlabcdefgijkmnopqrstuvwxyz",
		wanted: true,
	}
	got := isAlienSorted(example1.words, example1.order)
	assert.Equal(t, example1.wanted, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		words:  []string{"word", "world", "row"},
		order:  "worldabcefghijkmnpqstuvxyz",
		wanted: false,
	}
	got := isAlienSorted(example2.words, example2.order)
	assert.Equal(t, example2.wanted, got)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		words:  []string{"apple", "app"},
		order:  "abcdefghijklmnopqrstuvwxyz",
		wanted: false,
	}
	got := isAlienSorted(example3.words, example3.order)
	assert.Equal(t, example3.wanted, got)
}

func TestExample4(t *testing.T) {
	example4 := Example{
		words:  []string{"app", "apple"},
		order:  "abcdefghijklmnopqrstuvwxyz",
		wanted: true,
	}
	got := isAlienSorted(example4.words, example4.order)
	assert.Equal(t, example4.wanted, got)
}

func TestExample5(t *testing.T) {
	example5 := Example{
		words:  []string{"hello", "hello"},
		order:  "abcdefghijklmnopqrstuvwxyz",
		wanted: true,
	}
	got := isAlienSorted(example5.words, example5.order)
	assert.Equal(t, example5.wanted, got)
}
