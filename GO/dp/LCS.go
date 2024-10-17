package dp

import "fmt"

func LCS(X, Y string) ([][]string, [][]int) {
	m := len(X)
	n := len(Y)

	b := make([][]string, m+1)
	c := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		b[i] = make([]string, n+1)
		c[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		c[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		c[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = "left-up"
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = "up"
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = "left"
			}
		}
	}

	return b, c
}

func PrintLCS(b [][]string, X string, i, j int) {
	if i == 0 || j == 0 {
		return
	}

	if b[i][j] == "left-up" {
		PrintLCS(b, X, i-1, j-1)
		fmt.Printf("%c ", X[i-1])
	} else if b[i][j] == "up" {
		PrintLCS(b, X, i-1, j)
	} else {
		PrintLCS(b, X, i, j-1)
	}
}
