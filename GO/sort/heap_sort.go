package sort

type A struct {
	Arr      []int
	HeapSize int
}

func maxHeapify(a *A, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	// heapSize를 활용하여 범위 제한
	if left < a.HeapSize && a.Arr[i] < a.Arr[left] {
		largest = left
	}
	if right < a.HeapSize && a.Arr[largest] < a.Arr[right] {
		largest = right
	}

	if largest != i {
		swap(&a.Arr[i], &a.Arr[largest])
		maxHeapify(a, largest)
	}
}

func buildMaxHeap(a *A) {
	a.HeapSize = len(a.Arr)
	for i := len(a.Arr)/2 - 1; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

func HeapSort(a *A) {
	buildMaxHeap(a)
	for i := len(a.Arr) - 1; i >= 1; i-- {
		swap(&a.Arr[0], &a.Arr[i])
		a.HeapSize -= 1
		maxHeapify(a, 0)
	}
}
