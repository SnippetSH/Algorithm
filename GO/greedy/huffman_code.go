package greedy

import (
	"algorithm/helper"
	"container/heap"
	"fmt"
)

type Node struct {
	Char  rune  // 문자
	Freq  int   // 빈도수
	Left  *Node // 왼쪽 자식
	Right *Node // 오른쪽 자식
}

func (n Node) Cmp() int {
	return n.Freq
}

func (n Node) SetIndex(i int) {

}

func (n Node) Index() *int {
	return &n.Freq
}

func GenerateCode(code map[rune]string, node *Node, prefix string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		code[node.Char] = prefix
		return
	}

	GenerateCode(code, node.Left, prefix+"0")
	GenerateCode(code, node.Right, prefix+"1")
}

func HuffmanCode(freqs map[rune]int) (map[rune]string, *Node) {
	pq := make(helper.PriorityQueue[Node], 0)
	heap.Init(&pq)

	for char, freq := range freqs {
		heap.Push(
			&pq,
			&Node{
				Char:  char,
				Freq:  freq,
				Left:  nil,
				Right: nil,
			},
		)
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)

		heap.Push(
			&pq,
			&Node{
				Freq:  left.Freq + right.Freq,
				Left:  left,
				Right: right,
			},
		)
	}

	root := heap.Pop(&pq).(*Node)
	code := make(map[rune]string)
	GenerateCode(code, root, "")

	return code, root
}

func Encode(input string) (string, *Node) {
	freqs := make(map[rune]int)

	for _, char := range input {
		freqs[char] += 1
	}

	code, tree := HuffmanCode(freqs)
	i := 0
	for key, val := range code {
		fmt.Printf("%d| %c: %s \t", i, key, val)
		i++
	}
	fmt.Println()

	result := ""
	for _, char := range input {
		result += code[char]
	}

	return result, tree
}

func Decode(input string, tree *Node) string {
	if input == "" {
		return ""
	}

	result := ""
	current := tree

	for _, bit := range input {
		if bit == '0' {
			current = current.Left
		} else {
			current = current.Right
		}

		if current.Left == nil && current.Right == nil {
			result += string(current.Char)
			current = tree
		}
	}

	return result
}
