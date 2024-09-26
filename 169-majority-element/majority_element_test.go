package majority_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	nums []int
	want int
}

func TestExample1(t *testing.T) {
	example1 := Example{
		nums: []int{3, 2, 3},
		want: 3,
	}
	got := majorityElement(example1.nums)
	assert.Equal(t, example1.want, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		nums: []int{2, 2, 1, 1, 1, 2, 2},
		want: 2,
	}
	got := majorityElement(example2.nums)
	assert.Equal(t, example2.want, got)
}
