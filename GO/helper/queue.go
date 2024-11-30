package helper

type Comparator interface {
	Cmp() int
}

type PriorityQueue[T Comparator] []*T

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return (*pq[i]).Cmp() < (*pq[j]).Cmp()
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(*T))
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Queue[T any] []*T

func (q *Queue[T]) Pop() any {
	old := *q
	item := old[0]
	*q = old[1:]
	return item
}

func (q *Queue[T]) Push(x any) {
	*q = append(*q, x.(*T))
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Top() any {
	return (*q)[0]
}
