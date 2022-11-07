package sorting

func incrementSort(arr []int) []int {
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

func decrementSort(arr []int) []int {
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

func InsertionSort(arr []int, act int) []int {
	if act == 1 {
		return incrementSort(arr)
	} else {
		return decrementSort(arr)
	}
}
