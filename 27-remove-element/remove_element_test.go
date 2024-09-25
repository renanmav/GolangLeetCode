package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	nums []int
	val  int
}

func TestExample1(t *testing.T) {
	example1 := Example{
		nums: []int{3, 2, 2, 3},
		val:  3,
	}

	assert.Equal(t, 2, removeElement(example1.nums, example1.val))
}

func TestExample2(t *testing.T) {
	example2 := Example{
		nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:  2,
	}

	assert.Equal(t, 5, removeElement(example2.nums, example2.val))
}
