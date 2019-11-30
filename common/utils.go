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

func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}
