package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	nums         []int
	k            int
	expectedNums []int
}

func TestExample1(t *testing.T) {
	example1 := Example{
		nums:         []int{1, 1, 1, 2, 2, 3, 3},
		k:            6,
		expectedNums: []int{1, 1, 2, 2, 3, 3},
	}

	k := removeDuplicates(example1.nums)

	assert.Equal(t, example1.k, k)
	for i := 0; i < k; i++ {
		assert.Equal(t, example1.expectedNums[i], example1.nums[i])
	}
}

func TestExample2(t *testing.T) {
	example2 := Example{
		nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		k:            7,
		expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
	}

	k := removeDuplicates(example2.nums)

	assert.Equal(t, example2.k, k)
	for i := 0; i < k; i++ {
		assert.Equal(t, example2.expectedNums[i], example2.nums[i])
	}
}
