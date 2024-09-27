package best_time_to_buy_and_sell_stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	prices []int
	want   int
}

func TestExample1(t *testing.T) {
	example1 := Example{
		prices: []int{7, 1, 5, 3, 6, 4},
		want:   5,
	}
	got := maxProfit(example1.prices)
	assert.Equal(t, example1.want, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		prices: []int{7, 6, 4, 3, 1},
		want:   0,
	}
	got := maxProfit(example2.prices)
	assert.Equal(t, example2.want, got)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		prices: []int{1, 2},
		want:   1,
	}
	got := maxProfit(example3.prices)
	assert.Equal(t, example3.want, got)
}
