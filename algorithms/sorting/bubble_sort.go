package sorting

// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func BubbleSortString(list []string) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
