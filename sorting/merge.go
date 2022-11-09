package sorting

func incrementMerge(L []int, R []int) []int {
	var sorted []int

	i := 0
	j := 0

	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			sorted = append(sorted, L[i])
			i++
		} else {
			sorted = append(sorted, R[j])
			j++
		}
	}

	for ; i < len(L); i++ {
		sorted = append(sorted, L[i])
	}
	for ; j < len(R); j++ {
		sorted = append(sorted, R[j])
	}

	return sorted
}

func decrementMerge(L []int, R []int) []int {
	var sorted []int

	i := 0
	j := 0

	for i < len(L) && j < len(R) {
		if L[i] > R[j] {
			sorted = append(sorted, L[i])
			i++
		} else {
			sorted = append(sorted, R[j])
			j++
		}
	}

	for ; i < len(L); i++ {
		sorted = append(sorted, L[i])
	}
	for ; j < len(R); j++ {
		sorted = append(sorted, R[j])
	}

	return sorted
}

func MergeSort(arr []int, act int) []int {
	sorted := arr

	if len(sorted) < 2 {
		return sorted
	}

	m := len(arr) / 2
	L := MergeSort(sorted[:m], act)
	R := MergeSort(sorted[m:], act)

	if act == 1 {
		return incrementMerge(L, R)
	} else {
		return decrementMerge(L, R)
	}
}
