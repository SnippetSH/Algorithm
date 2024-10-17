package dp

import (
	"fmt"
	"math"
)

func ExtendeBottomUpCutRod(p []int, n int) ([]int, []int) {
	r := make([]int, n+1)
	s := make([]int, n+1)
	r[0] = 0

	for j := 1; j <= n; j++ {
		q := -math.MaxInt32
		for i := 1; i <= j; i++ {
			if q < p[i]+r[j-i] {
				q = p[i] + r[j-i]
				s[j] = i
			}
		}
		r[j] = q
	}

	return r, s
}

func PrintCutRod(s []int, n int) {
	for n > 0 {
		fmt.Printf("%d ", s[n])
		n -= s[n]
	}
}
