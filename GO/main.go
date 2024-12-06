package main

import (
	"fmt"
	"math/rand"
	"time"

	mst "algorithm/MST"
	"algorithm/disjoint"
	"algorithm/dp"
	"algorithm/graph"
	"algorithm/greedy"
	shortestpath "algorithm/shortest_path"
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
	fmt.Println("Original String's length:", len(input))
	fmt.Println("Encoded Ouput:", output)
	fmt.Println("Encoded String's length:", len(output))
	output = greedy.Decode(output, tree)
	fmt.Println("Decoded Ouput:", output)
}

func testDFS() {
	G := graph.NewGraph()
	G.AddEdge(1, 2)
	G.AddEdge(2, 5)
	G.AddEdge(5, 4)
	G.AddEdge(1, 4)
	G.AddEdge(4, 2)
	G.AddEdge(3, 5)
	G.AddEdge(3, 6)

	graph.DFS(G)
	for i := 1; i <= 6; i++ {
		v := G.V[i]
		fmt.Printf("%d, In: %d, Out: %d\n", v.ID, v.In, v.Out)
		if v.Predecessor != nil {
			fmt.Printf("%d, Predecessor: %d\n", v.ID, v.Predecessor.ID)
		}
	}
}

func testBFS() {
	G := graph.NewGraph()
	G.AddEdge(1, 5)
	G.AddEdge(5, 1)
	G.AddEdge(1, 2)
	G.AddEdge(2, 1)

	G.AddEdge(2, 6)
	G.AddEdge(6, 2)
	G.AddEdge(6, 3)
	G.AddEdge(3, 6)
	G.AddEdge(6, 7)
	G.AddEdge(7, 6)

	G.AddEdge(3, 7)
	G.AddEdge(7, 3)
	G.AddEdge(3, 4)
	G.AddEdge(4, 3)
	G.AddEdge(7, 4)
	G.AddEdge(4, 7)

	G.AddEdge(7, 8)
	G.AddEdge(8, 7)
	G.AddEdge(8, 4)
	G.AddEdge(4, 8)

	graph.BFS(G, 2)
	for i := 1; i <= 8; i++ {
		v := G.V[i]
		fmt.Printf("%d, D: %d\n", v.ID, v.In)
		if v.Predecessor != nil {
			fmt.Printf("%d, Predecessor: %d\n", v.ID, v.Predecessor.ID)
		}
	}
}

func testTopologicalSort() {
	t := &graph.TopologicalSorterContext{}
	t.SetSorter(&graph.DFSSorter{})

	G := graph.NewGraph()
	G.AddEdge(1, 4)
	G.AddEdge(1, 5)
	G.AddEdge(2, 5)
	G.AddEdge(4, 6)
	G.AddEdge(4, 5)
	G.AddEdge(6, 9)
	G.AddEdge(7, 6)
	G.AddEdge(7, 8)
	G.AddEdge(8, 9)

	result := t.Run(G)
	fmt.Println(result)
}

func testDisjointSet() {
	G := disjoint.NewGraph()
	c := G.NewVertex()
	h := G.NewVertex()
	e := G.NewVertex()
	b := G.NewVertex()
	f := G.NewVertex()
	d := G.NewVertex()
	g := G.NewVertex()

	fmt.Printf("c: %d \\", c.ID)
	fmt.Printf("h: %d \\", h.ID)
	fmt.Printf("e: %d \\", e.ID)
	fmt.Printf("b: %d \\", b.ID)
	fmt.Printf("f: %d \\", f.ID)
	fmt.Printf("d: %d \\", d.ID)
	fmt.Printf("g: %d \n", g.ID)

	disjoint.Union(G, c, h)
	disjoint.Union(G, c, e)
	disjoint.Union(G, b, h)
	disjoint.Union(G, f, d)
	disjoint.Union(G, d, g)

	fmt.Println(G.String())

	disjoint.Union(G, b, g)
	fmt.Println(G.String())
}

func testPrim() {
	G := mst.NewGraph()

	a := G.NewVertex("a")
	b := G.NewVertex("b")
	c := G.NewVertex("c")
	d := G.NewVertex("d")
	e := G.NewVertex("e")
	f := G.NewVertex("f")
	g := G.NewVertex("g")
	h := G.NewVertex("h")
	i := G.NewVertex("i")

	G.NewEdge(a, b, 4)
	G.NewEdge(a, h, 8)
	G.NewEdge(b, c, 8)
	G.NewEdge(c, i, 2)
	G.NewEdge(c, d, 7)
	G.NewEdge(d, e, 9)
	G.NewEdge(d, f, 14)
	G.NewEdge(e, f, 10)
	G.NewEdge(f, g, 2)
	G.NewEdge(f, c, 4)
	G.NewEdge(g, h, 1)
	G.NewEdge(h, i, 7)
	G.NewEdge(g, i, 6)
	G.NewEdge(b, h, 11)

	result := mst.Prim(G, a)
	total := 0
	for _, e := range result {
		total += e.Weight
		fmt.Printf("%s's weight is %d\n", e.Name, e.Weight)
	}
	fmt.Println("Total Weight: ", total)
}

func testKruskal() {
	G := mst.KNewGraph()

	a := G.NewVertex("a")
	b := G.NewVertex("b")
	c := G.NewVertex("c")
	d := G.NewVertex("d")
	e := G.NewVertex("e")
	f := G.NewVertex("f")
	g := G.NewVertex("g")
	h := G.NewVertex("h")
	i := G.NewVertex("i")

	G.NewEdge(a, b, 4)
	G.NewEdge(a, h, 8)
	G.NewEdge(b, c, 8)
	G.NewEdge(c, i, 2)
	G.NewEdge(c, d, 7)
	G.NewEdge(d, e, 9)
	G.NewEdge(d, f, 14)
	G.NewEdge(e, f, 10)
	G.NewEdge(f, g, 2)
	G.NewEdge(f, c, 4)
	G.NewEdge(g, h, 1)
	G.NewEdge(h, i, 7)
	G.NewEdge(g, i, 6)
	G.NewEdge(b, h, 11)

	result := mst.Kruskal(G)
	total := 0
	for _, e := range result {
		total += e.Weight
		fmt.Printf("%s's weight is %d\n", e.Name, e.Weight)
	}
	fmt.Println("Total Weight: ", total)
}

func testDijkstra() {
	g := shortestpath.NewGraph()

	s := g.NewVertex("s")
	t := g.NewVertex("t")
	x := g.NewVertex("x")
	y := g.NewVertex("y")
	z := g.NewVertex("z")

	g.NewEdge(s, t, 10)
	g.NewEdge(s, y, 5)
	g.NewEdge(t, y, 2)
	g.NewEdge(t, x, 1)
	g.NewEdge(y, t, 3)
	g.NewEdge(y, z, 2)
	g.NewEdge(y, x, 9)
	g.NewEdge(x, z, 4)
	g.NewEdge(z, x, 6)
	g.NewEdge(z, s, 7)

	shortestpath.Dijkstra(g, s)
	for _, v := range g.V {
		fmt.Printf("%s's shortest distance from s: %d\n", v.ID, v.Distance)
	}
}

func testBellmanFord() {
	g := shortestpath.NewGraph()

	s := g.NewVertex("s")
	t := g.NewVertex("t")
	x := g.NewVertex("x")
	y := g.NewVertex("y")
	z := g.NewVertex("z")

	g.NewEdge(s, t, 6)
	g.NewEdge(s, y, 7)

	g.NewEdge(t, y, 8)
	g.NewEdge(t, x, 5)
	g.NewEdge(t, z, -4)

	g.NewEdge(y, z, 9)
	g.NewEdge(y, x, -3)

	g.NewEdge(x, t, -2)

	g.NewEdge(z, x, 7)
	g.NewEdge(z, s, 2)

	shortestpath.BellmanFord(g, s)
	for _, v := range g.V {
		fmt.Printf("%s's shortest distance from s: %d\n", v.ID, v.Distance)
	}
}

func testFloyWarshall() {
	const inf = int(^uint32(0) >> 1)

	var distance = shortestpath.Distance{
		{0, 3, 8, inf, -4},
		{inf, 0, inf, 1, 7},
		{inf, 4, 0, inf, inf},
		{2, inf, -5, 0, inf},
		{inf, inf, inf, 6, 0},
	}

	D, Pi := shortestpath.FloydWarshall(distance)
	fmt.Println(D)

	shortestpath.PrintFloydWarshall(Pi, 0, 3)
}

func main() {
	// CallSortingAlgorithms()
	// CallDPAlgorithms()

	// CallGreedyAlgorithm()
	// testDFS()
	// testBFS()
	// testTopologicalSort()
	// testDisjointSet()
	// testPrim()
	// testKruskal()
	// testDijkstra()
	// testBellmanFord()
	testFloyWarshall()
}
