package sorted

func ShakerSort(slice []int) {
	n := len(slice) - 1
	for i := 0; i < n + 1; i++ {
		for j := i; j < n; j++ {
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}

		for k := n; k > i; k-- {
			if slice[k] < slice[k - 1] {
				slice[k], slice[k - 1] = slice[k - 1], slice[k]
			}
		}

		n--
	}
}
