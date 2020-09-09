package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"sort"
)

func countFrequency(str string) map[rune]int {
	res := map[rune]int{}
	for _, c := range str {
		res[c]++
	}
	return res
}

func buildHuffmanTree(m map[rune]int) *node {
	h := &nodeHeap{}
	for c, num := range m {
		n := &node{
			C:   c,
			Tot: num,
		}
		heap.Push(h, n)
	}
	sort.Sort(h)
	heap.Init(h)
	for h.Len() > 1 {
		a := h.Pop().(*node)
		b := h.Pop().(*node)
		n := &node{
			Left:  a,
			Right: b,
			Tot:   a.Tot + b.Tot,
		}
		heap.Push(h, n)
	}
	return h.Pop().(*node)
}

func buildBoolsFromTree(n *node, now []bool, dst *map[rune][]bool) {
	if n.isLeaf() {
		(*dst)[n.C] = append([]bool{}, now...)
		return
	}
	buildBoolsFromTree(n.Left, append(now, false), dst)
	buildBoolsFromTree(n.Right, append(now, true), dst)
}

func buildBinFromBools(str string, b *map[rune][]bool) []byte {
	s := []bool{}
	for _, c := range str {
		s = append(s, (*b)[c]...)
	}
	l := len(s)
	res := make([]byte, 0)
	for k := 0; k < 4; k++ {
		x := (l >> (8 * (3 - k))) & ((1 << 8) - 1)
		res = append(res, byte(x))
	}
	for k := 0; k < (l+7)/8; k++ {
		r := 0
		for i := 0; i < 8; i++ {
			r <<= 1
			if k*8+i < l && s[k*8+i] {
				r++
			}
		}
		res = append(res, byte(r))
	}
	return res
}

// encode emit binary and key mapping
func encode(str string) ([]byte, []byte, error) {
	if len(str) < 2 {
		return nil, nil, fmt.Errorf("text is too few")
	}
	m := countFrequency(str)
	t := buildHuffmanTree(m)
	b := &map[rune][]bool{}
	buildBoolsFromTree(t, make([]bool, 0), b)
	tb, err := json.Marshal(t)
	if err != nil {
		return nil, nil, err
	}
	d := buildBinFromBools(str, b)
	return tb, d, nil
}
