package sort

func partition(arr []int, p, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= x {
			i += 1
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[r])

	return i + 1
}

func QuickSort(arr []int, p, r int) {
	if p < r {
		q := partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}
