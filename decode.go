package main

import (
	"encoding/json"
	"fmt"
)

func getSeq(textb []byte) *[]bool {
	l := 0
	for k := 0; k < 4; k++ {
		l <<= 8
		l += int(textb[k])
	}
	textb = textb[4:]
	seq := make([]bool, 0)
	for k, b := range textb {
		for i := 0; i < 8 && k*8+i < l; i++ {
			x := b >> (7 - i)
			seq = append(seq, (x&1 == 1))
		}
	}
	return &seq
}

func decode(keyb []byte, textb []byte) (string, error) {
	var key node
	if err := json.Unmarshal(keyb, &key); err != nil {
		return "", err
	}
	if len(textb) < 5 {
		return "", fmt.Errorf("text is too few")
	}
	seq := getSeq(textb)
	res := ""
	var dfs func(p int, n *node) int
	dfs = func(p int, n *node) int {
		if n.isLeaf() {
			res += string(n.C)
			return p
		}
		if p >= len(*seq) {
			panic(res)
		}
		if b := (*seq)[p]; !b {
			return dfs(p+1, n.Left)
		}
		return dfs(p+1, n.Right)
	}
	for p := 0; p < len(*seq); p = dfs(p, &key) {
	}
	return res, nil
}
