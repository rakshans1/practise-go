package main

import "fmt"

func twoSum(nums []int, target int) []int {
	compliments := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := compliments[target-b]; ok {
			return []int{j, i}
		}
		compliments[b] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
