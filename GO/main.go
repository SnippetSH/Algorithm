package main

import (
	"fmt"
	"math/rand"
	"time"

	"algorithm/sort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	N := 100000
	arr := make([]int, N)

	for i := 0; i < len(arr); i++ {
		for {
			val := rand.Intn(N)
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
	/*
		s := 0
		var totalTime time.Duration = 0
		for i := 0; i < 100; i++ {
			carr := make([]int, 10000)
			copy(carr, arr)
			a := &sort.A{
				Arr:      carr,
				HeapSize: len(carr),
			}

			start := time.Now()
			sort.HeapSort(a)
			end := time.Since(start)
			fmt.Println("Excution Time: ", end)
			totalTime += end

			// fmt.Println(carr)
			s += carr[0]
		}
		fmt.Println(totalTime / 100)
	*/

	// array 할당
	carr1 := make([]int, N)
	copy(carr1, arr)
	a := &sort.A{
		Arr:      carr1,
		HeapSize: len(carr1),
	}
	// HeapSort
	start := time.Now()
	sort.HeapSort(a)
	end := time.Since(start).Microseconds()
	fmt.Println("HeapSort Excution Time: ", end, "µs")

	// array 할당
	carr2 := make([]int, N)
	copy(carr2, arr)
	// MergeSort
	start = time.Now()
	sort.MergeSort(carr2, 0, len(carr2)-1)
	end = time.Since(start).Microseconds()
	fmt.Println("MergeSort Excution Time: ", end, "µs")

	// array 할당
	carr3 := make([]int, N)
	copy(carr3, arr)
	// QuickSort
	start = time.Now()
	sort.QuickSort(carr3, 0, len(carr3)-1)
	end = time.Since(start).Microseconds()
	fmt.Println("QuickSort Excution Time: ", end, "µs")

	// array 할당
	carr4 := make([]int, N)
	copy(carr4, arr)
	// InsertionSort
	start = time.Now()
	sort.InsertionSort(carr4, len(carr4))
	end = time.Since(start).Milliseconds()
	fmt.Println("InsertionSort Excution Time: ", end, "ms")
}
