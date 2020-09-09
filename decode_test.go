package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Test_Decode(t *testing.T) {
	for try := 0; try < 10000; try++ {
		str := randStringRunes(2 + rand.Intn(1000))
		k, b, err := encode(str)
		assert.Nil(t, err)
		rstr, err := decode(k, b)
		assert.Nil(t, err)
		assert.Equal(t, str, rstr)
	}
}
