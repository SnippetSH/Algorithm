package main

import (
	"fmt"
	"math/rand"
	"time"

	"algorithm/dp"
	"algorithm/greedy"
	"algorithm/sort"
)

func MakeRandomArray(N int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, N)

	for i := 0; i < len(arr); i++ {
		for {
			val := r.Intn(N)
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

	return arr
}

func CallSortingAlgorithms() {
	N := 100000
	arr := MakeRandomArray(N)

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
	carr3 := make([]sort.Node[interface{}], N)
	for i := range carr3 {
		carr3[i].A = arr[i]
	}
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

	// array 할당
	carr5 := make([]int, N)
	result := make([]int, N)
	copy(carr5, arr)
	// CountingSort
	start = time.Now()
	sort.CountingSort(carr5, result, N)
	end = time.Since(start).Microseconds()
	fmt.Println("CountingSort Excution Time: ", end, "µs")

	// fmt.Println("Sorted Array by CountingSort: ", result)
}

func CallDPAlgorithms() {
	e := []int{2, 4}
	x := []int{3, 2}
	a := [][]int{
		{7, 9, 3, 4, 8, 4},
		{8, 5, 6, 4, 5, 7},
	}
	t := [][]int{
		{2, 1, 1, 3, 4},
		{2, 1, 2, 2, 1},
	}
	L, l := dp.FastestWay(a, t, e, x, 6)

	dp.PrintStation(L, l, 6)

	fmt.Print("\n___________\n\n")

	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	r, s := dp.ExtendeBottomUpCutRod(p, 10)
	fmt.Println(s)

	fmt.Println("result cost:", r)
	dp.PrintCutRod(s, 10)

	fmt.Print("\n___________\n\n")

	X := string("ABCBDAB")
	Y := string("BDCABA")
	b, c := dp.LCS(X, Y)

	m := len(X)
	n := len(Y)
	fmt.Println(c[m][n])
	dp.PrintLCS(b, X, m, n)

	fmt.Print("\n___________\n\n")

	P := []int{30, 35, 15, 5, 10, 20, 25}
	M, S := dp.MatrixChainOrder(P)

	for _, v := range M {
		fmt.Println("Matrix Result:", v)
	}

	dp.PrintOptimalParens(S, 0, 5)
}

func CallGreedyAlgorithm() {
	f := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	s := []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}

	nodes := make([]sort.Node[int], 11)
	for i := range nodes {
		nodes[i].A = f[i]
		nodes[i].B = s[i]
	}

	n := greedy.ActivitySelection(nodes)
	fmt.Println(n)

	fmt.Print("\n___________\n\n")

	input := "aaaaaabbbbheellohieveryoneccc"
	output, tree := greedy.Encode(input)
	fmt.Println("Encoded Ouput:", output)

	output = greedy.Decode(output, tree)
	fmt.Println("Decoded Ouput:", output)
}

func main() {
	// CallSortingAlgorithms()
	// CallDPAlgorithms()

	CallGreedyAlgorithm()
}
