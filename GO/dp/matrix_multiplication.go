package dp

import (
	"fmt"
	"math"
)

func MatrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p) - 1 // p의 값은 A1라는 행렬의 세로 * 가로 와, A2라는 행렬의 세로 * 가로... 의 연속 이기에 항상 행렬의 갯수보다 1이 많음.

	m := make([][]int, n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
		m[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		m[i][i] = 0
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt32

			for k := i; k < j; k++ {
				cost := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1]

				if cost < m[i][j] {
					m[i][j] = cost
					s[i][j] = k
				}
			}
		}
	}

	return m, s
}

func PrintOptimalParens(s [][]int, i, j int) {
	if i == j {
		fmt.Printf("A[%d]", i+1)
	} else {
		fmt.Print("(")
		PrintOptimalParens(s, i, s[i][j])
		PrintOptimalParens(s, s[i][j]+1, j)
		fmt.Print(")")
	}
}
