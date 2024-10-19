package sort

type Node[T any] struct {
	A int
	B T
}

func partition[T any](arr []Node[T], p, r int) int {
	x := arr[r].A
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j].A <= x {
			i += 1
			swap(&arr[i].A, &arr[j].A)
		}
	}
	swap(&arr[i+1].A, &arr[r].A)

	return i + 1
}

func QuickSort[T any](arr []Node[T], p, r int) {
	if p < r {
		q := partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}
