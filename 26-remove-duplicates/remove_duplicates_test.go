package remove_duplicates

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
		nums:         []int{1, 1, 2},
		k:            2,
		expectedNums: []int{1, 2},
	}

	removeDuplicates(example1.nums)

	assert.Equal(t, example1.k, len(example1.expectedNums))
	for i := 0; i < example1.k; i++ {
		assert.Equal(t, example1.nums[i], example1.expectedNums[i])
	}
}

func TestExample2(t *testing.T) {
	example2 := Example{
		nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		k:            5,
		expectedNums: []int{0, 1, 2, 3, 4},
	}

	removeDuplicates(example2.nums)

	assert.Equal(t, example2.k, len(example2.expectedNums))
	for i := 0; i < example2.k; i++ {
		assert.Equal(t, example2.nums[i], example2.expectedNums[i])
	}
}
