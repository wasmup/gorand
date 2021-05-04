package main

import (
	"strings"
	"testing"
)

var s string

func BenchmarkShuffleText2(b *testing.B) {
	a := strings.Fields(txt)
	for i := 0; i < b.N; i++ {
		s = shuffleText2(a)
	}
}
func BenchmarkShuffleText1(b *testing.B) {
	a := strings.Fields(txt)
	for i := 0; i < b.N; i++ {
		s = shuffleText1(a)
	}
}
