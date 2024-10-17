package sort

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
