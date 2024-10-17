package main

import (
	"fmt"
	"math/rand"
	"time"
)

type A struct {
	arr      []int
	heapSize int
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func maxHeapify(a *A, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	// heapSize를 활용하여 범위 제한
	if left < a.heapSize && a.arr[i] < a.arr[left] {
		largest = left
	}
	if right < a.heapSize && a.arr[largest] < a.arr[right] {
		largest = right
	}

	if largest != i {
		swap(&a.arr[i], &a.arr[largest])
		maxHeapify(a, largest)
	}
}

func buildMaxHeap(a *A) {
	a.heapSize = len(a.arr)
	for i := len(a.arr)/2 - 1; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

func heapSort(a *A) {
	buildMaxHeap(a)
	for i := len(a.arr) - 1; i >= 1; i-- {
		swap(&a.arr[0], &a.arr[i])
		a.heapSize -= 1
		maxHeapify(a, 0)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10000)

	for i := 0; i < len(arr); i++ {
		for {
			val := rand.Intn(10000)
			isDuplicate := false

			for j := 0; j < i; j++ {
				if val == arr[j] {
					isDuplicate = true
					break
				}
			}

			if !isDuplicate {
				arr[i] = val
				break
			}
		}
	}

	// fmt.Println(arr)
	s := 0
	var totalTime time.Duration = 0
	for i := 0; i < 100; i++ {
		carr := make([]int, 10000)
		copy(carr, arr)
		a := &A{
			arr:      carr,
			heapSize: len(carr),
		}

		start := time.Now()
		heapSort(a)
		end := time.Since(start)
		fmt.Println("Excution Time: ", end)
		totalTime += end

		// fmt.Println(carr)
		s += carr[0]
	}
	fmt.Println(totalTime / 100)
}
