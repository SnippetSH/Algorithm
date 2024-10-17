package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func merge(arr []int, p, q, r int) {
	n1 := q - p + 1 // alpha
	n2 := r - q     // beta
	/* alpha and beta are assumed already sorted */

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	// copy data to temparay array
	for i := 0; i < n1; i++ {
		L[i] = arr[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[q+j+1]
	}

	/*
		만약 하나의 배열에 대해 끝에 도달했을 때,
		남은 배열이 정상적으로 arr에 저장되는 것을 보장하기 위해
		inf를 사용함.
	*/
	inf := math.MaxInt32
	L[n1], R[n2] = inf, inf
	i, j := 0, 0

	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i += 1
		} else {
			arr[k] = R[j]
			j += 1
		}
	}
}

func mergeSort(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(arr, p, q)   // T(n/2)
		mergeSort(arr, q+1, r) // T(n/2)
		merge(arr, p, q, r)    // \theta(n)
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
		mergeSort(carr, 0, 10000-1)
		end := time.Since(start)
		fmt.Println("Excution Time: ", end)
		totalTime += end

		// fmt.Println(carr)
		s += carr[0]

	}
	fmt.Println(totalTime / 100)
}
