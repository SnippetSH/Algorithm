package sort

func InsertionSort(arr []int, len int) []int {
	for j := 1; j < len; j++ {
		key := arr[j]

		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i = i - 1
		}

		arr[i+1] = key
	}

	return arr
}
