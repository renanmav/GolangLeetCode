package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestExample1(t *testing.T) {
	example1 := Example{
		nums1: []int{1, 2, 3, 0, 0, 0},
		m:     3,
		nums2: []int{2, 5, 6},
		n:     3,
	}
	merge(example1.nums1, example1.m, example1.nums2, example1.n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, example1.nums1)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		nums1: []int{1},
		m:     1,
		nums2: []int{},
		n:     0,
	}
	merge(example2.nums1, example2.m, example2.nums2, example2.n)
	assert.Equal(t, []int{1}, example2.nums1)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		nums1: []int{0},
		m:     0,
		nums2: []int{1},
		n:     1,
	}
	merge(example3.nums1, example3.m, example3.nums2, example3.n)
	assert.Equal(t, []int{1}, example3.nums1)
}
