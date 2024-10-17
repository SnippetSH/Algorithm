package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(arr []int, len int) []int {
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
	for i := 0; i < 10; i++ {
		carr := make([]int, 10000)
		copy(carr, arr)

		start := time.Now()
		insertionSort(carr, 10000)
		fmt.Println("Excution Time: ", time.Since(start))

		fmt.Println(carr)

		break
	}
	fmt.Println(s)
}
