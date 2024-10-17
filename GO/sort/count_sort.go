package sort

func CountingSort(arr, brr []int, k int) {
	crr := make([]int, k+1)
	for i := 0; i <= k; i++ {
		crr[i] = 0
	}

	l := len(arr)
	for j := 0; j < l; j++ {
		crr[arr[j]] += 1
	}

	for i := 1; i < k; i++ {
		crr[i] = crr[i] + crr[i-1]
	}

	for j := l - 1; j >= 0; j-- {
		brr[crr[arr[j]]-1] = arr[j]
		crr[arr[j]] -= 1
	}
}

func RadixSort(arr []int, d int) {
	for i := 1; i < d; i++ {
		// use stable sort like counting sort to sort array on digit 'i'
	}
}
