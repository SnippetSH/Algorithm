package greedy

import (
	"algorithm/sort"
	"fmt"
)

func ActivitySelection(a []sort.Node[int]) int {
	sort.QuickSort(a, 0, len(a)-1)

	result := []int{1}
	n := a[0].A
	cnt := 1
	for i, node := range a[1:] {
		if node.B >= n {
			result = append(result, i+2)
			cnt++
			n = node.A
		}
	}

	fmt.Println(result)
	return cnt
}
