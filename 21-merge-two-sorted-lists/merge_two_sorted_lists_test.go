package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	list1 *ListNode
	list2 *ListNode
	want  *ListNode
}

func TestExample1(t *testing.T) {
	example1 := Example{
		list1: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		list2: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
	}
	got := mergeTwoLists(example1.list1, example1.list2)
	assert.Equal(t, example1.want, got)
}

func TestExample2(t *testing.T) {
	example2 := Example{
		list1: nil,
		list2: nil,
		want:  nil,
	}
	got := mergeTwoLists(example2.list1, example2.list2)
	assert.Equal(t, example2.want, got)
}

func TestExample3(t *testing.T) {
	example3 := Example{
		list1: nil,
		list2: &ListNode{
			Val: 0,
		},
		want: &ListNode{
			Val: 0,
		},
	}
	got := mergeTwoLists(example3.list1, example3.list2)
	assert.Equal(t, example3.want, got)
}
