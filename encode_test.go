package main

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Huffman(t *testing.T) {
	str := "abcabcaaaa"
	m := countFrequency(str)
	assert.Equal(t, 6, m['a'])
	assert.Equal(t, 2, m['b'])
	assert.Equal(t, 2, m['c'])
	tree := buildHuffmanTree(m)
	assert.Equal(t, 'a', tree.Right.C)
	assert.Equal(t, 'b', tree.Left.Left.C)
	assert.Equal(t, 'c', tree.Left.Right.C)
	dst := &map[rune][]bool{}
	buildBoolsFromTree(tree, make([]bool, 0), dst)
	assert.Equal(t, []bool{true}, (*dst)['a'])
	assert.Equal(t, []bool{false, false}, (*dst)['b'])
	assert.Equal(t, []bool{false, true}, (*dst)['c'])
	bytes, err := json.Marshal(dst)
	assert.Nil(t, err)
	assert.Equal(t, "{\"97\":[true],\"98\":[false,false],\"99\":[false,true]}", string(bytes))
	r := buildBinFromBools(str, dst)
	assert.Equal(t, 6, len(r))
	s := base64.RawStdEncoding.EncodeToString(r)
	assert.Equal(t, "AAAADox8", s)
}
