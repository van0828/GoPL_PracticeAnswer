package __3

import (
	"strings"
	"testing"
)

func BenchmarkString2Join(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"This", "is", "a", "benchmark", "test", "for", "join", "and", "plus"}
		strings.Join(input, " ")
	}
}

func BenchmarkString2Plus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"This", "is", "a", "benchmark", "test", "for", "join", "and", "plus"}
		var s, sep string
		for _, str := range input {
			s += sep + str
			sep = " "
		}
	}
}
