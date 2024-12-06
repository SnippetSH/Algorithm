package shortestpath

import "fmt"

type Distance [][]int
type Predecessor [][]int

const inf = int(^uint32(0) >> 1)

func FloydWarshall(start Distance) (Distance, Predecessor) {
	n := len(start)

	D := make(Distance, n)
	for i := 0; i < n; i++ {
		D[i] = make([]int, n)
	}

	pi := make(Predecessor, n)
	for i := 0; i < n; i++ {
		pi[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				D[i][j] = 0
				pi[i][j] = -1
			} else if start[i][j] < inf {
				D[i][j] = start[i][j]
				pi[i][j] = i
			} else {
				D[i][j] = inf
				pi[i][j] = -1
			}
		}
	}

	for k := 0; k < n; k++ {
		newD := make(Distance, n)
		for i := 0; i < n; i++ {
			newD[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if D[i][k] != inf && D[k][j] != inf {
					newD[i][j] = min(D[i][j], D[i][k]+D[k][j])
					if D[i][j] > D[i][k]+D[k][j] {
						pi[i][j] = pi[k][j]
					}
				} else {
					newD[i][j] = D[i][j]
				}
			}
		}
		D = newD
	}

	return D, pi
}

func PrintFloydWarshall(Pi Predecessor, i, j int) {
	if i == j {
		fmt.Println(i)
	} else if Pi[i][j] == -1 {
		fmt.Printf("no path from %d to %d exists", i, j)
	} else {
		PrintFloydWarshall(Pi, i, Pi[i][j])
		fmt.Println(j)
	}
}
