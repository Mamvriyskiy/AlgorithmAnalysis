package sorted

func HeapSort(slice []int) {
	n := len(slice)

	for i := n / 2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]

		heapify(slice, i, 0)
	}
}

func heapify(slice []int, n int, i int) {
	largest := i
	l := 2 * i + 1
	r := 2 * i + 2

	if l < n && slice[l] > slice[largest] {
		largest = l
	}

	if r < n && slice[r] > slice[largest] {
		largest = r
	}

	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]

		heapify(slice, n, largest)
	}
}