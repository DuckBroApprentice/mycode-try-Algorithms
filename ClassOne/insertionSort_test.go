package classone

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type quene struct {
		name string
		nums []int
		want []int
	}
	var test []quene
	test = []quene{
		{
			name: "test1",
			nums: []int{6, 2, 8, 4, 9, 10, 7, 3},
			want: []int{2, 3, 4, 6, 7, 8, 9, 10},
		},
		{
			name: "test2",
			nums: []int{-1, 5, 3, 4, 0},
			want: []int{-1, 0, 3, 4, 5},
		},
		{
			name: "test3",
			nums: []int{28, -6, -22, 8, 44, 17},
			want: []int{-22, -6, 8, 17, 28, 44},
		},
	}
	for _, tt := range test {
		InsertionSort(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("want: %v, nums: %v", tt.want, tt.nums)
		}
	}

}
