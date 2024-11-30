package sort

type Node[T any] struct {
	A int
	B T
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
