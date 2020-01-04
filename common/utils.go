package common

import "testing"

import "reflect"

func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}

func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}

func ChanToSlice(ch chan int) []int {
	out := []int{}
	for v := range ch {
		out = append(out, v)
	}
	return out
}
