package sort

func qswap[T any](a, b *Node[T]) {
	t := *a
	*a = *b
	*b = t
}

func partition[T any](arr []Node[T], p, r int) int {
	x := arr[r].A
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j].A <= x {
			i += 1
			qswap(&arr[i], &arr[j])
		}
	}
	qswap(&arr[i+1], &arr[r])

	return i + 1
}

func QuickSort[T any](arr []Node[T], p, r int) {
	if p < r {
		q := partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}
