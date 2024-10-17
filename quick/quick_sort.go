package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

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

func quickSort(arr []int, p, r int) {
	if p < r {
		q := partition(arr, p, r)
		quickSort(arr, p, q-1)
		quickSort(arr, q+1, r)
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

		start := time.Now()
		quickSort(carr, 0, 10000-1)
		end := time.Since(start)
		fmt.Println("Excution Time: ", end)
		totalTime += end

		// fmt.Println(carr)
		s += carr[0]

	}
	fmt.Println(totalTime / 100)
}
