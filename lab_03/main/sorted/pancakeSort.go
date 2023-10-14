package sorted 

func searchMaxIndex(slice []int, lenl int) int {
	indx := 0
	maxl := slice[0]

	for i := 0; i < lenl; i++ {
		if slice[i] > maxl {
			maxl = slice[i]
			indx = i
		}
	}

	return indx
}

func flipSlice(slice []int, n int) {
	for i := 0; i < n; i++ {
		slice[i], slice[n] = slice[n], slice[i]
		n--
	}
}

func PancakeSort(slice []int) {
	lenl := len(slice)

	if lenl > 1 {
		for i := lenl; i > 1; i-- { 
			maxIndex := searchMaxIndex(slice, lenl)

			if maxIndex != lenl - 1 {
				if maxIndex != 0 {
					flipSlice(slice, maxIndex)
				}
				flipSlice(slice, lenl - 1)
			}

			lenl--
		}
	}
}

