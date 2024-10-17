package dp

import (
	"fmt"
)

func FastestWay(a, t [][]int, e, x []int, n int) ([][]int, int) { // first: 각 노드에서 선택된 라인 번호, second: 최종 라인 번호
	s := make([][]int, 2)
	l := make([][]int, 2)

	for i := 0; i < 2; i++ {
		s[i] = make([]int, n+1)
		l[i] = make([]int, n+1)
	}

	s[0][0] = e[0] + a[0][0]
	s[1][0] = e[1] + a[1][0]

	for j := 1; j < n; j++ {
		if s[0][j-1] <= s[1][j-1]+t[1][j-1] {
			s[0][j] = s[0][j-1] + a[0][j]
			l[0][j] = 0
		} else {
			s[0][j] = s[1][j-1] + t[1][j-1] + a[0][j]
			l[0][j] = 1
		}

		if s[1][j-1] <= s[0][j-1]+t[0][j-1] {
			s[1][j] = s[1][j-1] + a[1][j]
			l[1][j] = 1
		} else {
			s[1][j] = s[0][j-1] + t[0][j-1] + a[1][j]
			l[1][j] = 0
		}
	}

	resultS := 0
	resultL := 0
	if s[0][n-1]+x[0] <= s[1][n-1]+x[1] {
		resultS = s[0][n-1] + x[0]
		resultL = 0
	} else {
		resultS = s[1][n-1] + x[1]
		resultL = 1
	}

	i := 0
	for i < 2 {
		fmt.Println("Each node's Minimum cost:", s[i])
		i++
	}
	fmt.Println("Final Cost:", resultS)
	return l, resultL
}

func PrintStation(nodeL [][]int, resultL, n int) {
	i := resultL
	fmt.Printf("line %d, station %d\n", i+1, n+1)

	for j := n - 1; j > 0; j-- {
		i = nodeL[i][j]
		fmt.Printf("line %d, station %d\n", i+1, j)
	}
}
