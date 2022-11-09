package sorting

func incrementBubble(arr []int) []int {
	sorted := arr

	for i, _ := range sorted {
		for j, _ := range sorted {
			if i == j {
				continue
			}

			if sorted[j] > sorted[i] { // Difference >/<
				temp := sorted[i]

				sorted[i] = sorted[j]
				sorted[j] = temp
			}
		}
	}

	return sorted
}

func decrementBubble(arr []int) []int {
	sorted := arr

	for i, _ := range sorted {
		for j, _ := range sorted {
			if i == j {
				continue
			}

			if sorted[j] < sorted[i] { // Difference >/<
				temp := sorted[i]

				sorted[i] = sorted[j]
				sorted[j] = temp
			}
		}
	}

	return sorted
}

func BubbleSort(arr []int, act int) []int {
	if act == 1 {
		return incrementBubble(arr)
	} else {
		return decrementBubble(arr)
	}
}
