// Package huffman is for generating huffman encoding code words from weights table.
package huffman

import "container/heap"

// Huffman computes Huffman encoding code words from given weight table.
// e.g. Huffman([0.5, 0.25, 0.25]) => ["0", "10", "11"]
func Huffman(weights []float64) []string {
	if len(weights) == 0 {
		return nil
	}
	if len(weights) == 1 {
		return []string{"0"}
	}

	var nodesHeap nodes
	for i, w := range weights {
		nodesHeap = append(nodesHeap, &node{i: i, weight: w})
	}
	heap.Init(&nodesHeap)

	for len(nodesHeap) > 1 {
		min1 := heap.Pop(&nodesHeap).(*node)
		min2 := heap.Pop(&nodesHeap).(*node)
		heap.Push(&nodesHeap, &node{
			weight: min1.weight + min2.weight,
			right:  min1,
			left:   min2,
		})
	}

	codes := make([]string, len(weights))
	var traverse func(*node, []byte)
	traverse = func(n *node, acc []byte) {
		if n.right == nil && n.left == nil {
			codes[n.i] = string(acc)
			return
		}
		traverse(n.right, append(acc, '0'))
		traverse(n.left, append(acc, '1'))
	}

	root := heap.Pop(&nodesHeap).(*node)
	traverse(root, nil)

	return codes
}

type node struct {
	i           int
	weight      float64
	right, left *node
}

type nodes []*node

func (n nodes) Less(x, y int) bool {
	return n[x].weight < n[y].weight
}

func (n nodes) Swap(x, y int) {
	n[x], n[y] = n[y], n[x]
}

func (n nodes) Len() int {
	return len(n)
}

func (n *nodes) Push(x interface{}) {
	*n = append(*n, x.(*node))
}

func (n *nodes) Pop() interface{} {
	old := *n
	ret := old[len(old)-1]
	*n = old[:len(old)-1]
	return ret
}
